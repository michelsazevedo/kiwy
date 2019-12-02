package request

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	baseUrl = "https://us-central1-rdsm-analytics-development.cloudfunctions.net/random"
)

var (
	fn   = 1
	pool = 1
)

type Result struct {
	DateTime string `json:"DateTime"`
	Time     int    `json:"Time"`
	Count    int    `json:"Count"`
}

type CollectionResult struct {
	Result []Result
}

func MakeParallelsRequests(numOfRequests int, ch chan CollectionResult) {
	defer close(ch)
	var results = []chan CollectionResult{}

	for i := 0; i < numOfRequests; i++ {
		results = append(results, make(chan CollectionResult))
		go MakeRequest(results[i])
	}

	for i := range results {
		for result := range results[i] {
			ch <- result
		}
	}
}

func MakeRequest(ch chan CollectionResult) {
	defer close(ch)
	res, err := http.Get(GetUrl())
	if err != nil {
		ch <- CollectionResult{}
	}

	defer res.Body.Close()

	jsonData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		ch <- CollectionResult{}
	}

	var collection CollectionResult

	err = json.Unmarshal([]byte(jsonData), &collection)
	if err != nil {
		ch <- CollectionResult{}
	}

	ch <- collection
}

func GetUrl() string {
	if pool >= 100 {
		pool = 1
		fn += 1
	}

	pool += 1

	log.Println("[URL]:", baseUrl+fn)
	return baseUrl + fn
}
