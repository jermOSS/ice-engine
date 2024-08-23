package defs

type GameFile struct {
	Rules *GameRules // Pointer to prevent copying of unnecessary data, rules are static so it shouldn't be done.
	Board BoardState
}

// Static struct of game Rules
type GameRules struct {
	MoveRule        MoveRuleType
	WhitePromotionY int64
	BlackPromotionY int64
}
