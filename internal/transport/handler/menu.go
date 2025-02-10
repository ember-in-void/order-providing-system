package handler

import (
	"encoding/json"
	"net/http"

	"frappuccino/internal/model"
)

func (h *HttpCustomHandler) MenuHandler(w http.ResponseWriter, r *http.Request) {
	h.Logger.Info("Menu Method handler")
	switch r.Method {
	case http.MethodGet:
		h.getMenu(w, r)
	case http.MethodPost:
		h.postMenu(w, r)
	default:
		SendJSONResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
		h.Logger.Error("Method not allowed")
	}
}

func (h *HttpCustomHandler) MenuByIdHandler(w http.ResponseWriter, r *http.Request) {
	h.Logger.Info("Menu by ID Method handler")
	switch r.Method {
	case http.MethodGet:
		h.getMenuById(w, r)
	case http.MethodPut:
		h.putMenuById(w, r)
	case http.MethodDelete:
		h.deleteMenuById(w, r)
	default:
		SendJSONResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
		h.Logger.Error("Method not allowed")
	}
}

// Получение всех элементов меню
func (h *HttpCustomHandler) getMenu(w http.ResponseWriter, r *http.Request) {
	h.Logger.Info("Listing all menu items")

	// Получение данных из сервисного слоя
	items, err := h.service.GetMenuItems()
	if err != nil {
		h.Logger.Error("Error fetching menu items: ", err)
		SendJSONResponse(w, http.StatusInternalServerError, "Failed to fetch menu items")
		return
	}

	// Преобразование в JSON и запись в ответ
	response, _ := json.Marshal(items)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// Добавление нового элемента меню
func (h *HttpCustomHandler) postMenu(w http.ResponseWriter, r *http.Request) {
	h.Logger.Info("Creating a new menu item")

	// Декодирование тела запроса
	var item model.MenuItem
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields() // Запрет неизвестных полей
	if err := decoder.Decode(&item); err != nil {
		h.Logger.Error("Invalid JSON payload: ", err)
		SendJSONResponse(w, http.StatusBadRequest, "Invalid JSON format")
		return
	}

	// Создание элемента меню через сервисный слой
	if err := h.service.AddMenuItem(item); err != nil {
		h.Logger.Error("Error creating menu item: ", err)
		SendJSONResponse(w, http.StatusInternalServerError, "Failed to create menu item")
		return
	}

	SendJSONResponse(w, http.StatusCreated, "Menu item created")
}

// Получение элемента меню по ID
func (h *HttpCustomHandler) getMenuById(w http.ResponseWriter, r *http.Request) {
	h.Logger.Info("Fetching menu item by ID")

	// Получение ID из URL
	id := r.URL.Query().Get("id")
	if id == "" {
		h.Logger.Error("Missing ID in request")
		SendJSONResponse(w, http.StatusBadRequest, "ID is required")
		return
	}

	// Получение элемента через сервисный слой
	item, err := h.service.GetMenuItemByID(id)
	if err != nil {
		h.Logger.Error("Error fetching menu item by ID: ", err)
		SendJSONResponse(w, http.StatusInternalServerError, "Failed to fetch menu item")
		return
	}
	if item == nil {
		SendJSONResponse(w, http.StatusNotFound, "Menu item not found")
		return
	}

	// Преобразование в JSON и запись в ответ
	response, _ := json.Marshal(item)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// Обновление элемента меню по ID
func (h *HttpCustomHandler) putMenuById(w http.ResponseWriter, r *http.Request) {
	h.Logger.Info("Updating menu item by ID")

	// Получение ID из URL
	id := r.URL.Query().Get("id")
	if id == "" {
		h.Logger.Error("Missing ID in request")
		SendJSONResponse(w, http.StatusBadRequest, "ID is required")
		return
	}

	// Декодирование тела запроса
	var item model.MenuItem
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		h.Logger.Error("Invalid request payload: ", err)
		SendJSONResponse(w, http.StatusBadRequest, "Invalid JSON format")
		return
	}
	item.ID = id // Установка ID из URL

	// Обновление элемента через сервисный слой
	if err := h.service.UpdateMenuItemById(item); err != nil {
		h.Logger.Error("Error updating menu item: ", err)
		SendJSONResponse(w, http.StatusInternalServerError, "Failed to update menu item")
		return
	}

	SendJSONResponse(w, http.StatusOK, "Menu item updated")
}

// Удаление элемента меню по ID
func (h *HttpCustomHandler) deleteMenuById(w http.ResponseWriter, r *http.Request) {
	h.Logger.Info("Deleting menu item by ID")

	// Получение ID из URL
	id := r.URL.Query().Get("id")
	if id == "" {
		h.Logger.Error("Missing ID in request")
		SendJSONResponse(w, http.StatusBadRequest, "ID is required")
		return
	}

	// Удаление элемента через сервисный слой
	if err := h.service.DeleteMenuItemById(id); err != nil {
		h.Logger.Error("Error deleting menu item: ", err)
		SendJSONResponse(w, http.StatusInternalServerError, "Failed to delete menu item")
		return
	}

	SendJSONResponse(w, http.StatusNoContent, "Menu item deleted")
}
