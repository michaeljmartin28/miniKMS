package grpcsrv

import (
	"context"

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
		Version:    uint32(resp.Version),
	}, nil
}

func (s *GRPCServer) Decrypt(ctx context.Context, req *kmsv1.DecryptRequest) (*kmsv1.DecryptResponse, error) {
	coreReq := core.DecryptRequest{
		KeyID:          req.KeyId,
		Ciphertext:     req.Ciphertext,
		AdditionalData: req.AdditionalData,
		Version:        int(req.Version),
	}

	resp, err := s.Engine.Decrypt(ctx, coreReq)
	if err != nil {
		return nil, mapErrorToGRPC(err)
	}

	return &kmsv1.DecryptResponse{
		Plaintext: resp.Plaintext,
	}, nil
}
