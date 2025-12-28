package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/mr-lasoma/chessgo/board"
)

var (
	BackgroundColor = rl.Color{
		R: 48,
		G: 46,
		B: 43,
		A: 255,
	}

	DarkBaseColor = rl.Color{
		R: 118,
		G: 150,
		B: 86,
		A: 255,
	}
	LightBaseColor = rl.Color{
		R: 238,
		G: 238,
		B: 210,
		A: 255,
	}

	Chessboard *board.Board
)

func main() {
	rl.InitWindow(1920, 1080, "Chessboard")
	rl.ToggleFullscreen()

	padding := rl.GetScreenHeight() - 400
	board.SetBoardConfig(padding, padding, (rl.GetScreenWidth()-padding)/2, (rl.GetScreenHeight()-padding)/2)

	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	Chessboard = board.NewBoard()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(BackgroundColor)
		Chessboard.Draw(DarkBaseColor, LightBaseColor)

		rl.EndDrawing()
	}
}
