package api

import "time"

type Product_Model struct {
	Product_ID  string
	Name        string
	Description string
	Fee         float32
	Create_Date time.Time
}
