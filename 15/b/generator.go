package main

const divider = 2147483647

type generator struct {
	seed       uint
	multipleOf uint
	factor     uint
}

func newGenerator(f, s, m uint) generator {
	return generator{
		factor:     f,
		seed:       s,
		multipleOf: m,
	}
}

func (g *generator) next() uint {
	var n uint
	for {
		n = (g.seed * g.factor) % divider
		g.seed = n
		if n%g.multipleOf == 0 {
			break
		}
	}
	return n
}

func (g *generator) nextLower16() uint {
	r := g.next()

	return r & 65535
}
