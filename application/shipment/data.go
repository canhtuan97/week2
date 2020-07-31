package shipment

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/canhtuan97/week2/connector"
	shipmentPb "github.com/canhtuan97/week2/protobuff/shipmentpb"
	"google.golang.org/grpc/metadata"
	"log"
)

func CreateShipment(ctx context.Context, request *shipmentPb.CreateShipmentRequest) (*shipmentPb.CreateShipmentResponse, error) {
	fmt.Println("CreateOrderRequest data running ...")
	headers, _ := metadata.FromIncomingContext(ctx)
	tokenCustomer := headers["authorization"]

	strData, _ := json.Marshal(request)
	fmt.Println("data; ", string(strData))

	var pbItems []connector.Items
	var pbTracks []connector.Tracks

	for _, index := range request.Items {
		pbItems = append(pbItems, connector.Items{
			OrderItemId: int(index.OrderItemId),
			Qty:         int(index.Qty),
		})
	}

	for _, index := range request.Tracks {
		pbTracks = append(pbTracks, connector.Tracks{
			TrackNumber: index.TrackNumber,
			Title:       index.Title,
			CarrierCode: index.CarrierCode,
		})
	}

	createShipmentRequest := connector.CreateShipmentRequest{
		Items:  pbItems,
		Tracks: pbTracks,
	}
	orderId := request.OrderId
	fmt.Println(orderId)
	client := connector.NewClient()
	data, err := client.Shipment.CreateShipment(orderId, tokenCustomer, createShipmentRequest)

	if err != nil {
		log.Fatalf(" loi cua minh %v", err)
	}
	fmt.Println(data)
	resp := shipmentPb.CreateShipmentResponse{
		ShipmentId: data,
	}
	return &resp, nil

}
