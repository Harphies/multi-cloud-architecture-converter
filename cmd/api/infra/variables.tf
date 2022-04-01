variable "confluent_cloud_api_key" {
  type = string
}
variable "confluent_cloud_api_secret" {
  type = string
}
variable "kafka_api_key" {
  type = string
}
variable "kafka_api_secret" {
  type = string
}

variable "topics" {
  type = list(string)
}

variable "topic_partition_count" {
  type = number
}

# Kafka Topic Properties: https://docs.confluent.io/platform/current/installation/configuration/topic-configs.html
variable "topic_config" {
  description = "Kafka topic configuration"
  type        = map(string)
  default = {
    "cleanup.policy"    = "delete"
    "max.message.bytes" = "2097164"
    "retention.ms"      = "604800000"
    "retention.bytes"   = "-1"
  }
}

