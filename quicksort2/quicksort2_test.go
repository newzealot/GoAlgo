package quicksort2

import (
	"reflect"
	"testing"
)

type TestData struct {
	TestArray []string

	want []string
}

func TestQuickSort(t *testing.T) {
	tests := []TestData{
		{[]string{"36", "75", "69", "60", "27", "15", "80", "55"},
			[]string{"15", "27", "36", "55", "60", "69", "75", "80"}},
	}
	for _, test := range tests {
		got := QuickSort(test.TestArray, 0, len(test.TestArray)-1)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("Got: %#v, Left: %#v, Right: %#v, Want: %#v\n", got, 0, len(test.TestArray)-1, test.want)
		}
	}
}
