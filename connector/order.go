package connector

import (
	"encoding/json"
	"fmt"
	"log"
)

type CreateOrderRequest struct {
	PaymentMethod  PaymentMethod  `json:"payment_method"`
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
	OrderID int `json:"order_id"`
}
type OrderService interface {
	CreateOrder(tokenCustomer []string,createOrderRequest CreateOrderRequest) (*CreateOrderResponse,error)
}

type Order struct {
	client *Client
}

func (c Order) CreateOrder(tokenCustomer []string,createOrderRequest CreateOrderRequest) (*CreateOrderResponse,error){

	url := c.client.UrlMagento + urlCreateOrder
	fmt.Println(url)
	dataConvert, err := json.Marshal(createOrderRequest)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.CreateRequest(url,tokenCustomer, dataConvert)
	if err != nil {
		log.Fatal(err)
	}

	createOrderResponse := CreateOrderResponse{}
	//-------------Đoạn này đang viết check lỗi
	//data, err := CheckResponse(resp)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//
	//if err == nil {
	//	log.Fatal(err)
	//}
	//--------------------------------------
	json.Unmarshal(resp, &createOrderResponse)
	fmt.Println(string(resp))
	return &createOrderResponse, nil
}