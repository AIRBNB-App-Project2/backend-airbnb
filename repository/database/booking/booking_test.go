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

		mockcust := entities.User{Name: "cust name", Email: "cust email", Password: "cust password"}
		rescust, errcust := user.New(db).Create(mockcust)
		if errcust != nil {
			t.Fatal()
		}

		star_date := time.Now().AddDate(0, 0, 2).UTC()
		end_date := time.Now().AddDate(0, 0, 5).UTC()
		mock4 := BookingReq{Start_date: star_date.String(), End_date: end_date.String()}

		res, err := repo.Create(rescust.User_uid, res2.Room_uid, mock4)
		assert.Nil(t, err)
		assert.NotNil(t, res)
		// log.Info(res)
	})

	t.Run("error in time parse star_date", func(t *testing.T) {
		mock1 := entities.User{Name: "user1 name", Email: "user2 email", Password: "user1 password"}
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

		mockcust := entities.User{Name: "cust name", Email: "cust2 email", Password: "cust password"}
		rescust, errcust := user.New(db).Create(mockcust)
		if errcust != nil {
			t.Fatal()
		}
		layoutInput := "02 Jan 2006"
		star_date := time.Now().AddDate(0, 0, 2).Format(layoutInput)
		end_date := time.Now().AddDate(0, 0, 5)

		mock4 := BookingReq{Start_date: star_date, End_date: end_date.String()}

		_, err := repo.Create(rescust.User_uid, res2.Room_uid, mock4)
		assert.NotNil(t, err)
		// log.Info(err)

	})

	t.Run("error in time parse end_date", func(t *testing.T) {
		mock1 := entities.User{Name: "user1 name", Email: "user3 email", Password: "user1 password"}
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

		mockcust := entities.User{Name: "cust name", Email: "cust3 email", Password: "cust password"}
		rescust, errcust := user.New(db).Create(mockcust)
		if errcust != nil {
			t.Fatal()
		}
		layoutInput := "02 Jan 2006"
		star_date := time.Now().AddDate(0, 0, 2)
		end_date := time.Now().AddDate(0, 0, 5).Format(layoutInput)

		mock4 := BookingReq{Start_date: star_date.String(), End_date: end_date}

		_, err := repo.Create(rescust.User_uid, res2.Room_uid, mock4)
		assert.NotNil(t, err)
		// log.Info(err)
	})

	t.Run("error the end_date mas larger than star_date", func(t *testing.T) {
		mock1 := entities.User{Name: "user1 name", Email: "user4 email", Password: "user1 password"}
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

		mockcust := entities.User{Name: "cust name", Email: "cust4 email", Password: "cust password"}
		rescust, errcust := user.New(db).Create(mockcust)
		if errcust != nil {
			t.Fatal()
		}

		star_date := time.Now().AddDate(0, 0, 5).UTC()
		end_date := time.Now().AddDate(0, 0, 2).UTC()

		mock4 := BookingReq{Start_date: star_date.String(), End_date: end_date.String()}

		_, err := repo.Create(rescust.User_uid, res2.Room_uid, mock4)
		assert.NotNil(t, err)
		// log.Info(err)
	})

	t.Run("you are owner", func(t *testing.T) {
		mock1 := entities.User{Name: "user1 name", Email: "user5 email", Password: "user1 password"}
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

		star_date := time.Now().AddDate(0, 0, 2).UTC()
		end_date := time.Now().AddDate(0, 0, 5).UTC()

		mock4 := BookingReq{Start_date: star_date.String(), End_date: end_date.String()}

		_, err := repo.Create(res1.User_uid, res2.Room_uid, mock4)
		assert.NotNil(t, err)
		// log.Info(err)
	})

	t.Run("the room is closed", func(t *testing.T) {
		mock1 := entities.User{Name: "user1 name", Email: "user6 email", Password: "user1 password"}
		res1, err1 := user.New(db).Create(mock1)
		if err1 != nil {
			t.Fatal()
		}
		mock2 := entities.Room{User_uid: res1.User_uid, City_id: 1, Name: "room1 name", Price: 100, Description: "room1 detail", Category: "superior", Status: "close"}
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

		mockcust := entities.User{Name: "cust name", Email: "cust6 email", Password: "cust password"}
		rescust, errcust := user.New(db).Create(mockcust)
		if errcust != nil {
			t.Fatal()
		}

		star_date := time.Now().AddDate(0, 0, 2).UTC()
		end_date := time.Now().AddDate(0, 0, 5).UTC()

		mock4 := BookingReq{Start_date: star_date.String(), End_date: end_date.String()}

		_, err := repo.Create(rescust.User_uid, res2.Room_uid, mock4)
		assert.NotNil(t, err)
	})

	t.Run("the date already picked up", func(t *testing.T) {
		mock1 := entities.User{Name: "user1 name", Email: "user7 email", Password: "user1 password"}
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

		mock1 = entities.User{Name: "cust name", Email: "cust7 email", Password: "cust password"}
		rescust, errcust := user.New(db).Create(mock1)
		if errcust != nil {
			t.Fatal()
		}

		star_date := time.Now().AddDate(0, 0, 2).UTC()
		end_date := time.Now().AddDate(0, 0, 5).UTC()

		mock4 := BookingReq{Start_date: star_date.String(), End_date: end_date.String(), Status: "paid"}
		if _, err := repo.Create(rescust.User_uid, res2.Room_uid, mock4); err != nil {
			t.Fatal()
			// log.Info(res.Booking_uid)
		}
		star_date = time.Now().AddDate(0, 0, 4).UTC()
		end_date = time.Now().AddDate(0, 0, 6).UTC()
		mock4 = BookingReq{Start_date: star_date.String(), End_date: end_date.String()}
		_, err := repo.Create(rescust.User_uid, res2.Room_uid, mock4)
		// log.Info(res.Booking_uid)
		assert.NotNil(t, err)
		// log.Info(err)
	})

}

// func TestUpdate(t *testing.T) {
// 	config := configs.GetConfig()
// 	db := utils.InitDB(config)
// 	repo := New(db)
// 	db.Migrator().DropTable(&entities.User{})
// 	db.Migrator().DropTable(&entities.Room{})
// 	db.Migrator().DropTable(&entities.Image{})
// 	db.Migrator().DropTable(&entities.Booking{})
// 	db.AutoMigrate(&entities.Booking{})
// 	db.AutoMigrate(&entities.Image{})

// 	t.Run("success run Update", func(t *testing.T) {
// 		mock1 := entities.User{Name: "user1 name", Email: "user1 email", Password: "user1 password"}
// 		res1, err1 := user.New(db).Create(mock1)
// 		if err1 != nil {
// 			t.Fatal()
// 		}
// 		mock2 := entities.Room{User_uid: res1.User_uid, City_id: 1, Name: "room1 name", Price: 100, Description: "room1 detail", Category: "superior"}
// 		res2, err2 := room.New(db).Create(mock2)
// 		if err2 != nil {
// 			t.Fatal()
// 		}

// 		mock3 := image.ImageReq{}

// 		for i := 0; i < 3; i++ {
// 			mock3.Array = append(mock3.Array, image.ImageInput{Url: (fmt.Sprintf("url%d", i+1))})
// 		}

// 		if err := image.New(db).Create(res2.Room_uid, mock3); err != nil {
// 			t.Fatal()
// 		}

// 		mock4 := entities.Booking{User_uid: res1.User_uid, Room_uid: res2.Room_uid, Start_date: datatypes.Date(time.Now().AddDate(0, 0, 2)), End_date: datatypes.Date(time.Now().AddDate(0, 0, 5))}
// 		res3, err3 := repo.Create(res1.User_uid, res2.Room_uid, mock4)
// 		if err3 != nil {
// 			t.Fatal()
// 		}
// 		mock5 := entities.Booking{User_uid: res1.User_uid, Room_uid: res2.Room_uid, Start_date: datatypes.Date(time.Now().AddDate(0, 0, 4)), End_date: datatypes.Date(time.Now().AddDate(0, 0, 14))}

// 		// mock5 := entities.Booking{User_uid: res1.User_uid, Room_uid: res2.Room_uid, Status: "reservation"}

// 		res, err := repo.Update(res1.User_uid, res3.Booking_uid, mock5)
// 		assert.Nil(t, err)
// 		assert.NotNil(t, res)
// 		// log.Info(res)
// 	})
// }

// func TestGetByID(t *testing.T) {
// 	config := configs.GetConfig()
// 	db := utils.InitDB(config)
// 	repo := New(db)
// 	db.Migrator().DropTable(&entities.User{})
// 	db.Migrator().DropTable(&entities.Room{})
// 	db.Migrator().DropTable(&entities.Image{})
// 	db.Migrator().DropTable(&entities.Booking{})
// 	db.AutoMigrate(&entities.Booking{})
// 	db.AutoMigrate(&entities.Image{})

// 	t.Run("success run Update", func(t *testing.T) {
// 		mock1 := entities.User{Name: "user1 name", Email: "user1 email", Password: "user1 password"}
// 		res1, err1 := user.New(db).Create(mock1)
// 		if err1 != nil {
// 			t.Fatal()
// 		}
// 		mock2 := entities.Room{User_uid: res1.User_uid, City_id: 1, Name: "room1 name", Price: 100, Description: "room1 detail", Category: "superior"}
// 		res2, err2 := room.New(db).Create(mock2)
// 		if err2 != nil {
// 			t.Fatal()
// 		}

// 		mock3 := image.ImageReq{}

// 		for i := 0; i < 3; i++ {
// 			mock3.Array = append(mock3.Array, image.ImageInput{Url: (fmt.Sprintf("url%d", i+1))})
// 		}

// 		if err := image.New(db).Create(res2.Room_uid, mock3); err != nil {
// 			t.Fatal()
// 		}

// 		mock4 := entities.Booking{User_uid: res1.User_uid, Room_uid: res2.Room_uid, Start_date: datatypes.Date(time.Now().AddDate(0, 0, 2)), End_date: datatypes.Date(time.Now().AddDate(0, 0, 5))}
// 		res3, err3 := repo.Create(res1.User_uid, res2.Room_uid, mock4)
// 		if err3 != nil {
// 			t.Fatal()
// 		}

// 		res, err := repo.GetById(res3.Booking_uid)
// 		assert.Nil(t, err)
// 		assert.NotNil(t, res)
// 		// log.Info(res)
// 	})
// }

// func TestDelete(t *testing.T) {
// 	config := configs.GetConfig()
// 	db := utils.InitDB(config)
// 	repo := New(db)
// 	db.Migrator().DropTable(&entities.User{})
// 	db.Migrator().DropTable(&entities.Room{})
// 	db.Migrator().DropTable(&entities.Image{})
// 	db.Migrator().DropTable(&entities.Booking{})
// 	db.AutoMigrate(&entities.Booking{})
// 	db.AutoMigrate(&entities.Image{})

// 	t.Run("success run Update", func(t *testing.T) {
// 		mock1 := entities.User{Name: "user1 name", Email: "user1 email", Password: "user1 password"}
// 		res1, err1 := user.New(db).Create(mock1)
// 		if err1 != nil {
// 			t.Fatal()
// 		}
// 		mock2 := entities.Room{User_uid: res1.User_uid, City_id: 1, Name: "room1 name", Price: 100, Description: "room1 detail", Category: "superior"}
// 		res2, err2 := room.New(db).Create(mock2)
// 		if err2 != nil {
// 			t.Fatal()
// 		}

// 		mock3 := image.ImageReq{}

// 		for i := 0; i < 3; i++ {
// 			mock3.Array = append(mock3.Array, image.ImageInput{Url: (fmt.Sprintf("url%d", i+1))})
// 		}

// 		if err := image.New(db).Create(res2.Room_uid, mock3); err != nil {
// 			t.Fatal()
// 		}

// 		mock4 := entities.Booking{User_uid: res1.User_uid, Room_uid: res2.Room_uid, Start_date: datatypes.Date(time.Now().AddDate(0, 0, 2)), End_date: datatypes.Date(time.Now().AddDate(0, 0, 5))}
// 		res3, err3 := repo.Create(res1.User_uid, res2.Room_uid, mock4)
// 		if err3 != nil {
// 			t.Fatal()
// 		}

// 		res, err := repo.Delete(res3.Booking_uid)
// 		assert.Nil(t, err)
// 		assert.Equal(t, true, res.DeletedAt.Valid)
// 		// log.Info(res)
// 	})
// }
