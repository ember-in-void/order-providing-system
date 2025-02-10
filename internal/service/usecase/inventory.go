package usecase

import (
	"errors"

	"frappuccino/internal/model"
)

// Получение всех элементов инвентаря
func (s *CustomService) GetInventoryItems() ([]model.InventoryItem, error) {
	items, err := s.NewRepo.GetInventoryItems()
	if err != nil {
		return nil, errors.New("failed to fetch inventory items: " + err.Error())
	}
	return items, nil
}

// Добавление нового элемента инвентаря
func (s *CustomService) AddInventoryItem(item model.InventoryItem) error {
	// Проверка данных
	if item.IngredientID == "" || item.Name == "" || item.Quantity < 0 || item.Unit == "" {
		return errors.New("invalid inventory item data: missing or incorrect fields")
	}

	// Добавление в репозиторий
	if err := s.NewRepo.AddInventoryItem(item); err != nil {
		return errors.New("failed to add inventory item: " + err.Error())
	}
	return nil
}

// Получение элемента инвентаря по ID
func (s *CustomService) GetInventoryItemByID(id string) (*model.InventoryItem, error) {
	if id == "" {
		return nil, errors.New("inventory item ID cannot be empty")
	}

	item, err := s.NewRepo.GetInventoryItemByID(id)
	if err != nil {
		return nil, errors.New("failed to fetch inventory item by ID: " + err.Error())
	}

	if item == nil {
		return nil, errors.New("inventory item not found")
	}

	return item, nil
}

// Обновление элемента инвентаря по ID
func (s *CustomService) UpdateInventoryItemById(item model.InventoryItem) error {
	// Проверка данных
	if item.IngredientID == "" || item.Name == "" || item.Quantity < 0 || item.Unit == "" {
		return errors.New("invalid inventory item data: missing or incorrect fields")
	}

	// Обновление в репозитории
	if err := s.NewRepo.UpdateInventoryItemById(item); err != nil {
		return errors.New("failed to update inventory item: " + err.Error())
	}

	return nil
}

// Удаление элемента инвентаря по ID
func (s *CustomService) DeleteInventoryItemById(id string) error {
	if id == "" {
		return errors.New("inventory item ID cannot be empty")
	}

	if err := s.NewRepo.DeleteInventoryItemById(id); err != nil {
		return errors.New("failed to delete inventory item: " + err.Error())
	}

	return nil
}
