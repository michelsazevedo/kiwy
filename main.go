package main

import (
	"github.com/michelsazevedo/kiwy/pkg/request"
	"log"
)

func main() {
	concorrentWorkers := 10

	result := make(chan request.Result)
	go request.MakeParallelsRequests(concorrentWorkers, result)

	for res := range result {
		log.Printf(
			"<Result table: %s key: %s, SysDate: %s sysTime: %f>",
			res.TableId, res.Key, res.SysDate, res.SysTime,
		)
	}
}
