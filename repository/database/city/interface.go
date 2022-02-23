package city

type City interface {
	GetAll() ([]CityResp, error)
}
