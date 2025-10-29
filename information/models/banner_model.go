package models

type Banner struct {
	ID          uint   `json:"id"`
	BannerName  string `json:"banner_name" validate:"required"`
	BannerImage string `json:"banner_image" validate:"required"`
	Description string `json:"description"`
}

func (Banner) TableName() string {
	return "banners"
}
