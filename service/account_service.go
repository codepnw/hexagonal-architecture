package service

import (
	"strings"
	"time"

	"github.com/codepnw/hexagonal/errs"
	"github.com/codepnw/hexagonal/logs"
	"github.com/codepnw/hexagonal/repository"
)

type accountService struct {
	accRepo repository.AccountRepository
}

func NewAccountService(accRepo repository.AccountRepository) AccountService {
	return accountService{accRepo: accRepo}
}

func (s accountService) NewAccount(customerID int, request NewAccountRequest) (*AccountResponse, error) {
	// Validate
	if request.Amount < 1000 {
		return nil, errs.NewErrValidation("amount at least 1,000")
	}

	if strings.ToLower(request.AccountType) != "saving" && strings.ToLower(request.AccountType) != "checking" {
		return nil, errs.NewErrValidation("account type should be saving or checking")
	}

	account := repository.Account{
		CustomerID:  customerID,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		AccountType: request.AccountType,
		Amount:      request.Amount,
		Status:      1,
	}

	newAcc, err := s.accRepo.Create(account)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewErrUnexpected()
	}

	response := AccountResponse{
		AccountID:   newAcc.AccountID,
		OpeningDate: newAcc.OpeningDate,
		AccountType: newAcc.AccountType,
		Amount:      newAcc.Amount,
		Status:      newAcc.Status,
	}

	return &response, nil
}

func (s accountService) GetAccounts(customerID int) ([]AccountResponse, error) {
	accounts, err := s.accRepo.GetAll(customerID)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewErrUnexpected()
	}

	responses := []AccountResponse{}

	for _, acc := range accounts {
		responses = append(responses, AccountResponse{
			AccountID:   acc.AccountID,
			OpeningDate: acc.OpeningDate,
			AccountType: acc.AccountType,
			Amount:      acc.Amount,
			Status:      acc.Status,
		})
	}
	return responses, nil
}
