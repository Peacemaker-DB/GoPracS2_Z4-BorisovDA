package httpapi

import "testing"

func TestNormalizePathStudents(t *testing.T) {
	got := NormalizePath("/students/1")
	if got != "/students/{id}" {
		t.Fatalf("expected /students/{id}, got %s", got)
	}
}

func TestNormalizePathOther(t *testing.T) {
	got := NormalizePath("/health")
	if got != "/health" {
		t.Fatalf("expected /health, got %s", got)
	}
}