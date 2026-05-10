package handler

import (
	"cafeteria-api/internal/domain"
	"cafeteria-api/internal/service"
	"encoding/json"
	"net/http"
	"strconv"
)

type ClienteHandler struct {
	svc *service.ClienteService
}

func NewClienteHandler(svc *service.ClienteService) *ClienteHandler {
	return &ClienteHandler{svc: svc}
}

func (h *ClienteHandler) Listar(w http.ResponseWriter, r *http.Request) {
	clientes, err := h.svc.ListarTodos()
	if err != nil {
		http.Error(w, "Erro ao buscar clientes", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(clientes)
}

func (h *ClienteHandler) Criar(w http.ResponseWriter, r *http.Request) {
	var c domain.Cliente
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}
	if err := h.svc.Criar(&c); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(c)
}

func (h *ClienteHandler) Deletar(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	if err := h.svc.Deletar(id); err != nil {
		http.Error(w, "Erro ao deletar cliente", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *ClienteHandler) Atualizar(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var c domain.Cliente
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}
	c.ID = id // Garante que o ID da URL seja o usado no banco

	if err := h.svc.Atualizar(&c); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(c)
}
