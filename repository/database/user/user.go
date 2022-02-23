package user

import (
	"be/entities"
	"be/utils"

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

func (repo *UserDb) GetById(uid string) (GetByIdResponse, error) {
	user := GetByIdResponse{}

	if err := repo.db.Model(&entities.User{}).Where("user_uid =?", uid).First(&user); err != nil {
		return user, err.Error
	}

	return user, nil
}

// func (repo *UserDb) Update(user_uid string, upUser entities.User) (entities.Room, error) {
// 	tx := repo.db.Begin()

// 	defer func() {
// 		if r := recover(); r != nil {
// 			tx.Rollback()
// 		}
// 	}()

// 	if err := tx.Error; err != nil {
// 		return entities.Room{}, err
// 	}

// 	resRoom1 := entities.Room{}

// 	if err := tx.Model(&entities.Room{}).Where("room_uid = ?", room_uid).Find(&resRoom1).Error; err != nil {
// 		tx.Rollback()
// 		return entities.Room{}, err
// 	}

// 	if resRoom1.User_uid != user_uid {
// 		tx.Rollback()
// 		return entities.Room{}, errors.New(gorm.ErrInvalidData.Error())
// 	}

// 	if res := tx.Model(&entities.Room{}).Where("room_uid = ?", room_uid).Delete(&resRoom1); res.RowsAffected == 0 {
// 		log.Info(res.RowsAffected)
// 		tx.Rollback()
// 		return entities.Room{}, errors.New(gorm.ErrRecordNotFound.Error())
// 	}
// 	resRoom1.DeletedAt = gorm.DeletedAt{}
// 	resRoom1.ID = 0
// 	if res := tx.Create(&resRoom1); res.Error != nil {
// 		tx.Rollback()
// 		return entities.Room{}, res.Error
// 	}

// 	if res := tx.Model(&entities.Room{}).Where("room_uid = ?", room_uid).Updates(entities.Room{Name: upRoom.Name, Category: upRoom.Category, Price: upRoom.Price, Detail: upRoom.Detail}); res.Error != nil {
// 		tx.Rollback()
// 		return entities.Room{}, res.Error
// 	}

// 	return resRoom1, tx.Commit().Error
// }

func (repo *UserDb) Delete(userUid string) error {

	var user entities.User

	if err := repo.db.Where("user_uid =?", userUid).First(&user); err != nil {
		return err.Error
	}
	repo.db.Delete(&user, userUid)
	return nil

}
