package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/dwisiswant0/slacksh/pkg/slacksh"
)

var port string

func init() {
	flag.StringVar(&port, "p", "", "Define port")
	flag.Parse()

	if port == "" {
		port = "8008"
	}

	fmt.Println("Listen port " + port)
}

func main() {
	http.HandleFunc("/", slacksh.Handler)
	log.Fatalln(http.ListenAndServe(":"+port, nil))
}
