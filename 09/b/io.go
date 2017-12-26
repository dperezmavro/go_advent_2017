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
