package connector

import (
	"encoding/json"
	"fmt"
	"log"
)

type CreateShipmentRequest struct {
	Items  []Items  `json:"items"`
	Tracks []Tracks `json:"tracks"`
}

type Items struct {
	OrderItemId int `json:"order_item_id"`
	Qty         int `json:"qty"`
}

type Tracks struct {
	TrackNumber string `json:"track_number"`
	Title       string `json:"title"`
	CarrierCode string `json:"carrier_code"`
}

type Shipment struct {
	client *Client
}

type ShipmentService interface {
	CreateShipment(orderId string, tokenCustomer []string, createShipmentRequest CreateShipmentRequest) (string, error)
}

func (c Shipment) CreateShipment(orderId string, tokenCustomer []string, createShipmentRequest CreateShipmentRequest) (string, error) {
	url := c.client.UrlMagento + urlCreateShipment[0] + orderId + urlCreateShipment[1]
	fmt.Println(url)
	dataConvert, err := json.Marshal(createShipmentRequest)

	fmt.Println("day la data push", string(dataConvert))
	resp, err := c.client.CreateRequest(url, tokenCustomer, nil)
	if err != nil {
		log.Fatal(err)
	}
	shipmentId := string(resp)

	fmt.Println("invoicId", shipmentId)
	fmt.Println(string(shipmentId))
	return shipmentId, nil
}
