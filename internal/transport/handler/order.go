package handler

import (
	"encoding/json"
	"net/http"

	"order-providing-system/internal/model"
)

func (h *HttpCustomHandler) OrderHandler(w http.ResponseWriter, r *http.Request) {
	h.Logger.Info("Order Method handler")
	switch r.Method {
	case http.MethodGet:
		h.GetOrder(w, r)
	case http.MethodPost:
		h.CreateOrder(w, r)
	default:
		SendJSONResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
		h.Logger.Error("Method not allowed")
	}
}

func (h *HttpCustomHandler) OrderByIdHandler(w http.ResponseWriter, r *http.Request) {
	h.Logger.Info("Order by ID Method handler")
	switch r.Method {
	case http.MethodGet:
		h.GetOrderById(w, r)
	case http.MethodPut:
		h.UpdateOrderById(w, r)
	case http.MethodDelete:
		h.DeleteOrderById(w, r)
	default:
		SendJSONResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
		h.Logger.Error("Method not allowed")
	}
}

// Получение всех заказов
func (h *HttpCustomHandler) GetOrder(w http.ResponseWriter, r *http.Request) {
	h.Logger.Info("Fetching all orders")

	orders, err := h.service.GetOrders()
	if err != nil {
		h.Logger.Error("Error fetching orders: ", err)
		SendJSONResponse(w, http.StatusInternalServerError, "Failed to fetch orders")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(orders)
}

// Создание нового заказа
func (h *HttpCustomHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	h.Logger.Info("Creating a new order")

	var order model.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		h.Logger.Error("Invalid JSON payload: ", err)
		SendJSONResponse(w, http.StatusBadRequest, "Invalid JSON format")
		return
	}

	if err := h.service.CreateOrder(order); err != nil {
		h.Logger.Error("Error creating order: ", err)
		SendJSONResponse(w, http.StatusInternalServerError, "Failed to create order")
		return
	}

	SendJSONResponse(w, http.StatusCreated, "Order created successfully")
}

// Получение заказа по ID
func (h *HttpCustomHandler) GetOrderById(w http.ResponseWriter, r *http.Request) {
	h.Logger.Info("Fetching order by ID")

	id := r.URL.Query().Get("id")
	if id == "" {
		h.Logger.Error("Missing ID in request")
		SendJSONResponse(w, http.StatusBadRequest, "ID is required")
		return
	}

	order, err := h.service.GetOrderById(id)
	if err != nil {
		h.Logger.Error("Error fetching order by ID: ", err)
		SendJSONResponse(w, http.StatusInternalServerError, "Failed to fetch order")
		return
	}

	if order == nil {
		SendJSONResponse(w, http.StatusNotFound, "Order not found")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(order)
}

// Обновление заказа по ID
func (h *HttpCustomHandler) UpdateOrderById(w http.ResponseWriter, r *http.Request) {
	h.Logger.Info("Updating order by ID")

	id := r.URL.Query().Get("id")
	if id == "" {
		h.Logger.Error("Missing ID in request")
		SendJSONResponse(w, http.StatusBadRequest, "ID is required")
		return
	}

	var order model.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		h.Logger.Error("Invalid JSON payload: ", err)
		SendJSONResponse(w, http.StatusBadRequest, "Invalid JSON format")
		return
	}
	order.ID = id

	if err := h.service.UpdateOrderById(order); err != nil {
		h.Logger.Error("Error updating order: ", err)
		SendJSONResponse(w, http.StatusInternalServerError, "Failed to update order")
		return
	}

	SendJSONResponse(w, http.StatusOK, "Order updated successfully")
}

// Удаление заказа по ID
func (h *HttpCustomHandler) DeleteOrderById(w http.ResponseWriter, r *http.Request) {
	h.Logger.Info("Deleting order by ID")

	id := r.URL.Query().Get("id")
	if id == "" {
		h.Logger.Error("Missing ID in request")
		SendJSONResponse(w, http.StatusBadRequest, "ID is required")
		return
	}

	if err := h.service.DeleteOrderById(id); err != nil {
		h.Logger.Error("Error deleting order: ", err)
		SendJSONResponse(w, http.StatusInternalServerError, "Failed to delete order")
		return
	}

	SendJSONResponse(w, http.StatusOK, "Order deleted successfully")
}

// Закрытие заказа по ID
func (h *HttpCustomHandler) OrderCloseHandler(w http.ResponseWriter, r *http.Request) {
	h.Logger.Info("Closing order by ID")

	id := r.URL.Query().Get("id")
	if id == "" {
		h.Logger.Error("Missing ID in request")
		SendJSONResponse(w, http.StatusBadRequest, "ID is required")
		return
	}

	if err := h.service.CloseOrderById(id); err != nil {
		h.Logger.Error("Error closing order: ", err)
		SendJSONResponse(w, http.StatusInternalServerError, "Failed to close order")
		return
	}

	SendJSONResponse(w, http.StatusOK, "Order closed successfully")
}
