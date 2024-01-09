package main
import  (
	"io/ioutil"
	"strings"
)
func ReadHangmanPositions(hangman.txt string) ([10]string, error) {
	var positions [10]string

	content, err := ioutil.ReadFile(hangman.txt)
	if err != nil {
		return positions, err
	}

	lines := strings.Split(string(content), "\n")
	for i := 0; i < 10 && i < len(lines); i++ {
		positions[i] = lines[i]
	}

	return positions, nil
}