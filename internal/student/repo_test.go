package student

import (
	"errors"
	"testing"
)

func TestRepoGetByIDFound(t *testing.T) {
	repo := NewRepo()
	st, err := repo.GetByID(1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if st.ID != 1 {
		t.Fatalf("expected id 1, got %d", st.ID)
	}
	if st.FullName == "" {
		t.Fatal("expected full name")
	}
}

func TestRepoGetByIDNotFound(t *testing.T) {
	repo := NewRepo()
	_, err := repo.GetByID(999)
	if !errors.Is(err, ErrStudentNotFound) {
		t.Fatalf("expected ErrStudentNotFound, got %v", err)
	}
}