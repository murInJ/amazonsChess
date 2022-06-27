package amazonsChess

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/fatih/color"
)

type Game struct {
	CurrentPlayer int                    `json:"current_player,omitempty"`
	CurrentState  *State                 `json:"current_state,omitempty"`
	Winner        int                    `json:"winner,omitempty"`
	Ai1Handler    func(*State) ChessMove `json:"ai_1_handler,omitempty"`
	Ai2Handler    func(*State) ChessMove `json:"ai_2_handler,omitempty"`
}

func NewGame(currentPlayer int) (*Game, error) {
	if currentPlayer != -1 && currentPlayer != 1 {
		return nil, errors.New("wrong currentPlayer(need -1 or 1)")
	}
	board := NewBoard()
	return &Game{
		CurrentPlayer: currentPlayer,
		CurrentState: &State{
			Board:         board,
			CurrentPlayer: currentPlayer,
		},
		Winner: 0,
	}, nil
}

//Reset reset a game, must call it befor a round start if not by calling function start
func (g *Game) Reset(currentPlayer int) error {
	if currentPlayer != -1 && currentPlayer != 1 {
		return errors.New("wrong currentPlayer(need -1 or 1)")
	}
	g.CurrentPlayer = currentPlayer
	g.CurrentState = &State{
		Board:         NewBoard(),
		CurrentPlayer: currentPlayer,
	}
	g.Winner = 0
	return nil
}

func (g *Game) GameOver() bool {
	red := 0
	blue := 0
	for i := 0; i < 100; i++ {
		if g.CurrentState.Board[i] == 1 || g.CurrentState.Board[i] == -1 {
			row := i / 10
			col := i % 10
			for j := 0; j < 8; j++ {
				tmpRow := row + DIR[j][0]
				tmpCol := col + DIR[j][1]
				tmpLoc := tmpRow*10 + tmpCol
				if tmpRow >= 0 && tmpRow < 10 && tmpCol >= 0 && tmpCol < 10 && g.CurrentState.Board[tmpLoc] == 0 {
					if g.CurrentState.Board[i] == 1 {
						red++
					} else {
						blue++
					}

				}
			}

		}
	}
	if red == 0 {
		g.Winner = -1
	} else if blue == 0 {
		g.Winner = 1
	} else {
		return false
	}

	return true
}

func (g *Game) LogGenerate() ([]byte, error) {
	var oneLog Log
	if g.GameOver() {
		oneLog = Log{
			GameState: *g.CurrentState,
			Status:    1,
			Winner:    g.Winner,
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

	fmt.Print("\x1b7") // 保存光标位置 保存光标和Attrs <ESC> 7
	for !g.GameOver() {
		var err error
		move := g.GetMove(g.CurrentState)
		if move.Equal(ChessMove{}) {
			g.CurrentState, _ = g.CurrentState.RandomMove()
		} else {
			g.CurrentState, err = g.CurrentState.StateMove(move)
			if err != nil {
				log.Fatal(err)
			}
		}
		if isShow {

			fmt.Print("\x1b8")
			fmt.Print("\x1b[2k") // 清空当前行的内容 擦除线<ESC> [2K
			g.CurrentState.PrintState()

			time.Sleep(50 * time.Millisecond)
		}
		logJson, err = g.LogGenerate()
		if err != nil {
			log.Fatal(err)
		}
		record = append(record, logJson)
	}

	var playerStr string
	if g.Winner == 1 {
		playerStr = color.New(color.FgHiRed).Sprintf("red")
	} else {
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
