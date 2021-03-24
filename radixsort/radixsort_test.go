package radixsort

import (
	"reflect"
	"testing"
)

type TestData struct {
	TestArray []int
	want      []int
}

func TestQuickSort(t *testing.T) {
	tests := []TestData{
		{[]int{23, 345, 5467, 12, 2345, 9852, 9},
			[]int{9, 12, 23, 345, 2345, 5467, 9852}},
	}
	for _, test := range tests {
		got := RadixSort(test.TestArray)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("Got: %#v, Left: %#v, Right: %#v, Want: %#v\n", got, 0, len(test.TestArray)-1, test.want)
		}
	}
}
