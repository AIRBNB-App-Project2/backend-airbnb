package user

import (
	"be/entities"
	"be/utils"
	"errors"

	"github.com/labstack/gommon/log"
	"github.com/lithammer/shortuuid"
	"gorm.io/gorm"
)

type UserDb struct {
	db *gorm.DB
}

func New(db *gorm.DB) *UserDb {
	return &UserDb{
		db: db,
	}
}

func (repo *UserDb) Create(user entities.User) (entities.User, error) {

	userInit := entities.User{}

	checkEmail := repo.db.Where("email = ?", user.Email).Find(&userInit)

	if checkEmail.RowsAffected != 0 {
		return entities.User{}, errors.New("email already exist")
	}

	var uid string

	for {
		uid = shortuuid.New()
		userFind := entities.User{}
		res := repo.db.Model(&entities.User{}).Where("user_uid = ?", uid).First(&userFind)
		if res.RowsAffected == 0 {
			break
		}
	}
	user.Password, _ = utils.HashPassword(user.Password)

	user.User_uid = uid

	if err := repo.db.Create(&user).Error; err != nil {
		return entities.User{}, err
	}

	return user, nil
}

func (repo *UserDb) GetById(user_uid string) (GetByIdResponse, error) {
	user := GetByIdResponse{}

	if err := repo.db.Model(&entities.User{}).Where("user_uid =?", user_uid).First(&user); err.Error != nil || err.RowsAffected == 0 {
		return user, err.Error
	}

	rooms := []RoomUserResp{}

	resR := repo.db.Model(entities.User{}).Where("rooms.deleted_at IS NULL").Select("rooms.room_uid as Room_uid, rooms.name as Name , description as Description , price as Price , status as Status").Where("rooms.user_uid = ?", user_uid).Joins("inner join rooms on rooms.user_uid = users.user_uid").Find(&rooms)
	if resR.Error != nil {
		return GetByIdResponse{}, errors.New(gorm.ErrRecordNotFound.Error())
	}

	user.Rooms = rooms

	bookings := []BookingUserResp{}

	resB := repo.db.Model(entities.User{}).Where("bookings.user_uid = ? AND bookings.status != 'waiting' AND bookings.deleted_at IS Null", user_uid).Select("bookings.booking_uid as Booking_uid, rooms.name as Name, rooms.description as Description,rooms.price as Price ,bookings.start_date as Start_date, bookings.end_date as End_date, DATEDIFF(bookings.end_date, bookings.start_date) as Days, DATEDIFF(bookings.end_date, bookings.start_date) * rooms.price as Price_total, bookings.status as Status").Joins("inner join bookings on bookings.user_uid = users.user_uid").Joins("inner join rooms on bookings.room_uid = rooms.room_uid").Find(&bookings)

	if resB.Error != nil {
		return GetByIdResponse{}, resB.Error
	}

	user.Bookings = bookings

	return user, nil
}

func (repo *UserDb) Update(user_uid string, upUser entities.User) (entities.User, error) {
	tx := repo.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return entities.User{}, err
	}

	resUser1 := entities.User{}

	if res := tx.Model(&entities.User{}).Where("User_uid = ?", user_uid).Find(&resUser1); res.Error != nil || res.RowsAffected == 0 {
		tx.Rollback()
		return entities.User{}, errors.New(gorm.ErrRecordNotFound.Error())
	}
	resUser1.ID = 0
	if resUser1.User_uid != user_uid {
		tx.Rollback()
		return entities.User{}, errors.New(gorm.ErrInvalidData.Error())
	}

	if res := tx.Model(&entities.User{}).Where("User_uid = ?", user_uid).Delete(&resUser1); res.RowsAffected == 0 {
		log.Info(res.RowsAffected)
		tx.Rollback()
		return entities.User{}, errors.New(gorm.ErrRecordNotFound.Error())
	}
	resUser1.DeletedAt = gorm.DeletedAt{}

	if res := tx.Create(&resUser1); res.Error != nil {
		tx.Rollback()
		return entities.User{}, res.Error
	}

	if res := tx.Model(&entities.User{}).Where("User_uid = ?", user_uid).Updates(entities.User{Name: upUser.Name, Email: upUser.Email, Password: upUser.Password}); res.Error != nil {
		tx.Rollback()
		return entities.User{}, res.Error
	}

	return resUser1, tx.Commit().Error
}

func (repo *UserDb) Delete(userUid string) (entities.User, error) {

	var user entities.User

	if res := repo.db.Model(&entities.User{}).Where("user_uid =?", userUid).Delete(&user); res.Error != nil || res.RowsAffected == 0 {
		return user, errors.New(gorm.ErrRecordNotFound.Error())
	}
	return user, nil

}
