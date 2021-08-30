# GKE cluster
resource "google_container_cluster" "primary" {
  enable_autopilot = true
  name             = "${data.google_project.home_api.project_id}-gke"
  location         = data.google_client_config.current.region

  master_authorized_networks_config {
    cidr_blocks {
      cidr_block   = "75.135.104.135/32"
      display_name = "kubectl-client"
    }
  }

  private_cluster_config {
    enable_private_nodes    = true
    enable_private_endpoint = false
    master_ipv4_cidr_block  = "172.168.0.0/28"
  }
}