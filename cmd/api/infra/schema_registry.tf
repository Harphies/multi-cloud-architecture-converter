# resource "confluentcloud_schema_registry" "test" {
#   environment_id   = confluentcloud_environment.environment.id
#   service_provider = "aws"
#   region           = "EU"

#   depends_on = [confluentcloud_kafka_cluster.test]
# }

# # api key for a kafka cluster
# resource "confluentcloud_api_key" "provider_test" {
#   cluster_id     = confluentcloud_kafka_cluster.test.id
#   environment_id = confluentcloud_environment.environment.id
# }

# # api key for schema registry
# resource "confluentcloud_api_key" "schema-registry" {
#   logical_clusters = [
#     confluentcloud_schema_registry.test.id
#   ]
#   environment_id = confluentcloud_environment.environment.id
# }
