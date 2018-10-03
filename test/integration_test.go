package test

import (
	"flag"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type testData struct {
	testTitle   string
	dna         []string
	deltaMutant int64
	deltaHuman  int64
}

var testTable = []testData{
	testData{
		testTitle:   "Null Dna",
		dna:         nil,
		deltaMutant: 0,
		deltaHuman:  0,
	},
	testData{
		testTitle:   "Mutant Dna",
		dna:         []string{"ATGCAA", "CAGTTT", "TTATTT", "AGAAGG", "CCACTA", "TCACTG"},
		deltaMutant: 1,
		deltaHuman:  0,
	},
	testData{
		testTitle:   "Human Dna",
		dna:         []string{"ATGC", "CAGT", "TATT", "AGAA"},
		deltaMutant: 0,
		deltaHuman:  1,
	},
	testData{
		testTitle:   "Invalid Dna Leters",
		dna:         []string{"ATGCAA", "invalid", "TTATTT", "AGAAGG", "CCACTA", "TCACTG"},
		deltaMutant: 0,
		deltaHuman:  0,
	},
	testData{
		testTitle:   "Invalid Dna length",
		dna:         []string{"ATGCAA", "CAGT", "TTATTT", "AGAAGG", "CCACTA", "TCACTG"},
		deltaMutant: 0,
		deltaHuman:  0,
	},
}

func TestMain(m *testing.M) {
	fmt.Printf("Test Main \n")
	envBaseUrl := os.Getenv("BASE_URL")
	if envBaseUrl != "" {
		baseURL = envBaseUrl
	}

	dna := []string{"ATGC", "CAGT", "TATT", "AGAA"}
	if err := postDna(&DnaModel{Dna: dna}); err != nil {
		fmt.Printf("ERROR: main - " + err.Error())
	}

	// Wait for dna to propagate
	// first time take longer to creat the table
	time.Sleep(5 * time.Second)

	flag.Parse()
	exitCode := m.Run()

	// Exit
	os.Exit(exitCode)
}

func TestDnaTable(t *testing.T) {
	for _, td := range testTable {
		fmt.Printf("Testing %s \n", td.testTitle)

		// Get initial stats
		stats, err := getStat()
		if err != nil {
			t.Error(err)
		}

		// post dna
		if err = postDna(&DnaModel{Dna: td.dna}); err != nil {
			t.Errorf("ERROR: " + td.testTitle + " - " + err.Error())
		}

		// Wait for dna to propagate
		time.Sleep(time.Second)

		// Get new stats
		newStats, err := getStat()
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, td.deltaHuman, newStats.CountHuman-stats.CountHuman, "Wrong Human count delta")
		assert.Equal(t, td.deltaMutant, newStats.CountMutant-stats.CountMutant, "Wrong Mutant count delta")
	}
}
