package handlers

import (
    "net/http"
    "net/http/httptest"
    "testing"
    // "fmt"
    "log"
    "github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {

  	// handlerMock := func(w http.ResponseWriter, r *http.Request) {
  		// http.Error(w, "something failed", http.StatusInternalServerError)
  	// }

  	req, err := http.NewRequest("GET", "/", nil)
  	if err != nil {
  		log.Fatal(err)
  	}

  	w := httptest.NewRecorder()
  	handler(w, req)

  	// fmt.Printf("END: %d - %s", w.Code, w.Body.String())

    assert.Equal(t, w.Code, 200, "Should return 200")
}

func TestReturnJsonArrayHandler(t *testing.T) {
  	req, err := http.NewRequest("GET", "/return-json", nil)
  	if err != nil {
  		log.Fatal(err)
  	}

  	w := httptest.NewRecorder()
  	returnJson(w, req)

  	// fmt.Printf("END: %d - %s", w.Code, w.Body.String())
    assert.Equal(t, w.Code, 200, "Should return 200")
}

func TestFetchJsonHandler(t *testing.T) {
  	req, err := http.NewRequest("GET", "/fetch-json", nil)
  	if err != nil {
  		log.Fatal(err)
  	}

  	w := httptest.NewRecorder()
  	fetchJson(w, req)

  	// fmt.Printf("END: %d - %s", w.Code, w.Body.String())
    assert.Equal(t, w.Code, 200, "Should return 200")
}
