package toml_test

import (
	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v2"
	"testing"
)

type Cell struct {
	A int
	B string
	C float32
	D bool
}

type TestStruct struct {
	A   int
	B   string
	C   float32
	D   bool
	Sub Cell
}

func TestParserToml(t *testing.T) {
	var TestToml = `
A = 1
B = "2"
C = 3.0
d = true

[Sub]
  A = 1
  B = "2"
  C = 3.0
  d = true
`
	var _struct TestStruct
	_, err := toml.Decode(TestToml, &_struct)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("%#v\n", _struct)
}

func TestParserYaml(t *testing.T) {
	var TestYaml = `

a: 1
B: "2"
c: 3.0
d: true
sub:
  a: 1
  b: "2"
  c: 3.0
  d: true
`
	var _struct TestStruct
	err := yaml.Unmarshal([]byte(TestYaml), &_struct)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("%#v\n", _struct)

}
