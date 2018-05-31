package webserver

import "net"
import "net/http"

/*
WebServer A simple web server
*/
type WebServer struct {
	http.Server
	address string
}

/*
New Constructor for the web server
*/
func New(host, port string, h http.Handler) *WebServer {
	var ws WebServer

	ws.Addr = net.JoinHostPort(host, port)
	ws.Handler = h

	return &ws
}

/*
New Constructor for the web server
*/
func (s *WebServer) Start() error {
	return s.ListenAndServe()
}
