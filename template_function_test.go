package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

type MyPage struct {
	Name string
}

func (mypage MyPage) SayHello(name string) string {
	return "Hello " + name + ", My name is " + mypage.Name
}
func TemplateFunction(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.New("FUNCTION").Parse(`{{.SayHello "Deril"}}`))
	t.ExecuteTemplate(writer, "FUNCTION", MyPage{
		Name: "Deril",
	})
}
func TestTemplateFunction(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()
	TemplateFunction(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
