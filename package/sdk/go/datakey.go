package kms

import (
	"context"
	"fmt"
	"net/http"
)

func (c *Client) GenerateDEK(ctx context.Context, keyID string, params GenerateDataParams) (*GenerateDataKeyResponse, error) {
	var resp GenerateDataKeyResponse
	route := fmt.Sprintf("/v1/keys/%s/generate-data-key", keyID)

	err := c.do(ctx, http.MethodPost, route, params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
func (c *Client) DecryptDEK(ctx context.Context, keyID string, params DecryptDataKeyParams) (*DecryptDataKeyResponse, error) {
	var resp DecryptDataKeyResponse
	route := fmt.Sprintf("/v1/keys/%s/decrypt-data-key", keyID)

	err := c.do(ctx, http.MethodPost, route, params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
