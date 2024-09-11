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
