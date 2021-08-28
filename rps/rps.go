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
)

var winMessages = []string{
	"Good job!",
	"Nice work!",
	"You should buy a lottery ticket",
}

var loseMessages = []string{
	"Too bad!",
	"Try again!",
	"This is just not your day.",
}

var drawMessages = []string{
	"Great minds think alike.",
	"Uh oh. Try again.",
	"Nobody wins, but you can try again.",
}

//Round holds the result and choices of a round
type Round struct {
	Message        string `json:"message"`
	ComputerChoice string `json:"computer_choice"`
	RoundResult    string `json:"round_result"`
}

//PlayRound get a player value and returns
// (Winner Value,Computer choice, round Result)
func PlayRound(playerValue int) Round {
	rand.Seed(time.Now().UnixNano())
	computerValue := rand.Intn(3)
	var message string
	var computerChoice, roundResult string

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

	messageIndex := rand.Intn(3)

	if playerValue == computerValue {
		roundResult = "It's a draw"
		message = drawMessages[messageIndex]
	} else if playerValue == (computerValue+1)%3 {
		roundResult = "Player wins!"
		message = winMessages[messageIndex]
	} else {
		roundResult = "Computer wins!"
		message = loseMessages[messageIndex]
	}

	return Round{
		Message:        message,
		ComputerChoice: computerChoice,
		RoundResult:    roundResult,
	}
}
