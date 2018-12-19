/*
 * Copyright 2018 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"syscall"
)

func main() {
	var rLimit syscall.Rlimit
	err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit)
	if err != nil {
		fmt.Println("Error Getting Rlimit ", err)
	}

	fmt.Println(rLimit)
	rLimit.Max = 100
	rLimit.Cur = 100
	err = syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit)
	if err != nil {
		fmt.Println("Error Setting Rlimit ", err)
	}
	err = syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit)
	if err != nil {
		fmt.Println("Error Getting Rlimit ", err)
	}
	fmt.Println("Rlimit Final", rLimit)
	var fs []*os.File
	for i := 0; i < 50; i++ {
		f, err := os.Open("/dev/null")
		if err != nil {
			panic(err)
		}
		fs = append(fs, f)
	}

	fmt.Println("put any key to done")
	reader := bufio.NewReader(os.Stdin)
	reader.ReadByte()
}
