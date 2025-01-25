package main

import "webl-fun/pkg/server"

func main() {
	s := server.New()
	s.Serve()
}
