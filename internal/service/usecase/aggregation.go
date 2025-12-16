package usecase

import (
	"errors"

	"order-providing-system/internal/model"
)

// Получение общей суммы продаж
func (s *CustomService) GetTotalSales() (float64, error) {
	totalSales, err := s.NewRepo.GetTotalSales()
	if err != nil {
		return 0, errors.New("failed to fetch total sales: " + err.Error())
	}
	return totalSales, nil
}

// Получение популярных товаров
func (s *CustomService) GetPopularItems() ([]model.MenuItem, error) {
	popularItems, err := s.NewRepo.GetPopularItems()
	if err != nil {
		return nil, errors.New("failed to fetch popular items: " + err.Error())
	}
	return popularItems, nil
}
