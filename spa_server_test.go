package spa_server

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestSanityCheckIndexExists(t *testing.T) {
	b, err := ioutil.ReadFile("public/index.html")
	if err != nil {
		t.Fatal("Error reading file: public/index.html")
	}

	knownString := []byte("Test")
	if bytes.Contains(b, knownString) != true {
		t.Fatal("File: public/index.html did not contain an expected string!")
	}
}

func TestSanityCheckJsFileExists(t *testing.T) {
	b, err := ioutil.ReadFile("public/js/test.js")
	if err != nil {
		t.Fatal("Error reading file: public/js/test.js")
	}

	knownString := []byte("Test")
	if bytes.Contains(b, knownString) != true {
		t.Fatal("File: public/js/test.js did not contain an expected string!")
	}
}

func TestSanityCheckLoginDirectory(t *testing.T) {
	f, err := os.Stat("public/login")
	if os.IsNotExist(err) {
		return
	}
	if f.IsDir() {
		t.Fatal("Path should not exist!")
	}
}

func TestSpaHandlerRoot(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := SpaHandler("public", "index.html")

	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected, _ := ioutil.ReadFile("public/index.html")
	rb := rr.Body.Bytes()
	if bytes.Compare(rb, expected) != 0 {
		t.Error("handler returned unexpected body")
	}
}

func TestSpaHandlerIndexShouldRedirect(t *testing.T) {
	req, err := http.NewRequest("GET", "/index.html", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := SpaHandler("public", "index.html")

	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusMovedPermanently {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	location := rr.Header().Get("Location")
	if location != "./" {
		t.Errorf("Handler returned wrong Location header: got %v wanted ./", location)
	}
}

func TestSpaHandlerDirectoryFound(t *testing.T) {
	req, err := http.NewRequest("GET", "/js", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := SpaHandler("public", "index.html")

	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected, _ := ioutil.ReadFile("public/index.html")
	rb := rr.Body.Bytes()
	if bytes.Compare(rb, expected) != 0 {
		t.Error("handler returned unexpected body")
	}
}

func TestSpaHandlerDirectoryNotFound(t *testing.T) {
	req, err := http.NewRequest("GET", "/login", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := SpaHandler("public", "index.html")

	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected, _ := ioutil.ReadFile("public/index.html")
	rb := rr.Body.Bytes()
	if bytes.Compare(rb, expected) != 0 {
		t.Error("handler returned unexpected body")
	}
}

func TestSpaHandlerFileFound(t *testing.T) {
	req, err := http.NewRequest("GET", "/js/test.js", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := SpaHandler("public", "index.html")

	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected, _ := ioutil.ReadFile("public/js/test.js")
	rb := rr.Body.Bytes()
	if bytes.Compare(rb, expected) != 0 {
		t.Error("handler returned unexpected body")
	}
}

func TestSpaHandlerFileNotFound(t *testing.T) {
	req, err := http.NewRequest("GET", "/js/missing.js", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := SpaHandler("public", "index.html")

	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected, _ := ioutil.ReadFile("public/index.html")
	rb := rr.Body.Bytes()
	if bytes.Compare(rb, expected) != 0 {
		t.Error("handler returned unexpected body")
	}
}
