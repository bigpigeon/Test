package main

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"log"
	"math/rand"
	"time"
)

func main() {
	config := nsq.NewConfig()

	w, _ := nsq.NewProducer("127.0.0.1:4150", config)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100; i++ {
		err := w.Publish("write_test", []byte(fmt.Sprint(rand.Int())))
		if err != nil {
			log.Panic("Could not connect")
		}
		time.Sleep(1 * time.Second)
	}

	w.Stop()

}
