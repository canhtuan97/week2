package connector

import (
	"encoding/json"
	"fmt"
	"log"
)

type CreateInvoiceRequest struct {
	Capture bool `json:"capture"`
	Notify  bool `json:"notify"`
}

type Invoice struct {
	client *Client
}

type InvoiceService interface {
	CreateInvoice(orderId string, tokenCustomer []string, createInvoiceRequest CreateInvoiceRequest) (string, error)
}

func (c Invoice) CreateInvoice(orderId string, tokenCustomer []string, createInvoiceRequest CreateInvoiceRequest) (string, error) {
	url := c.client.UrlMagento + urlCreateInvoice[0] + orderId + urlCreateInvoice[1]
	fmt.Println(url)
	dataConvert, err := json.Marshal(createInvoiceRequest)

	fmt.Println("day la data push", string(dataConvert))
	resp, err := c.client.CreateRequest(url, tokenCustomer, dataConvert)
	if err != nil {
		log.Fatal(err)
	}
	invoicId := string(resp)

	fmt.Println("invoicId", invoicId)
	fmt.Println(string(invoicId))
	return invoicId, nil
}
