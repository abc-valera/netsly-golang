package restHandler

import "net/http"

type Debug struct{}

func NewDebug() Debug {
	return Debug{}
}

func (Debug) Ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}
