package main

type dance struct {
	alphabet  []string
	charToPos map[string]int
	posToChar map[int]string
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
	mod := len(d.alphabet)
	for k, v := range d.charToPos {
		newIndex := (v + x) % mod
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

	d.posToChar[pA] = b
	d.posToChar[pB] = a

	d.charToPos[a] = pB
	d.charToPos[b] = pA
}
