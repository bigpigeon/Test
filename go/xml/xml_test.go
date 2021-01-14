/*
 * Copyright 2021 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package xml

import (
	"bytes"
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

type Address struct {
	City, State string
}
type Person struct {
	XMLName   xml.Name `xml:"person"`
	Id        int      `xml:"id,attr"`
	FirstName string   `xml:"name>first"`
	LastName  string   `xml:"name>last"`
	Age       int      `xml:"age"`
	Height    float32  `xml:"height,omitempty"`
	Married   bool
	Address
	Comment string `xml:",comment"`
}

func TestXmlParse(t *testing.T) {
	xmlData := `// Output:
<person id="13">
<name>
  <first>John</first>
        <last>Doe</last>
     </name>
<age>42</age>
<Married>false</Married>
<City>Hanga Roa</City>
<State>Easter Island&amp;abc</State>
<!-- Need more details. -->
</person>`
	person := Person{}
	decoder := xml.NewDecoder(strings.NewReader(xmlData))

	err := decoder.Decode(&person)
	assert.NoError(t, err)
	t.Logf("person %#v\n", person)
}

func TestEscapeText(t *testing.T) {
	var buff bytes.Buffer
	err := xml.EscapeText(&buff, []byte("http://localhost:8080?namespace=jia&id=abc"))
	require.NoError(t, err)
	t.Log(buff.String())
}
