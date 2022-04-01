# resource "confluentcloud_kafka_acl" "write-standard-cluster" {
#   kafka_cluster = confluentcloud_kafka_cluster.standard-cluster.id
#   resource_type = "TOPIC"
#   resource_name = "payment" #TOPIC NAME
#   pattern_type  = "LITERAL"
#   principal     = "User:sa-w73xmj"
#   host          = "*"
#   operation     = "WRITE"
#   permission    = "ALLOW"
#   http_endpoint = local.http_endpoint
#   credentials {
#     key    = var.kafka_api_key
#     secret = var.kafka_api_secret
#   }
#   depends_on = [confluentcloud_kafka_topic.this]
# }
