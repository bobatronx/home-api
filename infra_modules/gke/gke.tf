# GKE cluster
resource "google_container_cluster" "primary" {
  name     = "${data.google_project.home_api.project_id}-gke"
  location = data.google_client_config.current.region
  
  enable_autopilot = true
}