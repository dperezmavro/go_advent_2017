package main

import (
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

type Row []int

func diff(r Row) int {
	var max = 0
	var min = 0

	for i, v := range r {
		if i == 0 {
			min = r[0]
		}
		if v < min {
			min = v
		} else if v > max {
			max = v
		}
	}

	return max - min
}

func div(r Row) int {
	for _, v := range r {
		for _, k := range r {
			if k != 0 && v != k && v%k == 0 {
				return v / k
			}
		}
	}

	return 0
}

func buildRow(r string) Row {
	var row Row
	ints := strings.Split(r, "\t")
	for _, v := range ints {
		i, _ := strconv.Atoi(v)
		row = append(row, i)
	}
	sort.Ints(row)
	return row
}

func readInput(fname string) ([]Row, error) {
	var result []Row
	fileContents, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatal(err)
	}

	fileContentsStr := string(fileContents)

	for _, r := range strings.Split(fileContentsStr, "\n") {
		result = append(result, buildRow(r))
	}
	return result, nil
}

func partOne() {
	rows, err := readInput("input.txt")
	if err != nil {
		log.Fatalf("Error reading input %v\n", err)
	}

	var total = 0
	for _, v := range rows {
		total += diff(v)
	}

	log.Println("Part one", total)
}

func partTwo() {
	rows, err := readInput("input.txt")
	if err != nil {
		log.Fatalf("Error reading input %v\n", err)
	}

	var total = 0
	for _, v := range rows {
		total += div(v)
	}

	log.Println("Part two", total)
}

func main() {
	partOne()
	partTwo()
}
