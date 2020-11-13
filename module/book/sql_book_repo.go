package book

import "database/sql"

type SQLBookRepository struct {
	dummyReq *sql.DB
}

func (o *SQLBookRepository) GetByID(id uint) (*Book, error) {
	if !(id == 2 || id == 3) { //Assume only accept 2 and 3
		return nil, nil
	}
	return &Book{
		ID:   id,
		Data: "Dummy SQL",
	}, nil
}

func NewSQLBookRepository(db *sql.DB) *SQLBookRepository {
	return &SQLBookRepository{
		dummyReq: db,
	}
}
