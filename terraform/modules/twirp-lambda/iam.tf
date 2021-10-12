resource "aws_iam_role" "service_role" {
  name               = "${var.project}-role"
  assume_role_policy = data.aws_iam_policy_document.service_role_policy.json

  tags = var.tags
}
