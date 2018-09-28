package handler

import (
	"context"
	"strconv"

	"github.com/micro/go-log"

	mlmutant "github.com/rodrigodmd/ml-mutant"
	mutant "github.com/rodrigodmd/ml-mutant-srv/srv/mutant/proto/mutant"
	"github.com/rodrigodmd/ml-mutant-srv/srv/mutant/publisher"
)

type Mutant struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Mutant) IsMutant(ctx context.Context, req *mutant.Request, rsp *mutant.Response) error {
	log.Log("Received IsMutant request")
	log.Log(req.Dna)
	isMutant, err := mlmutant.IsMutant(req.Dna)

	if err != nil {
		log.Log(err)
		rsp.IsMutant = false
		rsp.Msg = err.Error()
		return err
	}

	go publisher.SendDna(&req.Dna, isMutant)
	rsp.IsMutant = isMutant
	rsp.Msg = "Resuls: " + strconv.FormatBool(isMutant)

	return nil
}
