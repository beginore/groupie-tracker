package pkg

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	TemplateDir = "../web/templates"
	os.Exit(m.Run())
}

func TestIndexHandler(t *testing.T) {
	TestHandler := []struct {
		Method string
		Path   string
		Answer int
	}{
		{
			Method: "POST",
			Path:   "http://localhost:8080/",
			Answer: 405,
		},
		{
			Method: "DELETE",
			Path:   "http://localhost:8080/",
			Answer: 405,
		},
		{
			Method: "GET",
			Path:   "http://localhost:8080/lol",
			Answer: 404,
		},
		{
			Method: "POST",
			Path:   "http://localhost:8080/ayeee",
			Answer: 404,
		},
	}

	GetAPI()
	for i := 0; i < len(TestHandler); i++ {
		request, err := http.NewRequest(TestHandler[i].Method, TestHandler[i].Path, nil)
		if err != nil {
			t.Fatal(err)
		}
		response := httptest.NewRecorder()
		handler := http.HandlerFunc(IndexHandler)
		handler.ServeHTTP(response, request)
		if status := response.Code; status != TestHandler[i].Answer {
			t.Errorf("Handler returned wrong status code: got %v want %v", status, TestHandler[i].Answer)
		}
	}
}

func TestArtistHandler(t *testing.T) {
	TestServer := []struct {
		Method string
		Path   string
		Answer int
	}{
		{
			Method: "GET",
			Path:   "http://localhost:8080/artist/",
			Answer: 400,
		},
		{
			Method: "GET",
			Path:   "http://localhost:8080/artist/-1",
			Answer: 400,
		},
		{
			Method: "GET",
			Path:   "http://localhost:8080/artist/53",
			Answer: 404,
		},
		{
			Method: "GET",
			Path:   "http://localhost:8080/artist/A",
			Answer: 400,
		},
		{
			Method: "POST",
			Path:   "http://localhost:8080/artist/",
			Answer: 405,
		},
		{
			Method: "DELETE",
			Path:   "http://localhost:8080/artist/",
			Answer: 405,
		},
		{
			Method: "GET",
			Path:   "http://localhost:8080/artist/48",
			Answer: 200,
		},
		{
			Method: "GET",
			Path:   "http://localhost:8080/artist/1",
			Answer: 200,
		},
	}
	for i := 0; i < len(TestServer); i++ {
		request, err := http.NewRequest(TestServer[i].Method, TestServer[i].Path, nil)
		if err != nil {
			t.Fatal(err)
		}
		response := httptest.NewRecorder()
		handler := http.HandlerFunc(ArtistHandler)
		handler.ServeHTTP(response, request)

		if status := response.Code; status != TestServer[i].Answer {
			t.Errorf("handler returned wrong status code: got %v want %v", status, TestServer[i].Answer)
		}
	}
}
