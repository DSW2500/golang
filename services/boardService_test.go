package services
import (
	"testing"
	"ticgame/components"
)

func TestNewBoardService(t *testing.T){
expected:= &BoardService{
	board: components.CreateBoard(3),
}
var size uint8
actual:= &Board{
		Cells: []*Cell{
			{Mark: NoMark},
			{Mark: NoMark}, {Mark: NoMark}, {Mark: NoMark}, {Mark: NoMark}, {Mark: NoMark}, {Mark: NoMark}, {Mark: NoMark}, {Mark: NoMark},},
		Size: 9,
	}
	var i uint8
	for i=0;i<size;i++{
	if actual.Board.Cells[i].Mark!=expected.Board.Cells[i].Mark {
		t.Error("Error")
	}
}
}