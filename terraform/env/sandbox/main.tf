module "sandbox" {
    source = "../../modules/twirp-lambda"

    environment     = "sandbox"
    tags            = local.tags
    project         = local.project
    function_name   = "twirp-lambda"
    lambda_filename = "../../../twirp/lambda/main.zip"
}
