data "aws_caller_identity" "current" {}

data "aws_region" "current" {}

data "aws_iam_roles" "administrator_access" {
  path_prefix = "/aws-reserved/sso.amazonaws.com/"
  name_regex  = "^AWSReservedSSO_AdministratorAccess_.*$"
}

locals {
  ## Construct the policy to include the current account root as a principal. Useful for testing. Not useful in production.
  policy = {
    "EnableIAMUserPermissions" = {
      sid    = "EnableIAMUserPermissions"
      effect = "Allow"
      actions = [
        "kms:*"
      ]
      resources = ["*"]
      principals = {
        "AWS" = ["arn:aws:iam::${data.aws_caller_identity.current.account_id}:root"]
      }
      condition = [
        {
          test     = "ArnEquals"
          variable = "aws:PrincipalArn"
          values   = ["arn:aws:iam::${data.aws_caller_identity.current.account_id}:root"]
        }
      ]
    }
    "AllowAccountKeyManagement" = {
      sid    = "AllowAccountKeyManagement"
      effect = "Allow"
      actions = [
        "kms:*"
      ]
      resources = ["arn:aws:kms:*:${data.aws_caller_identity.current.account_id}:key/*"]
      principals = {
        "AWS" = concat(
          tolist(data.aws_iam_roles.administrator_access.arns),
          ["arn:aws:iam::020127659860:role/github-actions-deploy-role-terraform"],
        )
      }
      condition = [
        {
          test     = "StringEquals"
          variable = "aws:PrincipalAccount"
          values   = [data.aws_caller_identity.current.account_id]
        },
        {
          test     = "ArnEquals"
          variable = "aws:PrincipalArn"
          values = concat(
            tolist(data.aws_iam_roles.administrator_access.arns),
            ["arn:aws:iam::020127659860:role/github-actions-deploy-role-terraform"],
          )
        }
      ]
    }
  }
}

module "kms_key" {
  # checkov:skip=CKV_TF_1: trusted registry source
  source                  = "terraform.registry.launch.nttdata.com/module_primitive/kms_key/aws"
  version                 = "~> 0.1"
  description             = "Terratest KMS Key"
  key_usage               = "ENCRYPT_DECRYPT"
  enable_key_rotation     = true
  rotation_period_in_days = 365
  multi_region            = true
  deletion_window_in_days = 7
  tags = {
    "Environment" = "Test"
  }
}

module "kms_key_policy" {
  source = "../../"

  key_id                             = module.kms_key.key_id
  policy                             = local.policy
  bypass_policy_lockout_safety_check = var.bypass_policy_lockout_safety_check

}
