package board

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	SquareCountX = 8
	SquareCountY = 8
)

var (
	BoardWidth  int
	BoardHeight int

	OffsetX int
	OffsetY int

	SquareSizeX int
	SquareSizeY int
)

func SetBoardConfig(width, height, offsetX, offsetY int) {
	BoardWidth = width
	BoardHeight = height

	OffsetX = offsetX
	OffsetY = offsetY

	SquareSizeX = BoardWidth / SquareCountX
	SquareSizeY = BoardHeight / SquareCountY
}

type Board struct {
	Squares [][]Square
}

func NewBoard() *Board {
	board := make([][]Square, SquareCountY)

	for y := 0; y < SquareCountY; y++ {
		row := make([]Square, SquareCountX)
		for x := 0; x < SquareCountX; x++ {
			row[x] = Square{
				x: OffsetX + (x * SquareSizeX),
				y: OffsetY + (y * SquareSizeY),
			}
		}
		board[y] = row
	}
	return &Board{
		Squares: board,
	}
}

func (b *Board) Draw(darkC, lightC rl.Color) {
	for y := range SquareCountY {
		for x := range SquareCountX {

			var color rl.Color
			if (y+x)%2 == 0 {
				color = darkC
			} else {
				color = lightC
			}

			b.Squares[y][x].Draw(SquareSizeX, SquareSizeY, color)
		}
	}
}
