package dto

import "sims_ppob/models"

type BannerResponse struct {
	BannerName  string `json:"banner_name"`
	BannerImage string `json:"banner_image"`
	Description string `json:"description"`
}

func ToBannerResponse(banners []models.Banner) []BannerResponse {
	var response []BannerResponse
	for _, banner := range banners {
		response = append(response, BannerResponse{
			BannerName:  banner.BannerName,
			BannerImage: banner.BannerImage,
			Description: banner.Description,
		})
	}
	return response
}
