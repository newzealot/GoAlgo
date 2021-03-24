package mergesort

import (
	"reflect"
	"testing"
)

type TestData struct {
	TestArray []string
	want      []string
}

func TestMergeSort(t *testing.T) {
	tests := []TestData{
		{[]string{"5", "4", "3", "1", "2"},
			[]string{"1", "2", "3", "4", "5"}},
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
		got := MergeSort(test.TestArray)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("Got: %#v, Want: %#v\n", got, test.want)
		}
	}
}
