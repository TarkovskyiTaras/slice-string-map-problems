package main

import "testing"

type testData struct {
	givenSliceOne []int
	givenSliceTwo []int
	want          bool
}

func TestCommonFirstOrLastEl(t *testing.T) {

	tests := []testData{
		{givenSliceOne: []int{1, 3, 5},
			givenSliceTwo: []int{1, 4, 6},
			want:          true},
		{givenSliceOne: []int{2, 3, 6},
			givenSliceTwo: []int{1, 4, 6},
			want:          true},
		{givenSliceOne: []int{2, 3, 4},
			givenSliceTwo: []int{5, 6, 7},
			want:          false},
	}

	for _, test := range tests {
		got := commonFirstOrLastElement(test.givenSliceOne, test.givenSliceTwo)
		if got != test.want {
			t.Errorf("commonFirstOrLastElement(%#v, %#v) = \"%t\", want: \"%t\"", test.givenSliceOne, test.givenSliceTwo, got, test.want)
		}
	}

}
