package transport

import (
	"net/http"

	"frappuccino/internal/transport/handler"
)

func SetupRouter(h *handler.HttpCustomHandler) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", h.HomeHandler)
	mux.HandleFunc("/menu", h.MenuHandler)
	mux.HandleFunc("/menu/", h.MenuHandler)
	mux.HandleFunc("/menu/{id}", h.MenuByIdHandler)
	mux.HandleFunc("/menu/{id}/", h.MenuByIdHandler)

	mux.HandleFunc("/order", h.OrderHandler)
	mux.HandleFunc("/order/", h.OrderHandler)
	mux.HandleFunc("/order/{id}", h.OrderByIdHandler)
	mux.HandleFunc("/order/{id}/", h.OrderByIdHandler)
	mux.HandleFunc("/order/{id}/close", h.OrderCloseHandler)
	mux.HandleFunc("/order/{id}/close/", h.OrderCloseHandler)

	mux.HandleFunc("/inventory", h.InventoryHandler)
	mux.HandleFunc("/inventory/", h.InventoryHandler)
	mux.HandleFunc("/inventory/{id}", h.InventoryByIdHandler)
	mux.HandleFunc("/inventory/{id}/", h.InventoryByIdHandler)

	mux.HandleFunc("/reports/total-sales", h.TotalSalesHandler)
	mux.HandleFunc("/reports/total-sales/", h.TotalSalesHandler)

	mux.HandleFunc("/reports/popular-items", h.PopularItemsHandler)
	mux.HandleFunc("/reports/popular-items/", h.PopularItemsHandler)

	h.Logger.Info("Router is ready")
	return mux
}

// Orders:
// POST /orders: Create a new order.
// GET /orders: Retrieve all orders.
// GET /orders/{id}: Retrieve a specific order by ID.
// PUT /orders/{id}: Update an existing order.
// DELETE /orders/{id}: Delete an order.
// POST /orders/{id}/close: Close an order.

// Menu Items:
// POST /menu: Add a new menu item.
// GET /menu: Retrieve all menu items.
// GET /menu/{id}: Retrieve a specific menu item.
// PUT /menu/{id}: Update a menu item.
// DELETE /menu/{id}: Delete a menu item.

// Inventory:
// POST /inventory: Add a new inventory item.
// GET /inventory: Retrieve all inventory items.
// GET /inventory/{id}: Retrieve a specific inventory item.
// PUT /inventory/{id}: Update an inventory item.
// DELETE /inventory/{id}: Delete an inventory item.

// Aggregations:
// GET /reports/total-sales: Get the total sales amount.
// GET /reports/popular-items: Get a list of popular menu items.
