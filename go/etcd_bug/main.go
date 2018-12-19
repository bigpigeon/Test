package main

import (
	"fmt"
	"go.etcd.io/etcd/client"
	"os"
)

func main() {
	_, err := client.New(
		client.Config{
			Endpoints: []string{"http://localhost:2379"},
		})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
