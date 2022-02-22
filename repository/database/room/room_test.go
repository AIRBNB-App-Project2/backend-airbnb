package room

import (
	"be/configs"
	"be/entities"
	"be/repository/database/user"
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
	db.AutoMigrate(&entities.Room{})

	t.Run("success run create", func(t *testing.T) {
		mock1 := entities.User{Name: "user1 name", Email: "user1 email", Password: "user1 password"}

		res1, err1 := user.New(db).Create(mock1)
		if err1 != nil {
			t.Fatal()

		}

		mock2 := entities.Room{User_uid: res1.User_uid, City_id: 1, Name: "room1 name", Price: 100, Detail: "room1 detail"}
		res, err := repo.Create(mock2)
		assert.Nil(t, err)
		assert.NotNil(t, res)
		log.Info(res)
	})
}

