package book

import (
	"faulty/errors"
)

type FindBook struct {
	repo IBookRepository
}

func (f *FindBook) Find(id uint) (*Book, error) {
	if id <= 0 { //ID could be zero
		return nil, errors.NewClientError("Parameter ID is zero")
	}
	book, err := f.repo.GetByID(id)
	if err != nil {
		return nil, errors.NewInternalError(err)
	}
	return book, nil
}

func NewFindBook(repo IBookRepository) *FindBook {
	return &FindBook{
		repo: repo,
	}
}
