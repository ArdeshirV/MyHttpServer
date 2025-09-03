package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

func main() {
	fmt.Println(Prompt("My HTTP Server"))
	myServer := new(myserver)
	myServer.GET("/hello/a/b/c/:num", myHelloHandler)
	http.ListenAndServe("localhost:64640", myServer)
}

type MyHandler func(server requestResponse)

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
	parts   []string
	handler MyHandler
	method  string
}

func NewServer() *myserver {
	return &myserver{
		routes: []Route{},
	}
}

func (s *myserver) route(method, route string, handler MyHandler) {
	s.routes = append(s.routes, Route{
		parts:   strings.Split(route, "/"),
		handler: handler,
		method:  method,
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

func (s *myserver) PUT(route string, handler MyHandler) {
	s.route(http.MethodPut, route, handler)
}

func (s *myserver) DELETE(route string, handler MyHandler) {
	s.route(http.MethodDelete, route, handler)
}

type requestResponse struct {
	request  *http.Request
	response http.ResponseWriter
}

func (s *requestResponse) Request() *http.Request {
	return s.request
}

func (s *requestResponse) Response() http.ResponseWriter {
	return s.response
}

func (s *myserver) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	server := &requestResponse{request: r, response: w}
	urlParts := strings.Split(r.URL.Path, "/")
	for _, route := range s.routes {
		if len(route.parts) != len(urlParts) {
			continue
		}
		for _, urlPart := range urlParts {
			for _, urlParam := range route.parts {
				if urlPart != urlParam && !strings.HasPrefix(urlParam, ":") {
					continue
				}
			}
		}

		if r.Method != route.method {
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "Method not allowd")
			return
		}

		route.handler(*server)
		return
	}
}

func myHelloHandler(s requestResponse) {
	fmt.Fprint(s.Response(), "<b>This is a new world!</b>")
}

func logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		time.Sleep(time.Second)
		next.ServeHTTP(w, r)
		fmt.Println("Took:", time.Since(t))
	})
}
