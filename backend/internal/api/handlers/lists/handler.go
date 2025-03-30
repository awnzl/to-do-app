package lists

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"

	"github.com/awnzl/to-do-app/internal/api/models"
	"github.com/awnzl/to-do-app/internal/service"
)

type Handler struct {
	svc service.TodoService
}

func NewHandler(svc service.TodoService) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) RegisterRoutes(r chi.Router) {
	r.Get("/", h.ListAll)
	r.Post("/", h.Create)
	r.Route("/{listID}", func(r chi.Router) {
		r.Get("/", h.GetByID)
		r.Put("/", h.Update)
		r.Delete("/", h.Delete)
		r.Get("/todos", h.ListTodos)
	})
}

func (h *Handler) ListAll(w http.ResponseWriter, r *http.Request) {
	lists, err := h.svc.ListLists(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(lists)
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	if name == "" {
		http.Error(w, "invalid list name", http.StatusBadRequest)
		return
	}

	var req models.CreateListRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	list, err := h.svc.CreateList(r.Context(), req.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(list)
}

func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	listID, err := uuid.Parse(chi.URLParam(r, "listID"))
	if err != nil {
		http.Error(w, "invalid list ID", http.StatusBadRequest)
		return
	}

	list, err := h.svc.GetList(r.Context(), listID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(list)
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	listID, err := uuid.Parse(chi.URLParam(r, "listID"))
	if err != nil {
		http.Error(w, "invalid list ID", http.StatusBadRequest)
		return
	}

	var req models.UpdateListRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	list, err := h.svc.GetList(r.Context(), listID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	list.Name = req.Name

	if err := h.svc.UpdateList(r.Context(), list); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	listID, err := uuid.Parse(chi.URLParam(r, "listID"))
	if err != nil {
		http.Error(w, "invalid list ID", http.StatusBadRequest)
		return
	}

	if err := h.svc.DeleteList(r.Context(), listID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) ListTodos(w http.ResponseWriter, r *http.Request) {
	listID, err := uuid.Parse(chi.URLParam(r, "listID"))
	if err != nil {
		http.Error(w, "invalid list ID", http.StatusBadRequest)
		return
	}

	todos, err := h.svc.ListTodos(r.Context(), listID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(todos)
}
