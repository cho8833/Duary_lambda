package util

type ApplicationError interface {
	Error() string
	StatusCode() int
}

type ValidateError struct {
	message string
}

func NewCustomApplicationError(message string) ValidateError {
	return ValidateError{message: message}
}
func (err ValidateError) Error() string {
	return err.message
}
func (err ValidateError) StatusCode() int {
	return 400
}

type BadRequestError struct {
}

func (err BadRequestError) Error() string {
	return "잘못된 요청입니다"
}
func (err BadRequestError) StatusCode() int {
	return 400
}

type DBSaveError struct {
}

func (err DBSaveError) Error() string {
	return "데이터를 저장하는데에 실패했습니다"
}
func (err DBSaveError) StatusCode() int {
	return 500
}

type InternalServerError struct{}

func (err InternalServerError) Error() string {
	return "알 수 없는 서버 오류"
}
func (err InternalServerError) StatusCode() int {
	return 500
}

type DBReadError struct{}

func (err DBReadError) Error() string   { return "정보를 가져오는데에 실패했습니다" }
func (err DBReadError) StatusCode() int { return 500 }

type DBUpdateError struct{}

func (err DBUpdateError) Error() string {
	return "정보를 업데이트하는데에 실패했습니다"
}
func (err DBUpdateError) StatusCode() int { return 500 }

type DBError struct{}

func (err DBError) Error() string   { return "작업에 실패했습니다" }
func (err DBError) StatusCode() int { return 500 }
