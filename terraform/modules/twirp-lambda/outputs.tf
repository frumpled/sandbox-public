output "invoke_arn" {
    value = aws_lambda_function.lambda.invoke_arn
}

output "base_url" {
  description = "Base URL for API Gateway stage."

  value = aws_apigatewayv2_stage.lambda.invoke_url
}