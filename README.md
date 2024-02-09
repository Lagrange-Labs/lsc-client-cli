# Client CLI

## Config

Set the config file `config.toml` in the root of the project with the proper values.
For now, we only support the `BN254` curve for the `BLSScheme`.

## Build & Run

```bash
    make build

    ./dist/lagrange-cli run -c ./config.toml
```
