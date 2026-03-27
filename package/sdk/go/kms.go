package kms

import (
	"context"
	"fmt"
	"net/http"
)

func (c *Client) CreateKey(ctx context.Context, params CreateKeyParams) (*Key, error) {

	var resp CreateKeyResponse
	err := c.do(ctx, http.MethodPost, "/v1/keys", params, &resp)
	if err != nil {
		return nil, err
	}
	return &Key{
		KeyID:     resp.KeyID,
		Version:   resp.Version,
		CreatedAt: resp.CreatedAt,
	}, nil
}

func (c *Client) EnableKey(ctx context.Context, keyID string) (*KeyMetadata, error) {
	var resp KeyMetadata

	route := fmt.Sprintf("/v1/keys/%s/enable", keyID)
	err := c.do(ctx, http.MethodPost, route, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) DisableKey(ctx context.Context, keyID string) (*KeyMetadata, error) {
	var resp KeyMetadata

	route := fmt.Sprintf("/v1/keys/%s/disable", keyID)
	err := c.do(ctx, http.MethodPost, route, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) RotateKey(ctx context.Context, keyID string) (*RotateKeyResponse, error) {
	var resp RotateKeyResponse

	route := fmt.Sprintf("/v1/keys/%s/rotate", keyID)
	err := c.do(ctx, http.MethodPost, route, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
