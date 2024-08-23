package defs

type Color = int8

const (
	Black   Color = -1
	Netural       = 0
	White         = 1
)

type PieceType = uint8

// Piece type, used to key into tables.
// It is important to update all tables depending on this when updating.
const (
	Amazon PieceType = iota
	ArchBishop
	Bishop
	Camel
	Centaur
	Chancellor
	Girrafe
	Guard
	Hawk
	King
	KnightRider
	Knight
	Obstacle
	Pawn
	Queen
	Rook
	RoayalCentaur
	RoyalQueen
	Void
	Zebra
)

type Piece interface {
	GetMoves() []MovePattern // Returns moves based on position, ignores surroundings
	GetPosition() Position
	GetColor() Color
	GetType() PieceType // 256 types supported
	CanBeCaptured() bool
	IsRoyal() bool // Is piece royal?
	Move(x int64, y int64)
	GetX() int64
	GetY() int64
	GetI() int64
	GetJ() int64
}
