package repositories

import (
	"database/sql"
	"sims_ppob/information/models"
)

type ServiceRepository interface {
	FindAllServices() ([]models.Services, error)
}

type ServiceRepositoryImpl struct {
	DB *sql.DB
}

// FindAllServices implements ServiceRepository.
func (s *ServiceRepositoryImpl) FindAllServices() ([]models.Services, error) {
	var data []models.Services
	query := "SELECT service_code, service_name, service_icon, service_tariff FROM services"
	rows, err := s.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var service models.Services
		if err := rows.Scan(&service.ServiceCode, &service.ServiceName, &service.ServiceIcon, &service.ServiceTariff); err != nil {
			return nil, err
		}
		data = append(data, service)
	}
	return data, nil
}

func NewServiceRepository(db *sql.DB) ServiceRepository {
	return &ServiceRepositoryImpl{
		DB: db,
	}
}
