package user

import (
	"be/configs"
	"be/entities"
	"be/utils"
	"testing"

	"github.com/labstack/gommon/log"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)
	db.Migrator().DropTable(&entities.User{})
	db.Migrator().DropTable(&entities.Room{})
	db.Migrator().DropTable(&entities.Image{})
	db.Migrator().DropTable(&entities.Booking{})
	db.AutoMigrate(&entities.User{})

	t.Run("success run create", func(t *testing.T) {
		mockUser := entities.User{Name: "anonim1", Email: "anonim1", Password: "anonim1"}
		res, err := repo.Create(mockUser)
		assert.Nil(t, err)
		assert.NotNil(t, res)
		log.Info(res)
	})

	t.Run("fail run create", func(t *testing.T) {
		mockUser1 := entities.User{Name: "anonim2", Email: "anonim2", Password: "anonim2"}
		if _, err := repo.Create(mockUser1); err != nil {
			t.Fatal()
		}
		mockUser := entities.User{Name: "anonim2", Email: "anonim2", Password: "anonim2"}
		res, err := repo.Create(mockUser)
		assert.NotNil(t, err)
		assert.Equal(t, entities.User{}, res)
	})
}

func TestGetById(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)
	db.Migrator().DropTable(&entities.User{})
	db.Migrator().DropTable(&entities.Room{})
	db.Migrator().DropTable(&entities.Image{})
	db.Migrator().DropTable(&entities.Booking{})
	db.AutoMigrate(&entities.User{})

	t.Run("success run GetById", func(t *testing.T) {
		mockUser1 := entities.User{Name: "anonim1", Email: "anonim1", Password: "anonim1"}
		res1, err1 := repo.Create(mockUser1)
		if err1 != nil {
			t.Fatal()
		}
		res, err := repo.GetById(res1.User_uid)
		assert.Nil(t, err)
		assert.NotNil(t, res)
		log.Info(res)
	})
}

