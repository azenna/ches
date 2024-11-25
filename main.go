package main

import "fmt"

type Color int

const (
	NoColor Color = iota
	White
	Black
)

type PieceType int

const (
	NoPiece PieceType = iota
	Rook
	Knight
	Bishop
	Queen
	King
	Pawn
)

type Piece struct {
	pieceType PieceType
	color     Color
}

var EmptySquare = Piece{NoPiece, NoColor}

type Board struct {
	state [8][8]Piece
}

func PrintBoard(board Board) {
	for _, row := range board.state {
		for _, piece := range row {
			PrintPiece(piece)
		}
		fmt.Println()
	}

}
func PrintPiece(piece Piece) {
	s := " "
	if piece.color == White {
		s = "W"
	} else if piece.color == Black {
		s = "B"
	}
	t := " "
	switch piece.pieceType {
	case Pawn:
		t = "P"
	case Rook:
		t = "R"
	case Knight:
		t = "N"
	case Bishop:
		t = "B"
	case King:
		t = "K"
	case Queen:
		t = "Q"
	}
	fmt.Printf("[%s%s]", s, t)
}
func PlacePiece(board *Board, i int, j int) {
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

	board.state[i][j] = Piece{pieceType, pieceColor}
}

func InitBoard(board *Board) {
	for i := range 8 {
		for j := range 8 {
			PlacePiece(board, i, j)
		}
	}
}

func main() {
	board := Board{[8][8]Piece{}}
	InitBoard(&board)
	PrintBoard(board)
}
