package pieces

import (
	"github.com/jermOSS/ice-engine/pkg/defs"
	"github.com/jermOSS/ice-engine/pkg/pieces/helpers"
)

type Bishop struct {
	helpers.Position
	helpers.ColoredPiece
	helpers.Capturable
	helpers.NonRoyal
}

func (self *Bishop) GetType() defs.PieceType {
	return defs.Bishop
}

func (self *Bishop) GetMoves() []defs.MovePattern {
	return []defs.MovePattern{
		{
			MoveType: defs.Regular,
			Pattern: defs.MoveDiagonalSlide{
				Start: defs.Position{
					X: self.X - 1,
					Y: self.Y + 1,
				},
				Direction: defs.TopLeft,
			},
		},
		{
			MoveType: defs.Regular,
			Pattern: defs.MoveDiagonalSlide{
				Start: defs.Position{
					X: self.X + 1,
					Y: self.Y + 1,
				},
				Direction: defs.TopRight,
			},
		},
		{
			MoveType: defs.Regular,
			Pattern: defs.MoveDiagonalSlide{
				Start: defs.Position{
					X: self.X - 1,
					Y: self.Y - 1,
				},
				Direction: defs.BottomLeft,
			},
		},
		{
			MoveType: defs.Regular,
			Pattern: defs.MoveDiagonalSlide{
				Start: defs.Position{
					X: self.X + 1,
					Y: self.Y - 1,
				},
				Direction: defs.BottomRight,
			},
		},
	}
}
