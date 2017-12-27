package main

import (
	"log"
	"strconv"
	"testing"
)

func TestNext(t *testing.T) {
	type test struct {
		name   string
		seed   uint
		factor uint
		want   []uint
	}

	tests := []test{
		test{
			name:   "A",
			seed:   65,
			factor: 16807,
			want:   []uint{1092455, 1181022009, 245556042, 1744312007, 1352636452},
		},
		test{
			name:   "B",
			seed:   8921,
			factor: 48271,
			want:   []uint{430625591, 1233683848, 1431495498, 137874439, 285222916},
		},
	}

	for _, te := range tests {
		g := newGenerator(te.factor, te.seed)
		for i, v := range te.want {
			res := g.next()
			if res != v {
				t.Errorf(
					"Wanted %d, but got %d for test index %d of generator %s",
					v,
					res,
					i,
					te.name,
				)
			}
		}
	}
}

func TestNextLower16(t *testing.T) {
	type test struct {
		name   string
		seed   uint
		factor uint
		want   []string
	}

	tests := []test{
		test{
			name:   "A",
			seed:   65,
			factor: 16807,
			want:   []string{"1010101101100111", "1111011100111001", "1110001101001010", "0001011011000111", "1001100000100100"},
		},
		test{
			name:   "B",
			seed:   8921,
			factor: 48271,
			want:   []string{"1101001100110111", "1000010110001000", "1110001101001010", "1100110000000111", "0010100000000100"},
		},
	}

	for _, te := range tests {
		g := newGenerator(te.factor, te.seed)
		for i, v := range te.want {
			res := g.nextLower16()
			converted, err := strconv.ParseUint(v, 2, 16)
			if err != nil {
				log.Fatal(err)
			}
			if res != uint(converted) {
				t.Errorf(
					"Wanted %s, but got %d for test index %d of generator %s",
					v,
					res,
					i,
					te.name,
				)
			}
		}
	}
}

func TestScore(t *testing.T) {
	type test struct {
		iterations int
		score      uint
	}

	tests := []test{
		test{
			iterations: 1,
			score:      0,
		},
		test{
			iterations: 2,
			score:      0,
		},

		test{
			iterations: 3,
			score:      1,
		},
		test{
			iterations: 4,
			score:      1,
		},
		test{
			iterations: 5,
			score:      1,
		},
		test{
			iterations: 40000001,
			score:      588,
		},
	}

	a := newGenerator(16807, 65)
	b := newGenerator(48271, 8921)

	for _, te := range tests {
		score := getScore(a, b, te.iterations)
		if score != te.score {
			t.Errorf("Wanted %d, got %d for %d iterations\n", te.score, score, te.iterations)
		}
	}
}
