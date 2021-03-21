package selectionsort

import (
	"reflect"
	"testing"
)

type TestData struct {
	TestArray []string
	want      []string
}

func TestLinearSearch(t *testing.T) {
	tests := []TestData{
		{[]string{},
			[]string{}},
		{[]string{"Calvin"},
			[]string{"Calvin"}},
		{[]string{"Zul", "Calvin"},
			[]string{"Calvin", "Zul"}},
		{[]string{"Zul", "Calvin", "Jing", "Toby", "Jeya", "Lakshimi", "Sean"},
			[]string{"Calvin", "Jeya", "Jing", "Lakshimi", "Sean", "Toby", "Zul"}},
		{[]string{"Zul", "Calvin", "Jing", "Toby", "Jeya", "Lakshimi", "Sean"},
			[]string{"Calvin", "Jeya", "Jing", "Lakshimi", "Sean", "Toby", "Zul"}},
		{[]string{"Cat", "Fish", "Chicken", "Cow", "Dog", "Horse", "Sheep"},
			[]string{"Cat", "Chicken", "Cow", "Dog", "Fish", "Horse", "Sheep"}},
	}
	for _, test := range tests {
		got := SelectionSort(test.TestArray)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("Got: %#v, Want: %#v\n", got, test.want)
		}
	}
}
