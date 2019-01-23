package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_http_client_testing(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(helloJSONWorld))
	defer ts.Close()

	client := NewMyClient(ts.Client(), ts.URL)

	msg, err := client.ViewMessage()

	require.NoError(t, err)
	assert.Equal(t, "Hello World", msg.Text)
}
