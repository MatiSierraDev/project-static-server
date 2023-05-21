package main

import (
	"fmt"
	"log"
	"net/http"
)

func HandleForm(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	fn := r.FormValue("name")
	fl := r.FormValue("lastname")

	fmt.Fprintf(w, "nombre: %s\n", fn)
	fmt.Fprintf(w, "Apellido: %s\n", fl)

}

func main() {
	sf := http.FileServer(http.Dir("./static"))
	http.Handle("/", sf)

	http.HandleFunc("/form", HandleForm)

	log.Printf("Runing server on port localhost:3030")
	err := http.ListenAndServe("127.0.0.1:3030", nil)
	if err != nil {
		log.Fatal("error:", err)
	}
}
