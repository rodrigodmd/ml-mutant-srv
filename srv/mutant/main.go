package main

import (
	"log"
	//"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"github.com/rodrigodmd/ml-mutant-srv/srv/mutant/handler"

	mutant "github.com/rodrigodmd/ml-mutant-srv/srv/mutant/proto/mutant"
	// To enable rabbitmq plugin uncomment
	//_ "github.com/micro/go-plugins/broker/rabbitmq"
	// To enable googlepubsub plugin uncomment
	//_ "github.com/micro/go-plugins/broker/googlepubsub"
	// To enable kafka plugin uncomment
	"github.com/rodrigodmd/ml-mutant-srv/srv/mutant/publisher"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.mutant"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	mutant.RegisterDnaHandler(service.Server(), new(handler.Mutant))

	// create publisher
	publisher.NewPublisher(service.Client())

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}
