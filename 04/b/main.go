package main

import (
	"io/ioutil"
	"log"
	"sort"
	"strings"

	"github.com/pkg/profile"
)

func readPassphrases(fname string) []string {
	var result []string
	fileContents, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatal(err)
	}

	fileContentsStr := string(fileContents)
	for _, line := range strings.Split(fileContentsStr, "\n") {
		if line == "" {
			continue
		}
		result = append(result, line)
	}

	return result
}

func processPassphrase(phrase []string) bool {
	m := make(map[string]bool)

	for _, v := range phrase {
		k := strings.Split(v, "")

		sort.Strings(k)
		newVal := strings.Join(k, "")

		if m[newVal] {
			return false
		}

		m[newVal] = true
	}
	return true
}

func main() {

	defer profile.Start(profile.MemProfile).Stop()

	passphrases := readPassphrases("input")
	validPassphrases := 0

	for _, v := range passphrases {
		parts := strings.Split(v, " ")

		if processPassphrase(parts) {
			validPassphrases++
		}
	}

	log.Println(validPassphrases)

}
