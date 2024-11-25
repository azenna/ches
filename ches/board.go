package ches

import "fmt"

type Board struct {
	State [8][8]Piece
}

func (board *Board) Init() {
	for i := range 8 {
		for j := range 8 {
			board.PlacePiece(i, j)
		}
	}
}

func (board *Board) Print() {
	for _, row := range board.State {
		for _, piece := range row {
			piece.Print()
		}
		fmt.Println()
	}

}

func (board *Board) PlacePiece(i int, j int) {
	pieceType := NoPiece
	pieceColor := NoColor

	if i <= 1 {
		pieceColor = White

	} else if i >= 6 {
		pieceColor = Black
	}

	if i == 1 || i == 6 {
		pieceType = Pawn

	} else if i == 0 || i == 7 {
		if j == 0 || j == 7 {
			pieceType = Rook
		} else if j == 1 || j == 6 {
			pieceType = Knight
		} else if j == 2 || j == 5 {
			pieceType = Bishop
		} else if (i == 0 && j == 3) || (i == 7 && j == 4) {
			pieceType = King
		} else if (i == 0 && j == 4) || (i == 7 && j == 3) {
			pieceType = Queen
		}
	}

	board.State[i][j] = Piece{pieceType, pieceColor}
}