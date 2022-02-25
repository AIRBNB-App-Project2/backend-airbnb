package booking

import "be/entities"

type Booking interface {
	Create(user_uid string, room_uid string, newBooking entities.Booking) (BookingCreateResp, error)
	Update(user_uid string, booking_uid string, newBooking entities.Booking) (BookingCreateResp, error)
	GetById(booking_uid string) (BookingGetByIdResp, error)
	Delete(booking_uid string) (entities.Booking, error)
}
