/*
 * Copyright 2020 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package thread

import (
	"fmt"
	"runtime"
	"sync"
	"syscall"
	"testing"
)

//#include <unistd.h>
//void sleep_one() {
//sleep(1);
//}
import "C"

type ThreadManager struct {
	Number int
	queue  chan func()
	close  chan struct{}
	once   sync.Once
}

func NewThreadManager(num int) *ThreadManager {
	_close := make(chan struct{})
	queue := make(chan func())
	for i := 0; i < num; i++ {
		go func() {
			runtime.LockOSThread()
			defer runtime.UnlockOSThread()
			for {
				select {
				case <-_close:
					return
				case task := <-queue:
					task()
				}
			}
		}()
	}
	return &ThreadManager{
		Number: num,
		queue:  queue,
		close:  _close,
	}
}

func (m *ThreadManager) Close() {
	m.once.Do(func() {
		close(m.close)
	})
}

func (m *ThreadManager) SyncCall(f func()) {
	wg := sync.WaitGroup{}
	wg.Add(1)
	m.queue <- func() {
		defer wg.Done()
		f()
	}
	wg.Wait()
}

func goThreadBasic(doSth func(callback func())) {
	threadMap := map[int]struct{}{}
	threadLock := sync.Mutex{}
	procs := runtime.GOMAXPROCS(0)
	fmt.Println("proc", procs, "cpu", runtime.NumCPU())
	wg := sync.WaitGroup{}

	for i := 0; i < 40; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			doSth(func() {
				tid := syscall.Gettid()
				fmt.Println("tid", tid)
				threadLock.Lock()
				threadMap[tid] = struct{}{}
				fmt.Println("len", len(threadMap))
				threadLock.Unlock()
			})
		}()
	}
	wg.Wait()
	fmt.Printf("current os threads %d\n", runtime.GOMAXPROCS(0))
}

func testCgoThread(t *testing.T) {
	goThreadBasic(func(callback func()) {
		v := 0
		for i := 0; i < 100; i++ {
			v++
		}
		C.sleep_one()
		callback()
	})

}

var m = NewThreadManager(runtime.NumCPU())

func testCgoThreadManager(t *testing.T) {
	goThreadBasic(func(callback func()) {
		m.SyncCall(func() {
			v := 0
			for i := 0; i < 100; i++ {
				v++
			}
			C.sleep_one()
			callback()
		})
	})
}
