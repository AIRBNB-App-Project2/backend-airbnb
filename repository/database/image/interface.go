package image

type Image interface {
	Create(room_uid string, image ImageReq) error
}
