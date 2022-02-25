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

	if err := repo.db.Model(&entities.User{}).Where("user_uid =?", user_uid).First(&user); err != nil {
		return user, err.Error
	}

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

	if err := tx.Model(&entities.User{}).Where("User_uid = ?", user_uid).Find(&resUser1).Error; err != nil {
		tx.Rollback()
		return entities.User{}, err
	}

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
	resUser1.ID = 0
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

	if res := repo.db.Model(&entities.User{}).Where("user_uid =?", userUid).Delete(&user); res.Error != nil {
		return user, res.Error
	}
	return user, nil

}
