package repository

import "errors"

type customerRepositoryMock struct {
	customers []Customer
}

func NewCustomerRepositoryMock() CustomerRepository {
	customers := []Customer{
		{CustomerID: 1001, Name: "Jack", City: "Pattaya", ZipCode: "32102", DateOfBirth: "1990-12-20", Status: 1},
		{CustomerID: 1002, Name: "Mac", City: "Bangkok", ZipCode: "11102", DateOfBirth: "2000-11-03", Status: 2},
		{CustomerID: 1003, Name: "John", City: "Surin", ZipCode: "66102", DateOfBirth: "2012-03-03", Status: 2},
	}
	return customerRepositoryMock{customers: customers}
}

func (m customerRepositoryMock) GetAll() ([]Customer, error) {
	return m.customers, nil
}

func (m customerRepositoryMock) GetById(id int) (*Customer, error) {
	for _, customer := range m.customers {
		if customer.CustomerID == id {
			return &customer, nil
		}
	}

	return nil, errors.New("customer not found")
}
