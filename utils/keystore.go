package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	nutils "github.com/Lagrange-Labs/lagrange-node/utils"
)

const keystoreDir = ".lagrange/keystore"

// KeyStore is a struct that holds the private and public keys.
type KeyStore struct {
	PrivKey string `json:"privKey"`
	PubKey  string `json:"pubKey"`
}

// NewKeyStore creates a new KeyStore instance.
func NewKeyStore(privKey, pubKey []byte) *KeyStore {
	return &KeyStore{
		PrivKey: nutils.Bytes2Hex(privKey),
		PubKey:  nutils.Bytes2Hex(pubKey),
	}
}

// SaveKeyStore saves the private and public keys to a file.
func SaveKeyStore(keyType string, ks *KeyStore) error {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get home directory: %w", err)
	}
	fileName := fmt.Sprintf("%s_%s.json", keyType, ks.PubKey)
	keyFilePath := filepath.Clean(filepath.Join(homePath, keystoreDir, fileName))
	if _, err := os.Stat(keyFilePath); err == nil {
		return fmt.Errorf("keystore file already exists: %s", keyFilePath)
	}
	if err := os.MkdirAll(filepath.Dir(keyFilePath), 0700); err != nil {
		return fmt.Errorf("failed to create keystore directory: %w", err)
	}
	ksBytes, err := json.Marshal(ks)
	if err != nil {
		return err
	}

	return os.WriteFile(keyFilePath, ksBytes, 0644)
}

// LoadKeyStores loads all the keystore files from the keystore directory.
func LoadKeyStores(keyType string) ([]*KeyStore, error) {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("failed to get home directory: %w", err)
	}
	keyDirPath := filepath.Clean(filepath.Join(homePath, keystoreDir))
	files, err := os.ReadDir(keyDirPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read keystore directory: %w", err)
	}
	var keyStores []*KeyStore
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if filepath.Ext(file.Name()) != ".json" {
			continue
		}
		if keyType != "" && file.Name()[:len(keyType)] != keyType {
			continue
		}
		filePath := filepath.Clean(filepath.Join(keyDirPath, file.Name()))
		ksBytes, err := os.ReadFile(filePath)
		if err != nil {
			return nil, fmt.Errorf("failed to read keystore file: %w", err)
		}
		var ks KeyStore
		if err := json.Unmarshal(ksBytes, &ks); err != nil {
			return nil, fmt.Errorf("failed to unmarshal keystore: %w", err)
		}
		keyStores = append(keyStores, &ks)
	}

	return keyStores, nil
}
