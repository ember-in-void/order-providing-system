package service

import "order-providing-system/internal/model"

type ServiceModule interface {
	OrderService
	MenuService
	InventoryService
	AggregationsService
}

type OrderService interface {
	GetOrders() ([]model.Order, error)
	CreateOrder(order model.Order) error
	GetOrderById(id string) (*model.Order, error)
	UpdateOrderById(order model.Order) error
	DeleteOrderById(id string) error
	CloseOrderById(id string) error
}

type MenuService interface {
	GetMenuItems() ([]model.MenuItem, error)
	AddMenuItem(item model.MenuItem) error
	GetMenuItemByID(id string) (*model.MenuItem, error)
	UpdateMenuItemById(item model.MenuItem) error
	DeleteMenuItemById(id string) error
}

type InventoryService interface {
	GetInventoryItems() ([]model.InventoryItem, error)
	AddInventoryItem(item model.InventoryItem) error
	GetInventoryItemByID(id string) (*model.InventoryItem, error)
	UpdateInventoryItemById(item model.InventoryItem) error
	DeleteInventoryItemById(id string) error
}

type AggregationsService interface {
	GetTotalSales() (float64, error)
	GetPopularItems() ([]model.MenuItem, error)
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
