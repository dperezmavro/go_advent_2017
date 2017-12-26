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

func TestParse(t *testing.T) {

	tests := []test{
		buildTest("{}", 1),
		buildTest("{{{}}}", 6),
		buildTest("{{},{}}", 5),
		buildTest("{{{},{},{{}}}}", 16),
		buildTest("{<a>,<a>,<a>,<a>}", 1),
		buildTest("{{<ab>},{<ab>},{<ab>},{<ab>}}", 9),
		buildTest("{{<!!>},{<!!>},{<!!>},{<!!>}}", 9),
		buildTest("{{<a!>},{<a!>},{<a!>},{<ab>}}", 3),
		buildTest("<>", 0),
		buildTest("<random characters>", 0),
		buildTest("<<<<<<<>", 0),
		buildTest("<{!>}>", 0),
		buildTest("<!!>", 0),
		buildTest("<!!!>>", 0),
		buildTest("<{o\"i!a,<{i<a>", 0),
	}

	for i, te := range tests {

		res := parse(te.input)
		if res != te.want {
			t.Errorf("Test %d failed with value %d instead of %d\n", i, res, te.want)
		}
	}
}
