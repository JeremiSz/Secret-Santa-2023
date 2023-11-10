package server

import (
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
