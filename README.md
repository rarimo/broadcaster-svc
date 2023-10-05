# broadcaster-svc

An auxiliary service designed to efficiently store events from different blockchains into the rarimo-core, utilizing queuing mechanism.

## Build

To build the service image locally, there is a shell script `build.sh` that can be used to build the image:

```bash
sh build.sh
```

It will build the image with the tag `broadcaster:latest` which could be used to run the service locally via
Docker or Docker-Compose.

## License
[MIT](./LICENSE)
