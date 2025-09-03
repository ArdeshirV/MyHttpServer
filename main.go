package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	fmt.Println(Prompt("My HTTP Server"))
}

type MyHandler func(server RequestResponse)

type MyServer interface {
  GET(route string, handler MyHandler)
  POST(route string, handler MyHandler)
  PATCH(route string, handler MyHandler)
  PUT(route string, handler MyHandler)
  DELETE(route string, handler MyHandler)
}

type RequestResponse interface {
  Request() *http.Request 
  Response() http.ResponseWriter
}

type myserver struct {
  routes []Route
}

type Route struct {
  parts []string
  handler MyHandler
  method string 
}

func NewServer() *myserver {
  return &myserver{
    routes: []Route{},
  }
}

func (s *myserver) route(method, route string, handler MyHandler) {
  s.routes = append(s.routes, Route{
    parts: strings.Split(route, "/"),
    handler: handler,
    method: method,
  })
}

func (s *myserver) GET(route string, handler MyHandler) {
  s.route(http.MethodGet, route, handler)
}

func (s *myserver) POST(route string, handler MyHandler) {
  s.route(http.MethodPost, route, handler)
}

func (s *myserver) PATCH(route string, handler MyHandler) {
  s.route(http.MethodPatch, route, handler)
}














