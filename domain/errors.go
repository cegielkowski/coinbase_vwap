package domain

import "errors"

var (
	ErrInternalServerError     = errors.New("some internal error happened")
	ErrWebSocketConnection     = errors.New("failed to connect to websocket")
	ErrWebSocketConnectionLost = errors.New("websocket connection lost")
	ErrMarshalJson             = errors.New("failed to marshal json")
	ErrUnmarshalJson           = errors.New("failed to unmarshal json")
)
