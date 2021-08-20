package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", manejador)
	http.ListenAndServe(":8080", nil)
}

//manejador controla las peticiones
func manejador(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola, %s, !este es un servidor de prueba wiki", r.URL.Path)
}
