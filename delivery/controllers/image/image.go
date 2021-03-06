package image

import (
	"be/delivery/controllers/templates"
	imageRepo "be/repository/database/image"
	"be/utils"
	"net/http"

	"github.com/labstack/gommon/log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/labstack/echo/v4"
)

type ImageController struct {
	repo imageRepo.Image
}

func New(repo imageRepo.Image) *ImageController {
	return &ImageController{
		repo: repo,
	}
}

func (ic *ImageController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		//default image

		// image.Room_uid = c.FormValue("room_uid")

		form, err1 := c.MultipartForm()
		if err1 != nil {
			return c.JSON(http.StatusBadRequest, templates.BadRequest(nil, "error in multipart form", nil))
		}
		files := form.File["files"]

		for _, file := range files {
			image := CreateImageRequesFormat{}
			src, err := file.Open()
			if err != nil {
				return err
			}
			log.Info(src)

			s, _ := session.NewSession(&aws.Config{
				Region: aws.String("ap-southeast-1"),
				Credentials: credentials.NewStaticCredentials(
					"AKIARYSBMFQ57G7DPE6Q",                     // id
					"lTRDdPA3ar/n9goTQeqs4olmccmRaJyY8JPo5z3k", // secret
					""),
			})

			fileName, _ := utils.UploadFileToS3(s, src, file)

			image.Url = "https://karen-givi-bucket.s3.ap-southeast-1.amazonaws.com/" + fileName

			log.Info(image.Url)

		}

		// log.Info(fileName)
		// user := UserCreateRequest{}
		// image := entities.Image{}

		// imageArrInput := []imageRepo.ImageInput{}

		// imageArrInput = append(imageArrInput, imageRepo.ImageInput{Url: image.Url})

		// imageReq := imageRepo.ImageReq{Array: imageArrInput}

		// err1 := ic.repo.Create(image.Room_uid, imageReq)

		// if err != nil {
		// 	return c.JSON(http.StatusInternalServerError, templates.InternalServerError(nil, "error internal server error fo create new image", err1))
		// }

		return c.JSON(http.StatusCreated, templates.Success(http.StatusCreated, "Success create new image", nil))
	}
}
