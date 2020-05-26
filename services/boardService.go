package services

import (
	"fmt"
	"ticgame/components"
)

//BoardService :
type BoardService struct {
	Board *components.Board
}

//NewBoardService : Returns a board which will place a mark in the specified position
func NewBoardService(board *components.Board, position int, mark string) *BoardService {
	err := board.Cells[position].SetMark(mark)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Here is your board!")
	var i uint8
	for i = 0; i < (board.Size * board.Size); i++ {
		fmt.Printf("\t%s", board.Cells[i])
	}
	//returns BoardInstance
	return &BoardService{
		Board: board,
	}
}
