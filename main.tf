locals {
  policy_input_is_document      = can(var.policy.json)
  policy_input_string_value     = try(trimspace(var.policy), null)
  policy_input_is_string        = (!local.policy_input_is_document && try(length(local.policy_input_string_value) > 0, false))
  policy_input_statement_values = try(values(var.policy), [])
  policy_input_is_statement_map = (!local.policy_input_is_document && !local.policy_input_is_string && length(local.policy_input_statement_values) > 0 && alltrue([
    for statement in local.policy_input_statement_values :
    can(statement.sid) &&
    can(statement.effect) &&
    can(statement.principals) &&
    can(statement.actions) &&
    can(statement.resources)
  ]))
}

# Always create exactly one instance so count does not depend on unknown var.policy (e.g. from another resource).
data "aws_iam_policy_document" "kms_policy" {
  dynamic "statement" {
    for_each = local.policy_input_is_statement_map ? var.policy : {}
    content {
      sid       = statement.value.sid
      effect    = statement.value.effect
      actions   = statement.value.actions
      resources = statement.value.resources
      dynamic "principals" {
        for_each = statement.value.principals
        content {
          type        = principals.key
          identifiers = principals.value
        }
      }
      dynamic "condition" {
        for_each = toset(lookup(statement.value, "condition", []))
        content {
          test     = condition.value.test
          variable = condition.value.variable
          values   = condition.value.values
        }
      }
    }
  }
}

locals {
  policy_json = local.policy_input_is_document ? var.policy.json : local.policy_input_is_statement_map ? data.aws_iam_policy_document.kms_policy.json : local.policy_input_string_value
}

resource "aws_kms_key_policy" "this" {
  key_id                             = var.key_id
  policy                             = local.policy_json
  bypass_policy_lockout_safety_check = var.bypass_policy_lockout_safety_check
}
