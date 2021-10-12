output "invoke_arn" {
    value = module.sandbox.invoke_arn
}

output "base_url" {
    description = "Base URL for API Gateway stage."
    value       = module.sandbox.base_url
}
