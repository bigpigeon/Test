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
	{

		pTime, err := time.Parse(time.RFC3339, "2002-10-02T10:00:00+08:00")
		assert.NoError(t, err)
		t.Log(pTime)
	}
}

func TestTimeUnix(t *testing.T) {
	{
		unixTime := 1542190917636
		ti := time.Unix(int64(unixTime), 0)
		t.Log(ti)
	}
	{

	}
}

func TestSince(t *testing.T) {
	start := time.Now()
	time.Sleep(2)
	t.Log(time.Since(start))
}
