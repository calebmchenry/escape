package main

import "webl-fun/pkg/server"

type Tick struct {
	ID      int    `json:"id"`
	Message string `json:"message"`
}

func main() {
	s := server.NewSocketServer()
	s.Serve()
}
