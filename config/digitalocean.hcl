// Configuration AutoGenerated by CloudQuery CLI
cloudquery {
  plugin_directory = "./cq/providers"
  policy_directory = "./cq/policies"

  provider "digitalocean" {
    version = "latest"
  }

  connection {
    // dsn = "tsdb://postgres:pass@localhost:5432/krishancy_api_development?sslmode=disable"
    dsn = "${DSN}"
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

provider "digitalocean" {
  configuration {
    // Better to pass the keys as environment variables instead of passing them in the configuration
    
    // API Token to access DigitalOcean resources 
    // See https://docs.digitalocean.com/reference/api/api-reference/#section/Authentication
    // token = "<YOUR_API_TOKEN_HERE>"
    // List of regions to fetch spaces from, if not given all regions are assumed
    // spaces_regions = ["nyc3", "sfo3", "ams3", "sgp1", "fra1"]
    // Spaces Access Key generated at https://cloud.digitalocean.com/settings/api/tokens
    // spaces_access_key = "<YOUR_SPACES_ACCESS_KEY>"
    // Spaces Access Key Id generated at https://cloud.digitalocean.com/settings/api/tokens
    // spaces_access_key_id = "<YOUR_SPACES_ACCESS_KEY_ID>"
    // SpacesDebugLogging allows enabling AWS S3 request logging on spaces requests
    // spaces_debug_logging = false
  }

  // list of resources to fetch
  resources = ["*"]
  // enables partial fetching, allowing for any failures to not stop full resource pull
  enable_partial_fetch = true
}
