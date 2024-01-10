package github.com/bahmed6/hangman-classic

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

type HangManData struct {
	Word             string
	ToFind           string
	Attempts         int
	HangmanPositions [10]string
}

func github.com/bahmed6/hangman-classic() {
	startWith := flag.String("startWith", "", "Specify a saved game file to start with.")
	flag.Parse()

	var hangmanData HangManData

	if *startWith != "" {

		err := loadGame(*startWith, &hangmanData)
		if err != nil {
			fmt.Println("Error loading saved game:", err)
			return
		}
		fmt.Println("Welcome Back, you have", hangmanData.Attempts, "attempts remaining.")
	} else {

		words, err := readWordsFromFile("words.txt")
		if err != nil {
			fmt.Println("Error reading words file:", err)
			return
		}

		wordToGuess := chooseRandomWord(words)
		initialDisplay := revealInitialLetters(wordToGuess)
		hangmanData = HangManData{
			Word:     initialDisplay,
			ToFind:   wordToGuess,
			Attempts: 10,
		}
		fmt.Println("Good Luck, you have", hangmanData.Attempts, "attempts.")
	}

	for {
		fmt.Println(hangmanData.Word)
		fmt.Print("Choose: ")

		var input string
		fmt.Scan(&input)

		if input == "STOP" {
			err := saveGame("save.txt", &hangmanData)
			if err != nil {
				fmt.Println("Error saving game:", err)
				return
			}
			fmt.Println("Game Saved in save.txt.")
			return
		}

		hangmanData = processInput(input, hangmanData)

		if hangmanData.Attempts == 0 || strings.Contains(hangmanData.Word, "_") {
			if strings.Contains(hangmanData.Word, "_") {
				fmt.Println("Out of attempts. The correct word was:", hangmanData.ToFind)
			} else {
				fmt.Println("Congrats!")
			}
			break
		}
	}
}

func loadGame(game string, data *HangManData) error {
	file, err := ioutil.ReadFile(game)
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, data)
	if err != nil {
		return err
	}

	return nil
}

func saveGame(save string, data *HangManData) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(save, jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func readWordsFromFile(word string) ([]string, error) {
	file, err := ioutil.ReadFile(word)
	if err != nil {
		return nil, err
	}

	words := strings.Split(string(file), "\n")
	return words, nil
}

func chooseRandomWord(words []string) string {
	rand.Seed(time.Now().UnixNano())
	return words[rand.Intn(len(words))]
}

func revealInitialLetters(word string) string {
	var display strings.Builder
	for i := 0; i < len(word)/2-1; i++ {
		display.WriteRune('_')
		display.WriteRune(' ')
	}
	display.WriteRune('\n')
	return display.String()
}

func processInput(input string, data HangManData) HangManData {
	if len(input) == 1 {

		letter := input[0]
		if strings.Contains(data.ToFind, string(letter)) {

			for i, char := range data.ToFind {
				if char == rune(letter) {

					display := []rune(data.Word)
					display[i*2] = char
					data.Word = string(display)
				}
			}
		} else {

			fmt.Println("Not present in the word,", data.Attempts-1, "attempts remaining")
			data.Attempts--
			displayHangman(data.Attempts)
		}
	} else if len(input) > 1 {

		if input == data.ToFind {

			data.Word = data.ToFind
		} else {

			fmt.Println("Incorrect word,", data.Attempts-2, "attempts remaining")
			data.Attempts -= 2
			displayHangman(data.Attempts)
		}
	}

	return data
}

func displayHangman(attempts int) {
	switch attempts {
	case 9:
		fmt.Println()
	case 8:
		fmt.Println("  +---+  ")
	case 7:
		fmt.Println("  +---+  ")
		fmt.Println("      |  ")
	case 6:
		fmt.Println("  +---+  ")
		fmt.Println("  |   |  ")
	case 5:
		fmt.Println("  +---+  ")
		fmt.Println("  |   |  ")
		fmt.Println("      O  ")
	case 4:
		fmt.Println("  +---+  ")
		fmt.Println("  |   |  ")
		fmt.Println("      O  ")
		fmt.Println("      |  ")
	case 3:
		fmt.Println("  +---+  ")
		fmt.Println("  |   |  ")
		fmt.Println("      O  ")
		fmt.Println("     /|  ")
	case 2:
		fmt.Println("  +---+  ")
		fmt.Println("  |   |  ")
		fmt.Println("      O  ")
		fmt.Println("     /|\\")
	case 1:
		fmt.Println("  +---+  ")
		fmt.Println("  |   |  ")
		fmt.Println("      O  ")
		fmt.Println("     /|\\")
		fmt.Println("     /   ")
	case 0:
		fmt.Println("  +---+  ")
		fmt.Println("  |   |  ")
		fmt.Println("      O  ")
		fmt.Println("     /|\\")
		fmt.Println("     / \\")
		fmt.Println("=========")
	}
}
