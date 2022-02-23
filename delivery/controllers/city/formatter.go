package city

import "be/repository/database/city"

type GetResponseFormat struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Data    []city.CityResp `json:"data"`
}
