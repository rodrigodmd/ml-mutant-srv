package subscriber

import (
	"context"

	"github.com/micro/go-log"

	mutant "github.com/rodrigodmd/ml-mutant-srv/mutant/proto/mutant"
)

type Mutant struct{}

func (e *Mutant) Handle(ctx context.Context, msg *mutant.Message) error {
	log.Log("Function Received message: ", msg.Id)
	log.Log("DNA: ", msg.Dna)
	log.Log("Is Mutant: ", msg.IsMutant)
	return nil
}

func Handler(ctx context.Context, msg *mutant.Message) error {
	log.Log("Function Received message: ", msg.Id)
	log.Log("DNA: ", msg.Dna)
	log.Log("Is Mutant: ", msg.IsMutant)
	return nil
}
