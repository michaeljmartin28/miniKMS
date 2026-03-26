package kms

import (
	"context"
	"net/http"
)

func (c *Client) CreateKey(ctx context.Context, params CreateKeyParams) (*Key, error) {

	var resp CreateKeyResponse
	err := c.do(ctx, http.MethodPost, "v1/keys", params, &resp)
	if err != nil {
		return nil, err
	}
	return &Key{
		KeyID:     resp.KeyID,
		Version:   resp.Version,
		CreatedAt: resp.CreatedAt,
	}, nil
}

func (c *Client) Encrypt(ctx context.Context)    {}
func (c *Client) Decrypt(ctx context.Context)    {}
func (c *Client) EnableKey(ctx context.Context)  {}
func (c *Client) DisableKey(ctx context.Context) {}
func (c *Client) RotateKey(ctx context.Context)  {}
