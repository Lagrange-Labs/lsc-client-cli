## Context

This is the simple command line tool that can generate the BLS signature for operator registration. 
To generate the signature, you need to provide the BLS private key and the digest hash:

```sh
go run main.go generate-bls-signature -k <BLS_PRIVATE_KEY> -d <DIGEST_HASH>
```

It will display the `BLSKeyWithProof` object line by line. You can copy and paste it to the operator registration request. Please refer to the [Operator Registration](../register-operator.sh#L104) script for more details.