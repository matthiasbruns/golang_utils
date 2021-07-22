package router

type ErrorResponse struct {
	Message      string  `json:"message"`
	DebugMessage *string `json:"debugMessage,omitempty"`
}
