package chess

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	SquareSizeX = 50
	SquareSizeY = 50
)

// Square is the struct that composes the chess board
// The x and y of the square stores the logic
// position on the chessboard (A1 -> [0; 0])
type Square struct {
	*Position
	Piece *Piece
}

// NewBaseSquare returns a square from an x and a y coordinates (logic ones)
func NewBaseSquare(x, y int) *Square {
	return &Square{
		Position: NewPosition(x, y),
		Piece:    nil,
	}
}

// NewSquareFromPiece returns a square from an x and a y coordinates (logic ones)
// and a chess piece (which implements the chess.Piece interface)
func NewSquareFromPiece(x, y int, piece *Piece) *Square {
	return &Square{
		Position: NewPosition(x, y),
		Piece:    piece,
	}
}

// GetLogicPosition gives the logical position of the chess board (A1 -> 0, 0)
func (s *Square) GetLogicPosition() *Position {
	return &Position{
		X: s.X,
		Y: s.Y,
	}
}

// GetPhysicalPosition gives the true position in the 2D space
func (s *Square) GetPhysicalPosition() *Position {
	return &Position{
		X: s.X * SquareSizeX,
		Y: s.Y * SquareSizeY,
	}
}

// Draw draws the square on the screen
func (s *Square) Draw(offsetX, offsetY int, c rl.Color) {
	pos := s.GetPhysicalPosition()

	rl.DrawRectangle(
		int32(pos.X+offsetX),
		int32(pos.Y+offsetY),
		int32(SquareSizeX),
		int32(SquareSizeY),
		c,
	)
}
