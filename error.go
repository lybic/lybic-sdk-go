package lybic

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e Error) Error() string {
	return "code: " + e.Code + ", message: " + e.Message
}
