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

func (repo *ImageDb) Create(room_uid string, image ImageReq) error {

	imgInit := entities.Image{}

	delImg := repo.db.Model(&entities.Image{}).Where("room_uid = ? AND url = 'https://test-upload-s3-rogerdev.s3.ap-southeast-1.amazonaws.com/6216503718eb9324b8213a1f.png'" , room_uid).Delete(&imgInit)

	if delImg.Error != nil {
		return delImg.Error
	}

	for i := 0; i < len(image.Array); i++ {

		if err := repo.db.Create(&entities.Image{Room_uid: room_uid, Url: image.Array[i].Url}).Error; err != nil {
			return err
		}

	}

	return nil
}
