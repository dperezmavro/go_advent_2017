package main

import (
	"strings"
)

var ignoreNext = false
var inGarbage = false

func parseCount(input string) int {
	stream := strings.Split(input, "")
	var score = 0

	for _, c := range stream {
		if ignoreNext == true {
			ignoreNext = false
			continue
		}

		switch c {
		case "<":
			if inGarbage {
				score++
			} else {
				inGarbage = true
			}
			break

		case ">":
			inGarbage = false
			break

		case "!":
			ignoreNext = true
			break
		default:
			if inGarbage {
				score++
				continue
			}
		}
	}
	return score
}
