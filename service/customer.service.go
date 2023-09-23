package service

import (
	"database/sql"

	"github.com/biskitsx/Hexagonal-Architecture-Go/errs"
	"github.com/biskitsx/Hexagonal-Architecture-Go/logs"
	"github.com/biskitsx/Hexagonal-Architecture-Go/repository"
)

type customerService struct {
	customerRepo repository.CustomerRepository
}

func NewCustomerService(customerRepo repository.CustomerRepository) CustomerService {
	return &customerService{
		customerRepo: customerRepo,
	}
}

func (s customerService) GetCustomers() ([]CustomerResponse, error) {
	customers, err := s.customerRepo.GetAll()
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	customerRes := []CustomerResponse{}
	for _, customer := range customers {
		subCustomerRes := CustomerResponse{
			CustomerID: customer.CustomerID,
			Name:       customer.Name,
			Status:     customer.Status,
		}
		customerRes = append(customerRes, subCustomerRes)
	}
	return customerRes, nil
}

func (s customerService) GetCustomer(id int) (*CustomerResponse, error) {
	customer, err := s.customerRepo.GetById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotfoundError("customer not found")
		}
		logs.Error(err)
		return nil, errs.NewUnexpectError()
	}
	customerRes := CustomerResponse{
		CustomerID: customer.CustomerID,
		Name:       customer.Name,
		Status:     customer.Status,
	}
	return &customerRes, nil
}
