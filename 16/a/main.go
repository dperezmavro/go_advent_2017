package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	alphabet := strings.Split("abcdefghijklmnop", "")
	// alphabet := strings.Split("abcde", "")
	actions := getActions("input")
	d := newDance(alphabet)
	fmt.Println(d)

	for _, action := range actions {

		actionParts := strings.Split(action, "")

		switch actionParts[0] {
		case "s":
			spin, _ := strconv.Atoi(actionParts[1])
			d.spin(spin)
			break
		case "x":
			a, _ := strconv.Atoi(actionParts[1])
			b, _ := strconv.Atoi(actionParts[3])
			d.exchange(a, b)
			break
		case "p":
			d.partner(actionParts[1], actionParts[3])
			break
		default:
			log.Println("Unknown instruction ", actions)
			break
		}
	}

	fmt.Println(d.sequence())
}
