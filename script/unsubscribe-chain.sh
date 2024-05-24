#!/bin/zsh

# --- Dependencies ---
# 1. zsh
# 2. `cast` - a binary from the foundry toolkit https://github.com/foundry-rs/foundry?tab=readme-ov-file
# 3. `openssl` - to generate your new ECDSA secp256k1 LAGR_KEY to be used only for the ZK Coprocessor AVS
# 4. `jq` - only needed for the demo; 
# see === MODIFY ME! === section below

# ==== MODIFY ME! ====
ETH_KEY=
OPERATOR_ADDR=
CHAIN_TO_UNSUBSCRIBE=420
export ETH_RPC_URL=https://rpc.holesky.ethpandaops.io

# --- Constants ---
HOLESKY_CHAINID=17000
SEPOLIA_CHAINID=11155111
LOCAL_CHAINID=1337
MAINNET_CHAINID=1

# Determine contract addresses based on the chain ID
set_contract_addresses() {
    local chain_id=$(cast chain-id)

    if [ -z "$chain_id" ]; then
        echo "Error: Unable to fetch chain ID."
        exit 1
    fi
    printf "Chain id: ${chain_id}\n"
    echo $chain_id

    if [ "$chain_id" -eq "$HOLESKY_CHAINID" ]; then
        echo "Registering Operator on Holesky Testnet"
        LAGRANGE_SERVICE_ADDR=0x18A74E66cc90F0B1744Da27E72Df338cEa0A542b
        LAGRANGE_COMMITTEE_ADDR=0x57BAf26C77BBBa3D3A8Bd620c8d74B44Bfda8818
        EIGEN_AVS_DIRECTORY_ADDR=0x055733000064333CaDDbC92763c58BF0192fFeBf
    elif [ "$chain_id" -eq "$MAINNET_CHAINID" ]; then
        LAGRANGE_SERVICE_ADDR=0x35F4f28A8d3Ff20EEd10e087e8F96Ea2641E6AA2
        LAGRANGE_COMMITTEE_ADDR=0xECc22f3EcD0EFC8aD77A78ad9469eFbc44E746F5
        EIGEN_AVS_DIRECTORY_ADDR=0x135DDa560e946695d6f155dACaFC6f1F25C1F5AF
    elif [ "$chain_id" -eq "$SEPOLIA_CHAINID" ]; then
        echo "Registering Operator on Sepolia Testnet"
        LAGRANGE_SERVICE_ADDR=0x34d8f7384Ddd4e8Ab447a150F9955f284882A43F
        LAGRANGE_COMMITTEE_ADDR=0x298baE5B68bA10e0dF7d98C5b19Aba7784d74F71
        EIGEN_AVS_DIRECTORY_ADDR=0xE5F13E61654363719407247BB7E06f539d3a9d32
    elif [ "$chain_id" -eq "$LOCAL_CHAINID" ]; then
        echo "Registering Operator on Sepolia Testnet"
        LAGRANGE_SERVICE_ADDR=0xBda41273d671bb33374F7b9C4Ae8746c712727f7
        LAGRANGE_COMMITTEE_ADDR=0xBF4E09354df24900e3d2A1e9057a9F7601fbDD06
        EIGEN_AVS_DIRECTORY_ADDR=0x6Bf0fF4eBa00E3668c0241bb1C622CDBFE55bbE0
    else
        echo "Unknown chain ID: $chain_id"
        exit 1
    fi
}

unsubscribe_chain() {
    cast send $LAGRANGE_SERVICE_ADDR "unsubscribe(uint32)" $CHAIN_TO_UNSUBSCRIBE --private-key $ETH_KEY

    printf "\nSuccessfully Unsubscribed!\n"
}

## --- Main Script ---
printf "Step 1: set_contract_addresses\n"
set_contract_addresses

sleep 3s
printf "Step 2: unsubscribe_chain\n"
unsubscribe_chain
