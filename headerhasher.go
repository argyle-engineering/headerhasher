// Package headerhasher is a traefik middleware that looks for a specified input header and attaches its sha256 hash as a new header.
package headerhasher

import (
	"context"
	"crypto/sha256"
	"fmt"
	"hash"
	"net/http"
)

// Config holds the plugin configuration.
type Config struct {
	InputHeader  string `json:"inputHeader,omitempty"`
	OutputHeader string `json:"outputHeader,omitempty"`
}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{
		InputHeader:  "Authorization",
		OutputHeader: "Authorization-Hashed",
	}
}

// HeaderHasher represents the plugin which looks for a header named `inputHeader`
// and if found calculates sha256 of its value and attaches it as `outputHeader`.
type HeaderHasher struct {
	inputHeader  string
	outputHeader string
	next         http.Handler
	hasher       hash.Hash
}

// New creates a new HeaderHasher plugin.
func New(_ context.Context, next http.Handler, config *Config, _ string) (http.Handler, error) {
	if config.InputHeader == "" || config.OutputHeader == "" {
		return nil, fmt.Errorf("inputHeader and outputHeader cannot be empty")
	}

	return &HeaderHasher{
		inputHeader:  config.InputHeader,
		outputHeader: config.OutputHeader,
		next:         next,
		hasher:       sha256.New(),
	}, nil
}

func (a *HeaderHasher) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	value := req.Header.Get(a.inputHeader)
	if value != "" {
		a.hasher.Write([]byte(value))
		hashedValue := fmt.Sprintf("%x", a.hasher.Sum(nil))
		a.hasher.Reset()

		req.Header.Set(a.outputHeader, hashedValue)
	}

	a.next.ServeHTTP(rw, req)
}
