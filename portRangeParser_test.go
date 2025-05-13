package main

import (
	"log"
	"slices"
	"testing"
)

type tt struct {
	in  string
	res []int
}

var ts = []tt{
	{"1", []int{1}},
	{" 3 ", []int{3}},
	{"1-2", []int{1, 2}},
	{" 1 - 2 ", []int{1, 2}},
	{"2-10", []int{2, 3, 4, 5, 6, 7, 8, 9, 10}},
	{"2,10", []int{2, 10}},
	{"2 ,10", []int{2, 10}},
	{" 2 , 10 ", []int{2, 10}},
}

func TestParsePortRange(t *testing.T) {
	for _, te := range ts {
		t.Run(te.in, func(t *testing.T) {
			// t.Log(te.in)
			ports, err := ParsePortRanges(te.in)
			if err != nil {
				log.Fatal(err)
			}

			if slices.Compare(ports, te.res) != 0 {
				t.Errorf("\nPortsRange:%v\nwant: %v", ports, te.res)
			}
		})
	}
}
