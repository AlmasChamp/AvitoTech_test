package handler

import (
	"avito/internal/repository"
	"net/http"
)

type Handler struct {
	Service repository.Service
}

func NewHandler(service repository.Service) repository.Handler {
	return &Handler{Service: service}
}

func (h *Handler) Register(mux *http.ServeMux) {
	mux.HandleFunc("/", h.RegUser)
	mux.HandleFunc("/balance", h.Balance)
	mux.HandleFunc("/credit", h.BalanceIncrease)
	mux.HandleFunc("/debit", h.BalanceDecrease)
	mux.HandleFunc("/transfer", h.BalanceTransfer)
}
