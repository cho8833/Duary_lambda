package util

type ApplicationError interface {
	Error() string
	StatusCode() int
}

type BadRequestError struct {
}

func (err BadRequestError) Error() string {
	return "잘못된 요청입니다"
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
