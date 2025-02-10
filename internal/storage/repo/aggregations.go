package repo

import (
	"database/sql"
	"log"

	"frappuccino/internal/model"
)

// Получение общей суммы продаж
func (r *Repo) GetTotalSales() (float64, error) {
	query := `SELECT SUM(total_amount) AS total_sales FROM orders WHERE status = 'close'`
	var totalSales float64

	err := r.db.QueryRow(query).Scan(&totalSales)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil // Если нет закрытых заказов, сумма продаж равна 0
		}
		log.Println("Error fetching total sales:", err)
		return 0, err
	}

	return totalSales, nil
}

// Получение популярных товаров
func (r *Repo) GetPopularItems() ([]model.MenuItem, error) {
	query := `
		SELECT mi.id, mi.name, SUM(oi.quantity) AS total_sold
		FROM menu_items mi
		JOIN order_items oi ON mi.id = oi.menu_item_id
		JOIN orders o ON oi.order_id = o.id
		WHERE o.status = 'close'
		GROUP BY mi.id, mi.name
		ORDER BY total_sold DESC
		LIMIT 10
	`

	rows, err := r.db.Query(query)
	if err != nil {
		log.Println("Error fetching popular items:", err)
		return nil, err
	}
	defer rows.Close()

	var items []model.MenuItem
	for rows.Next() {
		var item model.MenuItem
		var totalSold int

		if err := rows.Scan(&item.ID, &item.Name, &totalSold); err != nil {
			log.Println("Error scanning popular item row:", err)
			return nil, err
		}
		item.Description = "" // Можно оставить пустым или заполнить из другой таблицы, если потребуется
		item.Price = 0        // Цена не имеет значения в контексте популярных товаров
		item.Allergens = nil  // Оставить пустым
		item.Size = ""        // Оставить пустым

		items = append(items, item)
	}

	return items, nil
}
