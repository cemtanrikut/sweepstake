package api

import "time"

type User_Model struct {
	User_ID              string
	Name                 string
	Surname              string
	Email                string
	Phone                string
	Province             string
	District             string
	Address              string
	Citizenship_Number   string
	Is_Active            bool
	Create_Date          time.Time
	User_Draw_Data_Model User_Draw_Data_Model
}

type User_Draw_Data_Model struct {
	Join_Count int
	Win_Count  int
	Lost_Count int
}

type User_Joined_Draw struct {
	User_ID string
	Draw_ID string
	Fee     float32
	Status  bool
}
