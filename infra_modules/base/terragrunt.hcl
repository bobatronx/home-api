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
      project_id     = "teak-bebop-123023",
      project_number = "969632929420"
    }
  }
}