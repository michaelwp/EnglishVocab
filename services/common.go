package services

type ResponseStatus int

const (
	ERROR ResponseStatus = iota
	SUCCESS
)

type JsonResponse struct {
	Status      ResponseStatus `json:"status"`
	Message     string         `json:"message"`
	Translation string         `json:"translation"`
	Data        any            `json:"data,omitempty"`
}
