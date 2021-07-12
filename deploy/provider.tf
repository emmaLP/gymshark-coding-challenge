
terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.0"
    }
    random = {
      source  = "hashicorp/random"
      version = "~> 3.0"
    }
    archive = {
      source  = "hashicorp/archive"
      version = "~> 2"
    }
  }
  required_version = " 1.0.2"
  backend "s3" {
    bucket         = "emma-gymshark-challenge-terraform-state"
    key            = "gymshark-code-challenge.tfstate"
    dynamodb_table = "terraform-state-lock"
    region         = "eu-west-2"
  }
}

provider "aws" {
  region = var.region
}