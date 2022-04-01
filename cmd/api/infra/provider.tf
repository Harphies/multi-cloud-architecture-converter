terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "4.5.0"
    }
    confluentcloud = {
      source  = "confluentinc/confluentcloud"
      version = "0.4.0"
    }

  }
}

provider "aws" {
  region = "eu-west-2"
}

provider "confluentcloud" {
  api_key    = var.confluent_cloud_api_key
  api_secret = var.confluent_cloud_api_secret
}
