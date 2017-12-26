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

func processInput(input []string) []instruction {
	var instructions []instruction
	for _, v := range input {
		var i instruction
		var c conditional

		parts := strings.Split(v, " ")

		i.register = parts[0]
		i.action = parts[1]
		instrVal, _ := strconv.Atoi(parts[2])
		i.value = instrVal

		c.register = parts[4]
		c.operator = parts[5]
		condVal, _ := strconv.Atoi(parts[6])
		c.value = condVal

		i.condition = c

		instructions = append(instructions, i)
	}
	return instructions
}
