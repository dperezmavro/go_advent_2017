package main

const divider = 2147483647

type generator struct {
	seed   uint
	factor uint
}

func newGenerator(f uint, s uint) generator {
	return generator{
		factor: f,
		seed:   s,
	}
}

func (g *generator) next() uint {
	n := (g.seed * g.factor) % divider
	g.seed = n
	return n
}

func (g *generator) nextLower16() uint {
	r := g.next()

	return r & 65535
}
