package service

type Error struct {
	Code    int
	Message string
}

func (e *Error) Error() string {
	return e.Message
}

var ErrUserNameRepeat = &Error{
	Code:    400,
	Message: "用户名重复",
}

var ErrUserNotFound = &Error{
	Code:    404,
	Message: "用户不存在",
}
