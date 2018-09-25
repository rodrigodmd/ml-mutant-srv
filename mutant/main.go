package main

import (
	"log"
	//"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"github.com/rodrigodmd/ml-mutant-srv/mutant/handler"

	mutant "github.com/rodrigodmd/ml-mutant-srv/mutant/proto/mutant"
	// To enable rabbitmq plugin uncomment
	//_ "github.com/micro/go-plugins/broker/rabbitmq"
	// To enable googlepubsub plugin uncomment
	//_ "github.com/micro/go-plugins/broker/googlepubsub"
	// To enable kafka plugin uncomment
	"github.com/rodrigodmd/ml-mutant-srv/mutant/publisher"
)

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

	// create publisher
	publisher.NewPublisher(service.Client())

	// Register Struct as Subscriber
	//micro.RegisterSubscriber("go.micro.srv.mutant", service.Server(), new(subscriber.Mutant))

	// Register Function as Subscriber
	//micro.RegisterSubscriber("go.micro.srv.mutant", service.Server(), subscriber.Handler)

	///////////////////////////////////////////////
	// if err := broker.Init(); err != nil {
	// 	log.Fatalf("Broker Init error: %v", err)
	// }

	// if err := broker.Connect(); err != nil {
	// 	log.Fatalf("Broker Connect error: %v", err)
	// }

	// go pub()

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}
