package user

import (
	"be/entities"
	"be/utils"
	"errors"

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
func (repo *UserDb) GetById(userId int) (entities.User, error) {
	arrUser := entities.User{}

	if err := repo.db.Where("ID = ?", userId).First(&arrUser); err != nil {
		return arrUser, errors.New("User Not Found")
	}

	return arrUser, nil
}

func (repo *UserDb) Update(userId int, newUser entities.User) (entities.User, error) {

	var user entities.User
	if err := repo.db.First(&user, userId); err != nil {
		return entities.User{}, errors.New(" User Not Found")
	}

	if err := repo.db.Model(&user).Where("ID = ? ", userId).Updates(&newUser).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (repo *UserDb) Delete(userId int) error {

	var user entities.User

	if err := repo.db.First(&user, userId).Error; err != nil {
		return err
	}

	repo.db.Delete(&user, userId)
	return nil

}
