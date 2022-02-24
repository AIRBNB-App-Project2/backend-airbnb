package room

import (
	"be/entities"
	"errors"

	"github.com/labstack/gommon/log"
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

func (repo *RoomDb) Create(room entities.Room) (RoomCreateResp, error) {

	var uid string

	for {
		uid = shortuuid.New()
		find := entities.Room{}
		res := repo.db.Model(&entities.Room{}).Where("room_uid =  ?", uid).First(&find)
		if res.RowsAffected == 0 {
			break
		}
		if res.Error != nil {
			return RoomCreateResp{}, res.Error
		}
	}

	room.Room_uid = uid

	if err := repo.db.Create(&room).Error; err != nil {
		return RoomCreateResp{}, err
	}

	resImg := repo.db.Model(&entities.Image{}).Create(&entities.Image{Room_uid: room.Room_uid})
	if resImg.Error != nil {
		return RoomCreateResp{}, resImg.Error
	}

	resp := RoomCreateResp{}

	res := repo.db.Model(&entities.Room{}).Where("room_uid = ?", uid).Select("rooms.room_uid as Room_uid, users.name as Name_user, rooms.name as Name_room, category as Category, address as Address, cities.name as City, description as Description, price as Price").Joins("inner join users on users.user_uid = rooms.user_uid").Joins("inner join cities on rooms.city_id = cities.id").Find(&resp)
	if res.Error != nil {
		return RoomCreateResp{}, res.Error
	}

	return resp, nil
}

func (repo *RoomDb) Update(user_uid string, room_uid string, upRoom entities.Room) (entities.Room, error) {
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

	if err := tx.Model(&entities.Room{}).Where("room_uid = ?", room_uid).Find(&resRoom1).Error; err != nil {
		tx.Rollback()
		return entities.Room{}, err
	}

	if resRoom1.User_uid != user_uid {
		tx.Rollback()
		return entities.Room{}, errors.New(gorm.ErrInvalidData.Error())
	}

	if res := tx.Model(&entities.Room{}).Where("room_uid = ?", room_uid).Delete(&resRoom1); res.RowsAffected == 0 {
		log.Info(res.RowsAffected)
		tx.Rollback()
		return entities.Room{}, errors.New(gorm.ErrRecordNotFound.Error())
	}
	resRoom1.DeletedAt = gorm.DeletedAt{}
	resRoom1.ID = 0
	if res := tx.Create(&resRoom1); res.Error != nil {
		tx.Rollback()
		return entities.Room{}, res.Error
	}

	if res := tx.Model(&entities.Room{}).Where("room_uid = ?", room_uid).Updates(entities.Room{Name: upRoom.Name, Category: upRoom.Category, Price: upRoom.Price, Description: upRoom.Description}); res.Error != nil {
		tx.Rollback()
		return entities.Room{}, res.Error
	}

	return resRoom1, tx.Commit().Error
}

func (repo *RoomDb) GetAllRoom(length int, city, category, name, status string) ([]RoomGetAllResp, error) {

	respRoomAll := []RoomGetAllResp{}

	var condition string

	if city != "" {
		city = "cities.name LIKE '%" + city + "%'"
	}
	if category != "" {
		category = "AND category = " + category
	}
	if status != "" {
		status = "AND status = " + status
	}
	if name != "" {
		name = "AND rooms.name LIKE '%" + name + "%'"
	}

	condition = city + category + status + name

	choose := "rooms.room_uid as Room_uid, rooms.name as Name, price as Price, description as Description, status as Status, (SELECT images.url FROM images where images.room_uid = rooms.room_uid and images.deleted_at is null order by images.id limit 1) as Image"

	res := repo.db.Model(&entities.Room{}).Where(condition).Select(choose).Joins("inner join cities on rooms.city_id = cities.id").Limit(length).Order("rooms.name ASC").Find(&respRoomAll)

	if res.Error != nil {
		return []RoomGetAllResp{}, res.Error
	}

	return respRoomAll, nil
}

func (repo *RoomDb) GetById(room_uid string) (RoomGetByIdResp, error) {
	resp := RoomGetByIdResp{}

	res1 := repo.db.Model(&entities.Room{}).Where("room_uid = ?", room_uid).Select("rooms.room_uid as Room_uid, users.name as owner_room, rooms.name as Name, category as Category, address as Address, cities.name as City, description as Description, price as Price, status as Status").Joins("inner join users on users.user_uid = rooms.user_uid").Joins("inner join cities on rooms.city_id = cities.id").Find(&resp)

	if res1.Error != nil {
		return RoomGetByIdResp{}, res1.Error
	}

	images := []Images{}

	res2 := repo.db.Model(&entities.Image{}).Where("room_uid", room_uid).Find(&images)

	if res2.Error != nil {
		return RoomGetByIdResp{}, res2.Error
	}

	resp.Image = images

	return resp, nil
}


func (repo *RoomDb) Delete(room_uid string) (entities.Room, error) {
	roomInit := entities.Room{}

	resp := repo.db.Model(&entities.Room{}).Where("room_uid = ?", room_uid).Delete(&roomInit)

	if resp.RowsAffected == 0 {
		return entities.Room{}, errors.New(gorm.ErrRecordNotFound.Error())
	}

	return roomInit, nil
}