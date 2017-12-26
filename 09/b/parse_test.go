package main

import "testing"

type test struct {
	input string
	want  int
}

func buildTest(i string, w int) test {
	return test{
		input: i,
		want:  w,
	}
}

func TestParseCount(t *testing.T) {

	tests := []test{
		buildTest("<>", 0),
		buildTest("<random characters>", 17),
		buildTest("<<<<>", 3),
		buildTest("<{!>}>", 2),
		buildTest("<!!>", 0),
		buildTest("<!!!>>", 0),
		buildTest("<{o\"i!a,<{i<a>", 10),
	}

	for i, te := range tests {

		res := parseCount(te.input)
		if res != te.want {
			t.Errorf("Test %d(%s) failed with value %d instead of %d\n",
				i,
				te.input,
				res,
				te.want,
			)
		}
	}
}
