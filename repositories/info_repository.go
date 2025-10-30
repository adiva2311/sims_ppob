package repositories

import (
	"database/sql"
	"sims_ppob/models"
)

type InfoRepository interface {
	FindAllBanners() ([]models.Banner, error)
	FindAllServices() ([]models.Services, error)
}

type InfoRepositoryImpl struct {
	DB *sql.DB
}

// FindAllBanners implements InfoRepository.
func (i *InfoRepositoryImpl) FindAllBanners() ([]models.Banner, error) {
	var data []models.Banner
	query := "SELECT banner_name, banner_image, description FROM banners"
	rows, err := i.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var banner models.Banner
		if err := rows.Scan(&banner.BannerName, &banner.BannerImage, &banner.Description); err != nil {
			return nil, err
		}
		data = append(data, banner)
	}
	return data, nil
}

// FindAllServices implements InfoRepository.
func (i *InfoRepositoryImpl) FindAllServices() ([]models.Services, error) {
	var data []models.Services
	query := "SELECT service_code, service_name, service_icon, service_tariff FROM services"
	rows, err := i.DB.Query(query)
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

func NewInfoRepository(db *sql.DB) InfoRepository {
	return &InfoRepositoryImpl{
		DB: db,
	}
}
