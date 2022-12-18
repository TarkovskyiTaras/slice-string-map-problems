package main

import "testing"

type testData struct {
	givenString string
	want        bool
}

func TestXYBalanced(t *testing.T) {
	tests := []testData{
		{givenString: "tarasxxtarasxytarxttyxyttttt",
			want: true},
		{givenString: "tarasy",
			want: false},
		{givenString: "tarsxyxtar",
			want: false},
		{givenString: "tarasxy",
			want: true},
		{givenString: "tarasyxytarasxyxtarasxttywssdfsafs",
			want: true},
	}

	for _, test := range tests {
		got := xyBalance(test.givenString)
		if got != test.want {
			t.Error("error")
		}
	}
}
