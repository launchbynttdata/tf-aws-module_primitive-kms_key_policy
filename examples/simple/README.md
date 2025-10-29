# simple

<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | ~> 1.0 |
| <a name="requirement_aws"></a> [aws](#requirement\_aws) | ~> 5.0 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_aws"></a> [aws](#provider\_aws) | 5.100.0 |

## Modules

| Name | Source | Version |
|------|--------|---------|
| <a name="module_kms_key_policy"></a> [kms\_key\_policy](#module\_kms\_key\_policy) | ../../ | n/a |

## Resources

| Name | Type |
|------|------|
| [aws_kms_key.example](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/kms_key) | resource |
| [aws_caller_identity.current](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/caller_identity) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_bypass_policy_lockout_safety_check"></a> [bypass\_policy\_lockout\_safety\_check](#input\_bypass\_policy\_lockout\_safety\_check) | (Optional) A boolean flag to indicate whether to bypass the KMS key policy lockout safety check. Defaults to false. | `bool` | `false` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_policy_id"></a> [policy\_id](#output\_policy\_id) | The ID of the KMS Key Policy resource. |
<!-- END_TF_DOCS -->
