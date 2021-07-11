
terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.0"
    }
    acme = {
      source = "vancluever/acme"
    }
    kubernetes = {
      source  = "hashicorp/kubernetes"
      version = "~>2.0"
    }
    helm = {
      source  = "hashicorp/helm"
      version = "~> 2"
    }
    kustomization = {
      source  = "kbst/kustomization"
      version = "0.3.0"
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
  region  = var.region
}