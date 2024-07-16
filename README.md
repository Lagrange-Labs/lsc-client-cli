# LSC Client CLI

This CLI app provides functionalities such as key management, operator registration, chain subscription, monitoring configuration and automated docker deployment which are necessary to run the LSC Attestation Node.

## Lagrange State Committees

For a full breakdown of the Lagrange State Committee architecture, please refer to the below two documents:

1. [Lagrange Technical Overview Docs](https://docs.lagrange.dev/state-committees/overview)
2. [Lagrange State Committee Deep Dive](https://hackmd.io/@lagrange/lagrange-committee)

## Installing CLI from source

To begin with, install Go programming language following the steps mentioned in the [docs](https://docs.lagrange.dev/state-committees/run-node/prerequisite-installation).

- Clone [CLI](https://github.com/Lagrange-Labs/lsc-client-cli) repository

```bash
git clone https://github.com/Lagrange-Labs/lsc-client-cli.git
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

## Running a LSC Attestation Node

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
  - [Using Script](https://docs.lagrange.dev/state-committees/run-node/scripts)
- Deploy attestation node
  - [Deploy using CLI](https://docs.lagrange.dev/state-committees/run-node/deploy#deploy-using-cli)
  - [Deploy using templates](https://docs.lagrange.dev/state-committees/run-node/deploy#deploy-using-template)
- [Post deployment](#post-deployment)

### Commands

The below commands will allow a developer to run a node and attest to the state of supported chains.

- [generate-keystore](https://docs.lagrange.dev/state-committees/run-node/commands#generate-keystore)
- [import-keystore](https://docs.lagrange.dev/state-committees/run-node/commands#import-keystore)
- [export-keystore](https://docs.lagrange.dev/state-committees/run-node/commands#export-keystore)
- [export-public-key](https://docs.lagrange.dev/state-committees/run-node/commands#export-public-key)
- [register-operator](https://docs.lagrange.dev/state-committees/run-node/commands#register-operator)
- [deregister-operator](https://docs.lagrange.dev/state-committees/run-node/commands#deregister-operator)
- [update-bls-pub-key](https://docs.lagrange.dev/state-committees/run-node/commands#update-bls-pub-key)
- [update-signer-address](https://docs.lagrange.dev/state-committees/run-node/commands#update-signer-address)
- [add-bls-pub-key](https://docs.lagrange.dev/state-committees/run-node/commands#add-bls-pub-key)
- [remove-bls-pub-key](https://docs.lagrange.dev/state-committees/run-node/commands#remove-bls-pub-key)
- [subscribe-chain](https://docs.lagrange.dev/state-committees/run-node/commands#subscribe-chain)
- [unsubscribe-chain](https://docs.lagrange.dev/state-committees/run-node/commands#unsubscribe-chain)
- [generate-config](https://docs.lagrange.dev/state-committees/run-node/commands#generate-config)
- [generate-docker-compose](https://docs.lagrange.dev/state-committees/run-node/commands#generate-docker-compose)
- [deploy](https://docs.lagrange.dev/state-committees/run-node/commands#deploy)
- [generate-config-deploy](https://docs.lagrange.dev/state-committees/run-node/commands#generate-config-deploy)

### Post Deployment

- If you wish to setup Grafana dashboard for monitoring your attestation node, please review the steps mentioned in this [documentation](/monitoring/MONITORING.MD)

- The detailed information for the post-deployment monitoring can be found on our [documentation](https://docs.lagrange.dev/state-committees/run-node/monitoring) page.

**If you face any issues while running the LSC Attestation Node, please reach out to the Lagrange Labs team on [Discord](https://discord.lagrange.dev/).**
