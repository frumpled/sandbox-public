variable "environment" {
    type        = string
    description = "The infrastructure environemnt [staging/production]"
}

variable "tags" {
    type        = map(string)
    description = "A set of tags to help identify this project's infrastructure"
}

variable "project" {
    type        = string
    description = "A name to identify this project's resources"
}

variable "lambda_filename" {
    type        = string
    description = "Filename for the lambda source code zipfile"
}

variable "function_name" {
    type        = string
    description = "Fucntion name for the lambda"
}
