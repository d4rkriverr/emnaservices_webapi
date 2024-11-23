package expenses

import (
	"emnaservices/webapi/internal/datatype"
	"emnaservices/webapi/utils"
	"encoding/json"
	"net/http"
)

type Handler struct {
	service *Service
}

func newHandler(s *Service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) GetExpansesData(w http.ResponseWriter, r *http.Request) {
	date_from := r.URL.Query().Get("from")
	date_to := r.URL.Query().Get("to")

	data, err := h.service.GetExpensesWithRange(date_from, date_to)
	if err != nil {
		utils.RespondWithError(w, 400, err.Error())
		return
	}
	utils.RespondWithSuccess(w, data)
}

func (h *Handler) HandleCreateExpanses(w http.ResponseWriter, r *http.Request) {
	var newInvoice datatype.Transaction
	err := json.NewDecoder(r.Body).Decode(&newInvoice)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Set the response header to JSON and return the new invoice
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]any{"success": true})
}
