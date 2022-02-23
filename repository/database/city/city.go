package city

import (
	"be/entities"

	"gorm.io/gorm"
)

type CityDb struct {
	db *gorm.DB
}

func New(db *gorm.DB) *CityDb {
	return &CityDb{
		db: db,
	}
}

func (repo *CityDb) GetAll() ([]CityResp, error) {

	cityRespArr := []CityResp{}

	res := repo.db.Model(&entities.City{}).Unscoped().Find(&cityRespArr)

	if res.Error != nil {
		return []CityResp{}, res.Error
	}

	return cityRespArr, nil
}
