package response

import (
	"encoding/json"
	liberr "middleware/lib/error"
	"net/http"
)

type Body struct {
	Meta       Meta        `json:"meta"`
	Data       interface{} `json:"data"`
	Pagination Pagination  `json:"pagination"`
}

type Meta struct {
	Code   int             `json:"code"`
	Errors []liberr.Errors `json:"errors"`
	Url    string          `json:"url"`
	Method string          `json:"method"`
}

type Pagination struct {
	Page   int64 `json:"page"`
}

type Headers map[string]string

func Header(w http.ResponseWriter, status int, message interface{}, headers Headers) {
	for k, v := range headers {
		w.Header().Set(k, v)
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(message)
}

func MountEnvelope(data interface{}, code int, r *http.Request, err []liberr.Errors) Body {
	return Envelope(liberr.ErrorsAPI{
		Errors: err,
		Url:    string(r.URL.Path),
		Method: string(r.Method),
	}, code, data)
}

func Envelope(errAPI liberr.ErrorsAPI, status int, data interface{}) Body {
	envelope := Body{}

	envelope.Meta.Errors = errAPI.Errors
	envelope.Meta.Url = errAPI.Url
	envelope.Meta.Method = errAPI.Method
	envelope.Meta.Code = status
	envelope.Data = data

	return envelope
}