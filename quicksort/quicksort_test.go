package quicksort

import (
	"reflect"
	"testing"
)

type TestData struct {
	TestArray []string
	Start     int
	End       int
	want      []string
}

func TestQuickSort(t *testing.T) {
	tests := []TestData{
		{[]string{"4", "8", "2", "1", "5", "7", "6", "3"}, 0, 7,
			[]string{"1", "2", "3", "4", "5", "6", "7", "8"}},
	}
	for _, test := range tests {
		got := QuickSort(test.TestArray, test.Start, test.End)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("Got: %#v, Left: %#v, Right: %#v, Want: %#v\n", got, test.Start, test.End, test.want)
		}
	}
}
