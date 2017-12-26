package main

import "log"

func main() {
	input := readInput("input")

	for _, v := range input {
		log.Println(parseCount(v))
	}
}
