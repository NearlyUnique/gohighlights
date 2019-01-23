package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_bad_server(t *testing.T) {
	var actualURL string
	badResponder := func(w http.ResponseWriter, r *http.Request) {
		actualURL = r.URL.Path
		w.WriteHeader(http.StatusConflict)
	}

	ts := httptest.NewServer(http.HandlerFunc(badResponder))
	defer ts.Close()

	client := NewMyClient(ts.Client(), ts.URL)

	_, err := client.ViewMessage()

	require.Error(t, err)
	assert.IsType(t, err, ConflictError{})
	assert.Contains(t, actualURL, "/view")
}
