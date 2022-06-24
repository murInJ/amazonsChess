package amazonsChess

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/fatih/color"
	"log"
)

type Game struct {
	CurrentPlayer int
	CurrentState  *State
	winner        int
	record        []string
	Ai1Handler    func(*State) ChessMove
	Ai2Handler    func(*State) ChessMove
}

func (g *Game) NewGame(currentPlayer int) (*Game, error) {
	if currentPlayer != -1 && currentPlayer != 1 {
		return nil, errors.New("wrong currentPlayer(need -1 or 1)")
	}
	board := NewBoard()
	var record []string
	return &Game{
		CurrentPlayer: currentPlayer,
		CurrentState: &State{
			Board:         board,
			CurrentPlayer: currentPlayer,
		},
		winner: 0,
		record: record,
	}, nil
}

func (g *Game) Reset(currentPlayer int) error {
	if currentPlayer != -1 && currentPlayer != 1 {
		return errors.New("wrong currentPlayer(need -1 or 1)")
	}
	g.CurrentPlayer = currentPlayer
	g.CurrentState = &State{
		Board:         NewBoard(),
		CurrentPlayer: currentPlayer,
	}
	g.winner = 0
	return nil
}

func (g *Game) GameOver() bool {
	validNum := len(g.CurrentState.GetValid())
	if validNum == 0 {
		return true
	} else {
		return false
	}
}

func (g *Game) LogGenerate() ([]byte, error) {
	var oneLog Log
	if g.GameOver() {
		oneLog = Log{
			GameState: *g.CurrentState,
			Status:    1,
			Winner:    g.winner,
		}
	} else {
		oneLog = Log{
			GameState: *g.CurrentState,
			Status:    0,
			Winner:    0,
		}
	}

	logJson, err := json.Marshal(oneLog)
	if err != nil {
		return nil, err
	}
	return logJson, nil
}

func (g *Game) GetMove(state *State) ChessMove {
	if g.CurrentPlayer == -1 {
		if g.Ai1Handler == nil {
			return ChessMove{}
		}
		return g.Ai2Handler(state)
	} else {
		if g.Ai2Handler == nil {
			return ChessMove{}
		}
		return g.Ai2Handler(state)
	}
}

func (g *Game) Start(isShow bool) [][]byte {
	var record [][]byte
	var logJson []byte
	var err error

	err = g.Reset(g.CurrentPlayer)
	if err != nil {
		log.Fatal(err)
	}

	logJson, err = g.LogGenerate()
	if err != nil {
		log.Fatal(err)
	}
	record = append(record, logJson)

	for !g.GameOver() {
		var err error
		move := g.GetMove(g.CurrentState)
		if move.Equal(ChessMove{}) {
			g.CurrentState, _ = g.CurrentState.randomMove()
		} else {
			g.CurrentState, err = g.CurrentState.StateMove(move)
			if err != nil {
				log.Fatal(err)
			}
		}
		if isShow {
			g.CurrentState.PrintState()
		}
		logJson, err = g.LogGenerate()
		if err != nil {
			log.Fatal(err)
		}
		record = append(record, logJson)
	}

	var playerStr string
	if g.CurrentPlayer == -1 {
		g.winner = 1
		playerStr = color.New(color.FgHiRed).Sprintf("red")
	} else {
		g.winner = -1
		playerStr = color.New(color.FgHiBlue).Sprintf("blue")
	}
	fmt.Printf("winner is: %s\n", playerStr)

	logJson, err = g.LogGenerate()
	if err != nil {
		log.Fatal(err)
	}
	record = append(record, logJson)

	return record
}

func NewBoard() []int {
	board := make([]int, 100)
	board[3] = -1
	board[6] = -1
	board[30] = -1
	board[39] = -1
	board[60] = 1
	board[69] = 1
	board[93] = 1
	board[96] = 1
	return board
}
