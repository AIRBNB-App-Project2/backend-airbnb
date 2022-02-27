package booking

import "be/entities"

type Booking interface {
	Create(user_uid string, room_uid string, newBooking BookingReq) (BookingCreateResp, error)
	Update(user_uid string, booking_uid string, upBooking BookingReq) (BookingCreateResp, error)
	GetById(booking_uid string) (BookingGetByIdResp, error)
	Delete(booking_uid string) (entities.Booking, error)
	GetByIdMt(booking_uid string) (entities.Booking, error)
}
