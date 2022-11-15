package main

import "testing"

type testData struct {
	givenSlice []int
	want       []int
}

func twoSliceEqualChecker(sliceOne, sliceTwo []int) bool {
	if len(sliceOne) != len(sliceTwo) {
		return false
	}

	for i := range sliceOne {
		if sliceOne[i] != sliceTwo[i] {
			return false
		}
	}

	return true
}

func TestChangeThreeFollowingTwo(t *testing.T) {
	tests := []testData{
		{givenSlice: []int{1, 2, 3, 4, 5, 6, 2, 3, 10},
			want: []int{1, 2, 0, 4, 5, 6, 2, 0, 10}},
		{givenSlice: []int{10, -2, 2, 3, 124},
			want: []int{10, -2, 2, 0, 124}},
	}

	for _, test := range tests {
		got := changeThreeFollowingTwo(test.givenSlice)
		if twoSliceEqualChecker(got, test.want) != true {
			t.Errorf("changeThreeFollowingTwo(%#v) = \"%v\", want: \"%v\"", test.givenSlice, got, test.want)
		}
	}
}
