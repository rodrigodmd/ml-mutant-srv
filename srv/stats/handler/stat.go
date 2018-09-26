package handler

import (
	"context"
	"log"

	"github.com/rodrigodmd/ml-mutant-srv/srv/stats/db"

	stats "github.com/rodrigodmd/ml-mutant-srv/srv/stats/proto/stats"
)

type Dna struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Dna) Stats(ctx context.Context, req *stats.Request, rsp *stats.Response) error {
	log.Print("Received Stats request")
	st, _ := db.GetDnaStat()
	rsp.CountHumanDna = st.CountHuman
	rsp.CountMutantDna = st.CountMutant
	rsp.Ratio = st.Ratio
	return nil
}
