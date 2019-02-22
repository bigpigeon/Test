//
package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/client"
	"os"
)

func main() {
	c, err := client.New(
		client.Config{
			Endpoints: []string{
				"http://localhost:2379",
			},
		})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	v, err := c.GetVersion(context.Background())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(v.Server, v.Cluster)
	fmt.Println("ok")
}
