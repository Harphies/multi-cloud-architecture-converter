terraform {
  backend "s3" {
    region         = "eu-west-2"
    bucket         = "all-my-project-terraform-states"
    key            = "kafka-confluent-cloud.tfstate"
    encrypt        = true
    dynamodb_table = "TerraformStateLocks"
  }
}
