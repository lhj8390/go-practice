package main

import (
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"testing"
)

func TestHttpMockAPI(t *testing.T) {
	// block all HTTP requests
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder(
		"GET",
		"https://example.com",
		httpmock.NewStringResponder(200, "hello!"),
	)
	resp, _ := http.Get("https://example.com")

	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	assert.Equal(t, "hello!", string(body))
	assert.Equal(t, 200, resp.StatusCode)
}
