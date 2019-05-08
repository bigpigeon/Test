package main

import (
	"github.com/nsqio/go-nsq"
	"log"
)

func main() {

	config := nsq.NewConfig()
	q, err := nsq.NewConsumer("write_test", "ch", config)
	if err != nil {
		log.Panic("consumer error")
	}
	q.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		log.Printf("Got a message: %v\n", string(message.Body))
		return nil
	}))
	err = q.ConnectToNSQD("127.0.0.1:4150")
	if err != nil {
		log.Panic("Could not connect")
	}
	<-q.StopChan

}
