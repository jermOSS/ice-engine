package defs

type MovePawnJump struct {
	From            Position
	EnPassantSquare Position
	To              Position
}

type MoveCastlings struct {
	KingFrom Position
	KingTo   Position
	RookFrom Position
	RookTo   Position
}

type MoveJump struct {
	From Position
	To   Position
}

type MoveCapture struct {
	From          Position
	To            Position
	PieceCaptured Piece
}

type MovePromotion struct {
	From  Position
	To    Position
	Piece Piece
}
