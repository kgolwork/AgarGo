package main

import "AgarGo/server"

func main() {
	s := server.NewServer()
	s.Listen()
}
