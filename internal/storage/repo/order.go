package repo

import (
	"database/sql"
	"log"

	"order-providing-system/internal/model"
)

// Получение всех заказов
func (r *Repo) GetOrders() ([]model.Order, error) {
	query := `
		SELECT id, customer_name, status, created_at
		FROM orders
	`
	rows, err := r.db.Query(query)
	if err != nil {
		log.Println("Error fetching orders:", err)
		return nil, err
	}
	defer rows.Close()

	var orders []model.Order
	for rows.Next() {
		var order model.Order
		if err := rows.Scan(&order.ID, &order.CustomerName, &order.Status, &order.CreatedAt); err != nil {
			log.Println("Error scanning order:", err)
			return nil, err
		}

		order.Items, err = r.getOrderItems(order.ID)
		if err != nil {
			log.Println("Error fetching order items for order:", order.ID, err)
			return nil, err
		}

		orders = append(orders, order)
	}
	return orders, nil
}

// Добавление нового заказа
func (r *Repo) AddOrder(order model.Order) error {
	tx, err := r.db.Begin()
	if err != nil {
		log.Println("Error starting transaction:", err)
		return err
	}

	// Вставка основного заказа
	query := `
		INSERT INTO orders (id, customer_name, status, created_at)
		VALUES ($1, $2, $3, NOW())
	`
	_, err = tx.Exec(query, order.ID, order.CustomerName, "open")
	if err != nil {
		log.Println("Error inserting order:", err)
		tx.Rollback()
		return err
	}

	// Вставка элементов заказа
	for _, item := range order.Items {
		err := r.addOrderItem(tx, order.ID, item)
		if err != nil {
			log.Println("Error inserting order item:", err)
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

// Получение заказа по ID
func (r *Repo) GetOrderById(id string) (*model.Order, error) {
	var err error
	query := `
		SELECT id, customer_name, status, created_at
		FROM orders
		WHERE id = $1
	`
	row := r.db.QueryRow(query, id)

	var order model.Order
	if err = row.Scan(&order.ID, &order.CustomerName, &order.Status, &order.CreatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Println("Error fetching order by ID:", err)
		return nil, err
	}

	order.Items, err = r.getOrderItems(id)
	if err != nil {
		log.Println("Error fetching order items for order ID:", id, err)
		return nil, err
	}

	return &order, nil
}

// Обновление заказа по ID
func (r *Repo) UpdateOrderById(order model.Order) error {
	tx, err := r.db.Begin()
	if err != nil {
		log.Println("Error starting transaction:", err)
		return err
	}

	// Обновление основного заказа
	query := `
		UPDATE orders SET customer_name = $2, status = $3 WHERE id = $1
	`
	_, err = tx.Exec(query, order.ID, order.CustomerName, order.Status)
	if err != nil {
		log.Println("Error updating order:", err)
		tx.Rollback()
		return err
	}

	// Удаление старых элементов заказа
	deleteQuery := `
		DELETE FROM order_items WHERE order_id = $1
	`
	_, err = tx.Exec(deleteQuery, order.ID)
	if err != nil {
		log.Println("Error deleting order items:", err)
		tx.Rollback()
		return err
	}

	// Вставка обновленных элементов заказа
	for _, item := range order.Items {
		err := r.addOrderItem(tx, order.ID, item)
		if err != nil {
			log.Println("Error inserting order item:", err)
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

// Удаление заказа по ID
func (r *Repo) DeleteOrderById(id string) error {
	query := `
		DELETE FROM orders WHERE id = $1
	`
	_, err := r.db.Exec(query, id)
	if err != nil {
		log.Println("Error deleting order by ID:", err)
		return err
	}
	return nil
}

// Закрытие заказа по ID
func (r *Repo) CloseOrderById(id string) error {
	query := `
		UPDATE orders SET status = 'close' WHERE id = $1
	`
	_, err := r.db.Exec(query, id)
	if err != nil {
		log.Println("Error closing order:", err)
		return err
	}
	return nil
}

// Вспомогательные методы для работы с элементами заказа
func (r *Repo) getOrderItems(orderID string) ([]model.OrderItem, error) {
	query := `
		SELECT product_id, quantity
		FROM order_items
		WHERE order_id = $1
	`
	rows, err := r.db.Query(query, orderID)
	if err != nil {
		log.Println("Error fetching order items:", err)
		return nil, err
	}
	defer rows.Close()

	var items []model.OrderItem
	for rows.Next() {
		var item model.OrderItem
		if err := rows.Scan(&item.ProductID, &item.Quantity); err != nil {
			log.Println("Error scanning order item:", err)
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func (r *Repo) addOrderItem(tx *sql.Tx, orderID string, item model.OrderItem) error {
	query := `
		INSERT INTO order_items (order_id, product_id, quantity)
		VALUES ($1, $2, $3)
	`
	_, err := tx.Exec(query, orderID, item.ProductID, item.Quantity)
	if err != nil {
		log.Println("Error adding order item:", err)
		return err
	}
	return nil
}
