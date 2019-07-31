package main

import (
	"flag"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
)

type Data struct {
	FlagName string
	Flag     struct {
		Name string
	}
}

func main() {
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.StringP("flagname", "f", "", "help flag name")
	pflag.String("flag.name", "", "help flag name")
	pflag.Parse()

	err := viper.BindPFlags(pflag.CommandLine)
	if err != nil {
		panic(err)
	}
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.Print(viper.Get("flagname"))
	log.Print(viper.Get("flag.name"))
	var data Data
	err = viper.Unmarshal(&data)
	if err != nil {
		panic(err)
	}
	log.Printf("data %#v\n", data)
}
