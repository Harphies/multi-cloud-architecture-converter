output "bootrap_server" {
  value = [replace(confluentcloud_kafka_cluster.standard-cluster.bootstrap_endpoint, "SASL_SSL://", "")]
}

# output "sasl_username" {
#   value     = confluentcloud_api_key.this.key
#   sensitive = true
# }

# output "sasl_password" {
#   value     = confluentcloud_api_key.this.secret
#   sensitive = true
# }
