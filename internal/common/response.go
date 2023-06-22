package common

import (
	"net/http"

	"github.com/go-chi/render"
)

type ErrResponse struct {
	Code         string `json:"code"`
	ErrorMessage string `json:"message,omitempty"`
}

type Response struct {
	HTTPStatusCode int           `json:"-"`
	Status         string        `json:"status,omitempty"`
	Errors         []ErrResponse `json:"errors,omitempty"`
	Data           interface{}   `json:"data,omitempty"`
}

func (e *Response) Render(w http.ResponseWriter, req *http.Request) error {
	render.Status(req, e.HTTPStatusCode)
	return nil
}

func NewResponse(status string, statusCode int, errors []ErrResponse, data interface{}) *Response {
	return &Response{
		HTTPStatusCode: statusCode,
		Status:         status,
		Errors:         errors,
		Data:           data,
	}
}
