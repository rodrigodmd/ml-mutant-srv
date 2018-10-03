package test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

var baseURL = "http://localhost:8080/api/"

// Get stats request
func getStat() (*StatModel, error) {
	stat := StatModel{}
	if err := doRequest("GET", "stats", nil, &stat); err != nil {
		return nil, err
	}
	log.Printf("Human count %v", stat.CountHuman)
	log.Printf("Mutant count %v", stat.CountMutant)
	return &stat, nil
}

// Post dna request
func postDna(dna *DnaModel) error {
	if err := doRequest("POST", "mutant", dna, nil); err != nil {
		return err
	}
	return nil
}

// Generic request
func doRequest(method, apiUrl string, request, reponse interface{}) error {
	requBody, err := json.Marshal(request)

	req, err := http.NewRequest(method, baseURL+apiUrl, bytes.NewReader(requBody))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if reponse == nil {
		return nil
	}

	if err = json.Unmarshal(body, &reponse); err != nil {
		return err
	}
	return nil
}
