package model

import "time"

type Aggregations struct {
	TotalOrders    int       `json:"total_orders"`
	TotalRevenue   float64   `json:"total_revenue"`
	TotalInventory int       `json:"total_inventory"`
	ReportDate     time.Time `json:"report_date"`
}
