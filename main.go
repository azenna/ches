package main

import "github.com/azenna/ches/ches"

func main() {
	board := ches.Board{[8][8]ches.Piece{}}
	board.Init()
	board.Print()
}
