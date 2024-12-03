// Copyright (C) 2024 Storj Labs, Inc.
// See LICENSE for copying information.
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-oauth2/oauth2/v4/errors"
)

// MetaSearchClient proides a client for the metasearch REST service.
type MetaSearchClient struct {
	access *AccessOptions
	client *http.Client
}

func newMetaSearchClient(access *AccessOptions) *MetaSearchClient {
	client := &http.Client{}
	return &MetaSearchClient{
		access: access,
		client: client,
	}
}

func (c *MetaSearchClient) GetObjectMetadata(ctx context.Context, bucket string, key string) (meta map[string]interface{}, err error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.access.Server+"/metadata/"+bucket+"/"+key, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+c.access.Access)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, httpError(resp)
	}

	err = json.NewDecoder(resp.Body).Decode(&meta)
	if err != nil {
		return nil, fmt.Errorf("cannot decode metadata: %w", err)
	}

	return meta, nil
}

}

func httpError(resp *http.Response) error {
	switch resp.StatusCode {
	case http.StatusUnauthorized:
		return errors.New("unauthorized")
	case http.StatusNotFound:
		return errors.New("object not found")
	default:
		return fmt.Errorf("error response from server: %v", resp.Status)
	}
}
