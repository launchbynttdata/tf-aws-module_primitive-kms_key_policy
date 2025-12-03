# data_policy_document

<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | ~> 1.5 |
| <a name="requirement_aws"></a> [aws](#requirement\_aws) | ~> 5.0 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_aws"></a> [aws](#provider\_aws) | 5.100.0 |

## Modules

| Name | Source | Version |
|------|--------|---------|
| <a name="module_kms_key"></a> [kms\_key](#module\_kms\_key) | terraform.registry.launch.nttdata.com/module_primitive/kms_key/aws | ~> 0.1 |
| <a name="module_kms_key_policy"></a> [kms\_key\_policy](#module\_kms\_key\_policy) | ../../ | n/a |

## Resources

| Name | Type |
|------|------|
| [aws_caller_identity.current](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/caller_identity) | data source |
| [aws_iam_policy_document.example](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/iam_policy_document) | data source |
| [aws_iam_roles.administrator_access](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/iam_roles) | data source |
| [aws_region.current](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/region) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_bypass_policy_lockout_safety_check"></a> [bypass\_policy\_lockout\_safety\_check](#input\_bypass\_policy\_lockout\_safety\_check) | (Optional) A boolean flag to indicate whether to bypass the KMS key policy lockout safety check. Defaults to false. | `bool` | `false` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_policy_id"></a> [policy\_id](#output\_policy\_id) | The ID of the KMS Key Policy resource. |
| <a name="output_kms_key_region"></a> [kms\_key\_region](#output\_kms\_key\_region) | Region where the KMS key and policy are managed. |
| <a name="output_key_id"></a> [key\_id](#output\_key\_id) | The ID of the KMS Key associated with the policy. |
| <a name="output_key_arn"></a> [key\_arn](#output\_key\_arn) | The ARN of the KMS Key associated with the policy. |
<!-- END_TF_DOCS -->
