package time_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestFormat(t *testing.T) {
	pTime, err := time.Parse(time.RFC3339, "2002-10-02T10:00:00Z")
	assert.NoError(t, err)
	t.Log(pTime)
}
