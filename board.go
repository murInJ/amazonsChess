package amazonsChess

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

var DIR = [][]int{{0, 1}, {1, 0}, {1, 1}, {0, -1}, {-1, 0}, {-1, -1}, {-1, 1}, {1, -1}}

type State struct {
	Board         []int
	CurrentPlayer int
}

// NewState 3 6 30 39 60 69 93 96 loc
func (s *State) Str() string {
	b, err := json.Marshal(s)
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}

func NewState(board *[]int, currentPlayer int) *State {
	return &State{
		Board:         *board,
		CurrentPlayer: currentPlayer,
	}
}

//GetActionSpace get action space of diffrent direction by a loc
func (s *State) GetActionSpace(loc int) ([]int, error) {
	if loc < 0 || loc >= 100 {
		return nil, errors.New("illegal loc(need 0~99)")
	}
	var actionSpace []int
	for i := 0; i < 8; i++ {
		locN := loc
		tempRow := locN / 10
		tempCol := locN % 10
		for {
			tempRow += DIR[i][0]
			tempCol += DIR[i][1]
			if tempRow < 0 || tempRow >= 10 || tempCol < 0 || tempCol >= 10 {
				break
			}

			locN = 10*tempRow + tempCol
			if s.Board[locN] != 0 {
				break
			}

			actionSpace = append(actionSpace, locN)
		}
	}
	//if len(actionSpace) == 0{
	//	println(loc)
	//	println("")
	//}
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

				if len(obstacleList) == 0 {
					obstacleList = append(obstacleList, start)
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
			_ = copy(board, s.Board)
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
	fmt.Printf("current player: %s \n\n", playerStr)
}

func (s *State) RandomMove() (*State, error) {
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
		return color.New(color.BgWhite).Sprintf("%4d", loc)
	} else if number == -1 {
		return color.New(color.BgHiBlue).Sprintf("%4d", loc)
	} else if number == 1 {
		return color.New(color.BgHiRed).Sprintf("%4d", loc)
	} else if number == 2 {
		return color.New(color.BgHiBlack).Sprintf("%4s", ".")
	}
	return "ERR"
}
