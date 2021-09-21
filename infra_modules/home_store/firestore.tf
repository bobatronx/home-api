resource "google_app_engine_application" "app" {
  project     = data.google_project.home_api.project_id
  location_id = data.google_client_config.current.region
  database_type = "CLOUD_FIRESTORE"
}