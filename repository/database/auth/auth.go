package auth

import (
	"be/entities"

	"gorm.io/gorm"
)

type AuthDb struct {
	db *gorm.DB
}

func New(db *gorm.DB) *AuthDb {
	return &AuthDb{
		db: db,
	}
}

func (ad *AuthDb) Login(UserLogin entities.User) (entities.User, error) {
	user := entities.User{}
	if err := ad.db.Model(&entities.User{}).Where("email = ? AND password = ?", UserLogin.Email, UserLogin.Password).First(&user).Error; err != nil {
		return entities.User{}, err
	}

	return user, nil
}
