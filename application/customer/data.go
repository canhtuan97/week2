package customer

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/canhtuan97/week2/connector"
	"github.com/canhtuan97/week2/protobuff/customerpb"
	"google.golang.org/grpc/metadata"
	"log"
	"strconv"
)

func CreateCustomer(request *customerPb.CreateCustomerRequest) (*connector.Customer, error) {
	var pbAddresses []connector.Addresses
	for _, requestedAddress := range request.Customer.Addresses {
		pbAddresses = append(pbAddresses, connector.Addresses{
			DefaultShipping: requestedAddress.DefaultShipping,
			DefaultBilling:  requestedAddress.DefaultBilling,
			FirstName:       requestedAddress.FirstName,
			LastName:        requestedAddress.LastName,
			Region: connector.Region{
				RegionCode: requestedAddress.Region.RegionCode,
				Region:     requestedAddress.Region.Region,
				RegionId:   int(requestedAddress.Region.RegionId),
			},
			Postcode:  requestedAddress.Postcode,
			Street:    requestedAddress.Street,
			City:      requestedAddress.City,
			Telephone: requestedAddress.Telephone,
			CountryId: requestedAddress.CountryId,
		})
	}

	var customer connector.Customer
	customer.Email = request.Customer.Email
	customer.FirstName = request.Customer.FirstName
	customer.LastName = request.Customer.LastName
	customer.Addresses = pbAddresses

	createCustomerRequest := connector.CreateCustomerRequest{
		Customer: customer,
		Password: request.Password,
	}

	client := connector.NewClient()
	data, err := client.Customers.CreateCustomer(createCustomerRequest)
	if err != nil {
		log.Fatalf(" loi cua minh%v", err)
	}
	return data, nil
}

func GetAccessTokenCustomer(request *customerPb.GetAccessTokenCustomerRequest) (*customerPb.GetAccessTokenCustomerResponse, error) {
	client := connector.NewClient()
	url := client.UrlMagento + connector.UrlGetAccessToken
	fmt.Println(url)
	strData, _ := json.Marshal(request)
	fmt.Println("data; ", string(strData))

	resp, err := client.CreateRequestPostV2(url, strData)
	if err != nil {
		log.Fatal(err)
	}

	type GetTokenCustomerResponse struct {
		Token string `json:"token"`
	}
	getTokenCustomerResponse := GetTokenCustomerResponse{}
	json.Unmarshal(resp, &getTokenCustomerResponse)
	fmt.Println("day la data", string(resp))

	respData := &customerPb.GetAccessTokenCustomerResponse{
		AccessToken: getTokenCustomerResponse.Token,
	}
	return respData, nil
}

func GetQuoteIdCustomer(ctx context.Context, request *customerPb.GetQuoteIdCustomerRequest) (*customerPb.GetQuoteIdCustomerResponse, error) {
	fmt.Println("GetQuoteIdCustomer data running ...")
	headers, _ := metadata.FromIncomingContext(ctx)
	tokenCustomer := headers["authorization"]

	client := connector.NewClient()
	url := client.UrlMagento + connector.UrlQuote
	fmt.Println(url)

	resp, err := client.CreateRequest(url, tokenCustomer, nil)
	if err != nil {
		log.Fatal(err)
	}

	byteToInt, _ := strconv.Atoi(string(resp))
	fmt.Println(byteToInt)

	quoteIdResponse := QuoteIdResponse{}
	json.Unmarshal(resp, &quoteIdResponse)


	respClient := &customerPb.GetQuoteIdCustomerResponse{
		QuoteId: fmt.Println(byteToInt),
	}
	return respClient, nil

}
