package rps

import (
	"math/rand"
	"time"
)

const (
	// ROCK beats scissors. (scissors + 1) % 3 = 0
	ROCK = 0
	//PAPER beats rock. (rock + 1) % 3 = 1
	PAPER = 1
	//SCISSORS beats paper. (paper + 1) % 3 = 2
	SCISSORS = 2
	//PLAYERWINS indicate that a player wins the round
	PLAYERWINS = 1
	//COMPUTERWINS indicate that a computer wins the round
	COMPUTERWINS = 2
	//DRAW indicate that a no one win the round
	DRAW = 3
)

//PlayRound get a player value and returns
// (Winner Value,Computer choice, round Result)
func PlayRound(playerValue int) (winner int, computerChoice string, roundResult string) {
	rand.Seed(time.Now().UnixNano())
	computerValue := rand.Intn(3)

	switch computerValue {
	case ROCK:
		computerChoice = "Computer chose ROCK"
		break
	case PAPER:
		computerChoice = "Computer chose PAPER"
		break
	case SCISSORS:
		computerChoice = "Computer chose SCISSORS"
		break
	default:
	}

	if playerValue == computerValue {
		roundResult = "It's a draw"
		winner = DRAW
	} else if playerValue == (computerValue+1)%3 {
		roundResult = "Player wins!"
		winner = PLAYERWINS
	} else {
		roundResult = "Computer wins!"
		winner = COMPUTERWINS
	}

	return winner, computerChoice, roundResult
}
