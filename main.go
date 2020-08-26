package main

import (
	"fmt" 
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["key"]
	if !ok || len(keys[0]) < 1 {
		fmt.Fprintf(w, "Faltan parámetros en la URL")
		return
	}
	key := keys[0]
	fmt.Fprintf(w, "Bienvenido, ")
	fmt.Fprintf(w, "%s", key)
}

func main() {
	http.HandleFunc("/saludo", handler)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}
