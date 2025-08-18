package middleware

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
)

type WrapperWriter struct {
	http.ResponseWriter
	StatusCode int
}

func (w *WrapperWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.StatusCode = statusCode
}

func (ww *WrapperWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	hijacker, ok := ww.ResponseWriter.(http.Hijacker)
	if !ok {
		return nil, nil, fmt.Errorf("the wrapped ResponseWriter does not implement http.Hijacker")
	}
	return hijacker.Hijack()
}
