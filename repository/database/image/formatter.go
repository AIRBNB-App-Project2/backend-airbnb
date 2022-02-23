package image

type ImageInput struct {
	Url string `json:"url"`
}

type ImageReq struct {
	Array []ImageInput `json:"array"`
}
