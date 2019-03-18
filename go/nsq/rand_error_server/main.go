package main

import (
	"errors"
	"github.com/nsqio/go-nsq"
	"log"
	"math/rand"
)

func main() {

	config := nsq.NewConfig()
	q, _ := nsq.NewConsumer("write_test", "ch", config)
	q.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		log.Printf("Got a message: id %s %v\n", message.ID, string(message.Body))
		randVal := rand.Intn(2)
		log.Println("random val", randVal)
		return errors.New("rand error")
		return nil
	}))
	err := q.ConnectToNSQD("127.0.0.1:4150")
	if err != nil {
		log.Panic("Could not connect")
	}
	<-q.StopChan

}
