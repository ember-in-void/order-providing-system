package handler

import (
	"encoding/json"
	"net/http"
)

// Получение общей суммы продаж
func (h *HttpCustomHandler) TotalSalesHandler(w http.ResponseWriter, r *http.Request) {
	h.Logger.Info("Fetching total sales")

	// Получение данных из сервисного слоя
	totalSales, err := h.service.GetTotalSales()
	if err != nil {
		h.Logger.Error("Error fetching total sales: ", err)
		SendJSONResponse(w, http.StatusInternalServerError, "Failed to fetch total sales")
		return
	}

	// Формирование ответа
	response := map[string]float64{"total_sales": totalSales}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// Получение популярных товаров
func (h *HttpCustomHandler) PopularItemsHandler(w http.ResponseWriter, r *http.Request) {
	h.Logger.Info("Fetching popular items")

	// Получение данных из сервисного слоя
	popularItems, err := h.service.GetPopularItems()
	if err != nil {
		h.Logger.Error("Error fetching popular items: ", err)
		SendJSONResponse(w, http.StatusInternalServerError, "Failed to fetch popular items")
		return
	}

	// Формирование ответа
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(popularItems)
}

// Обработка нулевого запроса
func (h *HttpCustomHandler) HomeHandler(w http.ResponseWriter, r *http.Request) {
	h.Logger.Info("Home Method handler")
	SendJSONResponse(w, http.StatusOK, "Welcome to the Frappuccino API")
}

// Отправка JSON-ответа
func SendJSONResponse(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": message,
		"status":  status,
	})
}
