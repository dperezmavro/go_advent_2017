package main

import "log"

type dance struct {
	alphabet  []string
	charToPos map[string]int
	posToChar map[int]string
	mod       int
}

func newDance(alp []string) dance {
	cToP := make(map[string]int)
	pToC := make(map[int]string)

	for i, l := range alp {
		cToP[l] = i
		pToC[i] = l
	}

	return dance{
		alphabet:  alp,
		charToPos: cToP,
		posToChar: pToC,
		mod:       len(alp),
	}
}

func (d *dance) sequence() string {
	res := ""
	for i := 0; i < len(d.alphabet); i++ {
		res += d.posToChar[i]
	}

	return res
}

func (d *dance) spin(x int) {
	for k, v := range d.charToPos {
		newIndex := (v + x) % d.mod
		d.charToPos[k] = newIndex
		d.posToChar[newIndex] = k
	}
}

func (d *dance) exchange(a, b int) {
	cA := d.posToChar[a]
	cB := d.posToChar[b]

	d.posToChar[a] = cB
	d.posToChar[b] = cA

	d.charToPos[cA] = b
	d.charToPos[cB] = a
}

func (d *dance) partner(a, b string) {
	pA := d.charToPos[a]
	pB := d.charToPos[b]

	d.charToPos[a] = pB
	d.charToPos[b] = pA

	d.posToChar[pB] = a
	d.posToChar[pA] = b

}

func (d *dance) testMaps() {
	for v, p := range d.charToPos {
		if d.posToChar[p] != v {
			log.Printf("charToPos[%s] = %d != posToChar[%d] = %s\n",
				v, p, p, d.posToChar[p],
			)
		}
	}
}
