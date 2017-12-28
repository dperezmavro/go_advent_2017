package main

import (
	"log"
	"strconv"
	"strings"
)

func main() {
	alphabet := strings.Split("abcdefghijklmnop", "")
	actions := getActions("input")

	// alphabet := strings.Split("abcde", "")
	// actions := getActions("inputTest")
	d := newDance(alphabet)

	for _, action := range actions {

		actionParts := strings.Split(action, "")
		amounts := strings.Join(actionParts[1:], "")

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
	d.testMaps()
	log.Println("Sequence", d.sequence())
}
