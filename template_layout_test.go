package belajar_golang_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateLayout(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles(
		"./templates/index.gohtml",
		"./templates/header.gohtml",
		"./templates/content.gohtml",
		"./templates/footer.gohtml",
	))
	t.ExecuteTemplate(w, "index.gohtml", map[string]interface{}{
		"Title": "Template Layout",
		"Name":  "Fatih",
	})
}

func Test_TemplateLayout(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateLayout(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
