package main

import "log"

var input map[string][]string

func scan(pid string, canComm *map[string]bool) []string {
	var res []string
	for _, v := range input[pid] {
		if !(*canComm)[v] {
			res = append(res, v)
		}
	}

	return res
}

func findPids(toScan []string, canComm, isScanned *map[string]bool) {
	for _, pid := range toScan {
		if (*isScanned)[pid] {
			continue
		} else {
			(*isScanned)[pid] = true
		}

		pids := scan(pid, canComm)
		if len(pids) > 0 {
			for _, r := range pids {
				(*canComm)[r] = true
			}
			findPids(pids, canComm, isScanned)
		}
	}
}

func main() {
	var isScanned = make(map[string]bool)

	input = processInput(
		readInput("inputTest"),
	)

	var commsWZero = make(map[string]bool)
	commsWZero["0"] = true
	findPids(scan("0", &commsWZero), &commsWZero, &isScanned)

	log.Println(len(commsWZero))
}
