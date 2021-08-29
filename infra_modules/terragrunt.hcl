remote_state {
  backend = "gcs"    
  config = {
    bucket         = "home-api-infra-${get_env("project_number")}"
    prefix         = "${path_relative_to_include()}/terraform.tfstate"
  }
}