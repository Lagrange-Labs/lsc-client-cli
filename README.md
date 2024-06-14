# Lagrange CLI

This CLI app provides functionalities such as key management, operator registration, chain subscription, monitoring configuration and automated docker deployment which are necessary to run the Lagrange Attestation Node.

## Lagrange Labs State Committees

For a full breakdown of the Lagrange State Committee architecture, please refer to the below two documents:

1. [Lagrange Technical Overview Docs](https://docs.lagrange.dev/state-committees/overview)
2. [Lagrange State Committee Deep Dive](https://hackmd.io/@lagrange/lagrange-committee)

## Installing CLI from source

To begin with, install Go programming language following the steps mentioned in the [docs](https://docs.lagrange.dev/state-committees/run-node/prerequisite-installation).

- Clone [CLI](https://github.com/Lagrange-Labs/client-cli) repository

```bash
git clone https://github.com/Lagrange-Labs/client-cli.git
```

- Set CGO Flags

```bash
export CGO_CFLAGS="-O -D__BLST_PORTABLE__"
export CGO_CFLAGS_ALLOW="-O -D__BLST_PORTABLE__"
```

- Create binary

```bash
cd client-cli
mkdir -p dist
go build -o ./dist cmd/main.go
```

or

```bash
cd client-cli
sudo apt install make gcc
make build
```

## Running a Lagrange Attestation Node

Please refer to the detailed [documentation](https://docs.lagrange.dev/state-committees/overview).

- Networks:
  - Holesky
  - Mainnet
- [System Requirements](https://docs.lagrange.dev/state-committees/operator-guide/system-requirements)
- [Supported Chains](https://docs.lagrange.dev/state-committees/operator-guide/supported-chains)
- [Prerequisite Installations](https://docs.lagrange.dev/state-committees/run-node/prerequisite-installation)
- [Setup configuration file](https://docs.lagrange.dev/state-committees/run-node/configuration)
- [Commands](#commands)
- [Setup Keystore](https://docs.lagrange.dev/state-committees/run-node/setup-keystore)
- Register Operator
  - [Using CLI](https://docs.lagrange.dev/state-committees/run-node/register-operator)
  - [Using Script](#script)
- Deploy attestation node
  - [Deploy using CLI](https://docs.lagrange.dev/state-committees/run-node/deploy#deploy-using-cli)
  - [Deploy using templates](https://docs.lagrange.dev/state-committees/run-node/deploy#deploy-using-template)
- [Post deployment](#post-deployment)

### Commands

The below commands will allow a developer to run a node and attest to the state of supported chains.

- Generate Keystore: generates a new key pair for the given key type and password, and saves it in the keystore file. The key type can be either `ecdsa` or `bls`.

  ```bash
  lagrange-cli generate-keystore -t <Key Type> -p <Password File Path>

  # i.e. ./dist/lagrange-cli generate-keystore -t ecdsa -p ~/.lagrange/keystore/ecdsa_.pass
  ```

- Import Keystore: imports a key pair from the given private key and saves it in the keystore file. The key type can be either `ecdsa` or `bls`.

  ```bash
  lagrange-cli import-keystore -t <Key Type> -p <Password File Path> -k <Private Key>

  # i.e. ./dist/lagrange-cli import-keystore -t bls -p ~/.lagrange/keystore/bls_.pass -k 0x1234567890abcdef...
  ```

- Export Keystore: exports a private key from the keystore file for the given key type and password.

  ```bash
  lagrange-cli export-keystore -t <Key Type> -p <Password File Path> -f <Keystore File Path>

  # i.e. ./dist/lagrange-cli export-keystore -t ecdsa -p ~/.lagrange/keystore/ecdsa_.pass -f ~/.lagrange/keystore/bls_.key
  ```

- Export Public Key: exports a public key from the keystore file for the given key type.

  ```bash
  lagrange-cli export-public-key -t <Key Type> -f <Keystore File Path>

  # i.e. ./dist/lagrange-cli export-public-key -t ecdsa -f ~/.lagrange/keystore/ecdsa_.key
  ```

- Register Operator: registers an operator to the Lagrange State Committee. The network name can be either `mainnet` or `holesky`. The BLS key and Signer address are referenced from the config file.

  ```bash
  lagrange-cli register-operator -c <Config File Path> -n <Network Name>

  # i.e. ./dist/lagrange-cli register-operator -c ./config.toml -n mainnet
  ```

- Deregister Operator: deregisters an operator from the Lagrange State Committee. The network name can be either `mainnet` or `holesky`.

  ```bash
  lagrange-cli deregister-operator -c <Config File Path> -n <Network Name>

  # i.e. ./dist/lagrange-cli deregister-operator -c ./config.toml -n mainnet
  ```

- Update BLS Public Key: updates the BLS public key for the operator at the given index. The network name can be either `mainnet` or `holesky`. New BLS key is referenced from the config file.

  ```bash
  lagrange-cli update-bls-pub-key -c <Config File Path> -n <Network Name> -i <Key Index>

  # i.e. ./dist/lagrange-cli update-bls-pub-key -c ./config.toml -n mainnet -i 0
  ```

- Update Signer Address: updates the signer address for the operator. The network name can be either `mainnet` or `holesky`. New signer address is referenced from the config file.

  ```bash
  lagrange-cli update-signer-address -c <Config File Path> -n <Network Name>

  # i.e. ./dist/lagrange-cli update-signer-address -c ./config.toml -n mainnet
  ```

- Add BLS Public Key: adds a new BLS public key for the operator. The network name can be either `mainnet` or `holesky`. New BLS key is referenced from the config file.

  ```bash
  lagrange-cli add-bls-pub-key -c <Config File Path> -n <Network Name>

  # i.e. ./dist/lagrange-cli add-bls-pub-key -c ./config.toml -n mainnet
  ```

- Remove BLS Public Key: removes the BLS public key for the operator at the given index. The network name can be either `mainnet` or `holesky`.

  ```bash
  lagrange-cli remove-bls-pub-key -c <Config File Path> -n <Network Name> -i <Key Index>

  # i.e. ./dist/lagrange-cli remove-bls-pub-key -c ./config.toml -n mainnet -i 0
  ```

- Subscribe Chain: subscribes the operator to the given chain. The network name can be either `mainnet` or `holesky`. The chain name can be `optimism`, `arbitrum`, `base` etc.

  ```bash
  lagrange-cli subscribe-chain -c <Config File Path> -n <Network Name> -r <Chain Name>

  # i.e. ./dist/lagrange-cli subscribe-chain -c ./config.toml -n mainnet -r optimism
  ```

- Unsubscribe Chain: unsubscribes the operator from the given chain. The network name can be either `mainnet` or `holesky`. The chain name can be `optimism`, `arbitrum`, `base` etc.

  ```bash
  lagrange-cli unsubscribe-chain -c <Config File Path> -n <Network Name> -r <Chain Name>

  # i.e. ./dist/lagrange-cli unsubscribe-chain -c ./config.toml -n mainnet -r optimism
  ```

- Generate Config: generates a client config file. The network name can be either `mainnet` or `holesky`. The chain name can be `optimism`, `arbitrum`, `base` etc.

  ```bash
  lagrange-cli generate-config -c <Config File Path> -n <Network Name> -r <Chain Name>

  # i.e. ./dist/lagrange-cli generate-config -c ./config.toml -n mainnet -r optimism
  ```

- Deploy Node: creates a docker-compose file and deploys the docker container for the attestation node. The prometheus port is the port which binds the prometheus server to the host machine.

  ```bash
  lagrange-cli deploy -c <Client Config File Path> -i <Docker Image Name> -p <Prometheus Port>

  # i.e. ./dist/lagrange-cli deploy -c ~/.lagrange/config/client_mainnet_optimism_.toml -i lagrangelabs/lagrange-node:v0.4.0 -p 8080
  ```

- Generate Config & Deploy Node: generates a client config file and deploys the docker container for the attestation node. It combines the `generate-config` and `deploy` commands.

  ```bash
  lagrange-cli generate-config-deploy -c <Config File Path> -n <Network Name> -r <Chain Name> -i <Docker Image Name> -p <Prometheus Port>

  # i.e. ./dist/lagrange-cli generate-config-deploy -c ./config.toml -n mainnet -r optimism -i lagrangelabs/lagrange-node:v0.4.0 -p 8080
  ```

> NOTE: All file paths should be absolute or relative to the current working directory including config file paths and keystore file paths.

> NOTE: Please take a backup of your BLS and ECDSA key pairs. They are located at `~/.lagrange/keystore/<key_type>_<public_key_prefix>.key`.

> NOTE: We recommend using trusted providers like Alchemy, Quicknode, Infura if you don't run your own node. For the Beacon RPC endpoint, you should check if the given RPC provider supports the Beacon RPC API like `/eth/v1/beacon/genesis`. Quicknode supports this API.

> NOTE: You can check client config files in the `~/.lagrange/config/client_<network_name>_<chain_name>_<bls_pub_key_prefix>.toml` and docker-compose files in the `~/.lagrange/docker-compose_<network_name>_<chain_name>_<bls_pub_key_prefix>.yml`.

### Script

If you want to register an operator without running the commands manually, you can use scripts under the `script` folder. Please update the `# ==== MODIFY ME! ====` part in the script with your details. You can use `export-public-key` command to get the public key from the keystore file.

- To register an operator, run the below command:

  ```bash
  source ./script/register-operator.sh
  ```

- To subscribe to a chain, run the below command:

  ```bash
    source ./script/subscribe-chain.sh
  ```

### Post-Deployment

- To check the status of the attestation node from the docker container logs, run the below command:

  ```bash
  # to get the container id
  docker ps

  # if you are not seeing any running containers, then run the below command to see all suspended containers
  docker ps -a

  # to check the logs of the container
  docker logs <container_id>

  # to stop the container
  cd $HOME/.lagrange && docker compose -f <docker-compose-file> down --remove-orphans
  ```

- If you wish to setup Grafana dashboard for monitoring your attestation node, please review the steps mentioned in this [documentation](/monitoring/MONITORING.MD)

- If you experience the rpc provider issue, it can be due to the rate limit of the provider. You can control the `ConcurrentFetchers` in the `config.toml` file to manage the rate limit.

- If you face any issues while running the Lagrange Attestation Node, please reach out to the Lagrange Labs team on Telegram.
