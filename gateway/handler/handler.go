package handler

import (
	"context"
	"errors"
	"log"

	restful "github.com/emicklei/go-restful"
	mutant "github.com/rodrigodmd/ml-mutant-srv/mutant/proto/mutant"
	"github.com/micro/go-micro/client"
)

type Dna struct{}


var cl mutant.DnaService

func init(){
	// setup mutant Server Client
	cl = mutant.NewDnaService("go.micro.srv.mutant", client.DefaultClient)
}

func (s *Dna) Stat(req *restful.Request, rsp *restful.Response) {
	log.Print("Received Dna.Anything API request")
	rsp.WriteEntity(map[string]string{
		"message": "Hi, this is the Greeter API",
	})
}

func (s *Dna) IsMutant(req *restful.Request, rsp *restful.Response) {
	log.Print("Received Dna.Hello API request")

	dna := mutant.Request{}

	if err := req.ReadEntity(&dna); err != nil {
		log.Print("Failed read entity: %v", err)
		rsp.WriteError(500, err)
		return
	}

	response, err := cl.IsMutant(context.TODO(), &dna)
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
