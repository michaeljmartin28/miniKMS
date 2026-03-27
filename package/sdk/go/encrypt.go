package kms

import (
	"context"
	"fmt"
	"net/http"
)

func (c *Client) Encrypt(ctx context.Context, keyID string, params EncryptParams) (*EncryptResponse, error) {
	var resp EncryptResponse
	route := fmt.Sprintf("/v1/keys/%s/encrypt", keyID)

	err := c.do(ctx, http.MethodPost, route, params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) Decrypt(ctx context.Context, keyID string, params DecryptParams) (*DecryptResponse, error) {
	var resp DecryptResponse
	route := fmt.Sprintf("/v1/keys/%s/decrypt", keyID)

	err := c.do(ctx, http.MethodPost, route, params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
