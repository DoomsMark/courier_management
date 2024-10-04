package models

import "time"

type Trips struct{
	Trip_id int `json:"Shipment Leg"`
	Order_id int `json:"Order ID"`
	Status string `json:"Status"`
	Mode string `json:"Mode of transport"`
	Last_Update time.Time `json:"Last Updated"`
}