package format

import (
	"slices"
	"testing"
)

func TestDashSplit(t *testing.T) {
	var ports = []int{}
	err := dashSplit("1-2", &ports)
	if err != nil {
		t.Error(err)
	}
	if slices.Compare(ports, []int{1, 2}) != 0 {
		t.Errorf("dashSplit(1-2) = %d; want []int{1,2}", ports)
	}
}

func TestParse(t *testing.T) {
	type tt struct {
		in  string
		res []int
	}

	ts := []tt{
		{"1", []int{1}},
		{"1-2", []int{1, 2}},
		{"2-10", []int{2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{"2,10", []int{2, 10}},
	}

	for _, te := range ts {
		ports, err := Parse(te.in)
		if err != nil {
			t.Error(err)
		}
		if slices.Compare(ports, te.res) != 0 {
			t.Errorf("\nPortsRange:%v\nwant: %v", ports, te.res)
		}
	}

}
