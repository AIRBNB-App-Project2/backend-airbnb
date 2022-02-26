package booking

import (
	"be/entities"
	"errors"
	"time"

	"github.com/lithammer/shortuuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type BookingDb struct {
	db *gorm.DB
}

func New(db *gorm.DB) *BookingDb {
	return &BookingDb{
		db: db,
	}
}

func (repo *BookingDb) Create(user_uid string, room_uid string, newBooking BookingReq) (BookingCreateResp, error) {

	bookingCreate := entities.Booking{}

	layoutInput := "2006-01-02 15:04:05 +0000 UTC"
	s_date, errS := time.Parse(layoutInput, newBooking.Start_date)
	if errS != nil {
		return BookingCreateResp{}, errors.New("error in time parse start date")
	}
	e_date, errE := time.Parse(layoutInput, newBooking.End_date)
	if errE != nil {
		return BookingCreateResp{}, errors.New("error in time parse end date")
	}
	days := e_date.Sub(s_date).Hours() / 24
	// log.Info(days)
	if days < 0 {
		return BookingCreateResp{}, errors.New("error the end date must larger than start date")
	}

	var uid string

	for {
		uid = shortuuid.New()
		find := entities.Booking{}
		res := repo.db.Model(&entities.Booking{}).Where("booking_uid =  ?", uid).First(&find)
		if res.RowsAffected == 0 {
			break
		}
		if res.Error != nil {
			return BookingCreateResp{}, res.Error
		}
	}

	bookingCreate.Booking_uid = uid

	// validate owner
	roomOwn := entities.Room{}
	if err := repo.db.Model(entities.Room{}).Where("room_uid = ?", room_uid).Find(&roomOwn).Error; err != nil {
		return BookingCreateResp{}, err
	}
	if roomOwn.User_uid == user_uid {
		return BookingCreateResp{}, errors.New("you are owner")
	}

	if roomOwn.Status == "close" {
		return BookingCreateResp{}, errors.New("the room is closed")
	}

	bookingCreate.User_uid = user_uid
	bookingCreate.Room_uid = room_uid
	bookingCreate.Start_date = datatypes.Date(s_date)
	bookingCreate.End_date = datatypes.Date(e_date)
	bookingCreate.PaymentMethod = newBooking.PaymentMethod
	bookingCreate.Status = newBooking.Status
	// check reservation

	resRev := repo.db.Model(entities.Booking{}).Where("status = 'paid' AND start_date <= ? AND end_date >= ?", bookingCreate.End_date, bookingCreate.Start_date).Find(&bookingCreate)

	if resRev.RowsAffected != 0 {
		return BookingCreateResp{}, errors.New("the date already picked up")
	}

	res := repo.db.Model(&entities.Booking{}).Create(&bookingCreate)

	if res.Error != nil {
		return BookingCreateResp{}, res.Error
	}

	bookingres := BookingCreateResp{}

	resp := repo.db.Model(&entities.Booking{}).Where("booking_uid =?", bookingCreate.Booking_uid).Select("bookings.booking_uid as Booking_uid, rooms.name as Name, rooms.description as Description,rooms.price as Price ,bookings.start_date as Start_date, bookings.end_date as End_date, DATEDIFF(bookings.end_date, bookings.start_date) as Days, DATEDIFF(bookings.end_date, bookings.start_date) * rooms.price as Price_total").Joins("inner join rooms on bookings.room_uid = rooms.room_uid").First(&bookingres)
	if resp.Error != nil {
		return BookingCreateResp{}, resp.Error
	}
	layoutOutput := "2006-01-02T00:00:00+07:00"
	start_date, err := time.Parse(layoutOutput, bookingres.Start_date)
	if err != nil {
		return BookingCreateResp{}, err
	}
	outputDateFormat := "02 Jan 2006"
	bookingres.Start_date = start_date.Format(outputDateFormat)

	end_date, err := time.Parse(layoutOutput, bookingres.End_date)
	if err != nil {
		return BookingCreateResp{}, err
	}
	bookingres.End_date = end_date.Format(outputDateFormat)

	return bookingres, nil
}

func (repo *BookingDb) Update(user_uid string, booking_uid string, upBooking BookingReq) (BookingCreateResp, error) {

	bookingUpdate := entities.Booking{}

	layoutInput := "2006-01-02 15:04:05 +0000 UTC"
	s_date, errS := time.Parse(layoutInput, upBooking.Start_date)
	if errS != nil {
		return BookingCreateResp{}, errors.New("error in time parse start date")
	}
	e_date, errE := time.Parse(layoutInput, upBooking.End_date)
	if errE != nil {
		return BookingCreateResp{}, errors.New("error in time parse end date")
	}
	days := e_date.Sub(s_date).Hours() / 24
	// log.Info(days)
	if days < 0 {
		return BookingCreateResp{}, errors.New("error the end date must larger than start date")
	}

	bookingUpdate.Start_date = datatypes.Date(s_date)
	bookingUpdate.End_date = datatypes.Date(e_date)
	bookingUpdate.PaymentMethod = upBooking.PaymentMethod
	bookingUpdate.Status = upBooking.Status


	tx := repo.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return BookingCreateResp{}, err
	}

	bookingInit := entities.Booking{}

	if res := tx.Model(&entities.Booking{}).Where("booking_uid = ?", booking_uid).Find(&bookingInit); res.Error != nil {
		tx.Rollback()
		// log.Info(res)
		return BookingCreateResp{}, res.Error
	}

	if bookingUpdate.Status == "" {
		resRev := tx.Model(entities.Booking{}).Where("status = 'paid' AND start_date <= ? AND end_date >= ? AND booking_uid != ?", bookingUpdate.End_date, bookingUpdate.Start_date, booking_uid).Find(&bookingInit)
		// log.Info(resRev.RowsAffected)
		if resRev.RowsAffected != 0 {
			// log.Info("test")
			tx.Rollback()
			return BookingCreateResp{}, errors.New("the date already picked up")
		}
		// log.Info(bookingInit)
		roomOwn := entities.Room{}
		if err := repo.db.Model(entities.Room{}).Where("room_uid = ?", bookingInit.Room_uid).Find(&roomOwn).Error; err != nil {
			tx.Rollback()
			return BookingCreateResp{}, err
		}
		// log.Info(roomOwn.Status)
		if roomOwn.Status == "close" {
			tx.Rollback()
			return BookingCreateResp{}, errors.New("the room is closed")
		}
	}
	// log.Info(bookingInit)
	if bookingInit.User_uid != user_uid {
		tx.Rollback()
		// log.Info("test")
		return BookingCreateResp{}, errors.New(gorm.ErrInvalidData.Error())
	}

	if res := tx.Model(&entities.Booking{}).Where("booking_uid = ?", booking_uid).Delete(&bookingInit); res.RowsAffected == 0 {
		// log.Info("test")
		// log.Info(res.RowsAffected)
		tx.Rollback()
		return BookingCreateResp{}, errors.New(gorm.ErrRecordNotFound.Error())
	}
	bookingInit.DeletedAt = gorm.DeletedAt{}
	bookingInit.ID = 0
	if res := tx.Create(&bookingInit); res.Error != nil {
		// log.Info("test")
		return BookingCreateResp{}, res.Error
	}

	if res := tx.Model(&entities.Booking{}).Where("booking_uid = ?", booking_uid).Updates(entities.Booking{Start_date: bookingUpdate.Start_date, End_date: bookingUpdate.End_date, Status: bookingUpdate.Status, PaymentMethod: bookingUpdate.PaymentMethod}); res.Error != nil {
		// log.Info(res.Error)
		tx.Rollback()
		return BookingCreateResp{}, res.Error
	}

	tx.Commit()

	bookingres := BookingCreateResp{}

	resp := repo.db.Model(&entities.Booking{}).Where("booking_uid =?", booking_uid).Select("bookings.booking_uid as Booking_uid, rooms.name as Name, rooms.description as Description,rooms.price as Price ,bookings.start_date as Start_date, bookings.end_date as End_date, DATEDIFF(bookings.end_date, bookings.start_date) as Days, DATEDIFF(bookings.end_date, bookings.start_date) * rooms.price as Price_total").Joins("inner join rooms on bookings.room_uid = rooms.room_uid").First(&bookingres)
	if resp.Error != nil {
		// log.Info(resp.Error)
		return BookingCreateResp{}, resp.Error
	}

	return bookingres, nil

}

func (repo *BookingDb) GetById(booking_uid string) (BookingGetByIdResp, error) {

	bookingResp := BookingGetByIdResp{}

	resp := repo.db.Model(&entities.Booking{}).Where("booking_uid =?", booking_uid).Select("bookings.booking_uid as Booking_uid, rooms.name as Name, rooms.description as Description,rooms.price as Price ,bookings.start_date as Start_date, bookings.end_date as End_date, DATEDIFF(bookings.end_date, bookings.start_date) as Days, DATEDIFF(bookings.end_date, bookings.start_date) * rooms.price as Price_total, bookings.status as Status").Joins("inner join rooms on bookings.room_uid = rooms.room_uid").First(&bookingResp)
	if resp.Error != nil {
		// log.Info(resp.Error)
		return BookingGetByIdResp{}, resp.Error
	}

	return bookingResp, nil
}

func (repo *BookingDb) Delete(booking_uid string) (entities.Booking, error) {

	bookingResp := entities.Booking{}

	resp := repo.db.Model(&entities.Booking{}).Where("booking_uid =?", booking_uid).Delete(&bookingResp)

	if resp.RowsAffected == 0 {
		return entities.Booking{}, errors.New(gorm.ErrRecordNotFound.Error())
	}

	return bookingResp, nil
}
