package main

import "testing"

type testData struct {
	givenString string
	want        string
}

func TestDoubleChar(t *testing.T) {
	tests := []testData{
		{givenString: "Taras",
			want: "TTaarraass"},
		{givenString: "123",
			want: "112233"},
	}

	for _, test := range tests {
		got := doubleChar(test.givenString)
		if got != test.want {
			t.Errorf("doubleChar(%s) = %s, want: %s", test.givenString, got, test.want)
		}
	}
}
