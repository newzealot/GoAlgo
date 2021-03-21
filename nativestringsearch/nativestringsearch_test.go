package nativestringsearch

import (
	"testing"
)

type TestData struct {
	testArray string
	search    string
	want      int
}

func TestNaiveStringSearch(t *testing.T) {
	tests := []TestData{
		{"om", "omg", 0},
		{"omg", "omg", 1},
		{"zzomgzz", "omg", 1},
		{"omgzzomgzzomg", "omg", 3},
		{"asdfsadfsadf", "omg", 0},
	}
	for _, test := range tests {
		got := NativeStringSearch(test.testArray, test.search)
		if got != test.want {
			t.Errorf("Got1: %#v, Want: %#v", got, test.want)
		}
	}
}
