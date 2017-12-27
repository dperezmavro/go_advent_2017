package main

import (
	"log"
	"testing"
)

func TestNext(t *testing.T) {
	type test struct {
		name   string
		seed   int
		factor int
		want   []int
	}

	tests := []test{
		test{
			name:   "A",
			seed:   65,
			factor: 16807,
			want:   []int{1092455, 1181022009, 245556042, 1744312007, 1352636452},
		},
		test{
			name:   "B",
			seed:   8921,
			factor: 48271,
			want:   []int{430625591, 1233683848, 1431495498, 137874439, 285222916},
		},
	}

	for _, te := range tests {
		g := newGenerator(te.factor, te.seed)
		for i, v := range te.want {
			res := g.next()
			if res != v {
				t.Errorf("Wanted %d, but got %d for test index %d of generator %s", v, res, i, te.name)
			}
		}
		log.Println(te)
	}
}
