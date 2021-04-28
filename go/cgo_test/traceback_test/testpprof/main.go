/*
 * Copyright 2021 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package main

/*
extern void add100WTimes(void);
extern void alloc10MHeap(void);
*/
import "C"

import (
	_ "github.com/ianlancetaylor/cgosymbolizer"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"runtime/pprof"
	"strconv"
	"time"
)

func procCpu() {
	tmpFile, err := os.CreateTemp(os.TempDir(), "")
	if err != nil {
		panic(err)
	}
	defer tmpFile.Close()
	log.Println("cpu pprof file", tmpFile.Name())
	if err := pprof.StartCPUProfile(tmpFile); err != nil {
		log.Fatal("can't start CPU profile: ", err)
	}
	go cpuHogCFunction()
	runtime.Gosched()
	time.Sleep(100 * time.Millisecond)

	pprof.StopCPUProfile()
}

func procMemory() {
	tmpFile, err := os.CreateTemp(os.TempDir(), "")
	if err != nil {
		panic(err)
	}
	defer tmpFile.Close()
	log.Println("heap pprof file", tmpFile.Name())
	go memoryFunction()
	time.Sleep(1 * time.Second)
	runtime.Gosched()
	if err := pprof.WriteHeapProfile(tmpFile); err != nil {
		log.Fatal("can't alloc heap profile: ", err)
	}
	pid := os.Getpid()
	cmd := exec.Command("pmap", "-x", strconv.Itoa(pid))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func main() {
	procMemory()
}

func memoryFunction() {
	C.alloc10MHeap()
	data := make([]byte, 10<<20)
	rand.Read(data)
	for {
		time.Sleep(1 * time.Second)
	}
}

func cpuHogCFunction() {
	// Generates CPU profile samples including a Cgo call path.
	for {
		C.add100WTimes()
		runtime.Gosched()
	}
}
