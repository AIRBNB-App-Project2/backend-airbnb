package room

import (
	"be/entities"

	"github.com/lithammer/shortuuid"
	"gorm.io/gorm"
)

type RoomDb struct {
	db *gorm.DB
}

func New(db *gorm.DB) *RoomDb {
	return &RoomDb{
		db: db,
	}
}

func (repo *RoomDb) Create(room entities.Room) (entities.Room, error) {

	var uid string

	for {
		uid = shortuuid.New()
		find := entities.Room{}
		res := repo.db.Model(&entities.Room{}).Where("room_uid = ?", uid).First(&find)
		if res.RowsAffected == 0 {
			break
		}
		if res.Error != nil {
			return entities.Room{}, res.Error
		}
	}

	room.Room_uid = uid

	if err := repo.db.Create(&room).Error; err != nil {
		return entities.Room{}, err
	}

	return room, nil
}

func (repo *RoomDb) UpdateTranx(user_uid string, room_uid string,room entities.Room) (entities.Room, error) {
	tx := repo.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return entities.Room{}, err
	}

	resRoom1 := entities.Room{}

	if err := tx.Model(&entities.Room{}).Where("user_id = ? AND room_uid = ?", user_uid, room_uid).Find(&resRoom1).Error ; err != nil {
		tx.Rollback()
		return entities.Room{}, err
	}

	// if res := tx.m

}