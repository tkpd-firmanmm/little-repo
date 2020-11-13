package book

import (
	"errors"
)

type OtherServiceBookRepository struct{}

func (o *OtherServiceBookRepository) GetByID(id uint) (*Book, error) {
	if id == 5 { //Assume if given 5 it will crash
		return nil, errors.New("Repo Crash and Burn")
	}
	//Assume has everything
	return &Book{
		ID:   id,
		Data: "Dummy Other Service",
	}, nil
}

func NewOtherServiceBookRepository() *OtherServiceBookRepository {
	return &OtherServiceBookRepository{}
}
