package helpers

import "github.com/jermOSS/ice-engine/pkg/defs"

type ColoredPiece struct {
	Color bool
}

func (self *ColoredPiece) GetColor() defs.Color {
	if self.Color {
		return defs.White
	} else {
		return defs.Black
	}
}

type NeutralPiece struct {
	Color bool
}

func (self *NeutralPiece) GetColor() defs.Color {
	return defs.Netural
}
