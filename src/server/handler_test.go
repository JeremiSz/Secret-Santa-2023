package server

import (
	"net/http"
	"strconv"
	"testing"
)

func TestCreatingServer(t *testing.T) {
	server := NewServer()
	if server == nil {
		t.Error("server not created")
		return
	}
	if server.Addr != ":8080" {
		t.Error("server address not set")
	}
	if server.Handler == nil {
		t.Error("server handler not set")
	}
}

type testWriter struct {
	header http.Header
	body   []byte
	size   int
}

func (t *testWriter) Header() http.Header {
	return t.header
}
func (t *testWriter) Write([]byte) (int, error) {
	t.body = []byte("test")
	t.size = len(t.body)
	return t.size, nil
}
func (t *testWriter) WriteHeader(statusCode int) {
	t.header.Add("status", strconv.Itoa(statusCode))
}
