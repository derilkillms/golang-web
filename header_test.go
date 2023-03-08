package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// Untuk menangkap request Header yang dikirim oleh client, kita bisa mengambilnya di Request.Header. Header mirip seperti Query Parameter, isinya adalah map[string][]string. Tetapi berbeda dengan Query Parameter yang case sensitive, secara spesifikasi Header key tidaklah case sensitive

func RequestHeader(writer http.ResponseWriter, request *http.Request) {
	// contentType := request.Header.Get("content-type")

	writer.Header().Add("X-Powered-By", "derillab.com")
	fmt.Fprintf(writer, "ok")
}
func TestRequestHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/", nil)
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	RequestHeader(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
	fmt.Println(response.Header.Get("X-Powered-By"))
}

func FormPost(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()

	if err != nil {
		panic(err)
	}

	firstName := request.PostForm.Get("first_name")
	lastName := request.PostForm.Get("last_name")

	// jika menggunakan postformvalue maka kita tidak perlu melakuka parseform, kerana akan dilakukan otomatis
	// firstName := request.PostFormValue("first_name")
	// lastName := request.PostFormValue("last_name")

	fmt.Fprintf(writer, "hello %s %s", firstName, lastName)
}

func TestFormPost(t *testing.T) {
	requestBody := strings.NewReader("first_name=Muhammad&last_name=Deril")
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080", requestBody)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	recorder := httptest.NewRecorder()
	FormPost(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}
