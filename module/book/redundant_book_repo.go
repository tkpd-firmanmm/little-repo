package book

import (
	"errors"
	"strings"
)

type RedundantBookRepositoryParam struct {
	Repos []IBookRepository
}

type RedundantBookRepository struct {
	repos []IBookRepository
}

func (o *RedundantBookRepository) GetByID(id uint) (*Book, error) {
	errs := []string{}
	for _, repo := range o.repos {
		book, err := repo.GetByID(id)
		switch true {
		case err != nil:
			errs = append(errs, err.Error())
			fallthrough
		case book == nil:
			continue
		}
		return book, nil
	}
	return nil, errors.New("All strategy fail, " + strings.Join(errs, ", "))
}

func NewRedundantBookRepository(param *RedundantBookRepositoryParam) *RedundantBookRepository {
	return &RedundantBookRepository{
		repos: param.Repos,
	}
}
