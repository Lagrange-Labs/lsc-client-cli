package utils

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/Lagrange-Labs/lsc-node/core"
	"github.com/Lagrange-Labs/lsc-node/signer"
	"github.com/Lagrange-Labs/lsc-node/signer/types"
)

// SignerClient is a wrapper for the signer service client.
type SignerClient struct {
	types.SignerServiceClient
}

// NewSignerClient creates a new SignerClient instance.
func NewSignerClient(url string, cert *core.CertConfig) (*SignerClient, error) {
	opts := []grpc.DialOption{}
	if cert != nil {
		creds, err := core.LoadTLS(cert, false)
		if err != nil {
			return nil, err
		}
		opts = append(opts, grpc.WithTransportCredentials(credentials.NewTLS(creds)))
	} else {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	conn, err := grpc.NewClient(url, opts...)
	if err != nil {
		return nil, err
	}

	return &SignerClient{SignerServiceClient: types.NewSignerServiceClient(conn)}, nil
}

// Sign signs a message using the signer service.
func (c *SignerClient) Sign(msg []byte, accountID string, isG1 bool) ([]byte, error) {
	signMethod := signer.SignMethodG2
	if isG1 {
		signMethod = signer.SignMethodG1
	}

	resp, err := c.SignerServiceClient.Sign(context.Background(), &types.SignRequest{
		Message:    core.Bytes2Hex(msg),
		AccountId:  accountID,
		SignMethod: signMethod,
	})
	if err != nil {
		return nil, err
	}
	return core.Hex2Bytes(resp.Signature), nil
}

// GetPublicKey returns the public key for an account.
func (c *SignerClient) GetPublicKey(accountID string, isG1 bool) ([]byte, error) {
	signMethod := signer.PublicKeyMethodG2
	if isG1 {
		signMethod = signer.PublicKeyMethodG1
	}
	resp, err := c.SignerServiceClient.Sign(context.Background(), &types.SignRequest{
		AccountId:  accountID,
		SignMethod: signMethod,
	})
	if err != nil {
		return nil, err
	}
	return core.Hex2Bytes(resp.Signature), nil
}
