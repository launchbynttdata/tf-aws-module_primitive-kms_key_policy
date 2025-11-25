output "policy_id" {
  description = "The ID of the KMS Key Policy resource."
  value       = module.kms_key_policy.id
}

output "kms_key_region" {
  description = "Region where the KMS key and policy are managed."
  value       = data.aws_region.current.name
}

output "key_id" {
  description = "The ID of the KMS Key associated with the policy."
  value       = module.kms_key.key_id
}

output "key_arn" {
  description = "The ARN of the KMS Key associated with the policy."
  value       = module.kms_key.arn
}
