package main

import (
	"fmt"
	"log"

	//	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	// To enable rabbitmq plugin uncomment
	//_ "github.com/micro/go-plugins/broker/rabbitmq"
	// To enable googlepubsub plugin uncomment
	//_ "github.com/micro/go-plugins/broker/googlepubsub"
	// To enable kafka plugin uncomment
	//_ "github.com/micro/go-plugins/broker/kafka"
	//"github.com/micro/go-plugins/broker/kafka"
)

var (
	topic = "go.micro.topic.foo"
)

// Example of a shared subscription which receives a subset of messages
func sharedSub() {
	_, err := broker.Subscribe(topic, func(p broker.Publication) error {
		fmt.Println("[sub] received message:", string(p.Message().Body), "header", p.Message().Header)
		return nil
	}, broker.Queue("consumer"))
	if err != nil {
		fmt.Println(err)
	}
}

func main() {

	//////////////////////////////
	//cmd.Init()

	// b := kafka.NewBroker()
	//////////////////////////////
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.stats"),
		micro.Version("latest"),
		// micro.Broker(b),
	)

	// Initialise service
	service.Init()

	// Register Handler
	//example.RegisterExampleHandler(service.Server(), new(handler.Example))

	// Register Struct as Subscriber
	//micro.RegisterSubscriber("go.micro.srv.stats", service.Server(), new(subscriber.Example))

	// Register Function as Subscriber
	//micro.RegisterSubscriber("go.micro.srv.stats", service.Server(), subscriber.Handler)

	if err := broker.Init(); err != nil {
		log.Fatalf("Broker Init error: %v", err)
	}
	if err := broker.Connect(); err != nil {
		log.Fatalf("Broker Connect error: %v", err)
	}

	sharedSub()

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
