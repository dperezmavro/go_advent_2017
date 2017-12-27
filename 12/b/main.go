package main

import (
	"log"
	"strconv"
)

var input map[string][]string

func scanNewPidsFromInput(pid string, processGroup *map[string]bool) []string {
	var res []string
	for _, v := range input[pid] {
		if (*processGroup)[v] == false {
			res = append(res, v)
		}
	}

	return res
}

func findPids(toScan []string, processGroup *map[string]bool) {
	var newPids []string

	for _, pid := range toScan {
		if (*processGroup)[pid] {
			continue
		}

		pids := scanNewPidsFromInput(pid, processGroup)
		for _, v := range pids {
			if !(*processGroup)[pid] {
				newPids = append(newPids, v)
			}
		}

		(*processGroup)[pid] = true
	}

	if len(newPids) > 0 {
		findPids(newPids, processGroup)
	}
}

func getGroup(pid string) []string {
	var processGroup = make(map[string]bool)

	processGroup[pid] = true

	findPids(
		scanNewPidsFromInput(pid, &processGroup),
		&processGroup,
	)

	keys := make([]string, 0, len(processGroup))
	for k := range processGroup {
		keys = append(keys, k)
	}

	return keys
}

func main() {

	input = processInput(
		readInput("inputTest"),
	)

	for i := 0; i < 7; i++ {
		log.Println(i, getGroup(strconv.Itoa(i)))
	}

	// for v := range input {
	// 	groups := getGroup(v)
	// 	log.Println(v, groups)
	// 	// for _, g := range groups {
	// 	// 	delete(input, g)
	// 	// }
	// }
}
