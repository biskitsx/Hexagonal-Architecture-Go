package repository

import (
	"errors"
)

type customerRepositoryMock struct {
	customers []Customer
}

func NewCustomerRepositoryMock() CustomerRepository {
	customers := []Customer{
		{CustomerID: 1, Name: "suphakit", DateOfBirth: "10/07/2002", City: "Bangkok", ZipCode: "73140", Status: 5},
		{CustomerID: 2, Name: "genezy", DateOfBirth: "12/04/2003", City: "Ratchaburi", ZipCode: "73145", Status: 2},
	}
	return &customerRepositoryMock{customers: customers}
}

func (r customerRepositoryMock) GetAll() ([]Customer, error) {
	return r.customers, nil
}

func (r customerRepositoryMock) GetById(id int) (*Customer, error) {
	for _, customer := range r.customers {
		if customer.CustomerID == id {
			return &customer, nil
		}
	}
	return nil, errors.New("customer not found")
}
