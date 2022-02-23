package image

import (
	"be/configs"
	"be/entities"
	"be/repository/database/room"
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
	db.AutoMigrate(&entities.Image{})

	mock1 := entities.User{Name: "user1 name", Email: "user1 email", Password: "user1 password"}
	res1, err1 := user.New(db).Create(mock1)
	if err1 != nil {
		t.Fatal()
	}
	mock2 := entities.Room{User_uid: res1.User_uid, City_id: 1, Name: "room1 name", Price: 100, Description: "room1 detail"}
	res2, err2 := room.New(db).Create(mock2)
	if err2 != nil {
		t.Fatal()
	}

	t.Run("success run create", func(t *testing.T) {

		mock3 := ImageReq{}

		for i := 0; i < 3; i++ {
			mock3.Array = append(mock3.Array, ImageInput{Url: (fmt.Sprintf("url%d", i+1))})
		}
		// log.Info(mock3)

		err := repo.Create(res2.Room_uid, mock3)
		assert.Nil(t, err)

	})

}
