package binarysearch

import (
	"testing"
)

type TestData struct {
	TestArray []string
	search    string
	want      int
	ok        bool
}

func TestBinarySearch(t *testing.T) {
	tests := []TestData{
		{[]string{}, "h", 0, false},
		{[]string{"a"}, "a", 0, true},
		{[]string{"a"}, "b", 0, false},
		{[]string{"a", "b"}, "a", 0, true},
		{[]string{"a", "b"}, "b", 1, true},
		{[]string{"a", "b"}, "c", 0, false},
		{[]string{"a", "b", "c"}, "a", 0, true},
		{[]string{"a", "b", "c"}, "b", 1, true},
		{[]string{"a", "b", "c"}, "c", 2, true},
		{[]string{"a", "b", "c"}, "d", 0, false},
		{[]string{"a", "b", "c", "d", "e", "f", "g"}, "a", 0, true},
		{[]string{"a", "b", "c", "d", "e", "f", "g"}, "b", 1, true},
		{[]string{"a", "b", "c", "d", "e", "f", "g"}, "c", 2, true},
		{[]string{"a", "b", "c", "d", "e", "f", "g"}, "d", 3, true},
		{[]string{"a", "b", "c", "d", "e", "f", "g"}, "e", 4, true},
		{[]string{"a", "b", "c", "d", "e", "f", "g"}, "f", 5, true},
		{[]string{"a", "b", "c", "d", "e", "f", "g"}, "g", 6, true},
		{[]string{"a", "b", "c", "d", "e", "f", "g"}, "h", 0, true},
	}
	for _, test := range tests {
		got1, got2 := BinarySearch(test.TestArray, test.search)
		if got1 != test.want || got2 != test.ok {
			t.Errorf("Got1: %#v, Want: %#v\n, Got2: %#v, Ok: %#v", got1, test.want, got2, test.ok)
		}
	}
}
