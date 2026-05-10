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

// respondWithError envia uma resposta JSON padronizada para erros
func respondWithError(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

func (h *ClienteHandler) Listar(w http.ResponseWriter, r *http.Request) {
	// Captura os query parameters: /clientes?limit=10&offset=0
	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")

	limit, _ := strconv.Atoi(limitStr)
	if limit <= 0 { limit = 10 } // valor padrão

	offset, _ := strconv.Atoi(offsetStr)
	if offset < 0 { offset = 0 }

	// r.Context() carrega o sinal de cancelamento do cliente
	clientes, err := h.svc.Listar(r.Context(), limit, offset)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Erro ao buscar clientes")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(clientes)
}

func (h *ClienteHandler) Criar(w http.ResponseWriter, r *http.Request) {
	var c domain.Cliente
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		respondWithError(w, http.StatusBadRequest, "JSON inválido")
		return
	}

	if err := h.svc.Criar(r.Context(), &c); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(c)
}

func (h *ClienteHandler) Deletar(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "ID inválido")
		return
	}

	if err := h.svc.Deletar(r.Context(), id); err != nil {
		respondWithError(w, http.StatusInternalServerError, "Erro ao deletar cliente")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *ClienteHandler) Atualizar(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "ID inválido")
		return
	}

	var c domain.Cliente
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		respondWithError(w, http.StatusBadRequest, "JSON inválido")
		return
	}
	c.ID = id 

	if err := h.svc.Atualizar(r.Context(), &c); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(c)
}
