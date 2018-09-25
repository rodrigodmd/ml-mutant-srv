package publisher

import (
	"context"
	"time"

	"github.com/go-log/log"
	"github.com/google/uuid"
	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	mutant "github.com/rodrigodmd/ml-mutant-srv/mutant/proto/mutant"
)

var (
	topic = "go.micro.topic.mutant"
	pub   micro.Publisher
)

func NewPublisher(cl client.Client) {
	// create publisher
	pub = micro.NewPublisher(topic, cl)
}

func SendDna(dna *[]string, isMutant bool) {
	uuid, _ := uuid.NewUUID()

	ev := &mutant.Message{
		Id:        uuid.String(),
		Timestamp: time.Now().Unix(),
		Dna: *dna,
		IsMutant: isMutant,
	}

	log.Logf("publishing %+v\n", ev)

	// publish an event
	if err := pub.Publish(context.Background(), ev); err != nil {
		log.Logf("error publishing %v", err)
	}

}
