package repositories

import (
	"database/sql"
	"sims_ppob/information/models"
)

type BannerRepository interface {
	FindAllBanners() ([]models.Banner, error)
}

type BannerRepositoryImpl struct {
	DB *sql.DB
}

// FindAllBanners implements BannerRepository.
func (b *BannerRepositoryImpl) FindAllBanners() ([]models.Banner, error) {
	var data []models.Banner
	query := "SELECT banner_name, banner_image, description FROM banners"
	rows, err := b.DB.Query(query)
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

func NewBannerRepository(db *sql.DB) BannerRepository {
	return &BannerRepositoryImpl{
		DB: db,
	}
}
