package kms

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
)

func (c *Client) do(ctx context.Context, method string, path string, in any, out any) error {

	// Build the request
	urlPath, err := url.JoinPath(c.baseURL, path)
	if err != nil {
		return err
	}

	var body io.Reader
	if in != nil {
		b, err := json.Marshal(in)
		if err != nil {
			return err
		}
		body = bytes.NewReader(b)
	}

	// Send the request
	req, err := http.NewRequestWithContext(ctx, method, urlPath, body)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	var resp *http.Response
	resp, err = c.http.Do(req)

	if err != nil {
		return err
	}
	defer resp.Body.Close()
	// unmarshal into out

	if resp.StatusCode != http.StatusOK {
		return errors.New("error in response")
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(bodyBytes, out); err != nil {
		return err
	}

	return nil

}
