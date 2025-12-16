package repo

import (
	"encoding/json"
	"log"

	"order-providing-system/internal/model"
)

func (r *Repo) GetMenuItems() ([]model.MenuItem, error) {
	query := `SELECT id, name, description, price, allergens, size FROM menu_items`
	rows, err := r.db.Query(query)
	if err != nil {
		log.Println("Error fetching menu items:", err)
		return nil, err
	}
	defer rows.Close()

	var menuItems []model.MenuItem
	for rows.Next() {
		var item model.MenuItem
		var allergens []byte
		if err := rows.Scan(&item.ID, &item.Name, &item.Description, &item.Price, &allergens, &item.Size); err != nil {
			return nil, err
		}
		item.Allergens = ParseAllergens(allergens)
		menuItems = append(menuItems, item)
	}
	return menuItems, nil
}

func (r *Repo) AddMenuItem(item model.MenuItem) error {
	query := `INSERT INTO menu_items (id, name, description, price, allergens, size) 
	VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := r.db.Exec(query, item.ID, item.Name, item.Description, item.Price, StringifyAllergens(item.Allergens), item.Size)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repo) GetMenuItemByID(id string) (*model.MenuItem, error) {
	query := `SELECT id, name, description, price, allergens, size FROM menu_items WHERE id = $1`
	row := r.db.QueryRow(query, id)

	var item model.MenuItem
	var allergens []byte
	if err := row.Scan(&item.ID, &item.Name, &item.Description, &item.Price, &allergens, &item.Size); err != nil {
		return nil, err
	}
	item.Allergens = ParseAllergens(allergens)
	return &item, nil
}

func (r *Repo) UpdateMenuItemById(item model.MenuItem) error {
	query := `UPDATE menu_items SET name = $2, description = $3, price = $4, allergens = $5, size = $6 WHERE id = $1`
	_, err := r.db.Exec(query, item.ID, item.Name, item.Description, item.Price, StringifyAllergens(item.Allergens), item.Size)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repo) DeleteMenuItemById(id string) error {
	query := `DELETE FROM menu_items WHERE id = $1`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

// Универсальная функция для парсинга JSON в слайс
func ParseJSON[T any](data []byte) (T, error) {
	var result T
	err := json.Unmarshal(data, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}

// Универсальная функция для преобразования слайса в JSON
func ToJSON[T any](value T) []byte {
	data, err := json.Marshal(value)
	if err != nil {
		log.Fatalf("Error converting to JSON: %v", err)
	}
	return data
}

// Парсинг аллергенов
func ParseAllergens(data []byte) []string {
	allergens, err := ParseJSON[[]string](data)
	if err != nil {
		log.Printf("Error parsing allergens JSON: %v", err)
		return nil
	}
	return allergens
}

// Преобразование аллергенов в JSON
func StringifyAllergens(allergens []string) []byte {
	return ToJSON(allergens)
}

// Menu Items:
// POST /menu: Add a new menu item.
// GET /menu: Retrieve all menu items.
// GET /menu/{id}: Retrieve a specific menu item.
// PUT /menu/{id}: Update a menu item.
// DELETE /menu/{id}: Delete a menu item.
