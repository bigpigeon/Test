/*
 * Copyright 2020 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package oauth

import (
	"context"
	"fmt"
	"golang.org/x/oauth2"
	"testing"
)

func TestConf(t *testing.T) {
	ctx := context.Background()

	conf := &oauth2.Config{
		ClientID:     "test_123",
		ClientSecret: "password",
		Scopes:       []string{"SCOPE1", "SCOPE2"},
		Endpoint: oauth2.Endpoint{
			TokenURL: "https://provider.com/o/oauth2/token",
			AuthURL:  "https://provider.com/o/oauth2/auth",
		},
	}

	// Redirect user to consent page to ask for permission
	// for the scopes specified above.
	url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
	fmt.Printf("Visit the URL for the auth dialog: %v", url)

	// Use the authorization code that is pushed to the redirect
	// URL. Exchange will do the handshake to retrieve the
	// initial access token. The HTTP Client returned by
	// conf.Client will refresh the token as necessary.
	var code string
	if _, err := fmt.Scan(&code); err != nil {
		t.Fatal(err)
	}
	tok, err := conf.Exchange(ctx, code)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tok.AccessToken)

	client := conf.Client(ctx, tok)
	client.Get("...")
}
