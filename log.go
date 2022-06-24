package amazonsChess

// Log
// @Param GameState 棋盘状态
// @Param status 记录类型 0:未结束 1:已结束
// @Param winner 胜利方 (1、-1)
type Log struct {
	GameState State `json:"game_state"`
	Status    int   `json:"status"`
	Winner    int   `json:"winner"`
}
