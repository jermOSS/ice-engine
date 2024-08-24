package engine

import (
	"slices"

	"github.com/jermOSS/ice-engine/pkg/defs"
	"github.com/jermOSS/ice-engine/pkg/pieces"
	"github.com/jermOSS/ice-engine/pkg/pieces/helpers"
	"github.com/jermOSS/ice-engine/pkg/utils"
)

// Evaluation weights for existence of piece.
var StaticEvalTable = [...]int32{
	1300, // Amazon
	500,  // ArchBishop
	350,  // Bishop
	220,  // Camel
	240,  // Centaur
	530,  // Chancellor
	240,  // Girrafe
	200,  // Guard
	340,  // Hawk
	10,   // King
	700,  // KnightRider
	230,  // Knight
	0,    // Obstacle
	90,   // Pawn
	900,  // Queen
	420,  // Rook
	0,    // RoyalCentaur
	0,    // RoyalQueen
	0,    // Void
	260,  // Zebra
}

var PawnDistanceTable = [...]int64{
	60,
	55,
	30,
	25,
	20,
	18,
	14,
	10,
	8,
	6,
	4,
	2,
	1,
	0,
}

func insertPieceSorted(slice []defs.Piece, piece defs.Piece, condition func(p defs.Piece) bool) []defs.Piece {
	for i, p := range slice {
		if condition(p) {
			return slices.Insert(slice, i, piece)
		}
	}
	return append(slice, piece)
}

func generateMaps(pieces []defs.Piece) (
	XYMap map[int64]map[int64]defs.Piece,
	XMap map[int64][]defs.Piece,
	YMap map[int64][]defs.Piece,
	IMap map[int64][]defs.Piece,
	JMap map[int64][]defs.Piece,
) {
	// Make sure everything is initiatet
	XYMap = make(map[int64]map[int64]defs.Piece)
	XMap = make(map[int64][]defs.Piece)
	YMap = make(map[int64][]defs.Piece)
	IMap = make(map[int64][]defs.Piece)
	JMap = make(map[int64][]defs.Piece)
	for _, piece := range pieces {
		if XYMap[piece.GetX()] == nil {
			XYMap[piece.GetX()] = make(map[int64]defs.Piece)
		}
		XYMap[piece.GetX()][piece.GetY()] = piece

		XMap[piece.GetX()] = insertPieceSorted(XMap[piece.GetX()], piece, func(p defs.Piece) bool {
			return p.GetY() > piece.GetY()
		})

		YMap[piece.GetY()] = insertPieceSorted(YMap[piece.GetY()], piece, func(p defs.Piece) bool {
			return p.GetX() > piece.GetX()
		})

		IMap[piece.GetI()] = insertPieceSorted(IMap[piece.GetI()], piece, func(p defs.Piece) bool {
			return p.GetJ() > piece.GetJ()
		})

		JMap[piece.GetJ()] = insertPieceSorted(JMap[piece.GetJ()], piece, func(p defs.Piece) bool {
			return p.GetI() > piece.GetI()
		})
	}
	return
}

// Ignores that some pieces can be captured and some not, should serve only as rough estimeate
func getSlideMobility(move defs.MovePattern, XMap map[int64][]defs.Piece, YMap map[int64][]defs.Piece) (evaluation int64) {
	slide := move.Pattern.(defs.MoveSlide)
	evaluation = maxMobilityAddition
	switch slide.Direction {
	case defs.Top:
		for _, piece := range slices.Backward(XMap[slide.Start.X]) {
			if piece.GetY() >= slide.Start.Y {
				evaluation = utils.IntMin(evaluation, piece.GetY()-slide.Start.Y)
			} else {
				break
			}
		}
	case defs.Bottom:
		for _, piece := range XMap[slide.Start.X] {
			if piece.GetY() <= slide.Start.Y {
				evaluation = utils.IntMin(evaluation, slide.Start.Y-piece.GetY())
			} else {
				break
			}
		}
	case defs.Left:
		for _, piece := range YMap[slide.Start.Y] {
			if piece.GetX() <= slide.Start.X {
				evaluation = utils.IntMin(evaluation, slide.Start.X-piece.GetX())
			} else {
				break
			}
		}
	case defs.Right:
		for _, piece := range slices.Backward(YMap[slide.Start.Y]) {
			if piece.GetX() >= slide.Start.X {
				evaluation = utils.IntMin(evaluation, piece.GetX()-slide.Start.X)
			} else {
				break
			}
		}
	default:
		panic("Unsupported direction.")
	}
	return utils.IntMin(maxMobilityAddition, evaluation)
}

func getJumpMobility(move defs.MovePattern, color defs.Color, XYMap map[int64]map[int64]defs.Piece) int64 {
	jump := move.Pattern.(defs.MoveJump)
	if XYMap[jump.Position.X] == nil || XYMap[jump.Position.X][jump.Position.Y] == nil {
		if move.MoveType == defs.CheckAndCapture {
			return 0
		} else {
			return 1
		}
	}

	piece := XYMap[jump.Position.X][jump.Position.Y]
	if piece.GetColor() != color && move.MoveType != defs.MoveOnly && piece.CanBeCaptured() {
		return 1
	} else {
		return 0
	}
}

func getDiagonalSlideMobility(move defs.MovePattern, IMap map[int64][]defs.Piece, JMap map[int64][]defs.Piece) (evaluation int64) {
	slide := move.Pattern.(defs.MoveDiagonalSlide)
	slide_start_I := slide.Start.GetI()
	slide_start_J := slide.Start.GetJ()
	evaluation = maxMobilityAddition
	switch slide.Direction {
	case defs.TopLeft:
		for _, piece := range slices.Backward(JMap[slide_start_J]) {
			if piece.GetI() >= slide_start_I {
				evaluation = utils.IntMin(evaluation, piece.GetI()-slide_start_I)
			} else {
				break
			}
		}
	case defs.BottomRight:
		for _, piece := range JMap[slide_start_J] {
			if piece.GetI() <= slide_start_I {
				evaluation = utils.IntMin(evaluation, slide_start_I-piece.GetI())
			} else {
				break
			}
		}
	case defs.TopRight:
		for _, piece := range slices.Backward(IMap[slide_start_I]) {
			if piece.GetJ() >= slide_start_J {
				evaluation = utils.IntMin(evaluation, piece.GetJ()-slide_start_J)
			} else {
				break
			}
		}
	case defs.BottomLeft:
		for _, piece := range IMap[slide_start_I] {
			if piece.GetJ() <= slide_start_J {
				evaluation = utils.IntMin(evaluation, slide_start_J-piece.GetJ())
			} else {
				break
			}
		}
	default:
		panic("Unsupported direction.")
	}
	return utils.IntMin(maxMobilityAddition, evaluation)
}

func getPawnJumpMobility(move defs.MovePattern, XYMap map[int64]map[int64]defs.Piece) int64 {
	jump := move.Pattern.(defs.MovePawnJump)

	var bonus int64 = 0
	if XYMap[jump.EnPassant.X] == nil || XYMap[jump.EnPassant.X][jump.EnPassant.Y] == nil {
		bonus += 1
		if XYMap[jump.Position.X] == nil || XYMap[jump.Position.X][jump.Position.Y] == nil {
			bonus += 2
		}
	}
	return bonus
}

func getMobilityBonus(
	piece defs.Piece,
	XYMap map[int64]map[int64]defs.Piece,
	XMap map[int64][]defs.Piece,
	YMap map[int64][]defs.Piece,
	IMap map[int64][]defs.Piece,
	JMap map[int64][]defs.Piece,
) (bonus int64) {

	for _, move := range piece.GetMoves() {
		switch move.Pattern.(type) {
		case defs.MoveSlide:
			bonus += getSlideMobility(move, XMap, YMap)
		case defs.MoveJump:
			bonus += getJumpMobility(move, piece.GetColor(), XYMap)
		case defs.MoveDiagonalSlide:
			bonus += getDiagonalSlideMobility(move, IMap, JMap)
		case defs.MovePawnJump:
			bonus += getPawnJumpMobility(move, XYMap)
		}
	}
	return bonus
}

// TODO: Add pawn protection bonus
func getPawnEvaluation(piece defs.Piece, rules defs.GameRules) int64 {
	if piece.GetColor() == defs.White {
		return PawnDistanceTable[utils.IntMin(14, utils.IntAbs(piece.GetY()-rules.WhitePromotionY)-1)]
	} else {
		return PawnDistanceTable[utils.IntMin(14, utils.IntAbs(piece.GetY()-rules.BlackPromotionY)-1)]
	}
}

const maxMobilityAddition int64 = 10 // Maximum eval that gets counted for mobility bonus
const mobilityWeight int64 = 7
const kingSafetyWeight int64 = 20

func Evaluate(gamefile *defs.GameFile) (eval int64) {
	XYMap, XMap, YMap, IMap, JMap := generateMaps(gamefile.Board.GetPieces())
	for _, piece := range gamefile.Board.GetPieces() {
		eval += int64(StaticEvalTable[piece.GetType()] * int32(piece.GetColor()))
		switch {
		case piece.GetType() == defs.Pawn:
			eval += getPawnEvaluation(piece, *gamefile.Rules) * mobilityWeight * int64(piece.GetColor())
		case piece.IsRoyal():
			eval += getMobilityBonus(piece, XYMap, XMap, YMap, IMap, JMap) * mobilityWeight * int64(piece.GetColor())

			queen := &pieces.Queen{
				Position: helpers.Position{
					X: piece.GetX(),
					Y: piece.GetY(),
				},
				ColoredPiece: helpers.ColoredPiece{
					Color: piece.GetColor() == defs.White,
				},
			}
			// King safety deduction
			eval -= getMobilityBonus(queen, XYMap, XMap, YMap, IMap, JMap) * kingSafetyWeight * int64(piece.GetColor())
		default:
			eval += getMobilityBonus(piece, XYMap, XMap, YMap, IMap, JMap) * int64(piece.GetColor())
		}
	}

	if gamefile.Board.WhosTurn {
		return eval
	} else {
		return eval * -1
	}
}
