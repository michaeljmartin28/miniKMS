package kms

import (
	"context"
	"net/http"
	"time"
)

type Client struct {
	baseURL string
	http    *http.Client
}

func NewClient(baseURL string) *Client {

	return &Client{
		baseURL: baseURL,
		http: &http.Client{
			Timeout: time.Second * 10,
		},
	}
}

func (c *Client) CreateKey(ctx context.Context)   {}
func (c *Client) Encrypt(ctx context.Context)     {}
func (c *Client) Decrypt(ctx context.Context)     {}
func (c *Client) EnableKey(ctx context.Context)   {}
func (c *Client) DisableKey(ctx context.Context)  {}
func (c *Client) GenerateDEK(ctx context.Context) {}
func (c *Client) DecryptDEK(ctx context.Context)  {}
func (c *Client) RotateKey(ctx context.Context)   {}
