package util

type BadRequestError struct {
	error
}

func NewBadRequestError(err error) error {
	return BadRequestError{err}
}

type DBError struct {
	error
}

func NewDBError(err error) error {
	return DBError{err}
}
