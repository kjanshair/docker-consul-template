consul {
  address = "consul-dev:8500"
  retry {
    enabled = true
    attempts = 0
    backoff = "250ms"
    max_backoff = "1m"
  }
}

template {
  source = "/config/config.tpl"
  destination = "/config/config.json"
}
