package client

import "github.com/aalkan/plesk-go/pkg/models"

func (c *client) GetLicense(params *models.LicenseGet) (*models.LicenseResponse, error) {
	response := &models.LicenseResponse{}
	if err := c.request("POST", "/", params, response); err != nil {
		return nil, err
	}

	return response, nil
}
