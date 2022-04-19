package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode"
)

// the number to let user try to guess.
const try = 5

func main() {
	store := map[rune]bool{}
	randWord := randomWord()

	var tryLeft int

	for !wordGuessed(randWord, store) {
		guessing(randWord, store)
		input := readInput()
		if len(input) != 1 {
			fmt.Println("the input is wrong please just input one letter")
			continue
		}
		l := rune(input[0])
		if checkGuessing(randWord, l) {
			store[l] = true
		} else {
			tryLeft++
			if tryLeft == try {
				fmt.Println("!!!!!!! Game Over !!!!!!!")
				fmt.Printf(">>>>> %s \n", randWord)
				os.Exit(1)
			}
		}
	}

	fmt.Printf(" >>>>>>>> %s  \n ************* You Win ************* \n", randWord)
}

// randomWord is return a word randomly.
func randomWord() string {
	randseed := rand.NewSource(time.Now().UnixNano())
	gen := rand.New(randseed).Intn(219469) // length of the line
	var word string

	data, err := ioutil.ReadFile("english-words.txt")
	if err == nil {
		a := string(data)
		s := strings.Split(a, "\n")
		word = s[gen]
	}

	return strings.TrimSpace(word)
}

// guessing is return the correct statement.
func guessing(word string, guessLetter map[rune]bool) {
	for _, ch := range word {
		switch {
		case guessLetter[unicode.ToLower(rune(ch))]:
			fmt.Printf("%c", ch)
		default:
			fmt.Print("_")
		}
		fmt.Print(" ")
	}

	fmt.Println()
}

// readInput read one letter from the user.
func readInput() string {
	fmt.Print("> ")
	r := bufio.NewReader(os.Stdin)
	input, err := r.ReadString('\n')
	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(input)
}

// checkGuessing is response to check letter in the word.
func checkGuessing(word string, letter rune) bool {
	return strings.ContainsRune(word, letter)
}

// wordGuessed is checking the user input has guessed correctly or not.
func wordGuessed(word string, letter map[rune]bool) bool {
	for _, ch := range word {
		if !letter[ch] {
			return false
		}
	}

	return true
}
