package request

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	url = "https://us-east4-rdsm-analytics-development.cloudfunctions.net/Hello?instance=mushin-analytics&database=automation-roi&concurrency=1000"
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
	res, err := http.Get(url)
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
