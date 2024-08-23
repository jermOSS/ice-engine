package defs

import (
	"errors"
	"slices"
)

// All possible game conclusions
type GameConclusion = uint8

const (
	None GameConclusion = iota
	BlackWin
	Draw
	WhiteWin
)

type BoardState struct {
	Pieces          []Piece
	EnPassantSquare *Position
	MoveRuleCounter MoveRuleType
	GameConclusion  GameConclusion
	WhosTurn        bool // true = white, false = black
}

// Type alias for modularity, this value might be changed because there is space left in BoardState alignment.
// This value should fill it out. uint8 should be enough 99% percent of the time but there is not downside to doing this.
type MoveRuleType = uint32

// Get counter for X move rule
func (self *BoardState) GetMoveRuleCounter() MoveRuleType {
	return self.MoveRuleCounter
}

// Has game finished?
func (self *BoardState) IsOver() bool {
	return self.GameConclusion != None
}

func (self *BoardState) GetPieces() []Piece {
	return self.Pieces
}

func (self *BoardState) GetGameConclusion() uint8 {
	return self.GameConclusion
}

// Gets position of roayl piece, only one royal piece is supported at the moment
func (self *BoardState) GetRoaylPosition(color Color) (Position, error) {
	for _, piece := range self.Pieces {
		if piece.IsRoyal() && piece.GetColor() == color {
			return piece.GetPosition(), nil
		}
	}

	return Position{}, errors.New("No roayl piece found.")
}

// Gets positions of roayl piece, doesn't care about
func (self *BoardState) GetRoyalPositions() []Position {
	positions := make([]Position, 0)
	for _, piece := range self.Pieces {
		if piece.IsRoyal() {
			positions = append(positions, piece.GetPosition())
		}
	}
	return positions
}

func (self *BoardState) IncrementMoveRuleCounter() {
	self.MoveRuleCounter += 1
}

// Used to backtrack positions
func (self *BoardState) DecrementMoveRuleCounter() {
	self.MoveRuleCounter -= 1
}

func (self *BoardState) ResetMoveRuleCounter() {
	self.MoveRuleCounter = 0
}

func (self *BoardState) SetMoveRuleCounter(num MoveRuleType) {
	self.MoveRuleCounter = num
}

func (self *BoardState) GetTurnData() bool {
	return self.WhosTurn
}

func (self *BoardState) GetEnPassantSquare() *Position {
	return self.EnPassantSquare
}

func (self *BoardState) RemoveEnPassantSquare() {
	self.EnPassantSquare = nil
}

func (self *BoardState) SetEnPassantSquare(x int64, y int64) {
	*self.EnPassantSquare = Position{X: x, Y: y}
}

// Removes piece from board.
// Assumes there is only one piece at one position.
func (self *BoardState) RemovePieceBySquare(x int64, y int64) {
	for i, piece := range self.Pieces {
		position := piece.GetPosition()
		if position.X == x && position.Y == y {
			self.Pieces = slices.Delete(self.Pieces, i, i)
			return
		}
	}
}

// Adds a piece. Very naive should always check if adding piece is legal.
func (self *BoardState) AddPiece(piece Piece) {
	self.Pieces = append(self.Pieces, piece)
}

func (self *BoardState) EndGame(conclusion GameConclusion) {
	self.GameConclusion = conclusion
}
