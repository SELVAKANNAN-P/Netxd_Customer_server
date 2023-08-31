package controllers

import (
	"context"

	pro "github.com/SELVAKANNAN-P/Customer"
	"github.com/SELVAKANNAN-P/Netxd_Dal/interfaces"
	"github.com/SELVAKANNAN-P/Netxd_Dal/models"
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
