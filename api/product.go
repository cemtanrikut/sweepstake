package api

import "time"

type Product_Model struct {
	Product_ID  string
	Name        string
	Description string
	Fee         float32
	Create_Date time.Time
}

type Draw_Model struct {
	Draw_ID         string
	Product_Model   Product_Model
	Start_Date      time.Time
	End_Date        time.Time
	Is_Active       bool
	Min_Purchasers  int
	Total_Purchaser int
	Status          string
	Winner_User     string //User model must be here
}
