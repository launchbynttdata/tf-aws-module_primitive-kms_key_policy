variable "key_id" {
  description = "(Required) The ID of the KMS Key to attach the policy."
  type        = string
}

variable "policy" {
  description = "A JSON-formatted string that represents the key policy to attach to the KMS key."
  type = map(object({
    sid        = string
    effect     = string
    principals = map(list(string))
    actions    = list(string)
    resources  = list(string)
  }))
  default = null
}

variable "bypass_policy_lockout_safety_check" {
  description = "(Optional) A boolean flag to indicate whether to bypass the KMS key policy lockout safety check. Defaults to false."
  type        = bool
  default     = false
}
