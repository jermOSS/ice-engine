package pieces

import (
	"github.com/jermOSS/ice-engine/pkg/defs"
	"github.com/jermOSS/ice-engine/pkg/pieces/helpers"
)

type Queen struct {
	helpers.Position
	helpers.ColoredPiece
	helpers.Capturable
	helpers.NonRoyal
}

func (self *Queen) GetType() defs.PieceType {
	return defs.Queen
}

func (self *Queen) GetMoves() []defs.MovePattern {
	return []defs.MovePattern{
		{
			MoveType: defs.Regular,
			Pattern: defs.PatternSlide{
				Start: defs.Position{
					X: self.X,
					Y: self.Y + 1,
				},
				Direction: defs.Top,
			},
		},
		{
			MoveType: defs.Regular,
			Pattern: defs.PatternSlide{
				Start: defs.Position{
					X: self.X,
					Y: self.Y - 1,
				},
				Direction: defs.Bottom,
			},
		},
		{
			MoveType: defs.Regular,
			Pattern: defs.PatternSlide{
				Start: defs.Position{
					X: self.X - 1,
					Y: self.Y,
				},
				Direction: defs.Left,
			},
		},
		{
			MoveType: defs.Regular,
			Pattern: defs.PatternSlide{
				Start: defs.Position{
					X: self.X + 1,
					Y: self.Y,
				},
				Direction: defs.Right,
			},
		},
		{
			MoveType: defs.Regular,
			Pattern: defs.PatternDiagonalSlide{
				Start: defs.Position{
					X: self.X - 1,
					Y: self.Y + 1,
				},
				Direction: defs.TopLeft,
			},
		},
		{
			MoveType: defs.Regular,
			Pattern: defs.PatternDiagonalSlide{
				Start: defs.Position{
					X: self.X + 1,
					Y: self.Y + 1,
				},
				Direction: defs.TopRight,
			},
		},
		{
			MoveType: defs.Regular,
			Pattern: defs.PatternDiagonalSlide{
				Start: defs.Position{
					X: self.X - 1,
					Y: self.Y - 1,
				},
				Direction: defs.BottomLeft,
			},
		},
		{
			MoveType: defs.Regular,
			Pattern: defs.PatternDiagonalSlide{
				Start: defs.Position{
					X: self.X + 1,
					Y: self.Y - 1,
				},
				Direction: defs.BottomRight,
			},
		},
	}
}
