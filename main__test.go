package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Ankcorn/hello/handlers"
	"github.com/gin-gonic/gin"
)

type Person struct {
	Firstname string `json:"firstname,omitempty"`
}

func TestPeopleApi(t *testing.T) {
	data := []byte(`{"name": "dave"}`)
	req, _ := http.NewRequest("POST", "/people", bytes.NewBuffer(data))
	w := httptest.NewRecorder()

	r := gin.Default()
	r.POST("/people", handlers.AddPeople)
	r.ServeHTTP(w, req)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `[{"firstname":"test"},{"firstname":"dave"}]`
	if body := w.Body.String(); body != expected {
		t.Errorf("handler returned wrong body: got %v want %v",
			body, expected)
	}
}
