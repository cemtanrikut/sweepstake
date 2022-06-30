package api

import "time"

type Product_Model struct {
	Product_ID  string
	Name        string
	Description string
	Price       float32
	Fee         float32
	Create_Date time.Time
}

type Draw_Model struct {
	Draw_ID             string
	Product_Model       Product_Model
	Start_Date          time.Time
	End_Date            time.Time
	Is_Active           bool
	Min_Purchaser_Count int
	Total_Purchaser     int
	Status              string
	Winner_User         string //User model must be here
}

func getMinPurchaserCount(productModel Product_Model) float32 {
	productPrice := productModel.Price
	productFee := productModel.Fee
	expectedCount := productPrice + (productPrice * (productFee / 10))
	return expectedCount
}
