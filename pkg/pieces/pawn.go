package pieces

import (
	"github.com/jermOSS/ice-engine/pkg/defs"
	"github.com/jermOSS/ice-engine/pkg/pieces/helpers"
)

type Pawn struct {
	helpers.PositionHasMoved
	helpers.ColoredPiece
	helpers.Capturable
	helpers.NonRoyal
}

func (self *Pawn) GetType() defs.PieceType {
	return defs.Pawn
}

func (self *Pawn) GetMoves() []defs.MovePattern {
	// Maybe this should be moved to move generator?
	if self.Color {
		if !self.HasMoved {
			return []defs.MovePattern{
				{
					MoveType: defs.CheckAndCapture,
					Pattern: defs.MoveJump{
						Position: defs.Position{
							X: self.X - 1,
							Y: self.Y + 1,
						},
					},
				},
				{
					MoveType: defs.CheckAndCapture,
					Pattern: defs.MoveJump{
						Position: defs.Position{
							X: self.X + 1,
							Y: self.Y + 1,
						},
					},
				},
				{
					MoveType: defs.MoveOnly,
					Pattern: defs.MovePawnJump{
						Position: defs.Position{
							X: self.X,
							Y: self.Y + 2,
						},
						EnPassant: defs.Position{
							X: self.X,
							Y: self.Y + 1,
						},
					},
				},
			}
		} else {
			return []defs.MovePattern{
				{
					MoveType: defs.CheckAndCapture,
					Pattern: defs.MoveJump{
						Position: defs.Position{
							X: self.X - 1,
							Y: self.Y + 1,
						},
					},
				},
				{
					MoveType: defs.CheckAndCapture,
					Pattern: defs.MoveJump{
						Position: defs.Position{
							X: self.X + 1,
							Y: self.Y + 1,
						},
					},
				},
				{
					MoveType: defs.MoveOnly,
					Pattern: defs.MoveJump{
						Position: defs.Position{
							X: self.X,
							Y: self.Y + 1,
						},
					},
				},
			}
		}
	} else {
		if !self.HasMoved {
			return []defs.MovePattern{
				{
					MoveType: defs.CheckAndCapture,
					Pattern: defs.MoveJump{
						Position: defs.Position{
							X: self.X + 1,
							Y: self.Y - 1,
						},
					},
				},
				{
					MoveType: defs.CheckAndCapture,
					Pattern: defs.MoveJump{
						Position: defs.Position{
							X: self.X - 1,
							Y: self.Y - 1,
						},
					},
				},
				{
					MoveType: defs.MoveOnly,
					Pattern: defs.MovePawnJump{
						Position: defs.Position{
							X: self.X,
							Y: self.Y - 2,
						},
						EnPassant: defs.Position{
							X: self.X,
							Y: self.Y - 1,
						},
					},
				},
			}
		} else {
			return []defs.MovePattern{
				{
					MoveType: defs.CheckAndCapture,
					Pattern: defs.MoveJump{
						Position: defs.Position{
							X: self.X + 1,
							Y: self.Y - 1,
						},
					},
				},
				{
					MoveType: defs.CheckAndCapture,
					Pattern: defs.MoveJump{
						Position: defs.Position{
							X: self.X - 1,
							Y: self.Y - 1,
						},
					},
				},
				{
					MoveType: defs.MoveOnly,
					Pattern: defs.MoveJump{
						Position: defs.Position{
							X: self.X,
							Y: self.Y - 1,
						},
					},
				},
			}
		}
	}
}
