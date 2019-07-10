package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

var globalStr = "globalStr"

var globalData = TestData{
	Data: "globalData",
}

var globalPointData = &TestData{
	Data: "globalPointData",
}

type TestData struct {
	Data string
}

func test() {
	// global string test
	runtime.SetFinalizer(&globalStr, func(d *string) {
		fmt.Printf("%s dead\n", *d)
	})
	// local string test
	localStr := "localStr"
	runtime.SetFinalizer(&localStr, func(d *string) {
		fmt.Printf("%s dead\n", *d)
	})
	localSleepStr := "localSleepStr"
	runtime.SetFinalizer(&localSleepStr, func(d *string) {
		fmt.Printf("%s want dead\n", *d)
		time.Sleep(1 * time.Second)
		fmt.Printf("%s dead\n", *d)
	})
	// global
	runtime.SetFinalizer(&globalData, func(d *TestData) {
		fmt.Printf("%s dead\n", d.Data)
	})
	runtime.SetFinalizer(globalPointData, func(d *TestData) {
		fmt.Printf("%s dead\n", d.Data)
	})
}

func test2Data() *TestData {
	data := TestData{
		Data: "test 2 TestData" + fmt.Sprint(rand.Int()),
	}
	return &data
}

func test2() {
	data := test2Data()

	runtime.SetFinalizer(data, func(d *TestData) {
		fmt.Printf("%s dead\n", d.Data)
	})
}

func main() {
	test()
	runtime.GC()
	// local
	time.Sleep(1 * time.Millisecond)
	//test2()
	//runtime.GC()
	//// local
	//time.Sleep(1 * time.Millisecond)
}
