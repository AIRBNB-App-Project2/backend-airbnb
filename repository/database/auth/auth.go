package auth

import (
	"be/entities"
	"be/utils"
	"errors"

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
	ad.db.Model(entities.User{}).Where("email = ?", UserLogin.Email).First(&user)

	if match := utils.CheckPasswordHash(UserLogin.Password, user.Password); !match {
		return entities.User{}, errors.New("incorrect password")
	}
	return user, nil
}
