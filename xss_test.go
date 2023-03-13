package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Berbeda dengan bahasa pemrogaman lain seperi PHP, pada golang template, masalah XSS sudah diatasi secara otomatis, golang template memiliki fitur Auto Escape, dimana dia bisa mendeteksi data yang perlu ditampilkan di template, jika mengandung tag-tag html atau script, secara otomatis akan di escape.

func TemplateAutoEscape(writer http.ResponseWriter, request *http.Request) {
	myTemplates.ExecuteTemplate(writer, "post.gohtml", map[string]interface{}{
		"Title": "Template Auto Escape",
		"Body":  "<p>Ini adalah Body <script>alert('anda di hek')</script></p>",
	})
}
func TestTemplateAutoEscape(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()
	TemplateAutoEscape(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
