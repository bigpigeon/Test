/*
 * Copyright 2020 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package url_test

import (
	"github.com/stretchr/testify/require"
	"net/url"
	"testing"
)

func TestUrlParse(t *testing.T) {
	for _, u := range []string{
		"http://localhost:9200",
		"localhost",
		"/root/data/in/path",
		"current/data/in/path",
		"192.168.0.64/html/data/",
		"192.168.0.64:32000/html/data/",
	} {
		uri1, err := url.Parse(u)
		require.NoError(t, err)
		t.Logf("url %s uri %#v\n", u, uri1)
	}

}

func TestUrlParseRequestUri(t *testing.T) {
	for _, u := range []string{
		"http://localhost:9200",
		"192.168.0.64/html/data/",
		"192.168.0.64:32000/html/data/",
	} {
		uri1, err := url.ParseRequestURI(u)
		require.NoError(t, err)
		t.Logf("url %s uri %#v\n", u, uri1)
	}
}
