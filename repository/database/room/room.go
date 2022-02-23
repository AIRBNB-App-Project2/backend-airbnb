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

func (repo *RoomDb) GetAll(s, city, category, name, length, status string) ([]entities.Room, error) {
	var result []entities.Room
	var query string = "SELECT * FROM rooms "
	var orderBy string = ""
	var limit string = ""

	if s != "" {
		if city != "" {
			city = "city_id = '" + city + "' AND "
		}
		myQueries := city + " name LIKE ?"
		s = "%" + s + "%"
		// fmt.Println("ssssssssssssss", myQueries, s)
		if res := repo.db.Preload("Images").Preload("Bookings").Where(myQueries, s).Find(&result); res.Error != nil {
			return []entities.Room{}, res.Error
		}

		return result, nil

	}

	middle := ""
	if city != "" {
		query = "SELECT * FROM rooms WHERE city_id=" + city
	}
	if category != "" {
		category += "category =" + category
	}
	if limit != "" {
		limit += " LIMIT " + length
	}
	if category != "" && query != "" {
		middle += "AND"
	}

	myQueries := query + middle + category + orderBy

	if res := repo.db.Raw(myQueries).Find(&result); res.Error != nil {
		return []entities.Room{}, res.Error
	}

	return result, nil
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
