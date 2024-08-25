package pieces

import (
	"github.com/jermOSS/ice-engine/pkg/defs"
	"github.com/jermOSS/ice-engine/pkg/pieces/helpers"
)

type Rook struct {
	helpers.PositionHasMoved
	helpers.ColoredPiece
	helpers.Capturable
	helpers.NonRoyal
}

func (self *Rook) GetType() defs.PieceType {
	return defs.Rook
}

func (self *Rook) GetMoves() []defs.MovePattern {
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
	}
}
