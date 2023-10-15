package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// nb stands for number of bytes
		nb, err := fmt.Fprintf(w, "hello browser")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("Bytes Written: %d", nb)
	})
	err := http.ListenAndServe("Localhost:8080", nil)
	log.Fatal(err)
}
