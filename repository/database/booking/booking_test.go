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

	t.Run("error in time is in the past", func(t *testing.T) {
		mock1 := entities.User{Name: "user1 name", Email: "user30 email", Password: "user1 password"}
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

		mockcust := entities.User{Name: "cust name", Email: "cust30 email", Password: "cust password"}
		rescust, errcust := user.New(db).Create(mockcust)
		if errcust != nil {
			t.Fatal()
		}
		// layoutInput := "02 Jan 2006"
		star_date := time.Now().AddDate(0, 0, -5).UTC()
		end_date := time.Now().AddDate(0, 0, -2).UTC()

		mock4 := BookingReq{Start_date: star_date.String(), End_date: end_date.String()}

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

func TestUpdate(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)
	db.Migrator().DropTable(&entities.User{})
	db.Migrator().DropTable(&entities.Room{})
	db.Migrator().DropTable(&entities.Image{})
	db.Migrator().DropTable(&entities.Booking{})
	db.AutoMigrate(&entities.Booking{})
	db.AutoMigrate(&entities.Image{})

	t.Run("success run Update date", func(t *testing.T) {
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

		res4, err4 := repo.Create(rescust.User_uid, res2.Room_uid, mock4)
		if err4 != nil {
			t.Fatal()
		}

		star_date = time.Now().AddDate(0, 0, 4).UTC()
		end_date = time.Now().AddDate(0, 0, 6).UTC()

		mocUpdate := BookingReq{Start_date: star_date.String(), End_date: end_date.String()}

		res, err := repo.Update(rescust.User_uid, res4.Booking_uid, mocUpdate)
		assert.Nil(t, err)
		assert.NotNil(t, res)
		// log.Info(res, err)
	})

	t.Run("success run Update status", func(t *testing.T) {
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

		star_date := time.Now().AddDate(0, 0, 2).UTC()
		end_date := time.Now().AddDate(0, 0, 5).UTC()
		mock4 := BookingReq{Start_date: star_date.String(), End_date: end_date.String()}

		res4, err4 := repo.Create(rescust.User_uid, res2.Room_uid, mock4)
		if err4 != nil {
			t.Fatal()
		}

		mocUpdate := BookingReq{Status: "paid"}

		res, err := repo.Update(rescust.User_uid, res4.Booking_uid, mocUpdate)
		assert.Nil(t, err)
		assert.NotNil(t, res)
		log.Info(res)
	})

	t.Run("error in time parse start date", func(t *testing.T) {
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

		star_date := time.Now().AddDate(0, 0, 6).UTC()
		end_date := time.Now().AddDate(0, 0, 7).UTC()
		mock4 := BookingReq{Start_date: star_date.String(), End_date: end_date.String()}

		res4, err4 := repo.Create(rescust.User_uid, res2.Room_uid, mock4)
		if err4 != nil {
			t.Fatal()
		}

		star_date = time.Now().AddDate(0, 0, 8)
		end_date = time.Now().AddDate(0, 0, 9).UTC()

		mocUpdate := BookingReq{Start_date: star_date.String(), End_date: end_date.String()}

		_, err := repo.Update(rescust.User_uid, res4.Booking_uid, mocUpdate)
		// assert.Nil(t, err)
		assert.NotNil(t, err)
		// log.Info(err)
	})

	t.Run("error in time parse end date", func(t *testing.T) {
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

		star_date := time.Now().AddDate(0, 0, 8).UTC()
		end_date := time.Now().AddDate(0, 0, 9).UTC()
		mock4 := BookingReq{Start_date: star_date.String(), End_date: end_date.String()}

		res4, err4 := repo.Create(rescust.User_uid, res2.Room_uid, mock4)
		if err4 != nil {
			t.Fatal()
		}

		star_date = time.Now().AddDate(0, 0, 8).UTC()
		end_date = time.Now().AddDate(0, 0, 9)

		mocUpdate := BookingReq{Start_date: star_date.String(), End_date: end_date.String()}

		_, err := repo.Update(rescust.User_uid, res4.Booking_uid, mocUpdate)
		// assert.Nil(t, err)
		assert.NotNil(t, err)
		// log.Info(err)
	})

	t.Run("error in time is in the past", func(t *testing.T) {
		mock1 := entities.User{Name: "user1 name", Email: "user50 email", Password: "user1 password"}
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

		mockcust := entities.User{Name: "cust name", Email: "cust50 email", Password: "cust password"}
		rescust, errcust := user.New(db).Create(mockcust)
		if errcust != nil {
			t.Fatal()
		}

		star_date := time.Now().AddDate(0, 0, 50).UTC()
		end_date := time.Now().AddDate(0, 0, 51).UTC()
		mock4 := BookingReq{Start_date: star_date.String(), End_date: end_date.String()}

		res4, err4 := repo.Create(rescust.User_uid, res2.Room_uid, mock4)
		if err4 != nil {
			t.Fatal()
		}

		star_date = time.Now().AddDate(0, 0, -11).UTC()
		end_date = time.Now().AddDate(0, 0, -10).UTC()

		mocUpdate := BookingReq{Start_date: star_date.String(), End_date: end_date.String()}

		_, err := repo.Update(rescust.User_uid, res4.Booking_uid, mocUpdate)
		// assert.Nil(t, err)
		assert.NotNil(t, err)
		// log.Info(err)
	})

	t.Run("error the end date must larger than start date", func(t *testing.T) {
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

		mockcust := entities.User{Name: "cust name", Email: "cust5 email", Password: "cust password"}
		rescust, errcust := user.New(db).Create(mockcust)
		if errcust != nil {
			t.Fatal()
		}

		star_date := time.Now().AddDate(0, 0, 10).UTC()
		end_date := time.Now().AddDate(0, 0, 11).UTC()
		mock4 := BookingReq{Start_date: star_date.String(), End_date: end_date.String()}

		res4, err4 := repo.Create(rescust.User_uid, res2.Room_uid, mock4)
		if err4 != nil {
			t.Fatal()
		}

		star_date = time.Now().AddDate(0, 0, 11).UTC()
		end_date = time.Now().AddDate(0, 0, 10).UTC()

		mocUpdate := BookingReq{Start_date: star_date.String(), End_date: end_date.String()}

		_, err := repo.Update(rescust.User_uid, res4.Booking_uid, mocUpdate)
		// assert.Nil(t, err)
		assert.NotNil(t, err)
		// log.Info(err)
	})

	t.Run("error bookiInit is empyty", func(t *testing.T) {
		mock1 := entities.User{Name: "user1 name", Email: "user6 email", Password: "user1 password"}
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

		mockcust := entities.User{Name: "cust name", Email: "cust6 email", Password: "cust password"}
		rescust, errcust := user.New(db).Create(mockcust)
		if errcust != nil {
			t.Fatal()
		}

		star_date := time.Now().AddDate(0, 0, 12).UTC()
		end_date := time.Now().AddDate(0, 0, 13).UTC()
		mock4 := BookingReq{Start_date: star_date.String(), End_date: end_date.String()}

		_, err4 := repo.Create(rescust.User_uid, res2.Room_uid, mock4)
		if err4 != nil {
			t.Fatal()
		}

		star_date = time.Now().AddDate(0, 0, 12).UTC()
		end_date = time.Now().AddDate(0, 0, 13).UTC()

		mocUpdate := BookingReq{Start_date: star_date.String(), End_date: end_date.String()}

		_, err := repo.Update(rescust.User_uid, res2.Room_uid, mocUpdate)
		// assert.Nil(t, err)
		assert.NotNil(t, err)
		// log.Info(err)
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

		mockcust := entities.User{Name: "cust name", Email: "cust7 email", Password: "cust password"}
		rescust, errcust := user.New(db).Create(mockcust)
		if errcust != nil {
			t.Fatal()
		}

		star_date := time.Now().AddDate(0, 0, 100).UTC()
		end_date := time.Now().AddDate(0, 0, 150).UTC()
		mock4 := BookingReq{Start_date: star_date.String(), End_date: end_date.String(), Status: "paid"}

		_, err4 := repo.Create(rescust.User_uid, res2.Room_uid, mock4)
		if err4 != nil {
			t.Fatal()
		}
		// log.Info(res4)
		star_date = time.Now().AddDate(0, 0, 14).UTC()
		end_date = time.Now().AddDate(0, 0, 15).UTC()
		mock4 = BookingReq{Start_date: star_date.String(), End_date: end_date.String()}

		res5, err5 := repo.Create(rescust.User_uid, res2.Room_uid, mock4)
		if err5 != nil {
			t.Fatal()
		}

		star_date = time.Now().AddDate(0, 0, 140).UTC()
		end_date = time.Now().AddDate(0, 0, 150).UTC()

		mocUpdate := BookingReq{Start_date: star_date.String(), End_date: end_date.String()}

		_, err := repo.Update(rescust.User_uid, res5.Booking_uid, mocUpdate)
		// assert.Nil(t, err)
		assert.NotNil(t, err)
		// log.Info(err)
	})

	t.Run("the room is closed", func(t *testing.T) {
		mock1 := entities.User{Name: "user1 name", Email: "user8 email", Password: "user1 password"}
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

		mockcust := entities.User{Name: "cust name", Email: "cust8 email", Password: "cust password"}
		rescust, errcust := user.New(db).Create(mockcust)
		if errcust != nil {
			t.Fatal()
		}

		star_date := time.Now().AddDate(0, 0, 16).UTC()
		end_date := time.Now().AddDate(0, 0, 17).UTC()
		mock4 := BookingReq{Start_date: star_date.String(), End_date: end_date.String()}

		res4, err4 := repo.Create(rescust.User_uid, res2.Room_uid, mock4)
		if err4 != nil {
			t.Fatal()
		}

		if _, err := room.New(db).Update(res1.User_uid, res2.Room_uid, entities.Room{Status: "close"}); err != nil {
			// log.Info(err)
			t.Fatal()
		}

		star_date = time.Now().AddDate(0, 0, 18).UTC()
		end_date = time.Now().AddDate(0, 0, 19).UTC()

		mocUpdate := BookingReq{Start_date: star_date.String(), End_date: end_date.String()}

		_, err := repo.Update(rescust.User_uid, res4.Booking_uid, mocUpdate)
		// assert.Nil(t, err)
		assert.NotNil(t, err)
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
	db.AutoMigrate(&entities.Booking{})
	db.AutoMigrate(&entities.Image{})

	t.Run("success run get by id", func(t *testing.T) {
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

		res4, err4 := repo.Create(rescust.User_uid, res2.Room_uid, mock4)
		if err4 != nil {
			t.Fatal()
		}

		res, err := repo.GetById(res4.Booking_uid)
		assert.Nil(t, err)
		assert.NotNil(t, res)
		// log.Info(res)
	})

	t.Run("fail run get by id", func(t *testing.T) {
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

		star_date := time.Now().AddDate(0, 0, 2).UTC()
		end_date := time.Now().AddDate(0, 0, 5).UTC()
		mock4 := BookingReq{Start_date: star_date.String(), End_date: end_date.String()}

		_, err4 := repo.Create(rescust.User_uid, res2.Room_uid, mock4)
		if err4 != nil {
			t.Fatal()
		}

		_, err := repo.GetById(res2.Room_uid)
		assert.NotNil(t, err)
		// log.Info(err)
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
	db.AutoMigrate(&entities.Booking{})
	db.AutoMigrate(&entities.Image{})

	t.Run("success run delete", func(t *testing.T) {
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

		res4, err4 := repo.Create(rescust.User_uid, res2.Room_uid, mock4)
		if err4 != nil {
			t.Fatal()
		}

		res, err := repo.Delete(res4.Booking_uid)
		assert.Nil(t, err)
		assert.NotNil(t, res)
		// log.Info(res)
	})

	t.Run("fail run delete", func(t *testing.T) {
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

		star_date := time.Now().AddDate(0, 0, 2).UTC()
		end_date := time.Now().AddDate(0, 0, 5).UTC()
		mock4 := BookingReq{Start_date: star_date.String(), End_date: end_date.String()}

		_, err4 := repo.Create(rescust.User_uid, res2.Room_uid, mock4)
		if err4 != nil {
			t.Fatal()
		}

		_, err := repo.Delete(res2.Room_uid)
		assert.NotNil(t, err)
		// log.Info(err)
	})

}

func TestGetByIDMt(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)
	db.Migrator().DropTable(&entities.User{})
	db.Migrator().DropTable(&entities.Room{})
	db.Migrator().DropTable(&entities.Image{})
	db.Migrator().DropTable(&entities.Booking{})
	db.AutoMigrate(&entities.Booking{})
	db.AutoMigrate(&entities.Image{})

	t.Run("success run GetByIdMt", func(t *testing.T) {
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

		res4, err4 := repo.Create(rescust.User_uid, res2.Room_uid, mock4)
		if err4 != nil {
			t.Fatal()
		}

		res, err := repo.GetByIdMt(res4.Booking_uid)
		assert.Nil(t, err)
		assert.NotNil(t, res)
	})

	t.Run("fail run GetByIdMt", func(t *testing.T) {
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

		star_date := time.Now().AddDate(0, 0, 2).UTC()
		end_date := time.Now().AddDate(0, 0, 5).UTC()
		mock4 := BookingReq{Start_date: star_date.String(), End_date: end_date.String()}

		_, err4 := repo.Create(rescust.User_uid, res2.Room_uid, mock4)
		if err4 != nil {
			t.Fatal()
		}

		_, err := repo.GetByIdMt(rescust.User_uid)
		assert.NotNil(t, err)
	})
}
