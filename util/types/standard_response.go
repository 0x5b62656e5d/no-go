package types

type StandardResponse[T any] struct {
	Data    T       `json:"data"`
	Message *string `json:"message"`
	Error   *string `json:"error"`
	Success bool    `json:"success"`
}
