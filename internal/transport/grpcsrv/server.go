package grpcsrv

import (
	"context"
	"time"

	kmsv1 "github.com/michaeljmartin28/minikms/gen/kms/v1"
	"github.com/michaeljmartin28/minikms/internal/core"
)

type GRPCServer struct {
	kmsv1.UnimplementedKMSServer
	Engine *core.Engine
}

func NewGRPCServer(engine *core.Engine) *GRPCServer {
	return &GRPCServer{
		Engine: engine,
	}
}

func (s *GRPCServer) CreateKey(ctx context.Context, req *kmsv1.CreateKeyRequest) (*kmsv1.CreateKeyResponse, error) {
	coreReq := core.CreateKeyRequest{
		Algorithm: core.Algorithm(req.Algorithm),
		Name:      req.Name,
	}

	resp, err := s.Engine.CreateKey(ctx, coreReq)
	if err != nil {
		return nil, mapErrorToGRPC(err)
	}

	return &kmsv1.CreateKeyResponse{
		KeyId:     resp.KeyID,
		Version:   resp.Version,
		CreatedAt: resp.CreateAt.Format(time.RFC3339),
	}, nil
}

func (s *GRPCServer) Encrypt(ctx context.Context, req *kmsv1.EncryptRequest) (*kmsv1.EncryptResponse, error) {
	coreReq := core.EncryptRequest{
		KeyID:          req.KeyId,
		Plaintext:      req.Plaintext,
		AdditionalData: req.AdditionalData,
	}

	resp, err := s.Engine.Encrypt(ctx, coreReq)
	if err != nil {
		return nil, mapErrorToGRPC(err)
	}

	return &kmsv1.EncryptResponse{
		Ciphertext: resp.Ciphertext,
		Version:    resp.Version,
	}, nil
}

func (s *GRPCServer) Decrypt(ctx context.Context, req *kmsv1.DecryptRequest) (*kmsv1.DecryptResponse, error) {
	coreReq := core.DecryptRequest{
		KeyID:          req.KeyId,
		Ciphertext:     req.Ciphertext,
		AdditionalData: req.AdditionalData,
		Version:        req.Version,
	}

	resp, err := s.Engine.Decrypt(ctx, coreReq)
	if err != nil {
		return nil, mapErrorToGRPC(err)
	}

	return &kmsv1.DecryptResponse{
		Plaintext: resp.Plaintext,
	}, nil
}

func (s *GRPCServer) GenerateDataKey(ctx context.Context, req *kmsv1.GenerateDataKeyRequest) (*kmsv1.GenerateDataKeyResponse, error) {
	coreReq := core.GenerateDataKeyRequest{
		KeyID:          req.KeyId,
		AdditionalData: req.AdditionalData,
	}

	resp, err := s.Engine.GenerateDataKey(ctx, coreReq)
	if err != nil {
		return nil, mapErrorToGRPC(err)
	}

	return &kmsv1.GenerateDataKeyResponse{
		Plaintext:    resp.PlaintextDEK,
		EncryptedDek: resp.EncryptedDEK,
		Version:      resp.Version,
	}, nil
}

func (s *GRPCServer) DecryptDataKey(ctx context.Context, req *kmsv1.DecryptDataKeyRequest) (*kmsv1.DecryptDataKeyResponse, error) {
	coreReq := core.DecryptDataKeyRequest{
		KeyID:          req.KeyId,
		EncryptedDEK:   req.EncryptedDek,
		Version:        req.Version,
		AdditionalData: req.AdditionalData,
	}

	resp, err := s.Engine.DecryptDataKey(ctx, coreReq)
	if err != nil {
		return nil, mapErrorToGRPC(err)
	}

	return &kmsv1.DecryptDataKeyResponse{
		Plaintext: resp.PlaintextDEK,
	}, nil
}

func (s *GRPCServer) RotateKey(ctx context.Context, req *kmsv1.RotateKeyRequest) (*kmsv1.RotateKeyResponse, error) {

	resp, err := s.Engine.RotateKey(ctx, req.KeyId)
	if err != nil {
		return nil, mapErrorToGRPC(err)
	}

	return &kmsv1.RotateKeyResponse{
		Version: resp,
	}, nil
}

func (s *GRPCServer) EnableKey(ctx context.Context, req *kmsv1.EnableKeyRequest) (*kmsv1.KeyMetadata, error) {

	resp, err := s.Engine.EnableKey(ctx, req.KeyId)
	if err != nil {
		return nil, mapErrorToGRPC(err)
	}

	return &kmsv1.KeyMetadata{
		KeyId:         resp.KeyID,
		Enabled:       resp.State.IsEnabled(),
		CreatedAt:     resp.CreatedAt.Format(time.RFC3339),
		LatestVersion: resp.LatestVersion,
		Algorithm:     string(resp.Algorithm),
	}, nil
}

func (s *GRPCServer) DisableKey(ctx context.Context, req *kmsv1.DisableKeyRequest) (*kmsv1.KeyMetadata, error) {

	resp, err := s.Engine.DisableKey(ctx, req.KeyId)
	if err != nil {
		return nil, mapErrorToGRPC(err)
	}

	return &kmsv1.KeyMetadata{
		KeyId:         resp.KeyID,
		Enabled:       resp.State.IsEnabled(),
		CreatedAt:     resp.CreatedAt.Format(time.RFC3339),
		LatestVersion: resp.LatestVersion,
		Algorithm:     string(resp.Algorithm),
	}, nil
}
