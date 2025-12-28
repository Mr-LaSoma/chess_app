package piece

import "github.com/mr-lasoma/chessgo/board"

type Piece interface {
	Draw()
	GetValidMoves(cb board.Board) map[*board.Square]bool
	Move(cb board.Board)
}

type BasePiece struct {
	X int
	Y int
}

func (p *BasePiece) Move(cb board.Board) {
	return
}

func (p *BasePiece) GetValidMoves(cb board.Board) []board.Square {
	return nil
}

type Pawn struct {
	BasePiece
}
