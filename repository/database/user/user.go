package user

import (
	"be/entities"
	"be/utils"
	"fmt"

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
func (repo *UserDb) GetById(uid string) (entities.User, error) {
	user := entities.User{}

	if err := repo.db.Where("user_uid =?", uid).First(&user); err != nil {
		return user, err.Error
	}

	return user, nil
}

func (repo *UserDb) Update(userUid string, newUser entities.User) (entities.User, error) {

	var user entities.User
	fmt.Println(user)

	if err := repo.db.Model(&user).Where("user_uid = ? ", userUid).Updates(entities.User{Name: newUser.Name, Email: newUser.Email}).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (repo *UserDb) Delete(userUid string) error {

	var user entities.User

	if err := repo.db.Where("user_uid =?", userUid).First(&user); err != nil {
		return err.Error
	}
	repo.db.Delete(&user, user.ID)
	return nil

}
