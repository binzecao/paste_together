package common

type Error interface {
	error
	StatusCode() int
}

type AppError struct {
	Code int
	Err  error
}

func (se AppError) Error() string {
	return se.Err.Error()
}

func (se AppError) StatusCode() int {
	return se.Code
}
