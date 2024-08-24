package defs

type Direction = uint8

const (
	Top Direction = iota
	Bottom
	Left
	Right
)

type DirectionDiagonal = uint8

const (
	TopLeft DirectionDiagonal = iota
	TopRight
	BottomLeft
	BottomRight
)

type MoveType = uint8

const (
	Regular         MoveType = iota // Check, capture, move
	CheckAndCapture                 // Check and capture
	MoveOnly                        // Only move, no capture, no capture
)

// Singular square jump
type MoveJump struct {
	Position Position
}

// Simple slide, only vertical or horizontal, no jumping
type MoveSlide struct {
	Start     Position
	Direction uint8
}

// Simple diagonal slide
type MoveDiagonalSlide struct {
	Start     Position
	Direction uint8
}

type MoveHippogonal struct {
	Start Position
	XStep uint8
	YStep uint8
}

/*
I considered adding fields to this like King position, Rook position etc. In the end I decided against it.
The amount of fields needed would be too big and for this specific move, modularity isn't that important.
Castling behaviour should be hardcoded.
*/
type MoveCastling struct{}

type MovePawnJump struct {
	Position  Position
	EnPassant Position
}

type MovePattern struct {
	MoveType MoveType
	Pattern  interface{}
}
