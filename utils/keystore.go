package utils

import (
	"fmt"
	"os"
	"path/filepath"

	ecrypto "github.com/ethereum/go-ethereum/crypto"

	"github.com/Lagrange-Labs/lagrange-node/crypto"
	nutils "github.com/Lagrange-Labs/lagrange-node/utils"
)

const keystoreDir = ".lagrange/keystore"

// GenerateKeystore generates a new keystore file for the given key.
func GenerateKeystore(keyType, password string) error {
	switch keyType {
	case "ecdsa":
		privateKey, err := ecrypto.GenerateKey()
		if err != nil {
			return fmt.Errorf("failed to generate ECDSA key: %w", err)
		}
		privKey := ecrypto.FromECDSA(privateKey)
		addr := ecrypto.PubkeyToAddress(privateKey.PublicKey)
		return saveKeystore(keyType, password, addr.Bytes(), privKey)
	case "bls":
		blsScheme := crypto.NewBLSScheme(crypto.BN254)
		privKey, err := blsScheme.GenerateRandomKey()
		if err != nil {
			return fmt.Errorf("failed to generate BLS key: %w", err)
		}
		pubKey, err := blsScheme.GetPublicKey(privKey, false)
		if err != nil {
			return fmt.Errorf("failed to get BLS public key: %w", err)
		}
		return saveKeystore(keyType, password, pubKey, privKey)
	default:
		return fmt.Errorf("invalid key type: %s", keyType)
	}
}

// ImportFromPrivateKey imports a private key to the keystore.
func ImportFromPrivateKey(keyType, password string, privKey []byte) error {
	switch keyType {
	case "ecdsa":
		privateKey, err := ecrypto.ToECDSA(privKey)
		if err != nil {
			return fmt.Errorf("failed to convert private key to ECDSA: %w", err)
		}
		addr := ecrypto.PubkeyToAddress(privateKey.PublicKey)
		return saveKeystore(keyType, password, addr.Bytes(), privKey)
	case "bls":
		blsScheme := crypto.NewBLSScheme(crypto.BN254)
		pubKey, err := blsScheme.GetPublicKey(privKey, false)
		if err != nil {
			return fmt.Errorf("failed to get BLS public key: %w", err)
		}
		return saveKeystore(keyType, password, pubKey, privKey)
	default:
		return fmt.Errorf("invalid key type: %s", keyType)
	}
}

// ExportKeystore exports the private key from the keystore file.
func ExportKeystore(keyType, password, filePath string) ([]byte, error) {
	switch keyType {
	case "ecdsa":
		return crypto.LoadPrivateKey(crypto.CryptoCurve("ECDSA"), password, filePath)
	case "bls":
		return crypto.LoadPrivateKey(crypto.CryptoCurve("BN254"), password, filePath)
	default:
		return nil, fmt.Errorf("invalid key type: %s", keyType)
	}
}

func saveKeystore(keyType string, password string, pubKey, privKey []byte) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get user home directory: %w", err)
	}
	ksPath := filepath.Join(homeDir, keystoreDir, fmt.Sprintf("%s_%x.key", keyType, pubKey[:6]))
	var cryptoCurve crypto.CryptoCurve
	if keyType == "ecdsa" {
		cryptoCurve = crypto.CryptoCurve("ECDSA")
	} else if keyType == "bls" {
		cryptoCurve = crypto.CryptoCurve("BN254")
	}
	DisplayWarningMessage(keyType, nutils.Bytes2Hex(privKey), ksPath)
	return crypto.SaveKey(cryptoCurve, privKey, password, ksPath)
}

// ReadPrivateKey reads the private key from the keystore file.
func ReadPrivateKey(keyType, password, filePath string) ([]byte, error) {
	return crypto.LoadPrivateKey(crypto.CryptoCurve(keyType), password, filePath)
}
