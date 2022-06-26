package amazonsChess

import (
	"bytes"
	"encoding/gob"
)

type ChessMove struct {
	Start    int `json:"start,omitempty"`
	End      int `json:"end,omitempty"`
	Obstacle int `json:"obstacle,omitempty"`
}

func (m ChessMove) GetVal() []int {
	return []int{m.Start, m.End, m.Obstacle}
}

func NewChessMove(start, end, obstacle int) *ChessMove {
	return &ChessMove{
		Start:    start,
		End:      end,
		Obstacle: obstacle,
	}
}

func (m ChessMove) Equal(move ChessMove) bool {
	if m.Start == move.Start && m.End == move.End && m.Obstacle == move.Obstacle {
		return true
	}
	return false
}

// Clone 完整复制数据
func Clone(a, b interface{}) error {
	buff := new(bytes.Buffer)
	enc := gob.NewEncoder(buff)
	dec := gob.NewDecoder(buff)
	if err := enc.Encode(a); err != nil {
		return err
	}
	if err := dec.Decode(b); err != nil {
		return err
	}
	return nil
}
