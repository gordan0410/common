package dto

import (
	"net/http"
	"time"
)

type SendApiRequest struct {
	Method  string
	Url     string
	Header  map[string]string
	Body    []byte
	Timeout time.Duration
}

type SendApiResponse struct {
	HttpResponse *http.Response
	Data         []byte
	CostTime     time.Duration
}
