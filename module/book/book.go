package book

type Book struct {
	ID   uint
	Data string
}

type IBookRepository interface {
	GetByID(id uint) (*Book, error)
}
