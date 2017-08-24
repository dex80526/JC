package main

import (
	"fmt"
	"sync"
)

type Stat struct {
	Total        int64    `json:"total"`
	Average      int64    `json:"average"`

}

var lock  sync.Mutex

func NewStat() * Stat {
	return &Stat{0, 0}
}

//sync
func(s *Stat) UpdateCounter(atime int64) {
    lock.Lock()
	var nc = s.Total + 1
	var nv int64 = s.Total * s.Average  + atime
	s.Average = nv/nc  //todo do float number div and round to int
	s.Total = nc
	lock.Unlock()
	fmt.Printf("number requests: %d, average time: %d\n", s.Total, s.Average)
}

