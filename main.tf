data "aws_iam_policy_document" "kms_policy" {
  dynamic "statement" {
    for_each = var.policy
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
        for_each = toset(coalesce(statement.value.condition, []))
        content {
          test     = condition.value.test
          variable = condition.value.variable
          values   = condition.value.values
        }
      }
    }
  }
}

resource "aws_kms_key_policy" "this" {
  key_id                             = var.key_id
  policy                             = data.aws_iam_policy_document.kms_policy.json
  bypass_policy_lockout_safety_check = var.bypass_policy_lockout_safety_check
}
