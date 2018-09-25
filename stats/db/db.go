package db

import (
	"encoding/json"
	"log"

	"github.com/rodrigodmd/ml-mutant-srv/stats/elastic"
	proto "github.com/rodrigodmd/ml-mutant-srv/stats/proto/elastic"
)

const (
	DB_INDEX        = "dna"
	DNA_TYPE_NORMAL = "normal"
	DNA_TYPE_MUTANT = "mutant"
)

type StatsRsult struct {
	CountMutant int64
	CountHuman  int64
	Ratio       float32
}

func AddDna(dna *[]string, isMutant bool) error {
	data, _ := json.Marshal(dna)
	t := DNA_TYPE_NORMAL
	if isMutant {
		t = DNA_TYPE_MUTANT
	}
	req := proto.CreateRequest{
		Index: DB_INDEX,
		Type:  t,
		Data:  "{\"dna\":" + string(data) + "}",
	}
	if err := elastic.Create(&req); err != nil {
		log.Print(err)
		return err
	}

	GetDnaStat()
	return nil
}

func GetDnaStat() (*StatsRsult, error) {
	countMutant, _ := elastic.Count(&proto.CountRequest{
		Index: DB_INDEX,
		Type:  DNA_TYPE_MUTANT,
	})
	countNormal, _ := elastic.Count(&proto.CountRequest{
		Index: DB_INDEX,
		Type:  DNA_TYPE_NORMAL,
	})
	log.Print("Total so far Mutant and Normal ", countMutant, countNormal)

	return &StatsRsult{
		CountMutant: int64(countMutant),
		CountHuman:  int64(countNormal),
		Ratio:       float32(countMutant) / float32(countNormal),
	}, nil
}
