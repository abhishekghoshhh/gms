package http

import (
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCustomClient(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/text")
		io.WriteString(w, `hello`)
	}

	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()


	t.Run("should send the request", func(t *testing.T) {
		client := NewClient()

		requestConf := Request(server.URL, "/", "GET")
		bytes, err := client.Send(requestConf)

		assert.NoError(t, err)
		assert.Equal(t, []byte("hello"), bytes)
	})



	t.Run("should give error for invalid url", func(t *testing.T) {
		client := NewClient()

		requestConf := Request("", "/", "GET")
		_, err := client.Send(requestConf)

		assert.Error(t, err)
	})
}
