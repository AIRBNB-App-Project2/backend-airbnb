package image

import (
	"be/delivery/controllers/templates"
	"be/entities"
	"be/repository/database/image"
	"be/utils"
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type ImageController struct {
	repo image.Image
}

func New(repo image.Image) *ImageController {
	return &ImageController{
		repo: repo,
	}
}

func (ic *ImageController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {

		file, err := c.FormFile("file")
		if err != nil {
			return err
		}
		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		s, err := session.NewSession(&aws.Config{
			Region: aws.String("ap-southeast-1"),
			Credentials: credentials.NewStaticCredentials(
				"AKIAS4KA3W5H4Z73S3NR",                     // id
				"XVGjvN4ApOPqNFH95wfmpM06PpQfqiXdDhGuBcFp", // secret
				""),
		})

		fileName, _ := utils.UploadFileToS3(s, src, file)

		fmt.Println(fileName)
		// user := UserCreateRequest{}
		image := entities.Image{}

		if err := c.Bind(&image); err != nil {
			return c.JSON(http.StatusBadRequest, templates.BadRequest(nil, "There is some problem from input", err))
		}
		v := validator.New()
		if err := v.Struct(image); err != nil {
			return c.JSON(http.StatusBadRequest, templates.BadRequest(nil, "There is some problem from input", nil))
		}

		res, err := ic.repo.Create(entities.Image{})
		if err != nil {
			return c.JSON(http.StatusInternalServerError, templates.InternalServerError(nil, "error internal server error fo create new image", err))
		}

		return c.JSON(http.StatusCreated, templates.Success(http.StatusCreated, "Success create new image", res))
	}
}
