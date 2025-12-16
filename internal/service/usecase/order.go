package usecase

import (
	"errors"

	"order-providing-system/internal/model"
	
)

// Получение всех заказов
func (s *CustomService) GetOrders() ([]model.Order, error) {
	orders, err := s.NewRepo.GetOrders()
	if err != nil {
		return nil, errors.New("failed to fetch orders: " + err.Error())
	}
	return orders, nil
}

// Создание нового заказа
func (s *CustomService) CreateOrder(order model.Order) error {
	// Проверка данных
	if order.CustomerName == "" || len(order.Items) == 0 {
		return errors.New("invalid order data: missing customer name or items")
	}
	for _, item := range order.Items {
		if item.ProductID == "" || item.Quantity <= 0 {
			return errors.New("invalid order item data: missing product ID or quantity")
		}
	}

	// Создание заказа в репозитории
	if err := s.NewRepo.AddOrder(order); err != nil {
		return errors.New("failed to add order: " + err.Error())
	}
	return nil
}

// Получение заказа по ID
func (s *CustomService) GetOrderById(id string) (*model.Order, error) {
	if id == "" {
		return nil, errors.New("order ID cannot be empty")
	}

	order, err := s.NewRepo.GetOrderById(id)
	if err != nil {
		return nil, errors.New("failed to fetch order by ID: " + err.Error())
	}

	if order == nil {
		return nil, errors.New("order not found")
	}

	return order, nil
}

// Обновление заказа по ID
func (s *CustomService) UpdateOrderById(order model.Order) error {
	// Проверка данных
	if order.ID == "" || order.CustomerName == "" || len(order.Items) == 0 {
		return errors.New("invalid order data: missing required fields")
	}
	for _, item := range order.Items {
		if item.ProductID == "" || item.Quantity <= 0 {
			return errors.New("invalid order item data: missing product ID or quantity")
		}
	}

	// Обновление заказа в репозитории
	if err := s.NewRepo.UpdateOrderById(order); err != nil {
		return errors.New("failed to update order: " + err.Error())
	}

	return nil
}

// Удаление заказа по ID
func (s *CustomService) DeleteOrderById(id string) error {
	if id == "" {
		return errors.New("order ID cannot be empty")
	}

	if err := s.NewRepo.DeleteOrderById(id); err != nil {
		return errors.New("failed to delete order: " + err.Error())
	}

	return nil
}

// Закрытие заказа по ID
func (s *CustomService) CloseOrderById(id string) error {
	if id == "" {
		return errors.New("order ID cannot be empty")
	}

	if err := s.NewRepo.CloseOrderById(id); err != nil {
		return errors.New("failed to close order: " + err.Error())
	}

	return nil
}
