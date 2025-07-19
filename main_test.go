package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAlbums(t *testing.T) {
	router := setupRouter()
	req, _ := http.NewRequest("GET", "/albums", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected status 200, got %d", w.Code)
	}
}

func TestPostAlbum(t *testing.T) {
	router := setupRouter()
	body := []byte(`{"id":"4","title":"Test Album","artist":"Test Artist","price":12.34}`)
	req, _ := http.NewRequest("POST", "/albums", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("Expected status 201, got %d", w.Code)
	}
}

func TestGetAlbumByID(t *testing.T) {
	router := setupRouter()
	req, _ := http.NewRequest("GET", "/albums/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected status 200, got %d", w.Code)
	}
}
