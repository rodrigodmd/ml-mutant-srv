package main

import (
	"strconv"
	"strings"
	"time"

	"github.com/micro/cli"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
	"github.com/micro/go-plugins/broker/kafka"
	"github.com/rodrigodmd/ml-mutant-srv/srv/stats/elastic"
	"github.com/rodrigodmd/ml-mutant-srv/srv/stats/handler"
	stats "github.com/rodrigodmd/ml-mutant-srv/srv/stats/proto/stats"
	"github.com/rodrigodmd/ml-mutant-srv/srv/stats/subscriber"
)

func main() {
	waitSeconds := 0
	b := kafka.NewBroker()

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.stats"),
		micro.Version("latest"),
		micro.Broker(b),
		micro.Flags(
			cli.StringFlag{
				Name:   "elasticsearch_hosts",
				EnvVar: "ELASTICSEARCH_HOSTS",
				Usage:  "Comma separated list of elasticsearch hosts",
				Value:  "localhost:9200",
			},
			cli.StringFlag{
				Name:   "start_wait_time",
				EnvVar: "START_WAIT_TIME",
				Usage:  "Wait time to start microservice in seconds",
				Value:  "0",
			},
		),
		micro.Action(func(c *cli.Context) {
			waitSeconds, _ = strconv.Atoi(c.String("start_wait_time"))
			parts := strings.Split(c.String("elasticsearch_hosts"), ",")
			elastic.Hosts = parts
		}),
	)

	// Initialise service
	service.Init()

	// Initialise elasticsearch
	elastic.Init()

	// Register Handler
	stats.RegisterDnaHandler(service.Server(), new(handler.Dna))

	// register subscriber with queue, each message is delivered to a unique subscriber
	micro.RegisterSubscriber("go.micro.topic.mutant", service.Server(), subscriber.Handler, server.SubscriberQueue("queue.pubsub"))

	// Wait befor we start the service
	time.Sleep(time.Second * time.Duration(waitSeconds))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
