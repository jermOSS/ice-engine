package pieces

import (
	"github.com/jermOSS/ice-engine/pkg/defs"
	"github.com/jermOSS/ice-engine/pkg/pieces/helpers"
)

type Knight struct {
	helpers.Position
	helpers.ColoredPiece
	helpers.Capturable
	helpers.NonRoyal
}

func (self *Knight) GetType() defs.PieceType {
	return defs.Knight
}

func (self *Knight) GetMoves() []defs.MovePattern {
	return []defs.MovePattern{
		{
			MoveType: defs.Regular,
			Pattern: defs.MoveJump{
				Position: defs.Position{
					X: self.X - 1,
					Y: self.Y + 2,
				},
			},
		},
		{
			MoveType: defs.Regular,
			Pattern: defs.MoveJump{
				Position: defs.Position{
					X: self.X + 1,
					Y: self.Y + 2,
				},
			},
		},
		{
			MoveType: defs.Regular,
			Pattern: defs.MoveJump{
				Position: defs.Position{
					X: self.X - 2,
					Y: self.Y + 1,
				},
			},
		},
		{
			MoveType: defs.Regular,
			Pattern: defs.MoveJump{
				Position: defs.Position{
					X: self.X + 2,
					Y: self.Y + 1,
				},
			},
		},
		{
			MoveType: defs.Regular,
			Pattern: defs.MoveJump{
				Position: defs.Position{
					X: self.X - 2,
					Y: self.Y - 1,
				},
			},
		},
		{
			MoveType: defs.Regular,
			Pattern: defs.MoveJump{
				Position: defs.Position{
					X: self.X + 2,
					Y: self.Y - 1,
				},
			},
		},
		{
			MoveType: defs.Regular,
			Pattern: defs.MoveJump{
				Position: defs.Position{
					X: self.X - 1,
					Y: self.Y - 2,
				},
			},
		},
		{
			MoveType: defs.Regular,
			Pattern: defs.MoveJump{
				Position: defs.Position{
					X: self.X + 1,
					Y: self.Y - 2,
				},
			},
		},
	}
}
