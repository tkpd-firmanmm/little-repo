package errors

type ClientError struct {
	err string
}

func (c *ClientError) Error() string {
	return c.err
}

func NewClientError(err string) *ClientError {
	return &ClientError{
		err: err,
	}
}
