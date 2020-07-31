package connector

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	apiKey            = "71qhxcusonwayh5jeynhupjci7v2ecob"
	urlMagento        = "https://magento23demo.connectpos.com"
	urlCreateCustomer = "/rest/V1/customers"
	UrlGetAccessToken = "/rest/V1/integration/customer/token"
	urlAddItemToCart  = "/rest/default/V1/carts/mine/items"
	UrlQuote          = "/rest/V1/carts/mine"
	urlCreateOrder    = "/rest/V1/carts/mine/payment-information"
)

var urlEstimateShipping = [2]string{"/rest/V1/carts/", "/estimate-shipping-methods"}
var urlCreateInvoice = [2]string{"/rest/V1/order/", "/invoice"}
var urlCreateShipment = [2]string{"/rest/V1/order/", "/ship"}

type Client struct {
	Client     *http.Client
	ApiKey     string
	UrlMagento string
	Customers  CustomersService
	Carts      CartServices
	Order      OrderService
	Invoice    InvoiceService
	Shipment   ShipmentService
}

func NewClient() *Client {
	httpClient := http.DefaultClient
	c := &Client{Client: httpClient}
	c.ApiKey = apiKey
	c.UrlMagento = urlMagento
	c.Customers = &Customers{client: c}
	c.Carts = &Cart{client: c}
	c.Order = &Order{client: c}
	c.Invoice = &Invoice{client: c}
	c.Shipment = &Shipment{client: c}
	return c
}
func (c Client) CreateRequestPost(url string, body []byte) (*http.Response, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	bearer := "bearer " + apiKey
	req.Header.Set("Authorization", bearer)
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.Client.Do(req)

	if err != nil {
		return nil, err
	}
	return resp, nil

}

func (c Client) CreateRequestPostV2(url string, body []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	bearer := "bearer " + apiKey

	req.Header.Set("Authorization", bearer)
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.Client.Do(req)

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return data, nil

}

func (c Client) CreateRequest(url string, tokenCustomer []string, body []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	fmt.Println(string(tokenCustomer[0]))
	req.Header.Set("Authorization", string(tokenCustomer[0]))
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.Client.Do(req)

	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(resp.Body)

	return data, nil

}


func CheckResponseError(r *http.Response) error {
	if r.StatusCode >= 200 && r.StatusCode < 300 {
		return nil
	}

	bodyBytes, err := ioutil.ReadAll(r.Body)

	if err != nil {
		return err
	}

	responseError := &ResponseError{}

	if len(bodyBytes) > 0 {
		err := json.Unmarshal(bodyBytes, responseError)

		if err != nil {
			return ResponseDecodingError{
				Body:    bodyBytes,
				Message: err.Error(),
				Status:  r.StatusCode,
			}
		}
	}

	return responseError
}


type ResponseDecodingError struct {
	Body    []byte
	Message string
	Status  int
}

type ResponseError struct {
	RequestId    string `json:"request_id,omitempty"`
	TimeUsed     int    `json:"time_used,omitempty"`
	ErrorMessage string `json:"error_message,omitempty"`
}

func (e ResponseDecodingError) Error() string {
	return e.Message
}

func (e ResponseError) Error() string {
	if e.ErrorMessage != "" {
		return e.ErrorMessage
	}

	return "Unknown Error"
}