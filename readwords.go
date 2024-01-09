package main
import (
	"io/ioutil"
	"strings"
)
func ReadWordsFromFile(words.txt string) ([]string, error) {
	content, err := ioutil.ReadFile(words.txt)
	if err != nil {
		return nil, err
	}
	words := strings.Split(string(content), "\n")
	return words, nil
}