package request

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

const (
	urls = []string{
		"https://us-central1-rdsm-analytics-development.cloudfunctions.net/random1",
		"https://us-central1-rdsm-analytics-development.cloudfunctions.net/random2",
		"https://us-central1-rdsm-analytics-development.cloudfunctions.net/random3",
		"https://us-central1-rdsm-analytics-development.cloudfunctions.net/random4",
		"https://us-central1-rdsm-analytics-development.cloudfunctions.net/random5",
		"https://us-central1-rdsm-analytics-development.cloudfunctions.net/random6",
		"https://us-central1-rdsm-analytics-development.cloudfunctions.net/random7",
		"https://us-central1-rdsm-analytics-development.cloudfunctions.net/random8",
		"https://us-central1-rdsm-analytics-development.cloudfunctions.net/random9",
		"https://us-central1-rdsm-analytics-development.cloudfunctions.net/random10",
	}
)

type Result struct {
	TableId      string                 `json:"tableId"`
	Key          string                 `json:"key"`
	StartDate    time.Time              `json:"startDate"`
	EndDate      time.Time              `json:"endDate"`
	SysDate      time.Time              `json:"sysDate"`
	SysTime      float64                `json:"sysTime"`
	Count        int                    `json:"count"`
	ResultEvents map[string]interface{} `json:"resultEvents"`
}

func MakeParallelsRequests(numOfRequests int, ch chan Result) {
	defer close(ch)
	var results = []chan Result{}

	for i := 0; i < numOfRequests; i++ {
		results = append(results, make(chan Result))
		go MakeRequest(results[i])
	}

	for i := range results {
		for result := range results[i] {
			ch <- result
		}
	}
}

func MakeRequest(ch chan Result) {
	defer close(ch)
	res, err := http.Get(GetUrl())
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

func GetUrl() string {
	rand.Seed(time.Now().Unix())
	return urls[rand.Intn(len(urls))]
}
