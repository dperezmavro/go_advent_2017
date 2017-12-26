package main

import "log"

var counts = make(map[string]int)
var nodes = make(map[string]program)

func calculateCounts() {
	for v := range nodes {
		counts[v]++
		for _, n := range nodes[v].Children {
			counts[n]++
		}
	}
}

func main() {

	input := readInput("input")
	processInput(input)
	calculateCounts()

	minCount := 100000000000
	minName := "NONAME"

	for v, k := range counts {
		if k < minCount {
			minCount = k
			minName = v
		}
	}

	log.Println(minName)
}
