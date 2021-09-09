resource "google_container_registry" "registry" {
  project  = data.google_project.home_api.project_id
  location = "US"
}