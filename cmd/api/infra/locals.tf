locals {
  bootstrap_servers = [confluentcloud_kafka_cluster.standard-cluster.bootstrap_endpoint]
  http_endpoint     = confluentcloud_kafka_cluster.standard-cluster.http_endpoint
}
