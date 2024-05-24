# --- Dependencies ---
# 1. sh
# 2. `cast` - a binary from the foundry toolkit https://github.com/foundry-rs/foundry?tab=readme-ov-file
#           Version: tag nightly-f625d0fa7c51e65b4bf1e8f7931cd1c6e2e285e9
#                   cargo install --git https://github.com/foundry-rs/foundry --tag nightly-f625d0fa7c51e65b4bf1e8f7931cd1c6e2e285e9 --profile local --locked cast --force

# ==== MODIFY ME! ====
OPERATOR_PRIVATE_KEY=
OPERATOR_ADDR=
export ETH_RPC_URL= # Put rpc url here
PUBKEY_X= # Put pubkey X here
PUBKEY_Y= # Put pubkey Y here
SIG_EXPIRY_SECONDS=300

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

# The signature you create + send is valid for SIG_EXPIRY_SECONDS
calculate_expiry_timestamp() {
    (( EXPIRY_TIMESTAMP = $SIG_EXPIRY_SECONDS + $(cast block -f timestamp) ))
}

# A salt for your signature
generate_salt() {
    SALT=0x$(openssl rand -hex 32)
}

# Call read-only method on eigenlayer contract to get hash to sign for registration tx
calculate_registration_hash() {
    HASH=$(cast call $EIGEN_AVS_DIRECTORY_ADDR \
        "calculateOperatorAVSRegistrationDigestHash(address,address,bytes32,uint256)" \
        "$OPERATOR_ADDR" \
        "$LAGRANGE_SERVICE_ADDR" \
        "${SALT}" \
        "${EXPIRY_TIMESTAMP}")
    if [ -z "$HASH" ]; then
        echo "Error: Unable to calculate registration hash."
        exit 1
    fi
    printf "\nRegistration hash:\n${HASH}\n"
}

# Sign the registration hash with your OPERATOR_PRIVATE_KEY
sign_registration_hash() {
    printf "${HASH} \n"
    SIGNATURE=$(cast wallet sign --private-key $OPERATOR_PRIVATE_KEY $HASH --no-hash)
    printf "\nRegistration signature:${SIGNATURE}\n"
}

# === MODIFY ME ! ===
# Register with ZKMR contract using your signature and public key component of your new $LAGR_KEY
register_operator() {
    printf "cast send $LAGRANGE_SERVICE_ADDR \n"
    printf "register(address,uint256[2][],(bytes,bytes32,uint256)) \n"
    printf "$OPERATOR_ADDR \n"
    printf "[[${PUBKEY_X},${PUBKEY_Y}]] \n"
    printf "(${SIGNATURE},${SALT},${EXPIRY_TIMESTAMP}) \n"

    cast send $LAGRANGE_SERVICE_ADDR \
        "register(address,uint256[2][],(bytes,bytes32,uint256))" \
        "$OPERATOR_ADDR" \
        "[[${PUBKEY_X},${PUBKEY_Y}]]" \
        "(${SIGNATURE},${SALT},${EXPIRY_TIMESTAMP})" \
        --private-key $OPERATOR_PRIVATE_KEY

    printf "\nSuccessfully Registered!\n"
}

## --- Main Script ---
printf "Step 1: generate_bls_pubkey\n"
generate_bls_pubkey
sleep 3s
printf "Step 2: set_contract_addresses\n"
set_contract_addresses
sleep 3s
printf "Step 3: calculate_expiry_timestamp\n"
calculate_expiry_timestamp
sleep 3s
printf "Step 4: generate_salt\n"
generate_salt

printf "PUBKEY_X: $PUBKEY_X\n"
printf "PUBKEY_Y: $PUBKEY_Y\n"
printf "\nOperator Address: $OPERATOR_ADDR\n"
printf "Salt for signature: $SALT\n"
printf "Signature expires at: $EXPIRY_TIMESTAMP\n"

sleep 3s
printf "Step 5: calculate_registration_hash\n"
calculate_registration_hash
sleep 3s
printf "Step 6: sign_registration_hash\n"
sign_registration_hash
sleep 3s
printf "Step 7: register_operator\n"
register_operator
