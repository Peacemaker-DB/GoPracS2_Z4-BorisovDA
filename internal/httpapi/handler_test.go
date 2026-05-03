package httpapi

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"example.com/gopracs2-z4-borisovda/internal/student"
)

func TestHealthOK(t *testing.T) {
	handler := NewHandler(student.NewRepo())
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rr := httptest.NewRecorder()

	handler.Health(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rr.Code)
	}

	var body map[string]string
	if err := json.NewDecoder(rr.Body).Decode(&body); err != nil {
		t.Fatalf("invalid json: %v", err)
	}
	if body["status"] != "ok" {
		t.Fatalf("expected status ok, got %s", body["status"])
	}
}

func TestHealthMethodNotAllowed(t *testing.T) {
	handler := NewHandler(student.NewRepo())
	req := httptest.NewRequest(http.MethodPost, "/health", nil)
	rr := httptest.NewRecorder()

	handler.Health(rr, req)

	if rr.Code != http.StatusMethodNotAllowed {
		t.Fatalf("expected status 405, got %d", rr.Code)
	}
}

func TestGetStudentByIDOK(t *testing.T) {
	handler := NewHandler(student.NewRepo())
	req := httptest.NewRequest(http.MethodGet, "/students/1", nil)
	rr := httptest.NewRecorder()

	handler.GetStudentByID(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rr.Code)
	}

	var st student.Student
	if err := json.NewDecoder(rr.Body).Decode(&st); err != nil {
		t.Fatalf("invalid json: %v", err)
	}
	if st.ID != 1 {
		t.Fatalf("expected id 1, got %d", st.ID)
	}
}

func TestGetStudentByIDNotFound(t *testing.T) {
	handler := NewHandler(student.NewRepo())
	req := httptest.NewRequest(http.MethodGet, "/students/999", nil)
	rr := httptest.NewRecorder()

	handler.GetStudentByID(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Fatalf("expected status 404, got %d", rr.Code)
	}
}

func TestGetStudentByIDInvalidID(t *testing.T) {
	handler := NewHandler(student.NewRepo())
	req := httptest.NewRequest(http.MethodGet, "/students/test", nil)
	rr := httptest.NewRecorder()

	handler.GetStudentByID(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", rr.Code)
	}
}

func TestGetStudentByIDMethodNotAllowed(t *testing.T) {
	handler := NewHandler(student.NewRepo())
	req := httptest.NewRequest(http.MethodPost, "/students/1", nil)
	rr := httptest.NewRecorder()

	handler.GetStudentByID(rr, req)

	if rr.Code != http.StatusMethodNotAllowed {
		t.Fatalf("expected status 405, got %d", rr.Code)
	}
}