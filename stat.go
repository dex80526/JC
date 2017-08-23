package main

import (
	"fmt"
)

type Stat struct {
	Total        int64    `json:"total"`
	Average      int64    `json:"average"`

}


func NewStat() * Stat {
	return &Stat{0, 0}
}

//sync is handeld by channel (todo: we may still need sync for the ops here?)
func(s *Stat) UpdateCounter(atime int64) {

	var nc = s.Total + 1
	var nv int64 = s.Total * s.Average  + atime
	s.Average = nv/nc  //todo do float number div and round to int
	s.Total = nc
	fmt.Printf("number requests: %d, average time: %d\n", s.Total, s.Average)
}

