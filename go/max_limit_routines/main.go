/*
 * Copyright 2018. bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Task time.Duration

func (t Task) Run() {
	fmt.Println("sleep ", time.Duration(t))
	time.Sleep(time.Duration(t))
}

func Work(task Task, token chan struct{}) {
	task.Run()
	<-token
}

// max routines
var MaxRoutines = 100

type Routines struct {
	tokens chan struct{}
}

func (r Routines) Call(task Task) {
	// if routines > MaxRoutine , will block
	r.tokens <- struct{}{}
	go Work(task, r.tokens)
}

func main() {
	routines := Routines{tokens: make(chan struct{}, MaxRoutines)}

	for i := 0; true; i++ {
		task := Task(time.Duration(rand.Intn(10)+1) * time.Second)
		routines.Call(task)
		fmt.Println("n ", i)
	}
}
