package main

import (
	"context"
	"testing"
)

func TestContext(t *testing.T) {

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

}
