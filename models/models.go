package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	First_Name      *string            `json:"first_name" validate:"required,min=2,max=30"`
	Last_Name       *string            `json:"last_name" validate:"required,min=2,max=30"`
	Email           *string            `json:"email" validate:"required,email"`
	Password        *string            `json:"password" validate:"required,min=6,max=100"`
	Phone           *string            `json:"phone" validate:"required,min=10,max=15"`
	Token           *string            `json:"token"`
	Refresh_Token   *string            `json:"refresh_token"`
	Created_At      time.Time          `json:"created_at"`
	Updated_At      time.Time          `json:"updated_at"`
	User_ID         string             `json:"user_id"`
	UserCart        []ProductUser      `json:"user_cart"`
	Address_Details []Address          `json:"address_details"`
	Order_Status    []Order            `json:"order_status"`
}

type Product struct {
	Product_ID   primitive.ObjectID `bson:"_id,omitempty"`
	Product_Name *string            `json:"product_name"`
	Price        *uint64            `json:"price"`
	Rating       *uint64            `json:"rating"`
	Image        *string            `json:"image"`
}

type ProductUser struct {
	Product_ID   primitive.ObjectID `bson:"_id,omitempty"`
	Product_Name *string            `json:"product_name"`
	Price        *uint64            `json:"price"`
	Rating       *uint64            `json:"rating"`
	Image        *string            `json:"image"`
}

type Address struct {
	Address_ID primitive.ObjectID `bson:"_id,omitempty"`
	House      *string            `json:"house"`
	Street     *string            `json:"street"`
	City       *string            `json:"city"`
	State      *string            `json:"state"`
	Pincode    *string            `json:"pincode"`
}

type Order struct {
	Order_ID       primitive.ObjectID `bson:"_id,omitempty"`
	Ordered_Cart   []ProductUser      `json:"ordered_cart"`
	Ordered_At     time.Time          `json:"ordered_at"`
	Price          *uint64            `json:"price"`
	Discount       *uint64            `json:"discount"`
	Payment_Method Payment            `json:"payment_method"`
}
type Payment struct {
	Digit bool `json:"digit"`
	COD   bool `json:"cod"`
}
