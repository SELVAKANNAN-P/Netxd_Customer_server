package controllers

import (
	"Netxd_Project1/Netxd_Customer_dal/interfaces"
	"Netxd_Project1/Netxd_Customer_dal/models"
	pro "Netxd_Project1/customer"
	"context"
)

type RPCserver struct {
	pro.UnimplementedCustomerServiceServer
}

var (
	CustomerService interfaces.ICustomer
)

func (s *RPCserver) CreateCustomer(ctx context.Context, req *pro.Customer) (*pro.CustomerResponse, error) {
	 dbCustomer := &models.Customer{
		CustomerID: req.CustomerID,
		FirstName:  req.FirstName,
	}
	result, err := CustomerService.CreateCustomer(dbCustomer)
	if err != nil {
		return nil, err
	} else {
		responseCustomer := &pro.CustomerResponse{
			CustomerID: result.CustomerID,
		}
		return responseCustomer, nil
	}
}
