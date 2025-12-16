package usecase

import (
	"errors"

	"order-providing-system/internal/model"
)

// Получение всех элементов меню
func (s *CustomService) GetMenuItems() ([]model.MenuItem, error) {
	items, err := s.NewRepo.GetMenuItems()
	if err != nil {
		return nil, errors.New("failed to fetch menu items: " + err.Error())
	}
	return items, nil
}

// Добавление нового элемента меню
func (s *CustomService) AddMenuItem(item model.MenuItem) error {
	if item.ID == "" || item.Name == "" || item.Price <= 0 {
		return errors.New("invalid menu item data: missing required fields or invalid price")
	}
	if err := s.NewRepo.AddMenuItem(item); err != nil {
		return errors.New("failed to add menu item: " + err.Error())
	}
	return nil
}

// Получение элемента меню по ID
func (s *CustomService) GetMenuItemByID(id string) (*model.MenuItem, error) {
	if id == "" {
		return nil, errors.New("menu item ID cannot be empty")
	}
	item, err := s.NewRepo.GetMenuItemByID(id)
	if err != nil {
		return nil, errors.New("failed to fetch menu item by ID: " + err.Error())
	}
	if item == nil {
		return nil, errors.New("menu item not found")
	}
	return item, nil
}

// Обновление элемента меню по ID
func (s *CustomService) UpdateMenuItemById(item model.MenuItem) error {
	if item.ID == "" || item.Name == "" || item.Price <= 0 {
		return errors.New("invalid menu item data: missing required fields or invalid price")
	}
	if err := s.NewRepo.UpdateMenuItemById(item); err != nil {
		return errors.New("failed to update menu item: " + err.Error())
	}
	return nil
}

// Удаление элемента меню по ID
func (s *CustomService) DeleteMenuItemById(id string) error {
	if id == "" {
		return errors.New("menu item ID cannot be empty")
	}
	if err := s.NewRepo.DeleteMenuItemById(id); err != nil {
		return errors.New("failed to delete menu item: " + err.Error())
	}
	return nil
}

// GetMenuItems() ([]model.MenuItem, error)
// AddMenuItem(item model.MenuItem) error
// GetMenuItemByID(id string) (*model.MenuItem, error)
// UpdateMenuItemById(item model.MenuItem) error
// DeleteMenuItemsById(id string) error
