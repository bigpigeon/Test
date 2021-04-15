/*
 * Copyright 2021 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package main

/*
extern void add100WTimes(void);
*/
import "C"

import (
	_ "github.com/ianlancetaylor/cgosymbolizer"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

func main() {
	tmpFile, err := os.CreateTemp(os.TempDir(), "")
	if err != nil {
		panic(err)
	}
	defer tmpFile.Close()
	log.Println("pprof file", tmpFile.Name())
	if err := pprof.StartCPUProfile(tmpFile); err != nil {
		log.Fatal("can't start CPU profile: ", err)
	}
	go cpuHogCFunction()
	runtime.Gosched()
	time.Sleep(100 * time.Millisecond)

	pprof.StopCPUProfile()

}

func cpuHogCFunction() {
	// Generates CPU profile samples including a Cgo call path.
	for {
		C.add100WTimes()
		runtime.Gosched()
	}
}
