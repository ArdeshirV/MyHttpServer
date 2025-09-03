package main

import (
	"fmt"
	"net/http"
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
