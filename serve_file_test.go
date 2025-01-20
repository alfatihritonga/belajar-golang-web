package belajar_golang_web

import (
	_ "embed"
	"fmt"
	"net/http"
	"testing"
)

func ServeFile(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")

	switch page {
	case "exists":
		http.ServeFile(w, r, "./resources/exists.html")
	default:
		http.ServeFile(w, r, "./resources/notfound.html")
	}
}

func Test_ServeFileServer(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(ServeFile),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

//go:embed resources/exists.html
var resourceExists string

//go:embed resources/notfound.html
var resourceNotFound string

func ServeFileEmbed(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")

	switch page {
	case "exists":
		fmt.Fprint(w, resourceExists)
	default:
		fmt.Fprint(w, resourceNotFound)
	}
}
func Test_ServeFileEmbed(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(ServeFileEmbed),
	}

	fmt.Println("Server running http://localhost:8080")
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
