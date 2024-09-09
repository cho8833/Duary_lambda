package util

type ApplicationError interface {
	Error() string
	StatusCode() int
}

type BadRequestError struct {
	error
}

func NewBadRequestError(err error) ApplicationError {
	return BadRequestError{err}
}
func (err BadRequestError) StatusCode() int {
	return 400
}

type DBError struct {
	error
}

func NewDBError(err error) ApplicationError {
	return DBError{err}
}
func (err DBError) StatusCode() int {
	return 500
}
