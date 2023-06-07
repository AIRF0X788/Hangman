package main

import (
	"fmt"
	"os"

	"training.go/hangman/hangman"
	"training.go/hangman/hangman/words"
)

func main() {

	err := words.Load("words.txt")
	if err != nil {
		fmt.Printf("Bibliothèque introuvable !.... %v\n", err)
		os.Exit(1)

	}

	g := hangman.New(8, words.PickWord())
	
	hangman.DrawWelcome()

	guess := ""
	for{
		hangman.Draw(g, guess)
		
		switch g.State{
		case"gagne", "perdu":
		os.Exit(0)
		}

		l, err := hangman.ReadGuess()
		if err != nil{
			fmt.Printf("Impossible de d'utilisé cette lettre: %v", err)
			os.Exit(1)
		}
		guess = l
		g.MakeAGuess(guess)
	}
}
