// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package api

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type HTTPRequest struct {
	URL       string        `json:"url"`
	Method    HTTPMethod    `json:"method"`
	Body      *string       `json:"body"`
	Timestamp time.Time     `json:"timestamp"`
	Response  *HTTPResponse `json:"response"`
}

type HTTPResponse struct {
	StatusCode int     `json:"statusCode"`
	Body       *string `json:"body"`
}

type HTTPMethod string

const (
	HTTPMethodGet     HTTPMethod = "GET"
	HTTPMethodHead    HTTPMethod = "HEAD"
	HTTPMethodPost    HTTPMethod = "POST"
	HTTPMethodPut     HTTPMethod = "PUT"
	HTTPMethodDelete  HTTPMethod = "DELETE"
	HTTPMethodConnect HTTPMethod = "CONNECT"
	HTTPMethodOptions HTTPMethod = "OPTIONS"
	HTTPMethodTrace   HTTPMethod = "TRACE"
	HTTPMethodPatch   HTTPMethod = "PATCH"
)

var AllHTTPMethod = []HTTPMethod{
	HTTPMethodGet,
	HTTPMethodHead,
	HTTPMethodPost,
	HTTPMethodPut,
	HTTPMethodDelete,
	HTTPMethodConnect,
	HTTPMethodOptions,
	HTTPMethodTrace,
	HTTPMethodPatch,
}

func (e HTTPMethod) IsValid() bool {
	switch e {
	case HTTPMethodGet, HTTPMethodHead, HTTPMethodPost, HTTPMethodPut, HTTPMethodDelete, HTTPMethodConnect, HTTPMethodOptions, HTTPMethodTrace, HTTPMethodPatch:
		return true
	}
	return false
}

func (e HTTPMethod) String() string {
	return string(e)
}

func (e *HTTPMethod) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = HTTPMethod(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid HttpMethod", str)
	}
	return nil
}

func (e HTTPMethod) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
