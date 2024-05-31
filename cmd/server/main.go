package main

import "gosh/pkg/server"

const (
	host = "0.0.0.0"
	port = "23234"
)

func main() {
	server.Run(host, port)
}
