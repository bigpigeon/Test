/*
 * Copyright 2019 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package concurrent

import (
	"errors"
	"github.com/stretchr/testify/require"
	"sync"
	"testing"
)

func routineError(t *testing.T, wg *sync.WaitGroup) {
	defer wg.Done()
	require.NoError(t, errors.New("error"))
}

func TestRoutineError(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go routineError(t, &wg)
	wg.Wait()

}
