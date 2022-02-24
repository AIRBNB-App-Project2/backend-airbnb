package booking

import (
	"be/configs"
	"be/entities"
	"be/repository/database/image"
	"be/repository/database/room"
	"be/repository/database/user"
	"be/utils"
	"fmt"
	"testing"
	"time"

	"github.com/labstack/gommon/log"
	"github.com/stretchr/testify/assert"
	"gorm.io/datatypes"
)

func TestCreate(t *testing.T) {

	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)
	db.Migrator().DropTable(&entities.User{})
	db.Migrator().DropTable(&entities.Room{})
	db.Migrator().DropTable(&entities.Image{})
	db.Migrator().DropTable(&entities.Booking{})
	db.AutoMigrate(&entities.Booking{})
	db.AutoMigrate(&entities.Image{})

	t.Run("success run Create", func(t *testing.T) {
		mock1 := entities.User{Name: "user1 name", Email: "user1 email", Password: "user1 password"}
		res1, err1 := user.New(db).Create(mock1)
		if err1 != nil {
			t.Fatal()
		}
		mock2 := entities.Room{User_uid: res1.User_uid, City_id: 1, Name: "room1 name", Price: 100, Description: "room1 detail", Category: "superior"}
		res2, err2 := room.New(db).Create(mock2)
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

		mock4 := entities.Booking{User_uid: res1.User_uid, Room_uid: res2.Room_uid, Start_date: datatypes.Date(time.Now()), End_date: datatypes.Date(time.Now().AddDate(0, 0, 1))}

		res, err := repo.Create(res1.User_uid, res2.Room_uid, mock4)
		assert.Nil(t, err)
		assert.NotNil(t, res)
		log.Info(res)
	})

	t.Run("success handle overlapping", func(t *testing.T) {
		mock1 := entities.User{Name: "user2 name", Email: "user2 email", Password: "user2 password"}
		res1, err1 := user.New(db).Create(mock1)
		if err1 != nil {
			t.Fatal()
		}
		mock2 := entities.Room{User_uid: res1.User_uid, City_id: 100, Name: "room2 name", Price: 100, Description: "room2 detail", Category: "superior"}
		res2, err2 := room.New(db).Create(mock2)
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

		mock4 := entities.Booking{User_uid: res1.User_uid, Room_uid: res2.Room_uid, Start_date: datatypes.Date(time.Now().AddDate(0, 0, 2)), End_date: datatypes.Date(time.Now().AddDate(0, 0, 5)), Status: "reservation"}
		if _, err := repo.Create(res1.User_uid, res2.Room_uid, mock4); err != nil {
			t.Fatal()
		}
		mock5 := entities.Booking{User_uid: res1.User_uid, Room_uid: res2.Room_uid, Start_date: datatypes.Date(time.Now().AddDate(0, 0, 4)), End_date: datatypes.Date(time.Now().AddDate(0, 0, 6))}
		res, err := repo.Create(res1.User_uid, res2.Room_uid, mock5)
		assert.NotNil(t, err)
		assert.NotNil(t, res)
		// log.Info(res, err)

	})

}
