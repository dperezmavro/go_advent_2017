package main

const divider = 2147483647

type generator struct {
	seed   int
	factor int
}

func newGenerator(f int, s int) generator {
	return generator{
		factor: f,
		seed:   s,
	}
}

func (g *generator) next() int {
	n := (g.seed * g.factor) % divider
	g.seed = n
	return n
}
