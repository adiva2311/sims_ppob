package services

import (
	"database/sql"
	"sims_ppob/dto"
	"sims_ppob/repositories"
)

type InfoService interface {
	FindAllBanners() ([]dto.BannerResponse, error)
	FindAllServices() ([]dto.ServiceResponse, error)
}

type InfoServiceImpl struct {
	infoRepository repositories.InfoRepository
}

// FindAllBanners implements InfoService.
func (i *InfoServiceImpl) FindAllBanners() ([]dto.BannerResponse, error) {
	listBanner, err := i.infoRepository.FindAllBanners()
	if err != nil {
		return nil, err
	}

	return dto.ToBannerResponse(listBanner), nil
}

// FindAllServices implements InfoService.
func (i *InfoServiceImpl) FindAllServices() ([]dto.ServiceResponse, error) {
	listService, err := i.infoRepository.FindAllServices()
	if err != nil {
		return nil, err
	}
	return dto.ToServiceResponse(listService), nil
}

func NewInfoService(db *sql.DB) InfoService {
	return &InfoServiceImpl{
		infoRepository: repositories.NewInfoRepository(db),
	}
}
