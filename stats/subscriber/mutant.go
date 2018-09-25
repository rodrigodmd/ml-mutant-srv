package subscriber

import (
	"context"
	"log"

	"github.com/rodrigodmd/ml-mutant-srv/stats/db"
	stats "github.com/rodrigodmd/ml-mutant-srv/stats/proto/stats"
)

func Handler(ctx context.Context, msg *stats.Message) error {
	log.Print("Adding DNA to DB")
	return db.AddDna(&msg.Dna, msg.IsMutant)
}
