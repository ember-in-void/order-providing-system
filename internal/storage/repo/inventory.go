package repo

import (
	"database/sql"
	"log"

	"frappuccino/internal/model"
)

// Получение всех элементов инвентаря
func (r *Repo) GetInventoryItems() ([]model.InventoryItem, error) {
	query := `SELECT ingredient_id, name, quantity, unit FROM inventory`
	rows, err := r.db.Query(query)
	if err != nil {
		log.Println("Error fetching inventory items:", err)
		return nil, err
	}
	defer rows.Close()

	var items []model.InventoryItem
	for rows.Next() {
		var item model.InventoryItem
		if err := rows.Scan(&item.IngredientID, &item.Name, &item.Quantity, &item.Unit); err != nil {
			log.Println("Error scanning inventory item:", err)
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

// Добавление нового элемента инвентаря
func (r *Repo) AddInventoryItem(item model.InventoryItem) error {
	query := `INSERT INTO inventory (ingredient_id, name, quantity, unit) 
              VALUES ($1, $2, $3, $4)`
	_, err := r.db.Exec(query, item.IngredientID, item.Name, item.Quantity, item.Unit)
	if err != nil {
		log.Println("Error adding inventory item:", err)
		return err
	}
	return nil
}

// Получение элемента инвентаря по ID
func (r *Repo) GetInventoryItemByID(id string) (*model.InventoryItem, error) {
	query := `SELECT ingredient_id, name, quantity, unit FROM inventory WHERE ingredient_id = $1`
	row := r.db.QueryRow(query, id)

	var item model.InventoryItem
	if err := row.Scan(&item.IngredientID, &item.Name, &item.Quantity, &item.Unit); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Элемент не найден
		}
		log.Println("Error fetching inventory item by ID:", err)
		return nil, err
	}
	return &item, nil
}

// Обновление элемента инвентаря по ID
func (r *Repo) UpdateInventoryItemById(item model.InventoryItem) error {
	query := `UPDATE inventory SET name = $2, quantity = $3, unit = $4 WHERE ingredient_id = $1`
	_, err := r.db.Exec(query, item.IngredientID, item.Name, item.Quantity, item.Unit)
	if err != nil {
		log.Println("Error updating inventory item:", err)
		return err
	}
	return nil
}

// Удаление элемента инвентаря по ID
func (r *Repo) DeleteInventoryItemById(id string) error {
	query := `DELETE FROM inventory WHERE ingredient_id = $1`
	_, err := r.db.Exec(query, id)
	if err != nil {
		log.Println("Error deleting inventory item:", err)
		return err
	}
	return nil
}
