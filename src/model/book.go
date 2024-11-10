package model

import "fmt"

type Book struct {
	Title string
	Pages int
	Id    uint32
}

func (b Book) String() string {
	return fmt.Sprintf("Книга: %s", b.Title)
}
