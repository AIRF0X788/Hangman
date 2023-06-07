package hangman

import(
	"strings"
	"fmt"
	"bufio"
	"os"
)

var reader = bufio.NewReader(os.Stdin)

func ReadGuess() (guess string, err error){
	valid := false
	for !valid {
		fmt.Print("Quelle est votre lettre ? ")
		guess, err = reader.ReadString('\n')
		if err != nil {
			return guess, err
		}
		guess = strings.TrimSpace(guess)
		if len(guess) != 1{
			fmt.Printf("Mauvaise lettre entré ! lettre=%v, len%v\n", guess, len(guess))
			continue
		}
		valid = true
	}
	return guess, nil
}