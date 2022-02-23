package room

import (
	"be/configs"
	"be/entities"
	"be/repository/database/image"
	"be/repository/database/user"
	"be/utils"
	"fmt"
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

		mock2 := entities.Room{User_uid: res1.User_uid, City_id: 100, Name: "room1 name", Price: 100, Description: "room1 detail"}
		res, err := repo.Create(mock2)
		assert.Nil(t, err)
		assert.NotNil(t, res)
		log.Info(res)
	})
}

func TestUpdate(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)
	db.Migrator().DropTable(&entities.User{})
	db.Migrator().DropTable(&entities.Room{})
	db.Migrator().DropTable(&entities.Image{})
	db.Migrator().DropTable(&entities.Booking{})
	db.AutoMigrate(&entities.Room{})

	t.Run("success run Update", func(t *testing.T) {
		mock1 := entities.User{Name: "user1 name", Email: "user1 email", Password: "user1 password"}
		res1, err1 := user.New(db).Create(mock1)
		if err1 != nil {
			t.Fatal()
		}
		mock2 := entities.Room{User_uid: res1.User_uid, City_id: 1, Name: "room1 name", Price: 100, Description: "room1 detail"}
		res2, err2 := repo.Create(mock2)
		if err2 != nil {
			t.Fatal()
		}
		mock3 := entities.Room{Name: "room3 name", Price: 300, Description: "room3 detail"}
		res, err := repo.Update(res1.User_uid, res2.Room_uid, mock3)
		assert.Nil(t, err)
		assert.NotNil(t, res)
		// log.Info(res)
	})
}

func TestGetByID(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)
	db.Migrator().DropTable(&entities.User{})
	db.Migrator().DropTable(&entities.Room{})
	db.Migrator().DropTable(&entities.Image{})
	db.Migrator().DropTable(&entities.Booking{})
	db.AutoMigrate(&entities.Image{})

	t.Run("success run GetById", func(t *testing.T) {
		mock1 := entities.User{Name: "user1 name", Email: "user1 email", Password: "user1 password"}
		res1, err1 := user.New(db).Create(mock1)
		if err1 != nil {
			t.Fatal()
		}
		mock2 := entities.Room{User_uid: res1.User_uid, City_id: 1, Name: "room1 name", Price: 100, Description: "room1 detail", Category: "superior"}
		res2, err2 := repo.Create(mock2)
		if err2 != nil {
			t.Fatal()
		}

		mock3 := image.ImageReq{}

		for i := 0; i < 3; i++ {
			mock3.Array = append(mock3.Array, image.ImageInput{Url: (fmt.Sprintf("url%d", i+1))})
		}

		if err := image.New(db).Create(res2.Room_uid, mock3); err != nil {
			t.Fatal()
		}

		res, err := repo.GetById(res2.Room_uid)
		assert.Nil(t, err)
		assert.NotNil(t, res)
		log.Info(res)
	})
}
