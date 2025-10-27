variable "bypass_policy_lockout_safety_check" {
  description = "(Optional) A boolean flag to indicate whether to bypass the KMS key policy lockout safety check. Defaults to false."
  type        = bool
  default     = false
}
