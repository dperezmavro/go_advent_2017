package main

import "log"

func getScore(a, b generator, count int) uint {
	var score uint

	for i := 0; i < count; i++ {
		resA := a.nextLower16()
		resB := b.nextLower16()
		if resA == resB {
			score++
		}
	}

	return score
}

func main() {
	a := newGenerator(16807, 883, 4)
	b := newGenerator(48271, 879, 8)

	log.Println(getScore(a, b, 5000001))
}
