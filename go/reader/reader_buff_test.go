package reader

import (
	"bytes"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestReaderBuffer(t *testing.T) {
	data := "abcdefghijk"
	buffer := bytes.NewBuffer([]byte(data))

	buff := make([]byte, 5)
	for {
		n, err := buffer.Read(buff)
		t.Log(string(buff), n, err)
		require.NoError(t, err)
	}

}
