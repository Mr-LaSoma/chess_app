package chess

import (
	themes "github.com/mr-lasoma/chessgo/chess/boardthemes"
)

const (
	baseSquareCX = 8
	baseSquareCY = 8
)

var (
	chessboard *board = nil
)

type board struct {
	board [][]Square

	squareCountX int
	squareCountY int

	theme themes.Theme
}

// GetChessBoard returns the singleton of the chessboard
func GetChessBoard() *board {
	if chessboard == nil {
		chessboard = newChessBoard()
	}
	return chessboard
}

func (b *board) SetBasicConfig() {
	b.squareCountX = baseSquareCX
	b.squareCountY = baseSquareCY
}

// newChessBoard returns the chess board created with the basic config
func newChessBoard() *board {
	var chessb board

	board := make([][]Square, chessb.squareCountY)
	for y := range chessb.squareCountY {
		row := make([]Square, chessb.squareCountX)
		for x := range chessb.squareCountX {
			row[x] = *NewBaseSquare(x, y)
		}
		board[y] = row
	}

	chessb.board = board
	return &chessb
}
