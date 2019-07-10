package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func signPrint(i int) {
	sigs := make(chan os.Signal, 5)
	fmt.Println("signPrint", i)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	sig := <-sigs
	fmt.Println("catch signal", sig, i)
	signal.Reset(syscall.SIGINT, syscall.SIGTERM)

	pid := syscall.Getpid()
	syscall.Kill(pid, syscall.SIGINT)
}

func main() {
	go signPrint(1)
	//go signPrint(2)

	for {
		time.Sleep(1 * time.Second)
	}
}
