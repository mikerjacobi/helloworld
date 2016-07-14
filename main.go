package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("starting server")

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("req rcv2 world\n")
		fmt.Fprintf(w, "bar2worldworld\n")
	})

	log.Fatal(http.ListenAndServe("0.0.0.0:"+os.Getenv("PORT"), nil))
	fmt.Println("server started")
}
