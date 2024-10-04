package models

import "time"

type Orders struct{
	Id int `json:"Order ID"`
	User_id string `json:"User ID"`
	Pickup_address string `json:"Origin Address"`
	Pickup_pincode int32 `json:"Origin Pincode"`
	Pickup_country string `json:"Origin Country"`
	Dropoff_address string `json:"Destination Address"`
	Dropoff_pincode int32 `json:"Destination Pincode"`
	Dropoff_country string `json:"Destination Country"`
	Created_at time.Time 
	Weight float32 `json:"Weight (in kg)"`
	Status string `json:"Status"`
	}