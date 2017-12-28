package main

import (
	"log"
	"strconv"
	"strings"
)

func run(alph, input string) string {
	alphabet := strings.Split(alph, "")
	actions := getActions(input)

	d := newDance(alphabet)

	for _, action := range actions {

		actionParts := strings.Split(action, "")
		amounts := action[1:]

		switch actionParts[0] {

		case "s":
			spin, _ := strconv.Atoi(amounts)
			d.spin(spin)
			break

		case "x":
			parts := strings.Split(amounts, "/")
			a, _ := strconv.Atoi(parts[0])
			b, _ := strconv.Atoi(parts[1])
			d.exchange(a, b)
			break

		case "p":
			parts := strings.Split(amounts, "/")
			d.partner(parts[0], parts[1])
			break

		default:
			log.Println("Unknown instruction ", actions)
			break
		}
	}

	return d.sequence()
}

func main() {

	res := run("abcdefghijklmnop", "input")
	log.Println("Sequence", res)
}
