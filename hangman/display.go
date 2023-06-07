package hangman

import (
	"fmt"
)

func DrawWelcome() {
	fmt.Println("                                                   ")
	fmt.Println("---------------------------------------------------")
	fmt.Println("                                                   ")
	fmt.Println("                                                   ")
	fmt.Println("Bienvenue dans le HANGMAN *Ynov 2022 B1 Info Paris*")
	fmt.Println("                                                   ")
	fmt.Println("                                                   ")
	fmt.Println("---------------------------------------------------")
}
func Draw(g *Game, guess string) {
	drawTurns(g.TurnsLeft)
	drawState(g, guess)
}

func drawTurns(l int) {
	var draw string
	switch l {
	case 0:
		draw = `
    +---+  
    |   |  
    O   |  
   /|\  |  
   / \  |  
  R.I.P |  
  =========
		`

	case 1:
		draw = `
    +---+  
    |   |  
    O   |  
   /|\  |  
   /    |  
        |  
  =========
		`

	case 2:
		draw = `
    +---+  
    |   |  
    O   |  
   /|\  |  
        |  
        |  
  =========
		`

	case 3:
		draw = `
    +---+  
    |   |  
    O   |  
    |   |  
        |  
        |  
  ========= 
		`

	case 4:
		draw = `
    +---+  
    |   |  
    O   |  
        |  
        |  
        |  
  ========= 
		`

	case 5:
		draw = `
    +---+  
    |   |  
        |  
        |  
        |  
        |  
  ========= 
		`

	case 6:
		draw = `
    +---+  
        |  
        |  
        |  
        |  
        |  
  ========= 
		`

	case 7:
		draw = `
        |  
        |  
        |  
        |  
        |  
  ========= 
		`

	case 8:
		draw = `
		
		`
	}
	fmt.Println(draw)
}

func drawState(g *Game, guess string) {
	fmt.Print("Devinées: ")
	drawLetters(g.FoundLetters)

	fmt.Print("Utilisées: ")
	drawLetters(g.UsedLetters)

	switch g.State {
	case "correct":
		fmt.Print("Bien trouvé !  ")
	case "dejadevine":
		fmt.Printf("Lettre '%s' a déjà été utilisée ;( \n", guess)
	case "incorrect":
		fmt.Printf("Mauvais choix, '%s' n'est pas dans le mot\n")
	case "perdu":
		fmt.Printf("Perdu :(( Le mot était: ")
		drawLetters(g.Letters)
	case "gagne":
		fmt.Printf("GAGNE !!! Le mot était: ")
		drawLetters(g.Letters)
	}
}

func drawLetters(l []string) {
	for _, c := range l {
		fmt.Printf("%v", c)
	}
	fmt.Println()
}
