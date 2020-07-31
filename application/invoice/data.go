package invoice

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/canhtuan97/week2/connector"
	invoicePb "github.com/canhtuan97/week2/protobuff/invoicepb"

	"google.golang.org/grpc/metadata"
	"log"
)

func CreateOrder(ctx context.Context, request *invoicePb.CreateInvoiceRequest) (*invoicePb.CreateInvoiceResponse, error) {
	fmt.Println("CreateOrder data running ...")
	headers, _ := metadata.FromIncomingContext(ctx)
	tokenCustomer := headers["authorization"]

	strData, _ := json.Marshal(request)
	fmt.Println("data; ", string(strData))

	createInvoiceRequest := connector.CreateInvoiceRequest{
		Capture: request.Capture,
		Notify:  request.Notify,
	}
	orderId := request.OrderId
	client := connector.NewClient()
	data, err := client.Invoice.CreateInvoice(orderId, tokenCustomer, createInvoiceRequest)

	if err != nil {
		log.Fatalf(" loi cua minh %v", err)
	}
	fmt.Println(data)
	resp := invoicePb.CreateInvoiceResponse{
		InvoiceId: data,
	}
	return &resp, nil
}
