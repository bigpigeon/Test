/*
 * Copyright 2020 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package flags

import (
	"flag"
	"os"
	"testing"
)

var str = flag.String("testdata", "123", "")

func init() {
	flag.Parse()
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestFlag(t *testing.T) {
	t.Log(*str)

}
