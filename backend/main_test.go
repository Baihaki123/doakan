package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	. "doakan/handler"

	"github.com/gin-gonic/gin"
)

func TestLogin(t *testing.T) {

	var jsonStr = []byte(`{"username":"testakhir1@gmail.com","password":"password"}`)
	var userHandler = NewUserHandler(nil, nil)

	req, err := http.NewRequest("POST", "api/v1/login", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	g := gin.Default()
	userHandler.Login(&gin.Context{Request: req})

	g.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"username":"testakhir1@gmail.com","password":"password"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
