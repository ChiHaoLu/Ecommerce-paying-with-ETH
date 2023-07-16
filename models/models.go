package models

import(
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct{
	ID
	First_Name
	Last_Name
	Password
	Email
	Phone
	Token
	Refresh_Token
	Created_At
	Updated_At
	User_ID
	UserCart
	Address_Details
	Order_Status
}

type Product struct{


}

type ProductUser struct{

}

type Address struct{

}

type Order struct{

}

type Payment struct{
	
}