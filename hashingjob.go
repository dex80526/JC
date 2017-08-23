package main

import (
	"time"
	"fmt"
)

/*
hashing job: job id and input value for hashing
 */
type HashingJob struct {
	jid     int32
	input   string
}

func (j *HashingJob) Hash() []byte {
	defer TimeTrack(time.Now(), "hash(_)")
	var bytes = Sha512Hash(j.input)
	return bytes
}


func TimeTrack(start time.Time, name string) {
	var elapsed = time.Since(start)
	stats.UpdateCounter(elapsed.Nanoseconds())
	fmt.Printf("elapsed time for job: %s %d\n", name, elapsed)
}
