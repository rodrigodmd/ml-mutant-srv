package main

import (
	"log"
	"strconv"
	"time"

	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/broker/kafka"
	"github.com/rodrigodmd/ml-mutant-srv/srv/mutant/handler"
	mutant "github.com/rodrigodmd/ml-mutant-srv/srv/mutant/proto/mutant"
	"github.com/rodrigodmd/ml-mutant-srv/srv/mutant/publisher"
)

func main() {
	waitSeconds := 0
	b := kafka.NewBroker()

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.mutant"),
		micro.Version("latest"),
		micro.Broker(b),
		micro.Flags(
			cli.StringFlag{
				Name:   "start_wait_time",
				EnvVar: "START_WAIT_TIME",
				Usage:  "Wait time to start microservice in seconds",
				Value:  "0",
			},
		),
		micro.Action(func(c *cli.Context) {
			waitSeconds, _ = strconv.Atoi(c.String("start_wait_time"))
		}),
	)

	// Initialise service
	service.Init()

	// Register Handler
	mutant.RegisterDnaHandler(service.Server(), new(handler.Mutant))

	// create publisher
	publisher.NewPublisher(service.Client())

	// Wait befor we start the service
	time.Sleep(time.Second * time.Duration(waitSeconds))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}
