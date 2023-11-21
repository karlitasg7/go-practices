package rps

import (
	"math/rand"
	"strconv"
)

const (
	ROCK     = 0
	PAPER    = 1
	SCISSORS = 2
)

type Round struct {
	Message           string `json:"message"`
	ComputerChoice    string `json:"computer_choice"`
	RoundResult       string `json:"round_result"`
	ComputerChoiseInt int    `json:"computer_choice_int"`
	ComputerScore     string `json:"computer_score"`
	PlayerScore       string `json:"player_score"`
}

var winMessages = []string{
	"Well done",
	"Nice job",
	"Good luck next time",
}

var loseMessages = []string{
	"Better luck next time",
	"Try again",
	"Bad luck",
}

var drawMessages = []string{
	"Draw",
	"Oh no. Try again",
	"Nobody wins, but you can try again",
}

var ComputerScore, PlayerScore int

func PlayRound(playerChoice int) Round {
	computerValue := rand.Intn(3)

	var computerChoice, roundResult string
	var computerChoiceInt int

	switch computerValue {
	case ROCK:
		computerChoice = "Computer select ROCK"
		computerChoiceInt = ROCK
	case PAPER:
		computerChoice = "Computer select PAPER"
		computerChoiceInt = PAPER
	case SCISSORS:
		computerChoice = "Computer select SCISSORS"
		computerChoiceInt = SCISSORS
	}

	messageInt := rand.Intn(3)

	var message string

	if playerChoice == computerChoiceInt {
		roundResult = "draw"
		message = drawMessages[messageInt]
	} else if playerChoice == (computerValue+1)%3 {
		PlayerScore++
		roundResult = "User win"
		message = winMessages[messageInt]
	} else {
		ComputerScore++
		roundResult = "Computer win"
		message = loseMessages[messageInt]
	}

	return Round{
		Message:           message,
		ComputerChoice:    computerChoice,
		RoundResult:       roundResult,
		ComputerChoiseInt: computerChoiceInt,
		ComputerScore:     strconv.Itoa(ComputerScore),
		PlayerScore:       strconv.Itoa(PlayerScore),
	}

}
