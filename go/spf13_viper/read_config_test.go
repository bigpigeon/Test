package main

import (
	"bytes"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestReadConfig(t *testing.T) {
	viper.SetDefault("ContentDir", "content")
	viper.SetDefault("LayoutDir", "layouts")
	viper.SetConfigType("yaml") // or viper.SetConfigType("YAML")

	// any approach to require this configuration into your program.
	var yamlExample = []byte(`
Hacker: true
name: steve
hobbies:
- skateboarding
- snowboarding
- go
clothing:
  jacket: leather
  trousers: denim
age: 35
eyes : brown
beard: true
`)
	viper.AutomaticEnv()
	err := viper.ReadConfig(bytes.NewBuffer(yamlExample))
	require.NoError(t, err)

	t.Log(viper.Get("name")) // this would be "steve"
	t.Log(viper.Get("ContentDir"))
	t.Log(viper.Get("clothing.jacket"))
	t.Log(viper.Get("hobbies"))
	type DataBind struct {
		Name     string   `yaml:"name"`
		Hobbies  []string `yaml:"hobbies"`
		Clothing struct {
			Jacket   string `yaml:"jacket"`
			Trousers string `yaml:"trousers"`
		} `yaml:"clothing"`
	}
	var data DataBind
	err = viper.Unmarshal(&data)
	require.NoError(t, err)
	t.Logf("%#v\n", data)
}

func TestReadEnv(t *testing.T) {
	viper.SetEnvPrefix("spf") // will be uppercased automatically
	require.NoError(t, viper.BindEnv("id"))

	err := os.Setenv("SPF_ID", "13") // typically done outside of the app
	require.NoError(t, err)
	id := viper.Get("id") // 13
	t.Log(id)
	viper.AutomaticEnv()
	err = os.Setenv("SPF_NAME", "test")
	require.NoError(t, err)
	t.Log(viper.Get("name"))
}

func TestReadFlags(t *testing.T) {

}

//func TestReadRemote(t *testing.T) {
//	viper.AddRemoteProvider("etcd", "")
//}
