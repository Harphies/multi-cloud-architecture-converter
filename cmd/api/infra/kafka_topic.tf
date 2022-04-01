# Topic Properties: https://docs.confluent.io/platform/current/installation/configuration/topic-configs.html
resource "confluentcloud_kafka_topic" "this" {
  kafka_cluster    = confluentcloud_kafka_cluster.standard-cluster.id
  for_each         = toset(var.topics)
  topic_name       = each.value
  partitions_count = var.topic_partition_count
  http_endpoint    = local.http_endpoint
  config           = var.topic_config
  credentials {
    key    = var.kafka_api_key
    secret = var.kafka_api_secret
  }
}



