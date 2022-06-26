package amazonsChess

import (
	"bytes"
	"encoding/gob"
)

type ChessMove struct {
	start    int
	end      int
	obstacle int
}

func (m ChessMove) GetVal() []int {
	return []int{m.start,m.end,m.obstacle}
}

func NewChessMove(start ,end,obstacle int) *ChessMove{
	return &ChessMove{
		start:    start,
		end:      end,
		obstacle: obstacle,
	}
}

func (m ChessMove) Equal(move ChessMove) bool {
	if m.start == move.start && m.end == move.end && m.obstacle == move.obstacle {
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
