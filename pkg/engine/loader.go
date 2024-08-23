package engine

import (
	"errors"
	"strconv"
	"strings"
	"unicode"

	"github.com/jermOSS/ice-engine/pkg/defs"
	"github.com/jermOSS/ice-engine/pkg/pieces"
	"github.com/jermOSS/ice-engine/pkg/pieces/helpers"
)

// Loads position in UICN (Universal Infinite Chess Notation).
// Pretty forgiving, doesn't check for order or dulpication of values.
func LoadPackage(str string) (defs.BoardState, error) {
	// Create empty BoardState
	board := defs.BoardState{
		Pieces:          make([]defs.Piece, 0),
		EnPassantSquare: nil,
		MoveRuleCounter: 0,
		WhosTurn:        true, // White
	}

	// Initiate buffers
	has_moved := true
	piece_str_buffer := ""
	x_buffer := ""
	x_complete := false
	y_buffer := ""
	for _, c := range str + "|" {
		switch {
		case unicode.IsDigit(c) || c == '-':
			if x_complete {
				y_buffer += string(c)
			} else {
				x_buffer += string(c)
			}
		case unicode.IsLetter(c):
			piece_str_buffer += string(c)
		case c == ',':
			x_complete = true
		case c == '+':
			has_moved = false
		case c == '|':
			// Assert that all required values have been provided.
			if x_buffer == "" || y_buffer == "" || piece_str_buffer == "" {
				return board, errors.New("Failed to provide required UICN values.")
			}

			// Convert x buffer to int64
			x_int, err := strconv.Atoi(x_buffer)
			x := int64(x_int)
			if err != nil {
				return board, errors.New("X variable in UICN is not valid number.")
			}

			// Convert y buffer to int64
			y_int, err := strconv.Atoi(y_buffer)
			y := int64(y_int)
			if err != nil {
				return board, errors.New("Y variable in UICN is not valid number.")
			}

			piece_str_uppercase := strings.ToUpper(piece_str_buffer)

			// If uppercase and original version is the same, piece is white
			// Otherwise it is black
			color := piece_str_uppercase == piece_str_buffer

			// Convert from letters to piece struct
			var piece defs.Piece
			switch piece_str_uppercase {
			case "K":
				piece = &pieces.King{
					PositionHasMoved: helpers.PositionHasMoved{
						X:        x,
						Y:        y,
						HasMoved: has_moved,
					},
					ColoredPiece: helpers.ColoredPiece{
						Color: color,
					},
				}
			case "P":
				piece = &pieces.Pawn{
					PositionHasMoved: helpers.PositionHasMoved{
						X:        x,
						Y:        y,
						HasMoved: has_moved,
					},
					ColoredPiece: helpers.ColoredPiece{
						Color: color,
					},
				}
			case "N":
				piece = &pieces.Knight{
					Position: helpers.Position{
						X: x,
						Y: y,
					},
					ColoredPiece: helpers.ColoredPiece{
						Color: color,
					},
				}
			case "B":
				piece = &pieces.Bishop{
					Position: helpers.Position{
						X: x,
						Y: y,
					},
					ColoredPiece: helpers.ColoredPiece{
						Color: color,
					},
				}
			case "R":
				piece = &pieces.Rook{
					PositionHasMoved: helpers.PositionHasMoved{
						X:        x,
						Y:        y,
						HasMoved: has_moved,
					},
					ColoredPiece: helpers.ColoredPiece{
						Color: color,
					},
				}
			case "Q":
				piece = &pieces.Queen{
					Position: helpers.Position{
						X: x,
						Y: y,
					},
					ColoredPiece: helpers.ColoredPiece{
						Color: color,
					},
				}
			default:
				return board, errors.New("Invalid piece in UICN string.")
			}

			// Save piece
			board.Pieces = append(board.Pieces, piece)

			// Cleanup buffers for next piece
			has_moved = true
			piece_str_buffer = ""
			x_buffer = ""
			x_complete = false
			y_buffer = ""
		default:
			return board, errors.New("Unsuppurted character in UICN string.")
		}
	}
	return board, nil
}
