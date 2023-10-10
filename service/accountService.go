package service

import (
	"time"

	"github.com/AshakaE/banking/domain"
	"github.com/AshakaE/banking/dto"
	"github.com/AshakaE/banking/errors"
)

const dbTSLayout = "2006-01-02 15:04:05"

type AccountService interface {
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, *errors.AppError)
	MakeTransaction(dto.TransactionRequest) (*dto.TransactionResponse, *errors.AppError)
}

type DefaultAccountService struct {
	repo domain.AccountRepository
}

func (s DefaultAccountService) NewAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse, *errors.AppError) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	a := domain.Account{
		CustomerId:  req.CustomerId,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      "1",
	}

	newAccount, err := s.repo.Save(a)
	if err != nil {
		return nil, err
	}
	resp := newAccount.ToNewAccountResponseDto()
	return resp, nil
}

func (s DefaultAccountService) MakeTransaction(req dto.TransactionRequest) (*dto.TransactionResponse, *errors.AppError) {
	// incoming request validation
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	// server side validation for checking the available balance in the account
	if req.IsTransactionTypeWithdrawal() {
		account, err := s.repo.FindBy(req.AccountId)
		if err != nil {
			return nil, err
		}
		if account.CannnotWithdraw(req.Amount) {
			return nil, errors.NewValidationError("Insufficient balance in the account")
		}
	}

	// if all is well, build the domain object & save the transaction
	t := domain.Transaction{
		AccountId:       req.AccountId,
		Amount:          req.Amount,
		TransactionType: req.TransactionType,
		TransactionDate: time.Now().Format(dbTSLayout),
	}
	transaction, appError := s.repo.SaveTransaction(t)
	if appError != nil {
		return nil, appError
	}
	response := transaction.ToDto()
	return &response, nil
}

func NewAccountService(repo domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repo}
}

// type Shape interface {
//     Perimeter() float64
// }

// type Rectangle struct {
//     width  float64
//     height float64
// }

// func (r Rectangle) Perimeter() float64 {
//     return 2 * (r.width + r.height)
// }

// var r Shape = Rectangle{width: 3.0, height: 4.0}
// var perimeter = r.Perimeter()  // perimeter = 14.0