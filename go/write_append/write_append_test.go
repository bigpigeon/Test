/*
 * Copyright 2020 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package write_append

import (
	"bytes"
	"github.com/stretchr/testify/require"
	"golang.org/x/sys/unix"
	"io/ioutil"
	"os"
	"sync"
	"testing"
)

func TestWrite4096(t *testing.T) {

	for _, onceWrite := range []int{4096, 4097, 4099} {
		t.Log("start ", onceWrite)
		const tmpFile = "/tmp/out.tmp"
		const processes = 20
		const writeNum = 100
		os.Remove(tmpFile)
		waitReady := make(chan struct{})
		wg := sync.WaitGroup{}
		wg.Add(processes)
		for i := 0; i < processes; i++ {

			go func(c byte) {
				defer wg.Done()
				f, err := unix.Open(tmpFile, unix.O_APPEND|unix.O_WRONLY|unix.O_CREAT, 0644)
				require.NoError(t, err)
				defer func() {
					unix.Close(f)
				}()
				s := bytes.Repeat([]byte{c}, onceWrite-1)
				s = append(s, '\n')
				<-waitReady
				for i := 0; i < writeNum; i++ {
					n, err := unix.Write(f, s)
					if err != nil {
						panic(err)
					}
					if n != onceWrite {
						panic("not full write")
					}
				}
			}('A' + byte(i))
		}
		close(waitReady)
		wg.Wait()

		{
			data, err := ioutil.ReadFile(tmpFile)
			require.NoError(t, err)
			diffNum := 0
			for i := 0; i < len(data); i += onceWrite {
				for j := i + 1; j < i+onceWrite-1; j++ {
					if data[j] != data[i] {
						diffNum++
						break
					}
				}
			}
			t.Log("diff ", diffNum)
		}
	}
}
