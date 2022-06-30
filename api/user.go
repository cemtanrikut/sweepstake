package api

import "time"

type User_Model struct {
	User_ID              string
	User_SignIn          User_SignIn
	Address              User_Address
	Citizenship_Number   string
	Is_Active            bool
	Create_Date          time.Time
	User_Draw_Data_Model User_Draw_Data_Model
}

//Address info
type User_Address struct {
	Phone    string
	Province string
	District string
	Address  string
}

//Sign in info
type User_SignIn struct {
	Name     string
	Surname  string
	Email    string
	Password string
}

//Datas from draw
type User_Draw_Data_Model struct {
	Join_Count int
	Win_Count  int
	Lost_Count int
}

//User's joined draws
type User_Joined_Draw struct {
	User_ID string
	Draw_ID string
	Fee     float32
	Status  bool
}
