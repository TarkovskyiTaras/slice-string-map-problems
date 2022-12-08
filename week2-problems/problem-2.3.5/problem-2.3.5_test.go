package main

import "testing"

type testData struct {
	givenSlice []int
	want       bool
}

func TestNoTwoOrThreeInSlice(t *testing.T) {
	tests := []testData{
		{givenSlice: []int{1, 2, 4, 55, -8},
			want: false},
		{givenSlice: []int{4, 88, 3, 45},
			want: false},
		{givenSlice: []int{1, 2, 3},
			want: false},
		{givenSlice: []int{115, 33, 8, -2, 15},
			want: true},
	}

	for _, test := range tests {
		got := noTwoOrThreeInSlice(test.givenSlice)
		if got != test.want {
			t.Errorf("noTwoOrThreeInSlice(\"%#v\") = \"%v\", want: \"%v\"", test.givenSlice, got, test.want)
		}
	}
}
