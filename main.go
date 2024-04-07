package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `
			<h1>Hola Mundo desde mi app en golang!</h1>
			<p>Bienvenido a mi página</p>
			<a href="/about">About Me</a>
		`)
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `
			<h1>Hola desde la página about</h1>
			<a href="/">Volver al home</a>
		`)
	})

	fmt.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
