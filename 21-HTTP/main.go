package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Olá Mundo!"))
		fmt.Println("teste")
	})
	http.HandleFunc("/teste", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Olá Testeo!"))
		fmt.Println("teste")
	})

	log.Fatal(http.ListenAndServe(":5000", nil))
}
