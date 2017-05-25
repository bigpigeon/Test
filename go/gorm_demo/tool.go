package gormdemo

import (
	"testing"
)

func ErrorProcess(t *testing.T, err error) {
	if err != nil {
		t.Error(err)
	}
}
