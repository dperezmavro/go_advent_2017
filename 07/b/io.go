package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func readInput(fname string) []string {
	var result []string
	fileContents, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatal(err)
	}

	fileContentsStr := string(fileContents)

	for _, r := range strings.Split(fileContentsStr, "\n") {
		if r == "" {
			continue
		}
		result = append(result, r)
	}
	return result
}

func processInput(input []string) {
	for _, v := range input {
		parts := strings.Split(v, " ")

		var p program
		p.Name = parts[0]
		i, err := strconv.Atoi(
			strings.TrimLeft(
				strings.TrimRight(parts[1], ")"),
				"(",
			),
		)

		if err != nil {
			log.Fatal(err)
		}

		p.Weight = i

		if len(parts) > 2 {
			for j := 3; j < len(parts); j++ {
				p.Children = append(
					p.Children,
					strings.TrimRight(parts[j], ","),
				)
			}

		}
		nodes[p.Name] = p

	}
}
