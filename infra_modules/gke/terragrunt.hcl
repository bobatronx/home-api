include {
  path = find_in_parent_folders()
}

terraform {
  extra_arguments "retry_lock" {
    commands = [
      "init",
      "apply",
      "refresh",
      "import",
      "plan",
      "taint",
      "untaint"
    ]

    env_vars = {
      TF_VAR_project_id = "teak-bebop-123023",
      TF_VAR_project_number = "969632929420"
    }
  }
}