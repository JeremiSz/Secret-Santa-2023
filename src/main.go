package main

import (
	"log"
	"os"
	"secret_santa_2023/src/server"
)

func main() {
	var port string
	if len(os.Args) != 2 {
		port = "8080"
	} else {
		port = os.Args[1]
	}
	server := server.NewServer(port)
	log.Fatal(server.ListenAndServe())
}
