package room

import (
	"be/delivery/controllers/templates"
	"be/delivery/middlewares"
	"be/entities"
	imagerepo "be/repository/database/image"
	"be/repository/database/room"
	"be/utils"
	"net/http"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/go-playground/validator"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type RoomController struct {
	repo   room.Room
	repImg imagerepo.Image
}

func New(repo room.Room, repoImg imagerepo.Image) *RoomController {
	return &RoomController{
		repo:   repo,
		repImg: repoImg,
	}
}

func (cont *RoomController) GetById() echo.HandlerFunc {
	return func(c echo.Context) error {
		room_uid := c.Param("room_uid")

		res, err := cont.repo.GetById(room_uid)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, templates.InternalServerError(http.StatusInternalServerError, "Room not found "+err.Error(), nil))
		}

		return c.JSON(http.StatusOK, templates.Success(http.StatusOK, "Success Get Room", res))
	}
}

func (cont *RoomController) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {

		city := c.QueryParam("city")
		category := c.QueryParam("category")
		length, errL := strconv.Atoi(c.QueryParam("length"))
		// log.Info(city)
		if errL != nil {
			return c.JSON(http.StatusInternalServerError, templates.InternalServerError(http.StatusInternalServerError, "internal server error for converting to int "+errL.Error(), nil))
		}

		res, err := cont.repo.GetAllRoom(length, city, category)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, templates.InternalServerError(http.StatusInternalServerError, "internal server error for get all "+err.Error(), nil))
		}

		return c.JSON(http.StatusOK, templates.Success(http.StatusOK, "Success Get all Room", res))
	}
}
func (cont *RoomController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		v := validator.New()
		var room CreateRoomRequesFormat

		if err := c.Bind(&room); err != nil {
			return c.JSON(http.StatusBadRequest, templates.BadRequest(nil, "There is some problem from input", err))
		}

		if err := v.Struct(room); err != nil {
			return c.JSON(http.StatusBadRequest, templates.BadRequest(nil, "There is some problem from input", nil))
		}

		room.User_uid = middlewares.ExtractTokenUserUid(c)

		res, err := cont.repo.Create(entities.Room{User_uid: room.User_uid, City_id: room.City_id, Address: room.Address, Name: room.Name, Category: room.Category, Status: room.Status, Price: room.Price, Description: room.Description})

		if err != nil {
			return c.JSON(http.StatusInternalServerError, templates.InternalServerError(http.StatusInternalServerError, "error internal server for upload room form "+err.Error(), nil))
		}

		form, err1 := c.MultipartForm()
		if err1 != nil {
			return c.JSON(http.StatusBadRequest, templates.BadRequest(nil, "error in multipart form", nil))
		}
		files := form.File["files"]

		imageArrInput := []imagerepo.ImageInput{}
		for _, file := range files {
			image := entities.Image{}
			src, err1 := file.Open()
			if err1 != nil {
				return c.JSON(http.StatusBadRequest, templates.BadRequest(nil, "error in open file image", nil))
			}
			// log.Info(src)

			s, err2 := session.NewSession(&aws.Config{
				Region: aws.String("ap-southeast-1"),
				Credentials: credentials.NewStaticCredentials(
					"AKIAS4KA3W5H4Z73S3NR",                     // id
					"XVGjvN4ApOPqNFH95wfmpM06PpQfqiXdDhGuBcFp", // secret
					""),
			})

			if err2 != nil {
				return c.JSON(http.StatusInternalServerError, templates.InternalServerError(nil, "error in configuration with s3 aws", nil))
			}

			fileName, err3 := utils.UploadFileToS3(s, src, file)

			if err3 != nil {
				return c.JSON(http.StatusInternalServerError, templates.InternalServerError(nil, "error in upload image to s3 aes", nil))
			}

			image.Url = "https://test-upload-s3-rogerdev.s3.ap-southeast-1.amazonaws.com/" + fileName

			log.Info(image.Url)

			imageArrInput = append(imageArrInput, imagerepo.ImageInput{Url: image.Url})

		}

		imageReq := imagerepo.ImageReq{Array: imageArrInput}

		err4 := cont.repImg.Create(res.Room_uid, imageReq)

		if err4 != nil {
			return c.JSON(http.StatusInternalServerError, templates.InternalServerError(nil, "error in upload image "+err4.Error(), nil))
		}

		return c.JSON(http.StatusOK, templates.Success(http.StatusOK, "Success add Room", res))
	}
}
func (cont *RoomController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		roomParam := c.Param("room_uid")
		var room UpdateRoomRequesFormat

		if err := c.Bind(&room); err != nil {
			return c.JSON(http.StatusBadRequest, templates.BadRequest(nil, "There is some problem from input", err))
		}
		user_uid := middlewares.ExtractTokenUserUid(c)

		res, err := cont.repo.Update(user_uid, roomParam, entities.Room{Address: room.Address, Name: room.Name, Category: room.Category, Status: room.Status, Price: room.Price, Description: room.Description})

		if err != nil {
			return c.JSON(http.StatusInternalServerError, templates.InternalServerError(http.StatusInternalServerError, "error internal server for update room "+err.Error(), nil))
		}

		return c.JSON(http.StatusOK, templates.Success(http.StatusOK, "Success update Room", res))
	}
}

func (cont *RoomController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		room_uid := c.Param("room_uid")

		res, err := cont.repo.Delete(room_uid)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, templates.InternalServerError(http.StatusInternalServerError, "error internal server for delete room "+err.Error(), nil))
		}

		return c.JSON(http.StatusOK, templates.Success(http.StatusOK, "Success delete Room", res))
	}
}
