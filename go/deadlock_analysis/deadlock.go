/*
 * Copyright 2019 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package main

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os/exec"
)

func ExecCmd(appCmd string) (string, string, error) {
	var outStr, errStr string
	var content []byte
	var stdout, stderr io.ReadCloser
	var err error

	// prepare stdout/stderr
	cmd := exec.Command("bash", appCmd)
	stdout, err = cmd.StdoutPipe()
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		return outStr, errStr, err
	}
	defer stdout.Close()

	stderr, err = cmd.StderrPipe()
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		return outStr, errStr, err
	}
	defer stderr.Close()

	// launch process
	if err := cmd.Start(); err != nil {
		fmt.Printf("ERROR: %v\n", err)
		return outStr, errStr, err
	}

	// read stdout firstly
	content, err = ioutil.ReadAll(stdout)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		return outStr, errStr, err
	}
	outStr = string(content)

	// read stderr secondly
	content, err = ioutil.ReadAll(stderr)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		return outStr, errStr, err
	}
	errStr = string(content)

	// wait process
	if err = cmd.Wait(); err != nil {
		fmt.Printf("ERROR: %v\n", err)
		return outStr, errStr, err
	}

	if cmd.ProcessState.Success() {
		return outStr, errStr, nil
	} else {
		err = errors.New("Command Failed")
		fmt.Printf("ERROR: %v\n", err)
		return outStr, errStr, err
	}
}

func main() {
	stdout, stderr, err := ExecCmd("test.sh")
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}

	fmt.Printf("STDOUT: %s\n", stdout)
	fmt.Printf("STDERR: %s\n", stderr)
}
