/*
 * Copyright 2019 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */
package main

import (
	"flag"
	"fmt"
	"github.com/spf13/pflag"
	"strings"
)

var ip *int = flag.Int("ip", 1234, "help message for ip")
var help *bool = pflag.BoolP("help", "h", false, "help message for ip")
var addr = pflag.String("addr", "local", "help message for ip")
var addrName = pflag.StringP("addr-name", "a", "local", "help message for ip")

func WordSepNormalizeFunc(f *pflag.FlagSet, name string) pflag.NormalizedName {
	if strings.Contains(name, "_") {
		return pflag.NormalizedName(strings.Replace(name, "_", "-", -1))
	}
	return pflag.NormalizedName(name)
}
func main() {

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.CommandLine.SetNormalizeFunc(WordSepNormalizeFunc)
	err := pflag.CommandLine.MarkDeprecated("addr", "please use --addr-name instead")
	if err != nil {
		panic(err)
	}
	pflag.Parse()
	if *help {
		pflag.Usage()
		return
	}

	fmt.Println("ip", *ip)
	fmt.Println("addr ", *addrName)
}
