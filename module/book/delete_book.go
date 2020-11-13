package book

type DeleteBook struct {
	findBook *FindBook
}

func (d *DeleteBook) doDelete(book *Book) error {
	//Assume there is magic repo that perform delete operation by using book
	return nil
}

func (d *DeleteBook) DeleteByID(id uint) error {
	book, err := d.findBook.Find(id) //Must verify whether or not book exist
	if err != nil {
		return err
	}
	if err := d.doDelete(book); err != nil {
		return err
	}
	return nil
}

func NewDeleteBook(findBook *FindBook) *DeleteBook {
	return &DeleteBook{
		findBook: findBook,
	}
}
