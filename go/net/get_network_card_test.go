package net

import (
	"github.com/stretchr/testify/require"
	"net"
	"testing"
)

func TestGetNetworkCard(t *testing.T) {
	inets, err := net.Interfaces()
	require.NoError(t, err)
	for _, card := range inets {
		t.Logf("card %#v\n", card)
	}
}
