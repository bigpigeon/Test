package listen

import (
	"github.com/stretchr/testify/require"
	"net"
	"strings"
	"testing"
)

func TestListenAddr(t *testing.T) {
	li, err := net.Listen("tcp", "0.0.0.0:")
	require.NoError(t, err)
	t.Log(li.Addr())
	t.Log(li.Addr().Network())
	idx := strings.LastIndex(li.Addr().String(), ":")
	if idx != -1 {
		t.Log("port", li.Addr().String()[idx+1:])
	}
}

func TestListenPortAddr(t *testing.T) {
	{
		li, err := net.Listen("tcp", ":8080")
		require.NoError(t, err)
		t.Log(li.Addr())
		t.Log(li.Addr().Network())
		idx := strings.LastIndex(li.Addr().String(), ":")
		if idx != -1 {
			t.Log("port", li.Addr().String()[idx+1:])
		}
	}
	{

		li, err := net.Listen("tcp", "0.0.0.0:abcd")
		require.NoError(t, err)
		t.Log(li.Addr())
		t.Log(li.Addr().Network())
		idx := strings.LastIndex(li.Addr().String(), ":")
		if idx != -1 {
			t.Log("port", li.Addr().String()[idx+1:])
		}
	}
}
