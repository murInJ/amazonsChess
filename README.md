# amazonsChess
amazonsChess in go

## contributor
MurInJ

## function
amazonsChess have follows functions as so far:
- play one round game quickly
- randomly play
- get chess history after one round
- use AI written by yourself
- colorful terminal and hidden option\
 ![](./preview.gif)

## install
`go get github.com/murInJ/amazonsChess`

## qiuck start
The following code can achieve the effect in the preview
```go
game := amazonsChess.Game{CurrentPlayer: 1}

game.Start(true)

```