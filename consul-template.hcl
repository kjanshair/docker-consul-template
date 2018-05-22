consul = "consul-dev:8500"
# retry = "10s"
# max_stale = "10m"

template {
  source = "/config/config.tpl"
  destination = "/config/config.json"
}
