package main

import (

	"sync/atomic"
	"fmt"
)

type JobRepo struct {
	currentJid *int32
	jobs map[int32][]byte
}


func NewJobRepo() *JobRepo {
	var zero int32 = 0;
	var repo = JobRepo{&zero, make(map[int32][]byte)}
	return &repo
}

func (repo *JobRepo) Add(id int32, hash  []byte) {
	repo.jobs[id] = hash
}

func (repo *JobRepo) Get(id int32) []byte {
	//todo handle invalid id (i.e. not id is found i the table
	var  v = repo.jobs[id]
	fmt.Printf("hash for jid: %d, %v \n", id, v)
	return v
}

//sync with atmic ops
func (repo *JobRepo) NextId() int32 {
	atomic.AddInt32(repo.currentJid, 1)
	return atomic.LoadInt32(repo.currentJid)
}

