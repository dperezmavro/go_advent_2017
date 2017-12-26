package main

import (
	"strings"
)

var ignoreNext = false
var inGarbage = false

func parse(input string) int {
	stream := strings.Split(input, "")
	var depth = 0
	var score = 0

	for _, c := range stream {
		if ignoreNext == true {
			ignoreNext = false
			continue
		}

		switch c {
		case ",":
			if inGarbage {
				continue
			}
			break

		case "{":
			if inGarbage {
				continue
			}
			depth++
			score += depth

			break

		case "}":
			if inGarbage {
				continue
			}

			depth--

		case "<":
			inGarbage = true
			break

		case ">":
			inGarbage = false
			break

		case "!":
			ignoreNext = true
			break
		default:
			break
		}
	}
	return score
}
