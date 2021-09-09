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

  release_channel {
    channel = "REGULAR"
  }

  vertical_pod_autoscaling {
    enabled = true
  }
}
