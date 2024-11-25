package main

import "fmt"

type Color int

const (
	Black Color = iota
	White
	NoColor
)

type PieceType int

const (
	Pawn PieceType = iota
	Rook
	Knight
	Bishop
	Queen
	King
	NoPiece
)

type Piece struct {
	piecetype PieceType
	color     Color
}
 EmptySquare = Piece {NoPiece,NoColor}

func main() {
	fmt.Println("Hello world")
	fmt.Println("Hello world")

	for _ = range 10 {
		fmt.Println("in for loop")

	}
}

type Board struct {
	state [8][8]Piece
}
