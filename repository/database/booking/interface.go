package booking

import "be/entities"

type Booking interface {
	Create(user_uid string, room_uid string, newBooking entities.Booking) (BookingCreateResp, error)
}
