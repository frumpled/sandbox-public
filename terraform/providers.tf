provider "aws" {
	region  = "us-east-1"
}

#resource "aws_s3_bucket" "state_bucket" {
#	bucket = "tf-sandbox-public"
#	acl    = "private"
#
#	lifecycle {
#		prevent_destroy = true
#	}
#
#	tags = {
#		name    = "tf-${local.project}"
#		project = var.project
#	}
#}
#
#resource "aws_s3_bucket_public_access_block" "tf_state_bucket" {
#	bucket = aws_s3_bucket.state_bucket.id
#
#	block_public_acls       = true
#	block_public_policy     = true
#	ignore_public_acls      = true
#	restrict_public_buckets = true
#}
#
#terraform {
#	backend "s3" {
#		encrypt = true
#		bucket  = "tf-sandbox-public"
#		key     = "terraform/sandbox-public/terraform.tfstate"
#	}
#}

terraform {
	required_version = ">= 1.0.1"
}

