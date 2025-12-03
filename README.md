# tf-aws-module_primitive-kms_key_policy_policy

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

No modules.

## Resources

| Name | Type |
|------|------|
| [aws_kms_key_policy.this](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/kms_key_policy) | resource |
| [aws_iam_policy_document.kms_policy](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/iam_policy_document) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_key_id"></a> [key\_id](#input\_key\_id) | (Required) The ID of the KMS Key to attach the policy. | `string` | n/a | yes |
| <a name="input_policy"></a> [policy](#input\_policy) | Accepts one of the following:<br/>1. A map of statements (the legacy behaviour of this module) where each entry contains sid, effect, principals, actions, resources, and optional condition blocks.<br/>2. An aws\_iam\_policy\_document data source (the object itself) so the module can reuse its rendered JSON.<br/>3. A raw JSON string that should be attached to the KMS key as-is. | `any` | n/a | yes |
| <a name="input_bypass_policy_lockout_safety_check"></a> [bypass\_policy\_lockout\_safety\_check](#input\_bypass\_policy\_lockout\_safety\_check) | (Optional) A boolean flag to indicate whether to bypass the KMS key policy lockout safety check. Defaults to false. | `bool` | `false` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_id"></a> [id](#output\_id) | The ID of the KMS Key Policy resource. |
<!-- END_TF_DOCS -->
