data "aws_caller_identity" "current" {}

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
    }
  }
}

resource "aws_kms_key" "example" {
  description             = "Terratest KMS Key"
  key_usage               = "ENCRYPT_DECRYPT"
  policy                  = ""
  enable_key_rotation     = true
  rotation_period_in_days = 365
  multi_region            = false
  tags = {
    "Environment" = "Test"
  }
}

module "kms_key_policy" {
  source = "../../"

  key_id                             = aws_kms_key.example.key_id
  policy                             = local.policy
  bypass_policy_lockout_safety_check = var.bypass_policy_lockout_safety_check

}
