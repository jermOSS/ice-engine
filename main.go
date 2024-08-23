package main

import (
	"fmt"

	"github.com/jermOSS/ice-engine/pkg/defs"
	"github.com/jermOSS/ice-engine/pkg/engine"
)

func main() {
	board, err := engine.LoadPackage("P1,2+|P2,2+|P3,2+|P4,2+|P5,2+|P6,2+|P7,2+|P8,2+|p1,7+|p2,7+|p3,7+|p4,7+|p5,7+|p6,7+|p7,7+|p8,7+|R1,1+|R8,1+|r1,8+|r8,8+|N2,1|N7,1|n2,8|n7,8|B3,1|B6,1|b3,8|b6,8|Q4,1|q4,8|K5,1+|k5,8+")
	if err != nil {
		fmt.Println(err)
		return
	}
	gamefile := defs.GameFile{
		Rules: &defs.GameRules{
			MoveRule:        50,
			WhitePromotionY: 8,
			BlackPromotionY: 1,
		},
		Board: board,
	}

	fmt.Println(engine.Evaluate(&gamefile))
}
