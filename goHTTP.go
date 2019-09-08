package goHTTP

import (
	"net/http"
)

type ServerContext struct {
	routers [] *Router
}

type Server interface {
	RegisterRouter(router Router)
	ServeHTTP(http.ResponseWriter, *http.Request)
}

func NewServer() Server{
	return &ServerContext{
		routers: []*Router{},
	}
}

func (s *ServerContext) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for i := 0; i < len(s.routers); i++  {
		router := s.routers[i]
		if string(router.Method) == r.Method && isRouter(r.URL.Path, router.Path) {
			router.Handle(w, r)
		}
	}
}

func (s *ServerContext) RegisterRouter(router Router){
	s.routers = append(s.routers, &router)
}

type Method string

const (
	GET Method = "GET"
	POST Method = "POST"
	PUT Method = "PUT"
	PATCH Method = "PATCH"
	DELETE Method = "DELETE"
	HEAD Method = "HEAD"
	OPTION Method = "OPTION"
)

type Router struct {
	Path string
	Method Method
	Handle func(w http.ResponseWriter, r *http.Request)
}
