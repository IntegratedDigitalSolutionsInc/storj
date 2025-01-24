// Copyright (C) 2024 Storj Labs, Inc.
// See LICENSE for copying information.

package metasearch

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"storj.io/common/version"
	"storj.io/storj/satellite"
	"storj.io/uplink"
)

var (
	userAgent = "MetaSearch/" + version.Build.Version.String()
)

// Auth authenticates HTTP requests for metasearch
type Auth interface {
	Authenticate(ctx context.Context, r *http.Request) (project *uplink.Project, err error)
}

// HeaderAuth authenticates metasearch HTTP requests based on the Authorization header
type HeaderAuth struct {
	db satellite.DB
}

func NewHeaderAuth(db satellite.DB) *HeaderAuth {
	return &HeaderAuth{
		db: db,
	}
}

func (a *HeaderAuth) Authenticate(ctx context.Context, r *http.Request) (project *uplink.Project, err error) {
	// Parse authorization header
	hdr := r.Header.Get("Authorization")
	if hdr == "" {
		err = fmt.Errorf("%w: missing authorization header", ErrAuthorizationFailed)
		return
	}

	// Check for valid authorization
	if !strings.HasPrefix(hdr, "Bearer ") {
		err = fmt.Errorf("%w: invalid authorization header", ErrAuthorizationFailed)
		return
	}

	// Parse access token
	rawAccess := strings.TrimPrefix(hdr, "Bearer ")
	access, err := uplink.ParseAccess(rawAccess)
	if err != nil {
		err = fmt.Errorf("%w: %s", ErrAuthorizationFailed, err)
		return
	}

	config := uplink.Config{
		UserAgent: userAgent,
	}
	return config.OpenProject(ctx, access)
}
