package main

import (
	"fmt"
	"log"
	"time"

	//"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"

	"github.com/rodrigodmd/ml-mutant-srv/mutant/handler"
	"github.com/rodrigodmd/ml-mutant-srv/mutant/subscriber"

	mutant "github.com/rodrigodmd/ml-mutant-srv/mutant/proto/mutant"
	// To enable rabbitmq plugin uncomment
	//_ "github.com/micro/go-plugins/broker/rabbitmq"
	// To enable googlepubsub plugin uncomment
	//_ "github.com/micro/go-plugins/broker/googlepubsub"
	// To enable kafka plugin uncomment
)

var (
	topic = "go.micro.topic.foo"
)

func pub() {
	tick := time.NewTicker(100*time.Microsecond)
	i := 0
	for _ = range tick.C {
		msg := &broker.Message{
			Header: map[string]string{
				"id": fmt.Sprintf("%d", i),
			},
			Body: []byte(fmt.Sprintf("%d: %s", i, time.Now().String())),
		}
		if err := broker.Publish(topic, msg); err != nil {
			log.Printf("[pub] failed: %v", err)
		} else {
			fmt.Println("[pub] pubbed message:", string(msg.Body))
		}
		i++
	}
}

func main() {
	////////////////////////////////
	//cmd.Init()

	// b := kafka.NewBroker()

	/////////////////////////
	// New Service

	service := micro.NewService(
		micro.Name("go.micro.srv.mutant"),
		micro.Version("latest"),
		// micro.Broker(b),
	)

	// Initialise service
	service.Init()

	// Register Handler
	mutant.RegisterDnaHandler(service.Server(), new(handler.Mutant))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.mutant", service.Server(), new(subscriber.Mutant))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.mutant", service.Server(), subscriber.Handler)

	///////////////////////////////////////////////
	// if err := broker.Init(); err != nil {
	// 	log.Fatalf("Broker Init error: %v", err)
	// }

	// if err := broker.Connect(); err != nil {
	// 	log.Fatalf("Broker Connect error: %v", err)
	// }

	go pub()

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}
