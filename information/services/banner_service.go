package services

import (
	"database/sql"
	"sims_ppob/information/dto"
	"sims_ppob/information/repositories"
)

type BannerService interface {
	FindAllBanners() ([]dto.BannerResponse, error)
}

type BannerServiceImpl struct {
	BannerRepository repositories.BannerRepository
}

// FindAllBanners implements BannerService.
func (b *BannerServiceImpl) FindAllBanners() ([]dto.BannerResponse, error) {
	listBanner, err := b.BannerRepository.FindAllBanners()
	if err != nil {
		return nil, err
	}

	return dto.ToBannerResponse(listBanner), nil
}

func NewBannerService(db *sql.DB) BannerService {
	return &BannerServiceImpl{
		BannerRepository: repositories.NewBannerRepository(db),
	}
}
