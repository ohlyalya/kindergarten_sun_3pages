package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"kindergarten-sun-backend/internal/models"
	"kindergarten-sun-backend/internal/service"
)

type Handler struct {
	service *service.Service
}

func New(s *service.Service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/health", h.health)
	mux.HandleFunc("/api/home", h.home)
	mux.HandleFunc("/api/programs", h.programs)
	mux.HandleFunc("/api/parents", h.parents)
	mux.HandleFunc("/api/search", h.search)
	mux.HandleFunc("/api/applications", h.applications)
	mux.HandleFunc("/api/questions", h.questions)
	mux.HandleFunc("/docs/openapi.yaml", h.openapi)
	mux.HandleFunc("/swagger/", h.swagger)
	mux.HandleFunc("/swagger", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/swagger/", http.StatusMovedPermanently)
	})
}

func WithCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (h *Handler) health(w http.ResponseWriter, r *http.Request) {
	if !allow(w, r, http.MethodGet) {
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func (h *Handler) home(w http.ResponseWriter, r *http.Request) {
	if !allow(w, r, http.MethodGet) {
		return
	}
	writeJSON(w, http.StatusOK, h.service.Home())
}

func (h *Handler) programs(w http.ResponseWriter, r *http.Request) {
	if !allow(w, r, http.MethodGet) {
		return
	}
	writeJSON(w, http.StatusOK, h.service.Programs())
}

func (h *Handler) parents(w http.ResponseWriter, r *http.Request) {
	if !allow(w, r, http.MethodGet) {
		return
	}
	writeJSON(w, http.StatusOK, h.service.Parents())
}

func (h *Handler) search(w http.ResponseWriter, r *http.Request) {
	if !allow(w, r, http.MethodGet) {
		return
	}
	query := r.URL.Query().Get("q")
	writeJSON(w, http.StatusOK, h.service.Search(query))
}

func (h *Handler) applications(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		writeJSON(w, http.StatusOK, h.service.Applications())
	case http.MethodPost:
		var req models.ApplicationRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			writeJSON(w, http.StatusBadRequest, models.ErrorResponse{Error: "некорректный JSON"})
			return
		}
		app, err := h.service.CreateApplication(req)
		if err != nil {
			writeJSON(w, http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
			return
		}
		writeJSON(w, http.StatusCreated, app)
	default:
		methodNotAllowed(w)
	}
}

func (h *Handler) questions(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		writeJSON(w, http.StatusOK, h.service.Questions())
	case http.MethodPost:
		var req models.QuestionRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			writeJSON(w, http.StatusBadRequest, models.ErrorResponse{Error: "некорректный JSON"})
			return
		}
		question, err := h.service.CreateQuestion(req)
		if err != nil {
			writeJSON(w, http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
			return
		}
		writeJSON(w, http.StatusCreated, question)
	default:
		methodNotAllowed(w)
	}
}

func (h *Handler) openapi(w http.ResponseWriter, r *http.Request) {
	if !allow(w, r, http.MethodGet) {
		return
	}
	paths := []string{"docs/openapi.yaml", filepath.Join("..", "..", "docs", "openapi.yaml")}
	for _, path := range paths {
		if _, err := os.Stat(path); err == nil {
			w.Header().Set("Content-Type", "application/yaml; charset=utf-8")
			http.ServeFile(w, r, path)
			return
		}
	}
	writeJSON(w, http.StatusNotFound, models.ErrorResponse{Error: "openapi.yaml не найден"})
}

func (h *Handler) swagger(w http.ResponseWriter, r *http.Request) {
	if !allow(w, r, http.MethodGet) {
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, _ = w.Write([]byte(`<!doctype html>
<html lang="ru">
<head>
  <meta charset="utf-8" />
  <title>ГБДОУ «Солнце» API</title>
  <link rel="stylesheet" href="https://unpkg.com/swagger-ui-dist@5/swagger-ui.css" />
</head>
<body>
  <div id="swagger-ui"></div>
  <script src="https://unpkg.com/swagger-ui-dist@5/swagger-ui-bundle.js"></script>
  <script>
    window.onload = () => SwaggerUIBundle({ url: '/docs/openapi.yaml', dom_id: '#swagger-ui' });
  </script>
</body>
</html>`))
}

func allow(w http.ResponseWriter, r *http.Request, method string) bool {
	if r.Method != method {
		methodNotAllowed(w)
		return false
	}
	return true
}

func methodNotAllowed(w http.ResponseWriter) {
	writeJSON(w, http.StatusMethodNotAllowed, models.ErrorResponse{Error: "метод не поддерживается"})
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}

func NormalizePath(path string) string {
	return strings.TrimSpace(path)
}
