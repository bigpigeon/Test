/*
 * Copyright 2020 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package exec

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"os/exec"
	"testing"
	"time"
)

func TestStream(t *testing.T) {

	data := make([]byte, 100<<20)
	rand.Read(data)
	func() {
		defer func(_time time.Time) { fmt.Println("use time ", time.Now().Sub(_time)) }(time.Now())
		cmd := exec.Command("cat")
		cmd.Stdin = bytes.NewReader(data)
		cmd.Stdout = nil
		cmd.Run()
	}()
	func() {

		defer func(_time time.Time) { fmt.Println("use time ", time.Now().Sub(_time)) }(time.Now())
		cmd := exec.Command("cat")
		writer, err := cmd.StdinPipe()
		if err != nil {
			panic(err)
		}
		err = cmd.Start()
		if err != nil {
			panic(err)
		}
		_, err = writer.Write(data)
		if err != nil {
			panic(err)
		}
		writer.Close()
		//cmd.Stdin = bytes.NewReader(data)
		cmd.Stdout = nil
		cmd.Wait()
	}()
}
