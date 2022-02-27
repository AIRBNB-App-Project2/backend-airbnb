package room

import (
	"be/configs"
	"be/entities"
	"be/repository/database/image"
	"be/repository/database/user"
	"be/utils"
	"fmt"
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
	db.AutoMigrate(&entities.Image{})

	t.Run("success run create", func(t *testing.T) {
		mock1 := entities.User{Name: "user1 name", Email: "user1 email", Password: "user1 password"}

		res1, err1 := user.New(db).Create(mock1)
		if err1 != nil {
			t.Fatal()
		}

		mock2 := entities.Room{User_uid: res1.User_uid, City_id: 257, Name: "room1 name", Price: 100, Description: "room1 detail", Address: "room address"}
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
	db.AutoMigrate(&entities.Image{})

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

	t.Run("error record not found", func(t *testing.T) {
		mock1 := entities.User{Name: "user1 name", Email: "user2 email", Password: "user1 password"}
		res1, err1 := user.New(db).Create(mock1)
		if err1 != nil {
			t.Fatal()
		}
		mock2 := entities.Room{User_uid: res1.User_uid, City_id: 1, Name: "room1 name", Price: 100, Description: "room1 detail"}
		_, err2 := repo.Create(mock2)
		if err2 != nil {
			t.Fatal()
		}
		mock3 := entities.Room{Name: "room3 name", Price: 300, Description: "room3 detail"}
		_, err := repo.Update(res1.User_uid, res1.User_uid, mock3)
		// assert.Nil(t, err)
		assert.NotNil(t, err)
		// log.Info(err)
	})
}

func TestGetAllRoom(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)
	db.Migrator().DropTable(&entities.User{})
	db.Migrator().DropTable(&entities.Room{})
	db.Migrator().DropTable(&entities.Image{})
	db.Migrator().DropTable(&entities.Booking{})
	db.AutoMigrate(&entities.Room{})
	db.AutoMigrate(&entities.Image{})

	t.Run("success run GetAllRoom all", func(t *testing.T) {
		mockUser1 := entities.User{Name: "user1 name", Email: "user1 email", Password: "user1 password"}
		resu1, err1 := user.New(db).Create(mockUser1)
		if err1 != nil {
			t.Fatal()
		}
		mockroom1 := entities.Room{User_uid: resu1.User_uid, City_id: 257, Address: "JL.Dramaga", Name: "biasa name", Price: 100, Description: "room1 detail", Status: "open", Category: "standart"}
		resRoom1, errroom1 := repo.Create(mockroom1)
		if errroom1 != nil {
			t.Fatal()
		}
		mockroom2 := entities.Room{User_uid: resu1.User_uid, City_id: 232, Address: "JL.Dramaga", Name: "mewah name", Price: 100, Description: "room2 detail", Status: "open", Category: "superior"}
		_, errroom2 := repo.Create(mockroom2)
		if errroom2 != nil {
			t.Fatal()
		}
		mockroom3 := entities.Room{User_uid: resu1.User_uid, City_id: 212, Address: "JL.Dramaga", Name: "sederhana name", Price: 100, Description: "room1 detail", Status: "open", Category: "luxury"}
		_, errroom3 := repo.Create(mockroom3)
		if errroom3 != nil {
			t.Fatal()
		}
		mockroom4 := entities.Room{User_uid: resu1.User_uid, City_id: 200, Address: "JL.Dramaga", Name: "sederhana name", Price: 100, Description: "room1 detail", Status: "open", Category: "luxury"}
		_, errroom4 := repo.Create(mockroom4)
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

		res, err := repo.GetAllRoom(2, "", "superior")
		assert.Nil(t, err)
		assert.NotNil(t, res)
		// log.Info(res, len(res)

	})

	t.Run("invalid input", func(t *testing.T) {
		mockUser1 := entities.User{Name: "user1 name", Email: "user2 email", Password: "user1 password"}
		resu1, err1 := user.New(db).Create(mockUser1)
		if err1 != nil {
			t.Fatal()
		}
		mockroom1 := entities.Room{User_uid: resu1.User_uid, City_id: 257, Address: "JL.Dramaga", Name: "biasa name", Price: 100, Description: "room1 detail", Status: "open", Category: "standart"}
		resRoom1, errroom1 := repo.Create(mockroom1)
		if errroom1 != nil {
			t.Fatal()
		}
		mockroom2 := entities.Room{User_uid: resu1.User_uid, City_id: 232, Address: "JL.Dramaga", Name: "mewah name", Price: 100, Description: "room2 detail", Status: "open", Category: "superior"}
		_, errroom2 := repo.Create(mockroom2)
		if errroom2 != nil {
			t.Fatal()
		}
		mockroom3 := entities.Room{User_uid: resu1.User_uid, City_id: 212, Address: "JL.Dramaga", Name: "sederhana name", Price: 100, Description: "room1 detail", Status: "open", Category: "luxury"}
		_, errroom3 := repo.Create(mockroom3)
		if errroom3 != nil {
			t.Fatal()
		}
		mockroom4 := entities.Room{User_uid: resu1.User_uid, City_id: 200, Address: "JL.Dramaga", Name: "sederhana name", Price: 100, Description: "room1 detail", Status: "open", Category: "luxury"}
		_, errroom4 := repo.Create(mockroom4)
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

		res, err := repo.GetAllRoom(2, "nodata", "nodata")
		assert.Nil(t, err)
		assert.Equal(t, 0, len(res))
		// log.Info(err)

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
	db.AutoMigrate(&entities.Room{})
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
		// log.Info(res)
	})

	t.Run("fail run GetById", func(t *testing.T) {
		mock1 := entities.User{Name: "user1 name", Email: "user2 email", Password: "user1 password"}
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

		_, err := repo.GetById(res1.User_uid)
		// assert.Nil(t, err)
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
	db.AutoMigrate(&entities.Room{})
	db.AutoMigrate(&entities.Image{})

	t.Run("success Delete", func(t *testing.T) {
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

		res, err := repo.Delete(res2.Room_uid)
		assert.Nil(t, err)
		assert.Equal(t, true, res.DeletedAt.Valid)
	})

	t.Run("fail Delete", func(t *testing.T) {
		mock1 := entities.User{Name: "user1 name", Email: "user2 email", Password: "user1 password"}
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

		_, err := repo.Delete(res1.User_uid)
		assert.NotNil(t, err)
	})
}
