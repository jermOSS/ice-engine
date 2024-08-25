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

type PatternType = uint8

const (
	Regular         PatternType = iota // Check, capture, move
	CheckAndCapture                    // Check and capture
	MoveOnly                           // Only move, no capture, no capture
)

// Singular square jump
type PatternJump struct {
	Position Position
}

// Simple slide, only vertical or horizontal, no jumping
type PatternSlide struct {
	Start     Position
	Direction uint8
}

// Simple diagonal slide
type PatternDiagonalSlide struct {
	Start     Position
	Direction uint8
}

type PatternHippogonal struct {
	Start Position
	XStep uint8
	YStep uint8
}

/*
I considered adding fields to this like King position, Rook position etc. In the end I decided against it.
The amount of fields needed would be too big and for this specific move, modularity isn't that important.
Castling behaviour should be hardcoded.
*/
type PatternCastling struct{}

type PatternPawnJump struct {
	Position  Position
	EnPassant Position
}

// MovePattern of move, used to store behaivour of pieces.
type MovePattern struct {
	MoveType PatternType
	Pattern  interface{}
}
