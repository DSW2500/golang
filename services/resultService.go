package services

import "ticgame/components"

//ResultService :
type ResultService struct {
	BoardService *BoardService
	Result       string
}

//NewResultService :
func NewResultService(boardService *BoardService) *ResultService {
	return &ResultService{
		BoardService: boardService,
		Result:       "In Process",
	}
}

//GetResult :
func (resultService *ResultService) GetResult(player components.Player) string {
	// size := int(resultService.BoardService.Board.Size )
	board := resultService.BoardService.Board.Cells
	if board[0].Mark == board[4].Mark { // check the diagonal from left top to right bottom corner
		if board[4].Mark == board[8].Mark {
			if board[4].Mark != "-" {
				return "Winner! " + player.Name
			}
		}
	}
	if board[2].Mark == board[4].Mark { // check the diagonal from right top to left bottom corner
		if board[4].Mark == board[6].Mark {
			if board[6].Mark != "-" {
				return "Winner! " + player.Name
			}
		}
	}
	if board[0].Mark == board[3].Mark { // check the left column
		if board[3].Mark == board[6].Mark {
			if board[6].Mark != "-" {
				return "Winner! " + player.Name
			}
		}
	}
	if board[0].Mark == board[1].Mark { // check the top row
		if board[1].Mark == board[2].Mark {
			if board[0].Mark != "-" {
				return "Winner! " + player.Name
			}
		}
	}
	if board[2].Mark == board[5].Mark { // check right column
		if board[5].Mark == board[8].Mark {
			if board[2].Mark != "-" {
				return "Winner! " + player.Name
			}
		}
	}
	if board[6].Mark == board[7].Mark { // check bottom row
		if board[7].Mark == board[8].Mark {
			if board[8].Mark != "-" {
				return "Winner! " + player.Name
			}
		}
	}
	if board[1].Mark == board[4].Mark { // check middle column
		if board[4].Mark == board[7].Mark {
			if board[1].Mark != "-" {
				return "Winner! " + player.Name
			}
		}
	}
	if board[3].Mark == board[4].Mark { // check middle row
		if board[4].Mark == board[5].Mark {
			if board[3].Mark != "" {
				return "Winner! " + player.Name

			}
		}
	}
	//If Board is full and no winner has been declared, game ends in a draw!
	if resultService.BoardService.CheckBoardIsFull() {
		return "Draw!"
	}
	return "InProcess"
}

