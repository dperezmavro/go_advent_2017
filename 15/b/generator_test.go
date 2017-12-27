package main

import (
	"log"
	"strconv"
	"testing"
)

func TestNext(t *testing.T) {
	type test struct {
		name       string
		seed       uint
		multipleOf uint
		factor     uint
		want       []uint
	}

	tests := []test{
		test{
			name:       "A",
			seed:       65,
			factor:     16807,
			multipleOf: 4,
			want:       []uint{1352636452, 1992081072, 530830436, 1980017072, 740335192},
		},
		test{
			name:       "B",
			seed:       8921,
			factor:     48271,
			multipleOf: 8,
			want:       []uint{1233683848, 862516352, 1159784568, 1616057672, 412269392},
		},
	}

	for _, te := range tests {
		g := newGenerator(te.factor, te.seed, te.multipleOf)
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
		name       string
		seed       uint
		multipleOf uint
		factor     uint
		want       []string
	}

	tests := []test{
		test{
			name:       "A",
			seed:       65,
			factor:     16807,
			multipleOf: 4,
			want:       []string{"1001100000100100", "1011111010110000", "1101010001100100", "1010100110110000", "1001111001011000"},
		},
		test{
			name:       "B",
			seed:       8921,
			factor:     48271,
			multipleOf: 8,
			want:       []string{"1000010110001000", "1111010010000000", "1110100001111000", "0001010101001000", "1011101101010000"},
		},
	}

	for _, te := range tests {
		g := newGenerator(te.factor, te.seed, te.multipleOf)
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
			iterations: 1057,
			score:      1,
		},
		test{
			iterations: 5000001,
			score:      309,
		},
	}

	a := newGenerator(16807, 65, 4)
	b := newGenerator(48271, 8921, 8)

	for _, te := range tests {
		score := getScore(a, b, te.iterations)
		if score != te.score {
			t.Errorf("Wanted %d, got %d for %d iterations\n", te.score, score, te.iterations)
		}
	}
}
