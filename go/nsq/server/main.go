package main

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"log"
	"math/rand"
	"time"
)

func NewConsumer(topic, channel string, n int) (*nsq.Consumer, error) {
	q, err := nsq.NewConsumer(topic, channel, nsq.NewConfig())
	if err != nil {
		log.Panic("consumer error", err)
		return nil, err
	}
	q.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		message.Touch()
		time.Sleep(time.Millisecond * (time.Duration(rand.Intn(1000) + 1000)))
		log.Printf("channel %s process %d Got a message: %v\n", channel, n, string(message.Body))
		message.Finish()
		log.Println("has responded", message.HasResponded())
		return nil
	}))
	return q, nil
}

func main() {
	var qList []*nsq.Consumer
	for i := 0; i < 3; i++ {
		q, err := NewConsumer("write_test", "ch"+fmt.Sprint(i), 0)
		if err != nil {
			log.Panic(err)
		}
		err = q.ConnectToNSQD("127.0.0.1:4150")
		if err != nil {
			log.Panic("Could not connect", err.Error())
		}
		qList = append(qList, q)
	}

	for _, q := range qList {
		<-q.StopChan
	}
}
