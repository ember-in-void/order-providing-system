package storage

import "frappuccino/internal/model"

type RepoModule interface {
	OrderRepo
	MenuRepo
	InventoryRepo
	AggregationsRepo
}

type OrderRepo interface {
	GetOrders() ([]model.Order, error)
	AddOrder(order model.Order) error
	GetOrderById(id string) (*model.Order, error)
	UpdateOrderById(order model.Order) error
	DeleteOrderById(id string) error
	CloseOrderById(id string) error
}

type MenuRepo interface {
	GetMenuItems() ([]model.MenuItem, error)
	AddMenuItem(item model.MenuItem) error
	GetMenuItemByID(id string) (*model.MenuItem, error)
	UpdateMenuItemById(item model.MenuItem) error
	DeleteMenuItemById(id string) error
}

type InventoryRepo interface {
	GetInventoryItems() ([]model.InventoryItem, error)
	AddInventoryItem(item model.InventoryItem) error
	GetInventoryItemByID(id string) (*model.InventoryItem, error)
	UpdateInventoryItemById(item model.InventoryItem) error
	DeleteInventoryItemById(id string) error
}

type AggregationsRepo interface {
	GetTotalSales() (float64, error)
	GetPopularItems() ([]model.MenuItem, error)
}
