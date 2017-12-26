package main

import (
	"io/ioutil"
	"log"
	"strings"
)

func readPassphrases(fname string) []string{
	var result []string
	fileContents, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatal(err)
	}

	fileContentsStr := string(fileContents)
	for _, line := range strings.Split(fileContentsStr, "\n") {
		result = append(result, line)
	}


	return result
}

func processPassphrase(phrase []string) bool{
	m := make(map[string]bool)

	for _, v := range phrase {
		if m[v] {
			return false
		}

		m[v] = true
	}
	return true
}

func main(){
	passphrases := readPassphrases("input")
	validPassphrases := 0

	for _, v := range passphrases {
		parts := strings.Split(v, " ")

		if len(parts) == 1 {
			continue
		}

		if processPassphrase(parts) {
			validPassphrases++
			log.Println("valid", parts)
		}
	}

	log.Println(validPassphrases)

}