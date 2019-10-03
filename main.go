package main

import (
	"log"

	"github.com/michelsazevedo/kiwy/pkg/request"
)

func main() {
	concurrentWorkers := 10

	result := make(chan request.Result)
	go request.MakeParallelsRequests(concurrentWorkers, result)

	for res := range result {
		log.Printf(
			"<Result table: %s key: %s, SysDate: %s sysTime: %f>",
			res.TableId, res.Key, res.SysDate, res.SysTime,
		)
	}
}
