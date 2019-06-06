package models

type Error struct {
	api     string
	message string
	userID  string
	err     error
	req     interface{}
}

func (e Error) Req() interface{} {
	return e.req
}

func (e Error) Err() error {
	return e.err
}

func (e Error) UserID() string {
	return e.userID
}

func (e Error) Message() string {
	return e.message
}

func (e Error) Api() string {
	return e.api
}

func NewError(api string, message string, userID string, err error, req interface{}) *Error {
	return &Error{api: api, message: message, userID: userID, err: err, req: req}
}