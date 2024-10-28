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

	ports, err := Parse("1-3")
	if err != nil {
		t.Error(err)
	}

	if slices.Compare(ports, []int{1, 2, 3}) != 0 {
		t.Errorf("\nPortsRange:%v", ports)
	}

}
