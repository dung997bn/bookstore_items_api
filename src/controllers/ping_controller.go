package controllers

import "net/http"

var (
	//PingController type
	PingController pingControllerInterface = &pingController{}
)

const (
	pong = "pong"
)

type pingControllerInterface interface {
	Ping(w http.ResponseWriter, r *http.Request)
}

type pingController struct{}

func (p *pingController) Ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(pong))
}
