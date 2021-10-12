# Notes on configs: https://docs.aws.amazon.com/lambda/latest/dg/gettingstarted-limits.html
resource "aws_lambda_function" "lambda" {
  description = var.project

  filename         = var.lambda_filename
  function_name    = "${var.environment}-${var.project}-${var.function_name}"
  handler          = "main"
  memory_size      = 128 # Default & min = 128mb
  role             = aws_iam_role.service_role.arn
  runtime          = "go1.x"
  source_code_hash = filebase64sha256(var.lambda_filename)
  timeout          = 1 # seconds

  environment {
    variables = {
      TEST = "PASSED"
    }
  }

  tags = var.tags
}

resource "aws_cloudwatch_log_group" "lambda" {
  name = "/aws/lambda/${aws_lambda_function.lambda.function_name}"

  retention_in_days = 30
}

resource "aws_iam_role_policy_attachment" "lambda_policy" {
  role       = aws_iam_role.service_role.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
}
