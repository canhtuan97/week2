package connector

import (
	"encoding/json"
	"fmt"
	"log"
)

type CreateOrderRequest struct {
	PaymentMethod  PaymentMethod  `json:"paymentMethod"`
	BillingAddress BillingAddress `json:"billing_address"`
}

type PaymentMethod struct {
	Method string `json:"method"`
}

type BillingAddress struct {
	Email      string   `json:"email"`
	Region     string   `json:"region"`
	RegionId   int      `json:"region_id"`
	RegionCode string   `json:"region_code"`
	CountryId  string   `json:"country_id"`
	Street     []string `json:"street"`
	Postcode   string   `json:"postcode"`
	City       string   `json:"city"`
	Telephone  string   `json:"telephone"`
	FirstName  string   `json:"firstname"`
	LastName   string   `json:"lastname"`
}

type CreateOrderResponse struct {
	OrderID string `json:"order_id"`
}
type OrderService interface {
	CreateOrder(quoteId string,tokenCustomer []string,createOrderRequest CreateOrderRequest) (string,error)
}

type Order struct {
	client *Client
}

func (c Order) CreateOrder(quoteId string,tokenCustomer []string,createOrderRequest CreateOrderRequest) (string ,error){

	url := c.client.UrlMagento + urlCreateOrder
	fmt.Println(url)
	dataConvert, err := json.Marshal(createOrderRequest)


	fmt.Println("day la data push",string(dataConvert))
	resp, err := c.client.CreateRequest(url,tokenCustomer, dataConvert)
	if err != nil {
		log.Fatal(err)
	}
	orderID := string(resp)

	fmt.Println("orderID",orderID)
	fmt.Println(string(resp))
	return orderID, nil
}