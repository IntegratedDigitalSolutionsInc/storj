// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package metasearch

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"go.uber.org/zap"

	"storj.io/common/macaroon"
	"storj.io/storj/satellite"
	"storj.io/storj/satellite/metabase"
)

// Server implements the REST API for metadata search.
type Server struct {
	Logger      *zap.Logger
	SatelliteDB satellite.DB
	MetabaseDB  *metabase.DB
	Endpoint    string
}

type SearchRequest struct {
	Page  int    `json:"page"`
	Path  string `json:"path"`
	Query string `json:"query"`
	Meta  string `json:"metadata"`
}

// NewServer creates a new metasearch server process.
func NewServer(log *zap.Logger, db satellite.DB, metabase *metabase.DB, endpoint string) (*Server, error) {
	peer := &Server{
		Logger:      log,
		SatelliteDB: db,
		MetabaseDB:  metabase,
		Endpoint:    endpoint,
	}

	return peer, nil
}

func (a *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var err error
	var reqBody SearchRequest
	var resp interface{}

	// Parse authorization header
	hdr := r.Header.Get("Authorization")
	if hdr == "" {
		a.ErrorResponse(w, fmt.Errorf("%w: missing authorization header", ErrAuthorizationFailed))
		return
	}

	// Check for valid authorization
	if !strings.HasPrefix(hdr, "Bearer ") {
		a.ErrorResponse(w, fmt.Errorf("%w: invalid authorization header", ErrAuthorizationFailed))
		return
	}

	// Parse API token
	rawToken := strings.TrimPrefix(hdr, "Bearer ")
	apiKey, err := macaroon.ParseAPIKey(rawToken)
	if err != nil {
		a.ErrorResponse(w, fmt.Errorf("%w: %s", ErrAuthorizationFailed, err))
		return
	}
	a.Logger.Info("API key", zap.String("key", fmt.Sprint(apiKey)))

	// Decode request body
	if err = json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		a.ErrorResponse(w, fmt.Errorf("error decoding request body: %w", err))
		return
	}

	// Handle request
	switch {
	case r.Method == http.MethodPost:
		if reqBody.Query == "" {
			resp, err = a.ViewMetadata(&reqBody)
		} else {
			resp, err = a.QueryMetadata(&reqBody)
		}
	case r.Method == http.MethodPut:
		err = a.UpdateMetadata(&reqBody)
	case r.Method == http.MethodDelete:
		err = a.DeleteMetadata(&reqBody)
	default:
		err = fmt.Errorf("%w: unsupported method %s", ErrBadRequest, r.Method)
	}

	// Write response
	if err != nil {
		a.ErrorResponse(w, err)
		return
	}

	jsonBytes, err := json.Marshal(resp)
	if err != nil {
		a.ErrorResponse(w, fmt.Errorf("error marshalling response: %w", err))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (a *Server) Run() error {
	mux := http.NewServeMux()

	// Register the routes and handlers
	mux.Handle("/meta_search", a)
	mux.Handle("/meta_search/", a)

	// Run the server
	return http.ListenAndServe(a.Endpoint, mux)
}

func (a *Server) ViewMetadata(reqBody *SearchRequest) (meta map[string]interface{}, err error) {
	meta = map[string]interface{}{
		"view": "meta",
	}
	return
}

func (a *Server) QueryMetadata(reqBody *SearchRequest) (meta map[string]interface{}, err error) {
	meta = map[string]interface{}{
		"query": "meta",
	}
	return
}

func (a *Server) UpdateMetadata(reqBody *SearchRequest) (err error) {
	return nil
}

func (a *Server) DeleteMetadata(reqBody *SearchRequest) (err error) {
	return nil
}

// ErrorResponse writes an error response to the client.
func (a *Server) ErrorResponse(w http.ResponseWriter, err error) {
	a.Logger.Warn("error during API request", zap.Error(err))

	var e *ErrorResponse
	if !errors.As(err, &e) {
		e = ErrInternalError
	}

	resp, _ := json.Marshal(e)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(e.StatusCode)
	w.Write([]byte(resp))
}

func BadRequestHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("400 Bad Request"))
}

func InternalServerErrorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("500 Internal Server Error"))
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("404 Not Found"))
}
