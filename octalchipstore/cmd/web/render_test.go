package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestConfig_AddDefaultData(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)

	ctx := getCtx(req)

	req = req.WithContext(ctx)

	testApp.Session.Put(ctx, "flash", "flash")
	testApp.Session.Put(ctx, "warning", "warning")
	testApp.Session.Put(ctx, "error", "error")

	td := testApp.AddDefaultData(&TemplateData{}, req)

	if td.Flash != "flash" {
		t.Error("failed to get flash data")
	}

	if td.Warning != "warning" {
		t.Error("failed to get warning data")
	}

	if td.Error != "error" {
		t.Error("failed to get error data")
	}

}

func TestConfig_IsAuthenticated(t *testing.T) {
	// Create a new HTTP request
	req, _ := http.NewRequest("GET", "/", nil)

	// Create a context for the request
	ctx := getCtx(req)

	// Update the request with the new context
	req = req.WithContext(ctx)

	// Test when userID is not in the session
	if testApp.IsAuthenticated(req) {
		t.Error("user should not be authenticated")
	}

	// Store userID in the session
	testApp.Session.Put(ctx, "userID", "12345")

	// Test when userID is in the session
	if !testApp.IsAuthenticated(req) {
		t.Error("user should be authenticated")
	}
}

func TestConfig_render(t *testing.T) {

	pathToTemplates = "./templates"

	// Create a new HTTP response recorder
	rr := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	ctx := getCtx(req)
	req = req.WithContext(ctx)

	// Call the render method
	testApp.render(rr, req, "home.page.gohtml", &TemplateData{})

	// Check if the response was successful
	if rr.Code != http.StatusOK {
		t.Errorf("expected %d; got %d", http.StatusOK, rr.Code)
	}

}
