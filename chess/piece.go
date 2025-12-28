package chess

// Piece is the interface used for all the chess pieces
type Piece interface {
	Draw()
	GetValidMoves() []Position
	Move() string
	ChessString() string
}
