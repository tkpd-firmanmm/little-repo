package book

type RedisBookRepository struct{}

func (o *RedisBookRepository) GetByID(id uint) (*Book, error) {
	if id != 2 { //Assume only has 2
		return nil, nil
	}
	return &Book{
		ID:   2,
		Data: "Dummy Redis Book",
	}, nil
}

func NewRedisBookRepository() *RedisBookRepository {
	return &RedisBookRepository{}
}
