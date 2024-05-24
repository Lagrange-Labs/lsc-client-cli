# --- Dependencies ---
# 1. sh
# 2. `cast` - a binary from the foundry toolkit https://github.com/foundry-rs/foundry?tab=readme-ov-file
#           Version: tag nightly-f625d0fa7c51e65b4bf1e8f7931cd1c6e2e285e9
#                   cargo install --git https://github.com/foundry-rs/foundry --tag nightly-f625d0fa7c51e65b4bf1e8f7931cd1c6e2e285e9 --profile local --locked cast --force

# ==== MODIFY ME! ====
OPERATOR_PRIVATE_KEY=
OPERATOR_ADDR=
CHAIN_TO_UNSUBSCRIBE= # 10
export ETH_RPC_URL= # Put rpc url here

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
        EIGEN_AVS_DIRECTORY_ADDR=0x055733000064333CaDDbC92763c58BF0192fFeBf
    elif [ "$chain_id" -eq "$MAINNET_CHAINID" ]; then
        LAGRANGE_SERVICE_ADDR=0x35F4f28A8d3Ff20EEd10e087e8F96Ea2641E6AA2
        EIGEN_AVS_DIRECTORY_ADDR=0x135DDa560e946695d6f155dACaFC6f1F25C1F5AF
    elif [ "$chain_id" -eq "$SEPOLIA_CHAINID" ]; then
        echo "Registering Operator on Sepolia Testnet"
        LAGRANGE_SERVICE_ADDR=0x34d8f7384Ddd4e8Ab447a150F9955f284882A43F
        EIGEN_AVS_DIRECTORY_ADDR=0xE5F13E61654363719407247BB7E06f539d3a9d32
    elif [ "$chain_id" -eq "$LOCAL_CHAINID" ]; then
        echo "Registering Operator on Sepolia Testnet"
        LAGRANGE_SERVICE_ADDR=0xBda41273d671bb33374F7b9C4Ae8746c712727f7
        EIGEN_AVS_DIRECTORY_ADDR=0x6Bf0fF4eBa00E3668c0241bb1C622CDBFE55bbE0
    else
        echo "Unknown chain ID: $chain_id"
        exit 1
    fi
}

unsubscribe_chain() {
    cast send $LAGRANGE_SERVICE_ADDR "unsubscribe(uint32)" $CHAIN_TO_UNSUBSCRIBE --private-key $OPERATOR_PRIVATE_KEY

    printf "\nSuccessfully Unsubscribed!\n"
}

## --- Main Script ---
printf "Step 1: set_contract_addresses\n"
set_contract_addresses

sleep 3s
printf "Step 2: unsubscribe_chain\n"
unsubscribe_chain
