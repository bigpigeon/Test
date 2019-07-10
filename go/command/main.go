package main

import (
	"fmt"
	"os/exec"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			cmd := exec.Command("/bin/sleep", "5000")
			err := cmd.Start()
			if err != nil {
				panic(err)
			}
			fmt.Println("11")
		}()
	}
	wg.Wait()
}
