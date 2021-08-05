package main

import (
	"ObjectStorage/ch1/objects"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/objects/", objects.Handler)
	addr := os.Getenv("LISTEN_ADDRESS")
	fmt.Println(addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
