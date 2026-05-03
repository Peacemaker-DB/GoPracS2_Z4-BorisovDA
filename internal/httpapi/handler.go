package httpapi

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"example.com/gopracs2-z4-borisovda/internal/metrics"
	"example.com/gopracs2-z4-borisovda/internal/student"
)

type Handler struct {
	repo *student.Repo
}

func NewHandler(repo *student.Repo) *Handler {
	return &Handler{repo: repo}
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_ = json.NewEncoder(w).Encode(map[string]string{
		"status": "ok",
	})
}

func (h *Handler) GetStudentByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idText := strings.TrimPrefix(r.URL.Path, "/students/")
	if idText == "" || idText == r.URL.Path {
		http.Error(w, "student id is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseInt(idText, 10, 64)
	if err != nil {
		http.Error(w, "invalid student id", http.StatusBadRequest)
		return
	}

	start := time.Now()
	metrics.StudentRequestsTotal.WithLabelValues(idText).Inc()
	defer func() {
		metrics.StudentRequestDuration.WithLabelValues(idText).Observe(time.Since(start).Seconds())
	}()

	st, err := h.repo.GetByID(id)
	if err != nil {
		http.Error(w, "student not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_ = json.NewEncoder(w).Encode(st)
}