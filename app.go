package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", HelloServer)
	http.HandleFunc("/basic", CustomServer)
	http.HandleFunc("/jejexime", CustomServer)

	http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

func CustomServer(w http.ResponseWriter, r *http.Request) {
	// modificar por el archivo que quieren mostrar
	dat, err := os.ReadFile("files/" + r.URL.Path[1:] + ".html")
	if err != nil {
		fmt.Fprintf(w, "No se pudo encontrar el archivo, revise la ruta")
		return
	}
	fmt.Fprintf(w, string(dat))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
