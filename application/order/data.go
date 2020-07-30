package order

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/canhtuan97/week2/connector"
	orderPb "github.com/canhtuan97/week2/protobuff/orderpb"
	"google.golang.org/grpc/metadata"
	"log"
)

func CreateOrder(ctx  context.Context ,request *orderPb.CreateOrderRequest) (*orderPb.CreateOrderResponse, error) {
	fmt.Println("CreateOrderRequest data running ...")
	headers, _ := metadata.FromIncomingContext(ctx)
	tokenCustomer := headers["authorization"]

	strData, _ := json.Marshal(request)
	fmt.Println("data; ", string(strData))
	createOrderRequest := connector.CreateOrderRequest{
		PaymentMethod: connector.PaymentMethod{Method: request.PaymentMethod.Method},
		BillingAddress: connector.BillingAddress{
			Email:      request.BillingAddress.Email,
			Region:     request.BillingAddress.Region,
			RegionId:   int(request.BillingAddress.RegionId),
			RegionCode: request.BillingAddress.RegionCode,
			CountryId:  request.BillingAddress.CountryId,
			Street:     request.BillingAddress.Street,
			Postcode:   request.BillingAddress.Postcode,
			City:       request.BillingAddress.City,
			Telephone:  request.BillingAddress.Telephone,
			FirstName:  request.BillingAddress.Telephone,
			LastName:   request.BillingAddress.LastName,
		},
	}




	//var createOrderRequest connector.CreateOrderRequest
	//createOrderRequest.PaymentMethod.Method = request.PaymentMethod.Method
	//createOrderRequest.BillingAddress.Email = request.BillingAddress.Email
	//createOrderRequest.BillingAddress.Region = request.BillingAddress.Region
	//createOrderRequest.BillingAddress.RegionId = int(request.BillingAddress.RegionId)
	//createOrderRequest.BillingAddress.RegionCode = request.BillingAddress.RegionCode
	//createOrderRequest.BillingAddress.CountryId = request.BillingAddress.CountryId
	//createOrderRequest.BillingAddress.Street = request.BillingAddress.Street
	//createOrderRequest.BillingAddress.Postcode = request.BillingAddress.Postcode
	//createOrderRequest.BillingAddress.City = request.BillingAddress.City
	//createOrderRequest.BillingAddress.Telephone = request.BillingAddress.Telephone
	//createOrderRequest.BillingAddress.FirstName = request.BillingAddress.FirstName
	//createOrderRequest.BillingAddress.LastName = request.BillingAddress.LastName

	client := connector.NewClient()

	data, err := client.Order.CreateOrder(tokenCustomer,createOrderRequest)
	if err != nil {
		log.Fatalf(" loi cua minh %v",err)
	}
	resp := orderPb.CreateOrderResponse{
		OrderId: int32(data.OrderID),
	}
	return &resp,nil

}
