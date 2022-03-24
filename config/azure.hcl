// Configuration AutoGenerated by CloudQuery CLI
cloudquery {
  plugin_directory = "./cq/providers"
  policy_directory = "./cq/policies"

  provider "azure" {
    version = "latest"
  }

  connection {
    dsn = "tsdb://postgres:pass@localhost:5432/krishancy_api_development?sslmode=disable"
  }

  history {
    // Save data retention for 7 days
    retention = 7
    // Truncate our fetch by 6 hours per fetch
    truncation = 6
    // Tell Timescale to split our chunks in a daily interval (24 hours)
    interval = 1
  }
}

// Provider Configurations

provider "azure" {
  configuration {
    //  Optional. if you not specified, cloudquery tries to access all subscriptions available to tenant
    //  subscriptions = ["<YOU_SUBSCRIPTION_ID_HERE>"]
  }
  // list of resources to fetch
  resources = ["*"]
  // enables partial fetching, allowing for any failures to not stop full resource pull
  enable_partial_fetch = true
}
