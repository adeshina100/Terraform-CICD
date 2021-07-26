terraform {
  required_providers {
    awsamplify = {
      version = "~> 1.0"
      source  = "hashicorp.com/kislerdm/awsamplify"
    }
  }
}

# AWS Region where to deploy the app
provider "awsamplify" {
  region = "eu-west-1"
}
