package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net"

	"github.com/canhtuan97/week2/application/cart"
	"github.com/canhtuan97/week2/application/customer"
	"github.com/canhtuan97/week2/protobuff/cartpb"
	"github.com/canhtuan97/week2/protobuff/customerpb"
	"google.golang.org/grpc"

)

type server struct{}

func (s *server) EstimateShipping(ctx context.Context, request *cartPb.EstimateShippingRequest) (*cartPb.EstimateShippingResponse, error) {
	fmt.Println("EstimateShipping running ...")
	data , err := cart.EstimateShipping(request)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s *server) AddItemProductConfigurable(ctx context.Context, request *cartPb.AddItemProductConfigurableRequest) (*cartPb.AddItemProductConfigurableResponse, error) {
	data ,err := cart.AddItemProductConfigurable(request)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s *server) GetQuoteIdCustomer(ctx context.Context, request *customerPb.GetQuoteIdCustomerRequest) (*customerPb.GetQuoteIdCustomerResponse, error) {
	panic("implement me")
}

func (s *server) GetAccessTokenCustomer(ctx context.Context, request *customerPb.GetAccessTokenCustomerRequest) (*customerPb.GetAccessTokenCustomerResponse, error) {
	accessToken, err := customer.GetAccessTokenCustomer(request)
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(accessToken.Body)
	if err != nil {
		return nil, err
	}
	resp := &customerPb.GetAccessTokenCustomerResponse{
		AccessToken: string(data),
	}
	return resp, nil
}

func (s *server) AddItemProductSimple(ctx context.Context, request *cartPb.AddItemProductSimpleRequest) (*cartPb.AddItemProductSimpleResponse, error) {
	fmt.Println("AddItemProductSimple running ...")
	data, err := cart.AddItemProductSimple(request)
	if err != nil {
		return nil, err
	}

	resp := &cartPb.AddItemProductSimpleResponse{
		ItemId:      int32(data.ItemId),
		Sku:         data.Sku,
		Qty:         int32(data.Qty),
		Name:        data.Name,
		Price:       data.Price,
		ProductType: data.ProductType,
		QuoteId:     data.QuoteId,
	}
	return resp, nil
}

func (s *server) CreateCustomer(ctx context.Context, req *customerPb.CreateCustomerRequest) (*customerPb.CreateCustomerResponse, error) {
	log.Println("Create customer running...")
	data, err := customer.CreateCustomer(req)
	if err != nil {
		return nil, err
	}

	resp := &customerPb.CreateCustomerResponse{
		Id:                     int32(data.Id),
		GroupId:                int32(data.GroupId),
		DefaultBilling:         data.DefaultBilling,
		DefaultShipping:        data.DefaultShipping,
		CreatedAt:              data.CreatedAt,
		UpdatedAt:              data.UpdatedAt,
		CreatedIn:              data.CreatedIn,
		Email:                  data.Email,
		FirstName:              data.FirstName,
		LastName:               data.LastName,
		StoreId:                int32(data.StoreId),
		WebsiteId:              int32(data.WebsiteId),
		DisableAutoGroupChange: int32(data.DisableAutoGroupChange),
	}
	return resp, nil

}

func main() {

	lis, err := net.Listen("tcp", "0.0.0.0:50069")
	if err != nil {
		log.Fatalf("err while create listen %v", err)
	}

	s := grpc.NewServer()

	customerPb.RegisterCustomerServer(s, &server{})
	cartPb.RegisterAddItemProductServer(s, &server{})
	fmt.Println("Server running ...")

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("err while serve %v", err)
	}
}
