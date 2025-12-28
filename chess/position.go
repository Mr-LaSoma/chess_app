package chess

// Position is the struct used for all the units on the chess boards
type Position struct {
	X int
	Y int
}

// NewPosition constructor for the position struct
func NewPosition(x, y int) *Position {
	return &Position{
		X: x,
		Y: y,
	}
}
