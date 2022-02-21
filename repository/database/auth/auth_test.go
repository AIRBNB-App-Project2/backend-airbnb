package auth

import (
	"be/configs"
	"be/entities"
	"be/repository/database/user"
	"be/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)
	db.Migrator().DropTable(&entities.User{})
	db.Migrator().DropTable(&entities.Room{})
	db.Migrator().DropTable(&entities.Image{})
	db.Migrator().DropTable(&entities.Booking{})
	db.Migrator().DropTable(&entities.Order{})
	db.AutoMigrate(&entities.User{})

	t.Run("success run login", func(t *testing.T) {
		mockUser := entities.User{Name: "anonim123", Email: "anonim@123", Password: "anonim123"}
		_, err := user.New(db).Create(mockUser)
		if err != nil {
			t.Fail()
		}
		mockLogin := entities.User{Email: "anonim@123", Password: "anonim123"}
		res, err := repo.Login(mockLogin)
		assert.Nil(t, err)
		assert.Equal(t, "anonim@123", res.Email)
		assert.Equal(t, "anonim123", res.Password)
	})

	t.Run("fail run login", func(t *testing.T) {
		mockLogin := entities.User{Email: "anonim@456", Password: "anonim456"}
		_, err := repo.Login(mockLogin)
		assert.NotNil(t, err)
	})

}
