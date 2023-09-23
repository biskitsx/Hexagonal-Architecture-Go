package service

import (
	"strings"
	"time"

	"github.com/biskitsx/Hexagonal-Architecture-Go/errs"
	"github.com/biskitsx/Hexagonal-Architecture-Go/logs"
	"github.com/biskitsx/Hexagonal-Architecture-Go/repository"
)

type accountService struct {
	accountRepo repository.AccountRepository
}

func NewAccountService(accountRepo repository.AccountRepository) AccountService {
	return &accountService{
		accountRepo: accountRepo,
	}
}

func (s accountService) NewAccount(customerID int, req NewAccountRequest) (*AccountResponse, error) {
	// validate
	if req.Amount < 5000 {
		return nil, errs.NewValidationError("amout at least 5,000")
	}

	if strings.ToLower(req.AccountType) != "saving" && strings.ToLower(req.AccountType) != "checking" {
		return nil, errs.NewValidationError("account type should be saving or checking!")
	}

	account := repository.Account{
		CustomerID:  customerID,
		OpeningDate: time.Now().Format("2006-1-2 15:04:05"),
		Status:      1,
		AccountType: req.AccountType,
		Amount:      req.Amount,
	}

	newAccount, err := s.accountRepo.Create(account)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectError()
	}
	response := AccountResponse{
		AccountID:   newAccount.AccountID,
		OpeningDate: newAccount.OpeningDate,
		AccountType: newAccount.AccountType,
		Amount:      newAccount.Amount,
		Status:      newAccount.Status,
	}
	return &response, nil
}

func (s accountService) GetAccounts(customerID int) ([]AccountResponse, error) {
	accounts, err := s.accountRepo.GetAll(customerID)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectError()
	}

	responses := []AccountResponse{}
	for _, account := range accounts {
		responses = append(responses, AccountResponse{
			AccountID:   account.AccountID,
			OpeningDate: account.OpeningDate,
			AccountType: account.AccountType,
			Amount:      account.Amount,
			Status:      account.Status,
		})
	}
	return responses, nil

}
