package github.com/bahmed6/hangman-classic

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
) 

type HangManData struct {
	Word             string
	ToFind           string
	Attempts         int
	HangmanPositions [10]string
}


//selects a random word from a slice of words.
func RandomWord(words []string) string {
	return words[rand.Intn(len(words))]
}

//displays the hangman position based on the number of attempts left.
func DisplayHangmanPosition(attempts int, positions [10]string) {
	fmt.Println(positions[10-attempts])
}
//display the function of hangman-ascii art
func generateAsciiArt(text string) {
	asciiArt := map[rune]string {
		'!': '
_| 
_| 
_| 
   
_|' , 
    '='
   

_|  _| 
_|  _|' ,
        '#'
       
       
       
       
       
       

           
  _|  _|   
_|_|_|_|_| 
  _|  _|   
_|_|_|_|_| 
  _|  _|  ' ,
            '$'
           
           

       
  _|   
_|_|_| 
_|_|   
  _|_| 
_|_|_| 
  _| ' ,
        '%'  
       

           
_|_|    _| 
_|_|  _|   
    _|     
  _|  _|_| 
_|    _|_| ' ,
              '{'
           
 
  _| 
_|   
_|   
_|   
_|   
_|   
  _|' ,
      '}'
     

_|   
  _| 
  _| 
  _| 
  _| 
  _| 
_|  ' ,
      '*'
     

           
_|  _|  _| 
  _|_|_|   
_|_|_|_|_| 
  _|_|_|   
_|  _|  _|' ,
            '+'
           
           

           
    _|     
    _|     
_|_|_|_|_| 
    _|     
    _|' ,
        '_'     

           
_|_|_|_|_|' ,
            '/'
           


        _| 
      _|   
    _|     
  _|       
_| ' ,
    '0'       
           
           

       
  _|   
_|  _| 
_|  _| 
_|  _| 
  _| ' ,
       '1'      
       
       

     
  _| 
_|_| 
  _| 
  _| 
  _|' ,
      '2'
     
     

         
  _|_|   
_|    _| 
    _|   
  _|     
_|_|_|_| ' ,
           '3'

         
         

         
_|_|_|   
      _| 
  _|_|   
      _| 
_|_|_| ' ,
         '4'   
         
         

         
_|  _|   
_|  _|   
_|_|_|_| 
    _|   
    _|   ' ,
           '5'

         
         

         
_|_|_|_| 
_|       
_|_|_|   
      _| 
_|_|_| ',
        '6'   
         
         

         
  _|_|_| 
_|       
_|_|_|   
_|    _| 
  _|_| ' ,
         '7'  
         
         

           
_|_|_|_|_| 
        _| 
      _|   
    _|     
  _|      ' ,
            '8' 
           
           

         
  _|_|   
_|    _| 
  _|_|   
_|    _| 
  _|_|   ' ,
           '9'

         
         

         
  _|_|   
_|    _| 
  _|_|_| 
      _| 
_|_|_|   ' ,
           ':'

         
         

   
   
_| 
   
   
_| ' ,
     ';'   
   

     
     
  _| 
     
     
  _| 
_|   ' ,
       '<'
     

       
    _| 
  _|   
_|     
  _|   
    _| ' ,
         '='
       
       

           
           
_|_|_|_|_| 
           
_|_|_|_|_|' ,
            '>'
           
           
           

       
_|     
  _|   
    _| 
  _|   
_|  ' ,
      '@'    
       
       

       
  
       
       

                 
    _|_|_|_|_|   
  _|          _| 
_|    _|_|_|  _| 
_|  _|    _|  _| 
_|    _|_|_|_|   
  _|             
    _|_|_|_|_|_| ' ,
                   'A'

         
  _|_|   
_|    _| 
_|_|_|_| 
_|    _| 
_|    _| ' ,
           'B'

         
         

         
_|_|_|   
_|    _| 
_|_|_|   
_|    _| 
_|_|_|  ' ,
          'C'
         
         

         
  _|_|_| 
_|       
_|       
_|       
  _|_|_| ' ,
           'D'

         
         

         
_|_|_|   
_|    _| 
_|    _| 
_|    _| 
_|_|_|  ' ,
          'E'

         
         

         
_|_|_|_| 
_|       
_|_|_|   
_|       
_|_|_|_| ' ,
           'F'
         
         

         
_|_|_|_| 
_|       
_|_|_|   
_|       
_|       ' ,
           'G'
         
         

         
  _|_|_| 
_|       
_|  _|_| 
_|    _| 
  _|_|_| ' ,
           'H'

         
         

         
_|    _| 
_|    _| 
_|_|_|_| 
_|    _| 
_|    _| ' ,
           'I'
         
         

       
_|_|_| 
  _|   
  _|   
  _|   
_|_|_| ' ,
         'J'
       
       

         
      _| 
      _| 
      _| 
_|    _| 
  _|_|  ' ,
          'K' 
         
         

         
_|    _| 
_|  _|   
_|_|     
_|  _|   
_|    _| ' ,
           'L'
         
         

         
_|       
_|       
_|       
_|       
_|_|_|_| ' ,
          'M'
         
         

           
_|      _| 
_|_|  _|_| 
_|  _|  _| 
_|      _| 
_|      _| ' ,
             'N'
           
           

           
_|      _| 
_|_|    _| 
_|  _|  _| 
_|    _|_| 
_|      _| ' ,
             'O'
           
           

         
  _|_|   
_|    _| 
_|    _| 
_|    _| 
  _|_| ' ,
         'P'  
         
         

         
_|_|_|   
_|    _| 
_|_|_|   
_|       
_|      ' ,
          'Q'
         
         

           
  _|_|     
_|    _|   
_|  _|_|   
_|    _|   
  _|_|  _| ' ,
             'R'
           
           

         
_|_|_|   
_|    _| 
_|_|_|   
_|    _| 
_|    _| ' ,
           'S'
         
         

         
  _|_|_| 
_|       
  _|_|   
      _| 
_|_|_|  ' ,
          'T' 
         
         

           
_|_|_|_|_| 
    _|     
    _|     
    _|     
    _|   ' ,
           'u'     
           
           

         
_|    _| 
_|    _| 
_|    _| 
_|    _| 
  _|_|  ' ,
          'V'  
         
         

           
_|      _| 
_|      _| 
_|      _| 
  _|  _|   
    _|    ' ,
            'W'  
           
           

               
_|          _| 
_|          _| 
_|    _|    _| 
  _|  _|  _|   
    _|  _|   ' , 
               'X'  
               
               

           
_|      _| 
  _|  _|   
    _|     
  _|  _|   
_|      _| ' ,
             'Y' 
           
           

           
_|      _| 
  _|  _|   
    _|     
    _|     
    _|   ' ,
           'Z'  
           
           

           
_|_|_|_|_| 
      _|   
    _|     
  _|       
_|_|_|_|_| ' ,
             '['
           
           

_|_| 
_|   
_|   
_|   
_|   
_|   
_|_| ' ,
       '\'
     

           
_|         
  _|       
    _|     
      _|   
        _| ' ,
             ']'
           
           

_|_| 
  _| 
  _| 
  _| 
  _| 
  _| 
_|_| ' ,
       'a'

  _|_|_| 
_|    _| 
_|    _| 
  _|_|_| ' ,
           'b'
         
         

         
_|       
_|_|_|   
_|    _| 
_|    _| 
_|_|_|   ' ,
           'c'
         
         

         
         
  _|_|_| 
_|       
_|       
  _|_|_| ' ,
           'd'
         
         

         
      _| 
  _|_|_| 
_|    _| 
_|    _| 
  _|_|_| ' ,
           'e'
         
         

         
         
  _|_|   
_|_|_|_| 
_|       
  _|_|_| ' ,
           'f'
         
         

         
    _|_| 
  _|     
_|_|_|_| 
  _|     
  _|   ' ,
         'g'  
         
         

         
         
  _|_|_| 
_|    _| 
_|    _| 
  _|_|_| 
      _| 
  _|_|   ' ,
           'h'

         
_|       
_|_|_|   
_|    _| 
_|    _| 
_|    _| ' ,
           'i'
         
         

   
_| 
   
_| 
_| 
_| ' ,
     'j'
   
   

     
  _| 
     
  _| 
  _| 
  _| 
  _| 
_|   ' ,
       'k'

         
_|       
_|  _|   
_|_|     
_|  _|   
_|    _| ' ,
           'l'   
         
         

   
_| 
_| 
_| 
_| 
_| ' ,
     'm'
   
   

               
               
_|_|_|  _|_|   
_|    _|    _| 
_|    _|    _| 
_|    _|    _| ' ,
                 'n'
               
               

         
         
_|_|_|   
_|    _| 
_|    _| 
_|    _| ' ,
           'o'
         
         

         
         
  _|_|   
_|    _| 
_|    _| 
  _|_|  ' ,
          'p'  
         
         

         
         
_|_|_|   
_|    _| 
_|    _| 
_|_|_|   
_|       
_|    ' ,
        'q'   

         
         
  _|_|_| 
_|    _| 
_|    _| 
  _|_|_| 
      _| 
      _| ' ,
           'r'

         
         
_|  _|_| 
_|_|     
_|       
_|       ' ,
           's'
         
         

         
         
  _|_|_| 
_|_|     
    _|_| 
_|_|_|   ' ,
           't'
         
         

         
  _|     
_|_|_|_| 
  _|     
  _|     
    _|_| ' ,
           'u'
         
         

         
         
_|    _| 
_|    _| 
_|    _| 
  _|_|_| ' ,
           'v'
         
         

           
           
_|      _| 
_|      _| 
  _|  _|   
    _|    ' ,
            'w' 
           
           

                   
                   
_|      _|      _| 
_|      _|      _| 
  _|  _|  _|  _|   
    _|      _|    ' ,
                    'x' 
                   
                   

         
         
_|    _| 
  _|_|   
_|    _| 
_|    _|  ' ,
            'y'
         
         

         
         
_|    _| 
_|    _| 
_|    _| 
  _|_|_| 
      _| 
  _|_|   ' ,
           'z'

         
         
_|_|_|_| 
    _|   
  _|     
_|_|_|_| 

	}
	for _, char := range text {
		if art, ok := asciiArt[char]; ok {
			fmt.Println(art)
		} else {
			fmt.Println("Character not supported:", char)
		}
		fmt.Println()
	}
}
	

func github.com/bahmed6/hangman-classic() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: ./hangman words.txt hangman.txt")
		os.Exit(1)
	}

	words.txt := os.Args[1]
	hangman.txt := os.Args[2]

	// Read words from the file
	words, err := ReadWordsFromFile(words.txt)
	if err != nil {
		fmt.Println("Error reading words from file:", err)
		os.Exit(1)
	}

	// Read hangman positions from the file
	hangmanPositions, err := ReadHangmanPositions(hangman.txt)
	if err != nil {
		fmt.Println("Error reading hangman positions from file:", err)
		os.Exit(1)
	}

	// Select a random word
	wordToGuess := RandomWord(words)
	displayWord := strings.Repeat("_ ", len(wordToGuess)/2-1)
	fmt.Println("Good Luck, you have 10 attempts.")
	fmt.Println(displayWord)

	var hangmanData HangManData
	hangmanData.Word = displayWord
	hangmanData.ToFind = wordToGuess
	hangmanData.Attempts = 10
	hangmanData.HangmanPositions = hangmanPositions

	reader := bufio.NewReader(os.Stdin)

	for hangmanData.Attempts > 0 {
		fmt.Print("\nChoose a letter: ")
		input, _ := reader.ReadString('\n')
		guess := strings.ToUpper(strings.TrimSpace(input))

		if len(guess) == 1 && strings.Contains(wordToGuess, guess) {
			// Update the display word
			for i, char := range wordToGuess {
				if string(char) == guess {
					displayWord = displayWord[:i*2] + guess + displayWord[i*2+1:]
				}
			}
			hangmanData.Word = displayWord

			fmt.Println(displayWord)
			if displayWord == wordToGuess {
				fmt.Println("Congrats!")
				break
			}
		} else {
			fmt.Printf("Not present in the word, %d attempts remaining\n", hangmanData.Attempts-1)
			DisplayHangmanPosition(hangmanData.Attempts, hangmanData.HangmanPositions)
			hangmanData.Attempts--
		}
	}

	if hangmanData.Attempts == 0 {
		fmt.Printf("Sorry, you ran out of attempts. The correct word was %s.\n", wordToGuess)
	}
  reader:=bufio.NewReader(os.studin)
  fmt.print("Enter a string:")
  input, _ :=reader.ReadString('\n')
  text:= input [:len (input)-1]
  generateAsciiArt((text))
}
