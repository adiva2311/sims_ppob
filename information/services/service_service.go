package services

import (
	"database/sql"
	"sims_ppob/information/dto"
	"sims_ppob/information/repositories"
)

type ServiceService interface {
	FindAllServices() ([]dto.ServiceResponse, error)
}

type ServiceServiceImpl struct {
	ServiceRepository repositories.ServiceRepository
}

// FindAllServices implements ServiceService.
func (s *ServiceServiceImpl) FindAllServices() ([]dto.ServiceResponse, error) {
	listService, err := s.ServiceRepository.FindAllServices()
	if err != nil {
		return nil, err
	}
	return dto.ToserviceResponse(listService), nil
}

func NewServiceService(db *sql.DB) ServiceService {
	return &ServiceServiceImpl{
		ServiceRepository: repositories.NewServiceRepository(db),
	}
}
