package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPTransport(t *testing.T) {
	listenAddr := ":4000"
	tr := NewTCPTransport(listenAddr)
	assert.Equal(t, listenAddr, tr.listenAddress, listenAddr)

	// Server

	tr.ListenAndAccept()
	// tr.Start()
}