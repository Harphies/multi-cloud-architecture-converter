# Standard cluster creation
resource "confluentcloud_kafka_cluster" "standard-cluster" {
  display_name = "experiment_kafka_cluster"
  availability = "SINGLE_ZONE"
  cloud        = "AWS"
  region       = "eu-west-2"
  standard {}

  environment {
    id = confluentcloud_environment.this.id
  }
}
