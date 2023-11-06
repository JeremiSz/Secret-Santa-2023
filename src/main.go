package main

import (
	"log"
	"secret_santa_2023/src/server"
)

func main() {
	server := server.NewServer()
	log.Fatal(server.ListenAndServe())
}
