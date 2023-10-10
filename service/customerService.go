package service

import (
	"github.com/AshakaE/banking/domain"
	"github.com/AshakaE/banking/dto"
	"github.com/AshakaE/banking/errors"
)

type CustomerService interface {
	GetCustomer(string) (*dto.CustomerResponse, *errors.AppError)
	GetAllCustomers(string) ([]dto.CustomerResponse, *errors.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func convertToDtoSlice(customers []domain.Customer) []dto.CustomerResponse {
	dtoSlice := make([]dto.CustomerResponse, len(customers))
	for i, c := range customers {
		dtoSlice[i] = c.ToDto()
	}
	return dtoSlice
}


func (s DefaultCustomerService) GetAllCustomers(status string) ([]dto.CustomerResponse, *errors.AppError) {
	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else {
		status = ""
	}
	c, err := s.repo.FindAll(status)
	if err != nil {
		return nil, err
	}
	dtoToSlice := convertToDtoSlice(c)
	
	return dtoToSlice, nil
}

func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errors.AppError) {
	c, err := s.repo.FindById(id)
	if err != nil {
		return nil, err
	}

	resp := c.ToDto()

	return &resp, nil
}

func NewCustomerService(repo domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo}
}
