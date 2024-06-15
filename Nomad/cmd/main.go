package main

import (
	"bufio"
	"net/http"
	"os"
	"strings"

	"github.com/NathanPokharel/Nomad/views"
	"github.com/go-chi/chi"
)

var result [][]string

func Read() {
	file, err := os.Open("REGISTERY.txt")
	if err != nil {
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "-")
		result = append(result, parts)
	}
	if err := scanner.Err(); err != nil {
		panic("Error Reading File")
	}

}

func main() {
	Read()
	router := chi.NewMux()
	router.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		views.Home().Render(r.Context(), w)
	})
	router.Get("/products", func(w http.ResponseWriter, r *http.Request) {
		views.Products(result).Render(r.Context(), w)
	})
	router.Get("/checkout/{product_name}", func(w http.ResponseWriter, r *http.Request) {
		name := chi.URLParam(r, "product_name")
		views.CheckOut(name, result).Render(r.Context(), w)
	})
	router.Get("/snippets/{snippet_name}-{optional_values}", func(w http.ResponseWriter, r *http.Request) {
		snippet := chi.URLParam(r, "snippet_name")
		if snippet == "product" {
		}
	})
	http.ListenAndServe(":8080", router)
}
