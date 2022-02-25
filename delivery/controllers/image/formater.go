package image

type CreateImageRequesFormat struct {
	Room_uid string `form:"room_uid"`
	Url      string
}

type ImageInput struct {
	Url string `json:"url"`
}

type ImageReq struct {
	Array []ImageInput `json:"array"`
}
