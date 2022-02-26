package user

import (
	"be/configs"
	"be/entities"
	"be/repository/database/booking"
	"be/repository/database/image"
	"be/repository/database/room"
	"be/utils"
	"fmt"
	"testing"
	"time"

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
		// log.Info(err)
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
	db.AutoMigrate(&entities.Room{})
	db.AutoMigrate(&entities.Image{})
	db.AutoMigrate(&entities.Booking{})

	t.Run("success run GetById", func(t *testing.T) {
		mockUser1 := entities.User{Name: "anonim1", Email: "anonim1", Password: "anonim1"}
		res1, err1 := repo.Create(mockUser1)
		if err1 != nil {
			t.Fatal()
		}
		res, err := repo.GetById(res1.User_uid)
		assert.Nil(t, err)
		assert.NotNil(t, res)
		// log.Info(res)
	})

	t.Run("fail run GetById", func(t *testing.T) {
		mockUser1 := entities.User{Name: "anonim1", Email: "anonim2", Password: "anonim1"}
		_, err1 := repo.Create(mockUser1)
		if err1 != nil {
			t.Fatal()
		}
		_, err := repo.GetById("")
		assert.NotNil(t, err)
		// log.Info(res)
	})

	t.Run("success run user profile", func(t *testing.T) {
		mockUser1 := entities.User{Name: "user1 name", Email: "user1 email", Password: "user1 password"}
		resu1, err1 := repo.Create(mockUser1)
		if err1 != nil {
			t.Fatal()
		}
		mockroom1 := entities.Room{User_uid: resu1.User_uid, City_id: 257, Address: "JL.Dramaga", Name: "biasa name", Price: 100, Description: "room1 detail", Status: "open", Category: "standart"}
		resRoom1, errroom1 := room.New(db).Create(mockroom1)
		if errroom1 != nil {
			t.Fatal()
		}
		mockroom2 := entities.Room{User_uid: resu1.User_uid, City_id: 232, Address: "JL.Dramaga", Name: "mewah name", Price: 100, Description: "room2 detail", Status: "open", Category: "superior"}
		_, errroom2 := room.New(db).Create(mockroom2)
		if errroom2 != nil {
			t.Fatal()
		}
		mockroom3 := entities.Room{User_uid: resu1.User_uid, City_id: 212, Address: "JL.Dramaga", Name: "sederhana name", Price: 100, Description: "room1 detail", Status: "open", Category: "luxury"}
		_, errroom3 := room.New(db).Create(mockroom3)
		if errroom3 != nil {
			t.Fatal()
		}
		mockroom4 := entities.Room{User_uid: resu1.User_uid, City_id: 200, Address: "JL.Dramaga", Name: "sederhana name", Price: 100, Description: "room1 detail", Status: "open", Category: "luxury"}
		_, errroom4 := room.New(db).Create(mockroom4)
		if errroom4 != nil {
			t.Fatal()
		}
		mock3 := image.ImageReq{}

		for i := 0; i < 3; i++ {
			mock3.Array = append(mock3.Array, image.ImageInput{Url: (fmt.Sprintf("url%d", i+1))})
		}

		if err := image.New(db).Create(resRoom1.Room_uid, mock3); err != nil {
			t.Fatal()
		}

		// booking

		mock1B := entities.User{Name: "user1 name", Email: "user1B email", Password: "user1 password"}
		res1, err1 := repo.Create(mock1B)
		if err1 != nil {
			t.Fatal()
		}
		mock2B := entities.Room{User_uid: res1.User_uid, City_id: 1, Name: "room1 name", Price: 100, Description: "room1 detail", Category: "superior"}
		res2, err2 := room.New(db).Create(mock2B)
		if err2 != nil {
			t.Fatal()
		}

		mock3B := image.ImageReq{}

		for i := 0; i < 3; i++ {
			mock3B.Array = append(mock3.Array, image.ImageInput{Url: (fmt.Sprintf("url%d", i+1))})
		}
		if err := image.New(db).Create(res2.Room_uid, mock3B); err != nil {
			t.Fatal()
		}

		star_date := time.Now().AddDate(0, 0, 2).UTC()
		end_date := time.Now().AddDate(0, 0, 5).UTC()
		mock4 := booking.BookingReq{Start_date: star_date.String(), End_date: end_date.String()}

		_, err4 := booking.New(db).Create(resu1.User_uid, res2.Room_uid, mock4)
		if err4 != nil {
			t.Fatal()
		}

		res, err := repo.GetById(resu1.User_uid)
		assert.Nil(t, err)
		assert.NotNil(t, res)
		// log.Info(res, err)

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
	db.AutoMigrate(&entities.User{})

	t.Run("success run Update", func(t *testing.T) {
		mockUser1 := entities.User{Name: "anonim1", Email: "anonim1", Password: "anonim1"}
		res1, err1 := repo.Create(mockUser1)
		if err1 != nil {
			t.Fatal()
		}
		mockUpUser := entities.User{Name: "anonim2", Email: "anonim2", Password: "anonim2"}
		res, err := repo.Update(res1.User_uid, mockUpUser)
		assert.Nil(t, err)
		assert.NotNil(t, res)
		// log.Info(res)
	})

	t.Run("fail run Update", func(t *testing.T) {
		mockUser1 := entities.User{Name: "anonim1", Email: "anonim222", Password: "anonim1"}
		_, err1 := repo.Create(mockUser1)
		if err1 != nil {
			t.Fatal()
		}
		mockUpUser := entities.User{Name: "anonim2", Email: "anonim2", Password: "anonim2"}
		_, err := repo.Update("", mockUpUser)
		assert.NotNil(t, err)
		// log.Info(res)
	})
}

func TestDelete(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)
	db.Migrator().DropTable(&entities.User{})
	db.Migrator().DropTable(&entities.Room{})
	db.Migrator().DropTable(&entities.Image{})
	db.Migrator().DropTable(&entities.Booking{})
	db.AutoMigrate(&entities.User{})

	t.Run("success run create", func(t *testing.T) {
		mockUser1 := entities.User{Name: "anonim1", Email: "anonim1", Password: "anonim1"}
		res1, err1 := repo.Create(mockUser1)
		if err1 != nil {
			t.Fatal()
		}
		res, err := repo.Delete(res1.User_uid)
		assert.Nil(t, err)
		assert.Equal(t, true, res.DeletedAt.Valid)
	})

	t.Run("fail run create", func(t *testing.T) {
		mockUser1 := entities.User{Name: "anonim1", Email: "anonim2", Password: "anonim1"}
		_, err1 := repo.Create(mockUser1)
		if err1 != nil {
			t.Fatal()
		}
		_, err := repo.Delete("")
		assert.NotNil(t, err)
	})
}
