# broadcaster-svc

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

An auxiliary service designed to efficiently store events from different blockchains into the rarimo-core, utilizing
queuing mechanism. Frequently used in a couple with *-saver services to submit their transactions. 

## Build

To build the service image locally execute the following command (make sure shat project root contains the vendor package):
```shell
docker build . -t broadcaster-svc:latest
```

It will build the image with the tag `broadcaster:latest` which could be used to run the service locally via
Docker or Docker-Compose.

Also, use the following command to build the binary:
```shell
go build .
```

## Configuration

The following configuration .yaml file should be provided to launch your broadcaster service:
```yaml
log:
  disable_sentry: true
  level: debug

listener:
  addr: :80

## PostgreSQL database connection
db:
  url: "postgres://broadcaster:broadcaster@broadcaster-db/broadcaster?sslmode=disable"

key:
  ## Sender Rarimo private key in 0x.. hex format
  sender_prv_hex: "0x9...caa"
  chain_id: "rarimo_201411-2"
  # Base coin name to pay fee with
  coin_name: "urmo" 

cosmos:
  addr: "validator:9090"

```

You will also need some environment variables to run:

```yaml
- name: KV_VIPER_FILE
  value: /config/config.yaml # The path to your config file
```

## Run
To start the service (in vote mode) use the following command:
```shell
evm-identity-saver-svc run state-update-voter
```

