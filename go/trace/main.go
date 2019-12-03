/*
 * Copyright 2019 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"runtime"
	"runtime/trace"
	"sync"
	"time"
)

type stack []uintptr

func caller() *stack {
	const depth = 32
	var pcs [depth]uintptr
	n := runtime.Callers(1, pcs[:])
	var st stack = pcs[0:n]
	return &st
}

func printFrames() {
	st := caller()
	frames := runtime.CallersFrames(*st)
	for {
		f, more := frames.Next()
		if more == false {
			break
		}
		fmt.Printf("%s:%d \n in func %s \n", f.File, f.Line, f.Function)
		//if f.Func != nil {
		//	file, line := f.Func.FileLine(f.PC - 1)
		//	fmt.Printf("  in func %s line %d file %s\n", f.Func.Name(), line, file)
		//}
	}
}

func printTrace() {

	// create new channel of type int
	ch := make(chan int)

	// start new anonymous goroutine
	go func() {
		// send 42 to channel
		ch <- 42
	}()
	// read fro
}

func StartTrace(name string) func() {
	buff := bytes.Buffer{}
	err := trace.Start(&buff)
	if err != nil {
		panic(err)
	}
	return func() {
		trace.Stop()
		err := ioutil.WriteFile(name, buff.Bytes(), 0644)
		if err != nil {
			panic(err)
		}
	}
}

func traceAfterSleep() {
	time.Sleep(time.Second)

	tra := StartTrace("traceAfterSleep.out")
	defer tra()
	printTrace()
}

func traceInRoutine() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		tra := StartTrace("traceInRoutine.out")
		defer tra()
		fmt.Println("do sth")
	}()
	time.Sleep(time.Second)
	wg.Wait()
}

func main() {
	printFrames()
	traceAfterSleep()
	traceInRoutine()
}
