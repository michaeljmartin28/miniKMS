package kms

import (
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
