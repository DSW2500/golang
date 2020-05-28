package services

import (
	"errors"
	"ticgame/components"
)

//BoardService :
type BoardService struct {
	Board *components.Board
}

//NewBoardService : Returns a board which will place a mark in the specified position
func NewBoardService(board *components.Board) *BoardService {
	//returns BoardInstance
	return &BoardService{
		Board: board,
	}
}

//PutMarkInPosition :
func (boardService *BoardService) PutMarkInPosition(player *components.Player, position uint8) error {
	if boardService.CheckBoardIsFull() {
		return errors.New("Board is full, so no mark can be placed")
	}
	if boardService.Board.Cells[position].Mark == components.OMark || boardService.Board.Cells[position].Mark == components.XMark {
		return errors.New("Cell has already been marked, please try another cell")
	}
	boardService.Board.Cells[position].SetMark(player.Mark)
	return nil
}

//PrintBoard : Returns a slice of Cells which can then be printed in a loop
func (boardService *BoardService) PrintBoard() string {
	var out string = ""
	size := boardService.Board.Size * boardService.Board.Size
	var i uint8 = 0
	for ; i < size-1; i++ {
		out += boardService.Board.Cells[i].Mark
		if (i % boardService.Board.Size) == (boardService.Board.Size)-1 {
			out += "\n"
		}
	}
	return out
}

//CheckBoardIsFull : checks if the board is full
func (boardService *BoardService) CheckBoardIsFull() bool {
	flag := false
	var i uint8
	//Loops through all cells and checks for NoMark, if at the end of loop, all elements have been marked, then board is full
	for i = 0; i < boardService.Board.Size; i++ {
		if boardService.Board.Cells[i].Mark != components.NoMark {
			flag = true
		} else {
			flag = false
		}
	}
	return flag
}
