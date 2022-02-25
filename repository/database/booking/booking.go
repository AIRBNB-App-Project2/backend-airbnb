package booking

import (
	"be/entities"
	"errors"

	"github.com/lithammer/shortuuid"
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

func (repo *BookingDb) Create(user_uid string, room_uid string, newBooking entities.Booking) (BookingCreateResp, error) {

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

	newBooking.Booking_uid = uid

	// check reservation
	bookingInit := entities.Booking{}

	resRev := repo.db.Model(entities.Booking{}).Where("status = 'reservation' or status = 'onGoing' AND start_date <= ? AND end_date >= ?", newBooking.End_date, newBooking.Start_date).Find(&bookingInit)

	if resRev.RowsAffected != 0 {
		return BookingCreateResp{}, errors.New("the date already picked up")
	}

	// validate owner
	newBooking.User_uid = user_uid
	newBooking.Room_uid = room_uid
	res := repo.db.Model(&entities.Booking{}).Create(&newBooking)

	if res.Error != nil {
		return BookingCreateResp{}, res.Error
	}

	bookingres := BookingCreateResp{}

	resp := repo.db.Model(&entities.Booking{}).Where("booking_uid =?", newBooking.Booking_uid).Select("bookings.booking_uid as Booking_uid, rooms.name as Name, rooms.description as Description,rooms.price as Price ,bookings.start_date as Start_date, bookings.end_date as End_date, DATEDIFF(bookings.end_date, bookings.start_date) as Days, DATEDIFF(bookings.end_date, bookings.start_date) * rooms.price as Price_total").Joins("inner join rooms on bookings.room_uid = rooms.room_uid").First(&bookingres)
	if resp.Error != nil {
		return BookingCreateResp{}, resp.Error
	}

	// layoutIso := "2006-01-02T00:00:00+07:00"
	// start_date, err := time.Parse(layoutIso, bookingres.Start_date)
	// if err != nil {
	// 	return BookingCreateResp{}, err
	// }
	// bookingres.Start_date = start_date.Format(time.RFC822)

	// end_date, err := time.Parse(layoutIso, bookingres.End_date)
	// if err != nil {
	// 	return BookingCreateResp{}, err
	// }
	// bookingres.End_date = end_date.Format(time.RFC822)

	return bookingres, nil
}

func (repo *BookingDb) Update(user_uid string, booking_uid string, newBooking entities.Booking) (BookingCreateResp, error) {

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

	if newBooking.Status == "" {
		resRev := tx.Model(entities.Booking{}).Where("status = 'reservation' or status = 'onGoing' AND start_date <= ? AND end_date >= ? AND booking_uid != ?", newBooking.End_date, newBooking.Start_date, booking_uid).Find(&bookingInit)

		if resRev.RowsAffected != 0 {
			// log.Info("test")
			tx.Rollback()
			return BookingCreateResp{}, errors.New("the date already picked up")
		}
	}

	if res := tx.Model(&entities.Booking{}).Where("booking_uid = ?", booking_uid).Find(&bookingInit); res.Error != nil {
		tx.Rollback()
		// log.Info(res)
		return BookingCreateResp{}, res.Error
	}

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

	if res := tx.Model(&entities.Booking{}).Where("booking_uid = ?", booking_uid).Updates(entities.Booking{Start_date: newBooking.Start_date, End_date: newBooking.End_date, Status: newBooking.Status}); res.Error != nil {
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
