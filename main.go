package main

/*
Simple rest service implementation
It uses Channel to implement FIFO and "worker pool" for handling concurrent (hash) request
 */

import (
	"net/http"
	"log"
	"fmt"
	"time"
)
var (
	maxQueueSize = 10
	maxWorkers = 5
	port = "8080"
	jobChans chan HashingJob
	stats   *Stat
	jobRepo *JobRepo
)

func doWork(id int, j HashingJob) {

	fmt.Printf("worker%d: started, working \n", id)
	time.Sleep(5*time.Second)

	jobRepo.Add(j.jid, j.Hash())
	fmt.Printf("worker%d: completed job: %d!\n", id, j.jid)
}




func main() {

	//create job repo
	jobRepo = NewJobRepo()
	//create job stat
	stats = NewStat()
	// create job channel
	jobChans = make(chan HashingJob, maxQueueSize)

	// create workers
	for i := 1; i <= maxWorkers; i++ {
		go func(i int) {
			for j := range jobChans {
				doWork(i, j)
			}
		}(i)
	}
	//register handlers
	http.HandleFunc("/hash", HashCreate)
	http.HandleFunc("/hash/", HashShow)
	http.HandleFunc("/stats", Stats)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
