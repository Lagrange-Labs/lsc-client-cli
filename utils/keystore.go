package utils

import (
	"fmt"
	"os"
	"path/filepath"

	ecrypto "github.com/ethereum/go-ethereum/crypto"

	"github.com/Lagrange-Labs/lagrange-node/crypto"
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

func saveKeystore(keyType string, password string, pubKey, privKey []byte) error {
	ksPath := filepath.Join(keystoreDir, fmt.Sprintf("%s_%x.key", keyType, pubKey[:6]))
	var cryptoCurve crypto.CryptoCurve
	if keyType == "ecdsa" {
		cryptoCurve = crypto.CryptoCurve("ECDSA")
	} else if keyType == "bls" {
		cryptoCurve = crypto.CryptoCurve("BN254")
	}
	return crypto.SaveKey(cryptoCurve, privKey, password, ksPath)
}

// LoadKeyStores loads all the keystore files from the keystore directory.
func LoadKeyStores(keyType string) ([]string, error) {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("failed to get home directory: %w", err)
	}
	keyDirPath := filepath.Clean(filepath.Join(homePath, keystoreDir))
	files, err := os.ReadDir(keyDirPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read keystore directory: %w", err)
	}
	keystoreFiles := make([]string, 0)
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if filepath.Ext(file.Name()) != ".key" {
			continue
		}
		if keyType != "" && file.Name()[:len(keyType)] != keyType {
			continue
		}
		keystoreFiles = append(keystoreFiles, filepath.Clean(filepath.Join(keyDirPath, file.Name())))
	}

	return keystoreFiles, nil
}
