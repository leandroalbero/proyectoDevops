package main

import (
    "fmt"
    "net/http"
)
func handler(w http.ResponseWriter, r *http.Request) {
    keys, ok := r.URL.Query()["key"]
    if !ok || len(keys[0]) < 1{
        fmt.Fprintf(w,"Faltan parámetros en la URL")
        return
    }
    key := keys[0]
    fmt.Fprintf(w,"Bienvenido, ")
    fmt.Fprintf(w,string(key))
}

func main() {
    http.HandleFunc("/", handler)
    err := http.ListenAndServe(":80", nil)
    if(err != nil){
        return
    }
}