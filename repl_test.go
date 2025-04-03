package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCleanInput(t *testing.T) {
	type testcase struct {
		input    string
		expected []string
	}

	cases := []testcase{
		{
			input:    "Hello World",
			expected: []string{"hello", "world"},
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("cleanInput(%v)", c.input), func(t *testing.T) {
			got := cleanInput(c.input)
			want := c.expected

			if !reflect.DeepEqual(got, want) {
				t.Errorf("got: %v != want: %v", got, want)
			}

		})
	}
}
