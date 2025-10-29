output "policy_id" {
  description = "The ID of the KMS Key Policy resource."
  value       = module.kms_key_policy.id
}
