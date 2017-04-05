package storage

import (
	"context"
	"errors"

	"github.com/ains/niso"
)

// ExampleStorage is a backend used to persist data generated by the example OAuth2 server
type ExampleStorage struct {
	clients   map[string]*niso.ClientData
	authorize map[string]*niso.AuthorizationData
	access    map[string]*niso.AccessData
	refresh   map[string]*niso.RefreshTokenData
}

// NewExampleStorage returns a new ExampleStorage
func NewExampleStorage() *ExampleStorage {
	r := &ExampleStorage{
		clients:   make(map[string]*niso.ClientData),
		authorize: make(map[string]*niso.AuthorizationData),
		access:    make(map[string]*niso.AccessData),
		refresh:   make(map[string]*niso.RefreshTokenData),
	}

	r.clients["1234"] = &niso.ClientData{
		ClientID:     "1234",
		ClientSecret: "aabbccdd",
		RedirectURI:  "http://localhost:14000/appauth",
	}

	return r
}

// Close the resources the Storage potentially holds. (Implements io.Closer)
func (s *ExampleStorage) Close() error {
	return nil
}

// GetClientData fetches the data for a ClientData by id
// Should return NotFoundError, so an EInvalidClient error will be returned instead of EServerError
func (s *ExampleStorage) GetClientData(_ context.Context, id string) (*niso.ClientData, error) {
	if c, ok := s.clients[id]; ok {
		return c, nil
	}
	return nil, &niso.NotFoundError{Err: errors.New("client not found")}
}

// SaveAuthorizeData saves authorize data.
func (s *ExampleStorage) SaveAuthorizeData(_ context.Context, data *niso.AuthorizationData) error {
	s.authorize[data.Code] = data
	return nil
}

// GetAuthorizeData looks up AuthorizeData by a code.
//// ClientData information MUST be loaded together.
// Optionally can return error if expired.
func (s *ExampleStorage) GetAuthorizeData(_ context.Context, code string) (*niso.AuthorizationData, error) {
	if d, ok := s.authorize[code]; ok {
		return d, nil
	}
	return nil, errors.New("authorize not found")
}

// DeleteAuthorizeData revokes or deletes the authorization code.
func (s *ExampleStorage) DeleteAuthorizeData(ctx context.Context, code string) error {
	delete(s.authorize, code)
	return nil
}

// SaveAccessData writes AccessData to storage.
func (s *ExampleStorage) SaveAccessData(ctx context.Context, data *niso.AccessData) error {
	s.access[data.AccessToken] = data
	return nil
}

// GetRefreshTokenData retrieves refresh token data from the token string.
func (s *ExampleStorage) GetRefreshTokenData(ctx context.Context, token string) (*niso.RefreshTokenData, error) {
	if d, ok := s.refresh[token]; ok {
		return d, nil
	}
	return nil, errors.New("refresh token data not found")
}

// SaveRefreshTokenData saves refresh token data so it can be retrieved with GetRefreshTokenData
func (s *ExampleStorage) SaveRefreshTokenData(ctx context.Context, data *niso.RefreshTokenData) error {
	s.refresh[data.RefreshToken] = data
	return nil
}

// DeleteRefreshTokenData revokes or deletes a RefreshToken.
func (s *ExampleStorage) DeleteRefreshTokenData(ctx context.Context, token string) error {
	delete(s.refresh, token)
	return nil
}
