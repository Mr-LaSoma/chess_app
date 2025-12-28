package board

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/mr-lasoma/chessgo/piece"
)

// Square struct that contains every infos usefull for the render
type Square struct {
	x     int
	y     int
	Piece piece.Piece
}

func NewSquare(x, y int) *Square {
	return &Square{
		x: x,
		y: y,
	}
}

// DrawSquare draws the square on the x, y coord
func (s *Square) Draw(sizeX, sizeY int, color rl.Color) {
	rl.DrawRectangle(
		int32(s.x),
		int32(s.y),
		int32(sizeX),
		int32(sizeY),
		color,
	)
}

func (s *Square) GetLogicPos() (x int, y int) {
	return s.x / SquareSizeX, s.y / SquareSizeY
}

func (s *Square) GetTruePos() (x int, y int) {
	return s.x, s.y
}
