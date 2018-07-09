/*
 * Copyright 2018. bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package regexp_example

import (
	"regexp"
	"testing"
)

func TestRegexRepeat(t *testing.T) {
	re := regexp.MustCompile("`(\\w)`")
	match := re.FindAllString("`a`,`b`,`c`", 3)
	for _, m := range match {
		t.Log(m)
	}
}
