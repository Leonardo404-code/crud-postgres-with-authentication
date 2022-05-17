package controllers

import (
	"bytes"
	"crud-postgres/src/config"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetUsers(t *testing.T) {
	config.LoadDotEnvTests()

	req, err := http.NewRequest("GET", "/", nil)

	if err != nil {
		t.Fatalf("Error in open request: %v", err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(GetUsers)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestGetUser(t *testing.T) {
	req, err := http.NewRequest("GET", "/user", nil)

	if err != nil {
		t.Fatalf("Error in open request: %v", err)
	}

	req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTI3MDMxODEsImlkIjoxfQ.Y9lasDG1VvCiuaWp2HZ2Qq49l4hKXfHFLuCSyRichgU")

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(GetUser)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("error returned rwong status code: got %v, expected %v", status, http.StatusOK)
	}
}

func TestCreateUser(t *testing.T) {
	var jsonStr = []byte(`{"name":"user", "email":"user@gmail.com", "password": "123456789"}`)

	req, err := http.NewRequest("POST", "/", bytes.NewBuffer(jsonStr))

	if err != nil {
		t.Fatalf("error in create user %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(CreateUser)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("error returned wrong status code: got %v expected %v", status, http.StatusOK)
	}
}

func TestUpdateUser(t *testing.T) {
	var jsonStr = []byte(`{"name": "Leo", "email": "leo@gmail.com"}`)

	req, err := http.NewRequest("PATCH", "/", bytes.NewBuffer(jsonStr))

	q := req.URL.Query()

	q.Add("id", "1")

	req.URL.RawQuery = q.Encode()

	if err != nil {
		t.Fatalf("error in request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(UpdateUser)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("error returning wrong status code: got %v expected %v", status, http.StatusOK)
	}
}

func TestDeleteUser(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/", nil)

	q := req.URL.Query()

	q.Add("id", "1")

	req.URL.RawQuery = q.Encode()

	if err != nil {
		t.Fatalf("error in request: %v", err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(DeleteUser)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("error returned wrong status code: got %v expected %v", status, http.StatusOK)
	}
}
