package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"time"
)

const (
	url = "https://us-central1-rdsm-analytics-development.cloudfunctions.net/random"
)

type Result struct {
	TableId      string                 `json:"tableId"`
	Key          string                 `json:"key"`
	StartDate    time.Time              `json:"startDate"`
	EndDate      time.Time              `json:"endDate"`
	SysDate      time.Time              `json:"sysDate"`
	SysTime      float32                `json:"sysTime"`
	Count        int                    `json:"count"`
	ResultEvents map[string]interface{} `json:"resultEvents"`
}

func makeParallelsRequests(numOfRequests int, ch chan Result) {
	defer close(ch)
	var results = []chan Result{}

	for i := 0; i < numOfRequests; i++ {
		results = append(results, make(chan Result))
		go makeRequest(results[i])
	}

	for i := range results {
		for result := range results[i] {
			ch <- result
		}
	}
}

func makeRequest(ch chan Result) {
	defer close(ch)
	res, err := http.Get(url)
	if err != nil {
		ch <- Result{}
	}

	defer res.Body.Close()

	jsonData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		ch <- Result{}
	}

	var result Result

	err = json.Unmarshal([]byte(jsonData), &result)
	if err != nil {
		ch <- Result{}
	}

	ch <- result
}

func main() {
	log.Printf("Number Of CPU's ~> %d", runtime.NumCPU())

	result := make(chan Result)
	go makeParallelsRequests(800, result)

	var elapsedTime float32

	for res := range result {
		elapsedTime += res.SysTime

		log.Printf(
			"<Result table: %s key: %s, SysDate: %s sysTime: %f>",
			res.TableId, res.Key, res.SysDate, res.SysTime,
		)
	}
	log.Println("Avg: ", elapsedTime/200)
}
