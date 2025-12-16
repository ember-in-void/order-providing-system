package handler

import (
	"encoding/json"
	"net/http"

	"order-providing-system/internal/model"
)

func (h *HttpCustomHandler) InventoryHandler(w http.ResponseWriter, r *http.Request) {
	h.Logger.Info("Inventory Method handler")
	switch r.Method {
	case http.MethodGet:
		h.GetInventoryItems(w, r)
	case http.MethodPost:
		h.AddInventoryItem(w, r)
	default:
		SendJSONResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
		h.Logger.Error("Method not allowed")
	}
}

func (h *HttpCustomHandler) InventoryByIdHandler(w http.ResponseWriter, r *http.Request) {
	h.Logger.Info("Inventory by ID Method handler")
	switch r.Method {
	case http.MethodGet:
		h.GetInventoryItemById(w, r)
	case http.MethodPut:
		h.UpdateInventoryItemById(w, r)
	case http.MethodDelete:
		h.DeleteInventoryItemById(w, r)
	default:
		SendJSONResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
		h.Logger.Error("Method not allowed")
	}
}

// Получение всех элементов инвентаря
func (h *HttpCustomHandler) GetInventoryItems(w http.ResponseWriter, r *http.Request) {
	h.Logger.Info("Fetching all inventory items")

	items, err := h.service.GetInventoryItems()
	if err != nil {
		h.Logger.Error("Error fetching inventory items: ", err)
		SendJSONResponse(w, http.StatusInternalServerError, "Failed to fetch inventory items")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(items)
}

// Добавление нового элемента инвентаря
func (h *HttpCustomHandler) AddInventoryItem(w http.ResponseWriter, r *http.Request) {
	h.Logger.Info("Creating a new inventory item")

	var item model.InventoryItem
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		h.Logger.Error("Invalid JSON payload: ", err)
		SendJSONResponse(w, http.StatusBadRequest, "Invalid JSON format")
		return
	}

	if err := h.service.AddInventoryItem(item); err != nil {
		h.Logger.Error("Error creating inventory item: ", err)
		SendJSONResponse(w, http.StatusInternalServerError, "Failed to create inventory item")
		return
	}

	SendJSONResponse(w, http.StatusCreated, "Inventory item created")
}

// Получение элемента инвентаря по ID
func (h *HttpCustomHandler) GetInventoryItemById(w http.ResponseWriter, r *http.Request) {
	h.Logger.Info("Fetching inventory item by ID")

	id := r.URL.Query().Get("id")
	if id == "" {
		h.Logger.Error("Missing ID in request")
		SendJSONResponse(w, http.StatusBadRequest, "ID is required")
		return
	}

	item, err := h.service.GetInventoryItemByID(id)
	if err != nil {
		h.Logger.Error("Error fetching inventory item by ID: ", err)
		SendJSONResponse(w, http.StatusInternalServerError, "Failed to fetch inventory item")
		return
	}

	if item == nil {
		SendJSONResponse(w, http.StatusNotFound, "Inventory item not found")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(item)
}

// Обновление элемента инвентаря по ID
func (h *HttpCustomHandler) UpdateInventoryItemById(w http.ResponseWriter, r *http.Request) {
	h.Logger.Info("Updating inventory item by ID")

	id := r.URL.Query().Get("id")
	if id == "" {
		h.Logger.Error("Missing ID in request")
		SendJSONResponse(w, http.StatusBadRequest, "ID is required")
		return
	}

	var item model.InventoryItem
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		h.Logger.Error("Invalid JSON payload: ", err)
		SendJSONResponse(w, http.StatusBadRequest, "Invalid JSON format")
		return
	}
	item.IngredientID = id

	if err := h.service.UpdateInventoryItemById(item); err != nil {
		h.Logger.Error("Error updating inventory item: ", err)
		SendJSONResponse(w, http.StatusInternalServerError, "Failed to update inventory item")
		return
	}

	SendJSONResponse(w, http.StatusOK, "Inventory item updated")
}

// Удаление элемента инвентаря по ID
func (h *HttpCustomHandler) DeleteInventoryItemById(w http.ResponseWriter, r *http.Request) {
	h.Logger.Info("Deleting inventory item by ID")

	id := r.URL.Query().Get("id")
	if id == "" {
		h.Logger.Error("Missing ID in request")
		SendJSONResponse(w, http.StatusBadRequest, "ID is required")
		return
	}

	if err := h.service.DeleteInventoryItemById(id); err != nil {
		h.Logger.Error("Error deleting inventory item: ", err)
		SendJSONResponse(w, http.StatusInternalServerError, "Failed to delete inventory item")
		return
	}

	SendJSONResponse(w, http.StatusNoContent, "Inventory item deleted")
}
