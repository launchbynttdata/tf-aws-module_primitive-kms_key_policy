variable "key_id" {
  description = "(Required) The ID of the KMS Key to attach the policy."
  type        = string
}

variable "policy" {
  description = <<EOT
Accepts one of the following:
1. A map of statements (the legacy behaviour of this module) where each entry contains sid, effect, principals, actions, resources, and optional condition blocks.
2. An aws_iam_policy_document data source (the object itself) so the module can reuse its rendered JSON.
3. A raw JSON string that should be attached to the KMS key as-is.
EOT
  type        = any

  validation {
    condition = (
      var.policy != null &&
      anytrue([
        can(var.policy.json),
        try(length(trimspace(var.policy)) > 0, false),
        (
          length(try(keys(var.policy), [])) > 0 &&
          alltrue([
            for statement in try(values(var.policy), []) :
            can(statement.sid) &&
            can(statement.effect) &&
            can(statement.principals) &&
            can(statement.actions) &&
            can(statement.resources)
          ])
        )
      ])
    )

    error_message = "policy must be an aws_iam_policy_document object, a non-empty JSON string, or a map of statements with sid, effect, principals, actions, and resources."
  }
}

variable "bypass_policy_lockout_safety_check" {
  description = "(Optional) A boolean flag to indicate whether to bypass the KMS key policy lockout safety check. Defaults to false."
  type        = bool
  default     = false
}
