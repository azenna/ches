package ches

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
	PieceType PieceType
	Color     Color
}

func (piece Piece) Print() {
	s := " "
	if piece.Color == White {
		s = "W"
	} else if piece.Color == Black {
		s = "B"
	}
	t := " "
	switch piece.PieceType {
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
