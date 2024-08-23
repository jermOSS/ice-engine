package pieces

import (
	"github.com/jermOSS/ice-engine/pkg/defs"
	"github.com/jermOSS/ice-engine/pkg/pieces/helpers"
)

type King struct {
	helpers.PositionHasMoved
	helpers.ColoredPiece
	helpers.Uncapturable
	helpers.Royal
}

func (self *King) GetType() defs.PieceType {
	return defs.King
}

func (self *King) GetMoves() []defs.MovePattern {
	return []defs.MovePattern{
		{
			MoveType: defs.Regular,
			Pattern: defs.MoveJump{
				Position: defs.Position{
					X: self.X - 1,
					Y: self.Y - 1,
				},
			},
		},
		{
			MoveType: defs.Regular,
			Pattern: defs.MoveJump{
				Position: defs.Position{
					X: self.X,
					Y: self.Y - 1,
				},
			},
		},
		{
			MoveType: defs.Regular,
			Pattern: defs.MoveJump{
				Position: defs.Position{
					X: self.X + 1,
					Y: self.Y - 1,
				},
			},
		},
		{
			MoveType: defs.Regular,
			Pattern: defs.MoveJump{
				Position: defs.Position{
					X: self.X - 1,
					Y: self.Y,
				},
			},
		},
		{
			MoveType: defs.Regular,
			Pattern: defs.MoveJump{
				Position: defs.Position{
					X: self.X + 1,
					Y: self.Y,
				},
			},
		},
		{
			MoveType: defs.Regular,
			Pattern: defs.MoveJump{
				Position: defs.Position{
					X: self.X - 1,
					Y: self.Y + 1,
				},
			},
		},
		{
			MoveType: defs.Regular,
			Pattern: defs.MoveJump{
				Position: defs.Position{
					X: self.X,
					Y: self.Y + 1,
				},
			},
		},
		{
			MoveType: defs.Regular,
			Pattern: defs.MoveJump{
				Position: defs.Position{
					X: self.X + 1,
					Y: self.Y + 1,
				},
			},
		},
	}
}
