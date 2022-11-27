package main

import "testing"

type testData struct {
	givenArray [length]int
	want       bool
}

func TestFirstLast6(t *testing.T) {
	tests := []testData{
		{givenArray: [length]int{6, 1, 2, 3, 4, 5}, want: true},
		{givenArray: [length]int{1, 2, 3, 4, 5, 6}, want: true},
		{givenArray: [length]int{6, 1, 2, 3, 4, 6}, want: true},
		{givenArray: [length]int{1, 2, 4, 6, 5, 15}, want: false},
	}

	for _, test := range tests {
		got := firstLast6(test.givenArray)

		if got != test.want {
			t.Errorf("firstLast6(%#v) = \"%t\", want: \"%t\"", test.givenArray, got, test.want)
		}
	}

}
