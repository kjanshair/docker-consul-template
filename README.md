## consul-template with Golang configuration

This is a simple repository for understanding dynamic file configuration with Consul's KV (Key Value) store. This repository uses a super simple Golang project, located inside the `src` folder, which reads JSON configuration from `config.json` file located at `src/config/config.json`. But you can use this in your project the way you want.

There are 2 Consul specific files `consul-template.hcl` at the root and `config.tpl` template file located at `src/config/config.tpl`. `consul-template.hcl` which `consul-template` uses to get `config.tpl` as input and produce `config.json` file as the output.

## Running the solution

If you have Docker and `docker-compose` installed on your system, simply run `docker-compose up -d` at the root and start editing the Consul KV store to reflect the changes in the `config.json` file.