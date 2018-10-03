package handler

import (
	"context"
	"errors"
	"log"

	"github.com/emicklei/go-restful"
	"github.com/micro/go-micro/client"
	mutant "github.com/rodrigodmd/ml-mutant-srv/srv/mutant/proto/mutant"
	stats "github.com/rodrigodmd/ml-mutant-srv/srv/stats/proto/stats"
	"math"
)

type Dna struct{}

var (
	clMutant mutant.DnaService
	clStats  stats.DnaService
)

func init() {
	// setup mutant Server Client
	clMutant = mutant.NewDnaService("go.micro.srv.mutant", client.DefaultClient)
	clStats = stats.NewDnaService("go.micro.srv.stats", client.DefaultClient)
}

func (s *Dna) Stat(req *restful.Request, rsp *restful.Response) {
	log.Print("Received Stat API request")

	response, err := clStats.Stats(context.TODO(), &stats.Request{})
	if err != nil {
		log.Print(err)
		rsp.WriteError(500, err)
		return
	}
	log.Print(response)

	if math.IsNaN(float64(response.Ratio)) || response.CountHumanDna == 0 || response.CountMutantDna == 0 {
		response.Ratio = 0
	}

	statsResponse := StatsResponse{
		CountHumanDna:  response.CountHumanDna,
		CountMutantDna: response.CountMutantDna,
		Ratio:          response.Ratio,
	}
	rsp.WriteEntity(statsResponse)
}

func (s *Dna) Mutant(req *restful.Request, rsp *restful.Response) {
	log.Print("Received Mutant API request")

	dna := mutant.Request{}

	if err := req.ReadEntity(&dna); err != nil {
		log.Printf("Failed read entity: %v", err)
		rsp.WriteError(500, err)
		return
	}

	response, err := clMutant.IsMutant(context.TODO(), &dna)
	if err != nil {
		rsp.WriteError(500, err)
		return
	}

	if response == nil || !response.IsMutant {
		rsp.WriteError(403, errors.New("Forbidden"))
	} else {
		rsp.WriteEntity("Success")
	}

}
