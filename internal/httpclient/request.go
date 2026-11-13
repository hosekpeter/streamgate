package httpclient

import (
	"io"
	"net/http"
	"net/url"
)

type Request struct {
	Method  string
	BaseURL string
	Path    string
	Query   url.Values
	Header  http.Header
	Body    io.Reader
}
