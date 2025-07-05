package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, World!uiuiui")
	})

	fmt.Println("Starting server at http://localhost:8080")
	http.ListenAndServe(":8080", r) // ← net/http の代わりに r を渡すaa
}
