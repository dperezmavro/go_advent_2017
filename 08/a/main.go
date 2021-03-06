package main

import "log"

var registers = make(map[string]int)
var largest = 0

func main() {
	input := readInput("input")
	instructions := processInput(input)

	for _, i := range instructions {
		i.apply(&registers)
	}

	log.Println(registers)

	for i, v := range registers {
		if v > largest {
			log.Println(i, v)
			largest = v
		}
	}
}
