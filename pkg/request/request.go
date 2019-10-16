package request

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	url = "https://us-central1-rdsm-analytics-development.cloudfunctions.net/random1"
)

type Result struct {
	Query      string    `json:"query"`
	TableId    string    `json:"tableId"`
	TenantId   string    `json:"tenantId"`
	WorkflowId string    `json:"workflowId"`
	StartDate  time.Time `json:"startDate"`
	EndDate    time.Time `json:"endDate"`
	SysDate    time.Time `json:"sysDate"`
	SysTime    float64   `json:"sysTime"`
	Count      int       `json:"count"`
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
