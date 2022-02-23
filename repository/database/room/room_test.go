package room

import (
	"be/configs"
	"be/entities"
	"be/repository/database/user"
	"be/utils"
	"testing"

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
		// log.Info(res)
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

	t.Run("success run create", func(t *testing.T) {
		mock1 := entities.User{Name: "user1 name", Email: "user1 email", Password: "user1 password"}
		res1, err1 := user.New(db).Create(mock1)
		if err1 != nil {
			t.Fatal()
		}
		mock2 := entities.Room{User_uid: res1.User_uid, City_id: 1, Name: "room1 name", Price: 100, Detail: "room1 detail"}
		res2, err2 := repo.Create(mock2)
		if err2 != nil {
			t.Fatal()
		}
		mock3 := entities.Room{Name: "room3 name", Price: 300, Detail: "room3 detail"}
		res, err := repo.Update(res1.User_uid, res2.Room_uid, mock3)
		assert.Nil(t, err)
		assert.NotNil(t, res)
		// log.Info(res)
	})
}
func TestGetAll(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)
	db.Migrator().DropTable(&entities.User{})
	db.Migrator().DropTable(&entities.Room{})
	db.Migrator().DropTable(&entities.Image{})
	db.Migrator().DropTable(&entities.Booking{})
	db.AutoMigrate(&entities.Room{})

	t.Run("success run get all", func(t *testing.T) {

		//mock User
		mockUser1 := entities.User{Name: "user1 name", Email: "user1 email", Password: "user1 password"}
		resu1, err1 := user.New(db).Create(mockUser1)
		if err1 != nil {
			t.Fatal()
		}
		mockUser2 := entities.User{Name: "user2 name", Email: "user2 email", Password: "user1 password"}
		resu2, err2 := user.New(db).Create(mockUser2)
		if err2 != nil {
			t.Fatal()
		}
		mockUser3 := entities.User{Name: "user3 name", Email: "user3 email", Password: "user1 password"}
		_, err3 := user.New(db).Create(mockUser3)
		if err3 != nil {
			t.Fatal()
		}
		//==================

		city := "1"
		var category string = ""
		var name string = ""
		var length string = ""

		var status string = ""
		mockroom1 := entities.Room{User_uid: resu2.User_uid, City_id: 1, Name: "room1 name", Price: 100, Detail: "room1 detail"}
		resroom1, errroom1 := repo.Create(mockroom1)
		if errroom1 != nil {
			t.Fatal()
		}
		mockroom2 := entities.Room{User_uid: "jkajhskjdsa", City_id: 2, Name: "room2 name", Price: 100, Detail: "room1 detail"}
		resroom2, errroom2 := repo.Create(mockroom2)
		if errroom2 != nil {
			t.Fatal()
		}

		mockroom3 := entities.Room{User_uid: "jkajhskjdsa", City_id: 2, Name: "room3 name", Price: 100, Detail: "room1 detail"}
		_, errroom3 := repo.Create(mockroom3)
		if errroom3 != nil {
			t.Fatal()
		}
		var mockAll []entities.Room

		mockAll, _ = repo.GetAll(city, category, name, length, status)
		assert.Equal(t, mock2.City_id, mockAll[1].City_id)
		// assert.Equal(t, res)
		// log.Info(res)
	})
}
