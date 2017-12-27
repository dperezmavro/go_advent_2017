package main

import (
	"io/ioutil"
	"log"
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

func processInput(input []string) map[string][]string {
	var res = make(map[string][]string)
	for _, line := range input {
		parts := strings.Split(line, " <-> ")
		comms := strings.Split(parts[1], ", ")
		res[parts[0]] = comms
	}
	return res
}
