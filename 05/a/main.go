package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func readInput(fname string) []int {
	var result []int
	fileContents, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatal(err)
	}

	fileContentsStr := string(fileContents)

	for _, r := range strings.Split(fileContentsStr, "\n") {
		if r == "" {
			continue
		}
		i, err := strconv.Atoi(r)

		if err != nil {
			log.Fatal(err)
		}

		result = append(result, i)
	}
	return result
}

func main() {

	list := readInput("input")

	l := len(list)
	currentIndex := 0
	nom := 0

	for 0 <= currentIndex && currentIndex < l {
		nom++
		newIndex := currentIndex + list[currentIndex]
		list[currentIndex]++
		currentIndex = newIndex
	}

	log.Println(nom)
}
