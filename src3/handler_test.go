package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_http_testing(t *testing.T) {
	rr := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)

	// just call the handler directly
	helloJSONWorld(rr, req)

	var msg map[string]string
	err := json.Unmarshal(rr.Body.Bytes(), &msg)

	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "Hello World", msg["message"])
}
