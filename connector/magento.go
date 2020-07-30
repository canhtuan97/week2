package connector

import (
	"bytes"
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

type Client struct {
	Client     *http.Client
	ApiKey     string
	UrlMagento string
	Customers  CustomersService
	Carts      CartServices
	Order      OrderService
}

func NewClient() *Client {
	httpClient := http.DefaultClient
	c := &Client{Client: httpClient}
	c.ApiKey = apiKey
	c.UrlMagento = urlMagento
	c.Customers = &Customers{client: c}
	c.Carts = &Cart{client: c}
	c.Order = &Order{client: c}
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

func CheckResponse(res *http.Response) ([]byte, error) {
	if res.Status == "400" {
		return nil, nil
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}
