package subscriber

import (
	"context"

	"github.com/micro/go-log"

	stats "github.com/rodrigodmd/ml-mutant-srv/stats/proto/stats"
)

type Mutant struct{}

func (e *Mutant) Handle(ctx context.Context, msg *stats.Message) error {
	log.Log("Function Received message: ", msg.Id)
	log.Log("DNA: ", msg.Dna)
	log.Log("Is Mutant: ", msg.IsMutant)
	return nil
}

func Handler(ctx context.Context, msg *stats.Message) error {
	log.Log("Function Received message: ", msg.Id)
	log.Log("DNA: ", msg.Dna)
	log.Log("Is Mutant: ", msg.IsMutant)
	return nil
}
