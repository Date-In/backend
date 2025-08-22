package middleware

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
)

type WrapperWriter struct {
	http.ResponseWriter
	StatusCode    int
	headerWritten bool
}

func NewWrapperWriter(w http.ResponseWriter) *WrapperWriter {
	return &WrapperWriter{
		ResponseWriter: w,
		StatusCode:     http.StatusOK,
	}
}

func (w *WrapperWriter) WriteHeader(statusCode int) {
	if w.headerWritten {
		return
	}
	w.ResponseWriter.WriteHeader(statusCode)
	w.StatusCode = statusCode
	w.headerWritten = true
}

func (w *WrapperWriter) Write(b []byte) (int, error) {
	if !w.headerWritten {
		w.WriteHeader(http.StatusOK)
	}
	return w.ResponseWriter.Write(b)
}

func (ww *WrapperWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	hijacker, ok := ww.ResponseWriter.(http.Hijacker)
	if !ok {
		return nil, nil, fmt.Errorf("the wrapped ResponseWriter does not implement http.Hijacker")
	}
	return hijacker.Hijack()
}
