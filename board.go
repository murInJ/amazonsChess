package main

import (
	"errors"
	"fmt"
	"github.com/fatih/color"
	"log"
	"math/rand"
	"time"
)

var DIR = []int{1, 10, 11, -1, -10, -11}

type State struct {
	Board         []int
	CurrentPlayer int
}

// NewState 3 6 30 39 60 69 93 96

func (s *State) GetActionSpace(loc int) ([]int, error) {
	if loc < 0 || loc >= 100 {
		return nil, errors.New("illegal loc(need 0~99)")
	}
	var actionSpace []int
	for i := 0; i < 8; i++ {
		loc_n := loc
		for loc_n+DIR[i] != 0 {
			loc_n += DIR[i]
			actionSpace = append(actionSpace, loc_n)
		}
	}
	return actionSpace, nil
}

func (s *State) GetValid() []ChessMove {
	var validChess []ChessMove
	for c := 0; c < 100; c++ {
		if s.Board[c] == s.CurrentPlayer {
			start := c
			endList, err := s.GetActionSpace(start)

			if err != nil {
				log.Fatal(err)
			}

			for _, end := range endList {
				obstacleList, err := s.GetActionSpace(end)

				if err != nil {
					log.Fatal(err)
				}

				for _, obstacle := range obstacleList {
					validChess = append(validChess, ChessMove{
						start:    start,
						end:      end,
						obstacle: obstacle,
					})
				}
			}

		}
	}
	return validChess
}

func (s *State) StateMove(move ChessMove) (*State, error) {
	validChess := s.GetValid()
	for _, validMove := range validChess {
		if validMove.Equal(move) {
			board := make([]int, 100)
			_ = Clone(board, s.Board)
			board[move.start] = 0
			board[move.end] = s.CurrentPlayer
			board[move.obstacle] = 2

			var currentPlayer int
			if s.CurrentPlayer == -1 {
				currentPlayer = 1
			} else {
				currentPlayer = -1
			}

			return &State{
				Board:         board,
				CurrentPlayer: currentPlayer,
			}, nil
		}
	}
	return nil, errors.New("illegal Move")
}

func (s *State) PrintState() {
	for index, value := range s.Board {
		fmt.Printf("%3d", value)
		fmt.Fprintf(color.Output, "%s", num2colorStr(value, index))
		if index%10 == 9 {
			fmt.Println()
		}
	}
	var playerStr string
	if s.CurrentPlayer == 1 {
		playerStr = color.New(color.FgHiRed).Sprintf("red")
	} else {
		playerStr = color.New(color.FgHiBlue).Sprintf("blue")
	}
	fmt.Printf("current player: %s\n", playerStr)
}

func (s *State) randomMove() (*State, error) {
	valid := s.GetValid()
	if len(valid) == 0 {
		return nil, errors.New("terminal state")
	}
	rand.Seed(time.Now().Unix()) //产生Seed
	move := valid[rand.Intn(len(valid))]
	state, err := s.StateMove(move)
	if err != nil {
		log.Fatal(err)
	}
	return state, nil
}
func num2colorStr(number int, loc int) string {
	if number == 0 {
		return color.New(color.FgHiWhite).Sprintf("%d", loc)
	} else if number == -1 {
		return color.New(color.FgHiBlue).Sprint("#")
	} else if number == 1 {
		return color.New(color.FgHiRed).Sprint("#")
	} else if number == 2 {
		return color.New(color.FgHiBlack).Sprint("#")
	}
	return "ERR"
}
