# CLI APP

This CLI app provides the functionalities needed to run the a Lagrange Attestation Node.

## Lagrange Labs State Committees Attestation Node

Lagrange Labs State Committees provide a mechanism for generating succinct zero-knowledge state proofs for optimistic rollups based on the use of either staked or restaked collateral. Each state committee is a group of attestors/validators that have either staked an optimistic rollup’s native token or dualstaked with EigenLayer. Each state committee node attests to the execution and finality of transaction batches submitted by optimistic sequencers to Ethereum.

Whenever a batch consisting of rollup blocks is considered either safe (OP stack) or has had its corresponding transaction batch settled on Ethereum (Mainnet), each node is required to attest to the batch of blocks using its BLS key.

Broadly, each signature is executed on a tuple containing 3 essential elements:

```
struct batch {
    var batch_header,
    var current_committee,
    var next_committee
}
```

For a full breakdown of the architecture, please refer to the below two documents:

1. [Lagrange Technical Overview Docs](https://lagrange-labs.gitbook.io/lagrange-labs/lagrange-state-committees/commitees-overview)
2. [Lagrange State Committee Deep Dive](https://hackmd.io/@lagrange/lagrange-committee)

## Running a Lagrange Client Node

For the purpose of demonstrating how to run a Lagrange Attestation Node, we've created a Lagrange Holesky Network. The below commands will allow a developer to run a node and attest to the state of `Optimism` and `Base` chains.

### Chains

- Optimism: `10`
- Base: `8453`

### Minimum Recommended Hardware

- VCPUs: `2`
- RAM: `4 GB`
- Storage: `8 GB`
- AWS instance type: `t3.medium`
- Architecture: 64-bit ARM instance

> Note: The commands in this README file are for 64-bit ARM AWS ubuntu instance.

### Pre-requisite Installations

1. Golang

```bash
sudo apt-get update
sudo apt-get -y upgrade
wget https://go.dev/dl/go1.21.9.linux-amd64.tar.gz
sudo tar -xvf go1.21.9.linux-amd64.tar.gz -C /usr/local
echo "export PATH=$PATH:/usr/local/go/bin" >> ~/.profile
export GOROOT=/usr/local/go
source ~/.profile
go version
```

2. Docker and Docker Compose

```bash
sudo apt-get update
sudo install -m 0755 -d /etc/apt/keyrings
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /etc/apt/keyrings/docker.gpg
sudo chmod a+r /etc/apt/keyrings/docker.gpg
echo "deb [arch=\"$(dpkg --print-architecture)\" signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/ubuntu $(. /etc/os-release && echo \"$VERSION_CODENAME\") stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
sudo apt-get update
sudo apt-get install -y docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin make gcc
echo '{ "log-opts": { "max-size": "10m", "max-file": "100" } }' | sudo tee /etc/docker/daemon.json
sudo usermod -aG docker $USER
newgrp docker
sudo systemctl restart docker
```

## Steps

1. Send your EigenLayer operator Ethereum address to Lagrange Labs team for allowlisting. ([Form](https://docs.google.com/forms/d/1oJq1QddKb1Sa_Pe_C1Sa-9p_jN4jBEV2ARI3-M8yb8c/edit))

2. Clone the [Lagrange CLI repository](https://github.com/Lagrange-Labs/client-cli) to your machine.

```bash
git clone https://github.com/Lagrange-Labs/client-cli.git
```

3. Set environment variables and download dependencies.

```bash
export CGO_CFLAGS="-O -D__BLST_PORTABLE__"
export CGO_CFLAGS_ALLOW="-O -D__BLST_PORTABLE__"
cd client-cli
go mod download
```

4. Create a go binary.

```bash
sudo apt install make gcc

make build
```

5. Update `EthereumRPCURL` in `config.toml` file.

`EthereumRPCURL`: RPC endpoint for `Holesky` network.

Currently, we only support the `BN254` curve for the `BLSScheme`.

6. Run the following command to register your operator to Lagrange State Committees AVS.

```bash
./dist/lagrange-cli run -c ./config.toml
```

- Enter `g` in the prompt to start the registration.

- Each Lagrange Attestation Node requires a BLS key pair to participate in the committee. Generate the BLS key pairs or insert your own keys in the prompt.

  - Prompt: "Do you want to add a new BLS Key Pair?" - Enter `y` if you want to add a BLS key pair
  - Prompt: "Do you want to generate a new Key Pair?" - Enter `y` if you want the CLI to randomly generate a key pair or enter your own BLS private key.
  - Please re-run this step to add more BLS key pairs if you want to run more nodes.

- Each EigenLayer operator that participates into Lagrange State Committees AVS needs to generate one signer ECDSA private key.

  > Note: This signer ECDSA key is only used to verify the identity and signature. It is not used in governance and/or asset flow.

  - Enter `y` in the prompt for CLI to randomly generate a signer private key if you haven't already created it before.

- Enter `y` to register to Lagrange State Committee if you haven't already registered.

  - Enter ECDSA key of your EigenLayer operator to complete registration.

7. Upon completion of the registration process, subscribe to the chain that you want to run the Lagrange Attestation Node.

- Enter `s` in the prompt to start the subscription process.

- Enter EigenLayer operator ECDSA key.

- Enter CHAIN_ID to subscribe to that chain, [please check the above chain ids](#chains).

> Note: Each Operator can be subscribed to multiple chains.

8. Generate config for your Lagrange Attestation Node which will be used to deploy the container.

- Enter `c` in the prompt to start generating the config.

- Enter the operator address.

- Enter the chain name which you want to run the Lagrange Attestation Node, [please check the above chains](#chains).

- Enter L1 RPC endpoint, which will be ETH mainnet endpoint for this testnet.

- Enter L2 RPC endpoint, which will be the rollup chain's endpoint.

- Enter the Beacon RPC endpoint, which is used to fetch blobs data.

- Enter gRPC URL.

  - Optimism: `44.210.11.64:9090`
  - Base: `3.209.124.237:9090`

- Select and enter index of the key pair from the list of available BLS key pairs (registered in step #6)

> Note: We recommend using trusted providers like Alchemy, Quicknode, Infura if you don't run your own node. For the Beacon RPC endpoint, you should check if the given RPC provider supports the Beacon RPC API like `/eth/v1/beacon/genesis`. Quicknode supports this API.

9. Deploy and run the Lagrange Attestation Node.

- Enter `r` in the prompt.

- Select index of the config file from the list of available configs to deploy the docker container.

10. Repeat `config` and `run` commands to deploy more attestation nodes. If you subscribe to `k` chains and you have `n` BLS key-pairs then there will be `n*k` containers running.

### Post-Deployment

- To check the status of the attestation node from the docker container logs, run the below command:

```bash
# to get the container id
sudo docker ps

# if you are not seeing any running containers, then run the below command to see all suspended containers
sudo docker ps -a

# to check the logs of the container
sudo docker logs <container_id>
```

> Note: If you are seeing this error message `time= level=error msg=failed to join the network: failed to join the network: pc error: code = Unknown desc = the operator is not a committee member`, it means that the given attestation node setup is successful and it will start attesting from the next epoch. The current epoch period of the state committee rotation is 24 hours for Holesky testnet.

- If you face any issues while running the Lagrange Attestation Node, please reach out to the Lagrange Labs team on Telegram.
