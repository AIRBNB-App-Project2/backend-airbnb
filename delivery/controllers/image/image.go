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

		image := CreateImageRequesFormat{}
		image.Room_uid = c.FormValue("room_uid")

		file, err := c.FormFile("file")
		if err != nil {
			return err
		}
		src, err := file.Open()
		if err != nil {
			return err
		}
		log.Info(src)

		defer src.Close()

		s, _ := session.NewSession(&aws.Config{
			Region: aws.String("ap-southeast-1"),
			Credentials: credentials.NewStaticCredentials(
				"AKIAS4KA3W5H4Z73S3NR",                     // id
				"XVGjvN4ApOPqNFH95wfmpM06PpQfqiXdDhGuBcFp", // secret
				""),
		})

		fileName, _ := utils.UploadFileToS3(s, src, file)

		// log.Info(fileName)
		// user := UserCreateRequest{}
		// image := entities.Image{}
		image.Url = "https://test-upload-s3-rogerdev.s3.ap-southeast-1.amazonaws.com/" + fileName

		imageArrInput := []imageRepo.ImageInput{}

		imageArrInput = append(imageArrInput, imageRepo.ImageInput{Url: image.Url})

		imageReq := imageRepo.ImageReq{Array: imageArrInput}

		err1 := ic.repo.Create(image.Room_uid, imageReq)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, templates.InternalServerError(nil, "error internal server error fo create new image", err1))
		}

		return c.JSON(http.StatusCreated, templates.Success(http.StatusCreated, "Success create new image", nil))
	}
}
