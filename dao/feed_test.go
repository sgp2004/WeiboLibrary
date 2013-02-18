package dao

import "fmt"
import "testing"

func TestTimeLine(t *testing.T) {
	fmt.Println("-----start-----")
	fmt.Println(TimeLine(0, 10))
	fmt.Println("-----end-------")
	info := new(BorrowInfo)
	info.BookId = 29
	info.OwnerId = 32
	info.BorrowerId = 30
	CreateInfo(info)
	fmt.Println("-----after insert-----")
	infos := TimeLine(0, 10)
	for _, v := range infos {
		fmt.Println(v)
	}
}
