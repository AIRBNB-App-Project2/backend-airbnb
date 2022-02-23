package image

import (
	"be/entities"

	"gorm.io/gorm"
)

type ImageDb struct {
	db *gorm.DB
}

func New(db *gorm.DB) *ImageDb {
	return &ImageDb{
		db: db,
	}
}

func (repo *ImageDb) Create(image entities.Image) (entities.Image, error) {

	// if err := repo.db.Create(&entities.Image{Room_uid: room_uid, Url: image.Array[i].Url}).Error; err != nil {
	// 	return err
	// }

	return entities.Image{}, nil
}
