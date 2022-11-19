package main

import "testing"

type testData struct {
	givenSlice []int
	want       bool
}

func TestUnluckyOne(t *testing.T) {
	tests := []testData{
		{givenSlice: []int{}, want: false},
		{givenSlice: []int{1}, want: false},
		{givenSlice: []int{1, 4}, want: false},
		{givenSlice: []int{1, 3}, want: true},
		{givenSlice: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, want: false},
		{givenSlice: []int{1, 3, 4, 5, 6, 7, 8, 9}, want: true},
		{givenSlice: []int{22, 1, 3, 4, 5, 6, 7, 8, 9}, want: true},
		{givenSlice: []int{22, 1, 4, 5, 6, 7, 8, 8}, want: false},
		{givenSlice: []int{1, 2, 3, 4, 5, 6, 1, 3}, want: true},
		{givenSlice: []int{1, 2, 3, 4, 5, 6, 1, 4}, want: false},
		{givenSlice: []int{1, 2, 3, 4, 5, 6, 3, 1}, want: false},
	}

	for _, test := range tests {
		got := unluckyOne(test.givenSlice)
		if got != test.want {
			t.Errorf("unluckyOne(%#v) = %t, want: %t", test.givenSlice, got, test.want)
		}
	}
}
