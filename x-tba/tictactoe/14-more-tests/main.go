package main

import "fmt"

/*
~ TICTACTOE GAME IN GO ~
+ This example uses the very basics of the Go language.
+ The goal is learning all the basics.
*/

const maxTurns = 9

var (
	won, tie bool      // is there any winner or a tie?
	turn     int       // total valid turns played
	player   = player1 // current player

	cells     [maxTurns]string // used to draw the board: contains the players' moves
	lastPos   int              // last played position
	wrongMove bool             // was the last move wrong?
)

// main is only responsible for the game loop, that's it.
func main() {
	printBoard()
	wait()

	for nextTurn() {
		wait()
	}
}

// nextTurn prints the board for the next turn and checks for the winning conditions.
// if win or tie: returns false, otherwise true.
func nextTurn() bool {
	play()
	printBoard()

	fmt.Printf("\n>>> PLAYER %q PLAYS to %d\n", player, lastPos+1)

	// the switch below is about winning and tie conditions.
	// so it is good have checkWinOrTie() as a simple statement.
	// totally optional.
	switch checkWinOrTie(); {
	default:
		switchPlayer()
		printStatus()

	case wrongMove:
		fmt.Printf(">>> CELL IS OCCUPIED: PLAY AGAIN!\n")
		wrongMove = false // reset for the next turn

	case won, tie:
		if won {
			fmt.Println(">>> WINNER:", player)
		} else {
			fmt.Println(">>> TIE!")
		}
		return false
	}
	return true
}

func wait() {
	fmt.Println()
	fmt.Scanln()
}

// printStatus prints the current status of the game
// it cannot access to the names (vars, consts, etc) inside any other func
func printStatus() {
	fmt.Println()

	progress := (1 - (float64(turn) / maxTurns)) * 100
	fmt.Printf("Current Turn           : %d\n", turn)
	fmt.Printf("Is there a winner      : %t\n", won)
	fmt.Printf("Turns left             : %.1f%%\n", progress)
}
