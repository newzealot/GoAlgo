package linearsearch

import (
	"testing"
)

type TestData struct {
	TestArray []string
	search    string
	want      int
	ok        bool
}

func TestLinearSearch(t *testing.T) {
	tests := []TestData{
		{[]string{"Zul", "Calvin", "Jing", "Toby", "Jeya", "Lakshimi", "Sean"}, "Calvin", 1, true},
		{[]string{"Zul", "Calvin", "Jing", "Toby", "Jeya", "Lakshimi", "Sean"}, "Venki", 0, false},
		{[]string{"Cat", "Fish", "Chicken", "Cow", "Dog", "Horse", "Sheep"}, "Cat", 0, true},
		{[]string{"Cat", "Fish", "Chicken", "Cow", "Dog", "Horse", "Sheep"}, "Spider", 0, false},
	}
	for _, test := range tests {
		got1, got2 := LinearSearch(test.TestArray, test.search)
		if got1 != test.want || got2 != test.ok {
			t.Errorf("Got1: %#v, Want: %#v\n, Got2: %#v, Ok: %#v", got1, test.want, got2, test.ok)
		}
	}
}
