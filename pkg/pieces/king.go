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
	if self.HasMoved {
		return []defs.MovePattern{
			{
				MoveType: defs.Regular,
				Pattern: defs.PatternJump{
					Position: defs.Position{
						X: self.X - 1,
						Y: self.Y - 1,
					},
				},
			},
			{
				MoveType: defs.Regular,
				Pattern: defs.PatternJump{
					Position: defs.Position{
						X: self.X,
						Y: self.Y - 1,
					},
				},
			},
			{
				MoveType: defs.Regular,
				Pattern: defs.PatternJump{
					Position: defs.Position{
						X: self.X + 1,
						Y: self.Y - 1,
					},
				},
			},
			{
				MoveType: defs.Regular,
				Pattern: defs.PatternJump{
					Position: defs.Position{
						X: self.X - 1,
						Y: self.Y,
					},
				},
			},
			{
				MoveType: defs.Regular,
				Pattern: defs.PatternJump{
					Position: defs.Position{
						X: self.X + 1,
						Y: self.Y,
					},
				},
			},
			{
				MoveType: defs.Regular,
				Pattern: defs.PatternJump{
					Position: defs.Position{
						X: self.X - 1,
						Y: self.Y + 1,
					},
				},
			},
			{
				MoveType: defs.Regular,
				Pattern: defs.PatternJump{
					Position: defs.Position{
						X: self.X,
						Y: self.Y + 1,
					},
				},
			},
			{
				MoveType: defs.Regular,
				Pattern: defs.PatternJump{
					Position: defs.Position{
						X: self.X + 1,
						Y: self.Y + 1,
					},
				},
			},
		}
	} else {
		return []defs.MovePattern{
			{
				MoveType: defs.Regular,
				Pattern: defs.PatternJump{
					Position: defs.Position{
						X: self.X - 1,
						Y: self.Y - 1,
					},
				},
			},
			{
				MoveType: defs.Regular,
				Pattern: defs.PatternJump{
					Position: defs.Position{
						X: self.X,
						Y: self.Y - 1,
					},
				},
			},
			{
				MoveType: defs.Regular,
				Pattern: defs.PatternJump{
					Position: defs.Position{
						X: self.X + 1,
						Y: self.Y - 1,
					},
				},
			},
			{
				MoveType: defs.Regular,
				Pattern: defs.PatternJump{
					Position: defs.Position{
						X: self.X - 1,
						Y: self.Y,
					},
				},
			},
			{
				MoveType: defs.Regular,
				Pattern: defs.PatternJump{
					Position: defs.Position{
						X: self.X + 1,
						Y: self.Y,
					},
				},
			},
			{
				MoveType: defs.Regular,
				Pattern: defs.PatternJump{
					Position: defs.Position{
						X: self.X - 1,
						Y: self.Y + 1,
					},
				},
			},
			{
				MoveType: defs.Regular,
				Pattern: defs.PatternJump{
					Position: defs.Position{
						X: self.X,
						Y: self.Y + 1,
					},
				},
			},
			{
				MoveType: defs.Regular,
				Pattern: defs.PatternJump{
					Position: defs.Position{
						X: self.X + 1,
						Y: self.Y + 1,
					},
				},
			},
			{
				MoveType: defs.MoveOnly,
				Pattern:  defs.PatternCastling{},
			},
		}
	}
}
