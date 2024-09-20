package service

import (
	"database/sql"

	"github.com/codepnw/hexagonal/errs"
	"github.com/codepnw/hexagonal/logs"
	"github.com/codepnw/hexagonal/repository"
)

type customerService struct {
	custRepo repository.CustomerRepository
}

func NewCustomerService(custRepo repository.CustomerRepository) CustomerService {
	return customerService{custRepo: custRepo}
}

func (s customerService) GetCustomers() ([]CustomerResponse, error) {
	customers, err := s.custRepo.GetAll()
	if err != nil {
		logs.Error(err)
		return nil, errs.NewErrUnexpected()
	}

	custResponses := []CustomerResponse{}
	for _, customer := range customers {
		response := CustomerResponse{
			CustomerID: customer.CustomerID,
			Name:       customer.Name,
			Status:     customer.Status,
		}
		custResponses = append(custResponses, response)
	}

	return custResponses, nil
}

func (s customerService) GetCustomer(id int) (*CustomerResponse, error) {
	customer, err := s.custRepo.GetById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewErrNotFound("customer not found")
		}

		logs.Error(err)
		return nil, errs.NewErrUnexpected()
	}

	response := CustomerResponse{
		CustomerID: customer.CustomerID,
		Name:       customer.Name,
		Status:     customer.Status,
	}

	return &response, nil
}
