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
	piecetype PieceType
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
	switch piece.piecetype {
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
func main() {
	board := Board{[8][8]Piece{}}
	PrintBoard(board)
}
