package booking

import (
	"be/delivery/controllers/templates"
	"be/delivery/middlewares"
	"be/entities"
	"be/repository/database/booking"
	"net/http"
	"time"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"gorm.io/datatypes"
)

type BookingController struct {
	repo booking.Booking
}

func New(repo booking.Booking) *BookingController {
	return &BookingController{
		repo: repo,
	}
}

func (cont *BookingController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		v := validator.New()
		var book CreateBookingRequesFormat

		if err := c.Bind(&book); err != nil {
			return c.JSON(http.StatusBadRequest, templates.BadRequest(nil, "There is some problem from input", err))
		}

		if err := v.Struct(book); err != nil {
			return c.JSON(http.StatusBadRequest, templates.BadRequest(nil, "There is some problem from input", nil))
		}

		book.User_uid = middlewares.ExtractTokenId(c)

		//parse string tu date time.Time
		layoutFormat := "2006-01-02 15:04:05"
		start_date, _ := time.Parse(layoutFormat, book.Start_date)
		end_date, _ := time.Parse(layoutFormat, book.End_date)

		res, err := cont.repo.Create(book.User_uid, book.Room_uid, entities.Booking{Start_date: datatypes.Date(start_date), End_date: datatypes.Date(end_date)})

		if err != nil {
			return c.JSON(http.StatusInternalServerError, templates.InternalServerError(http.StatusInternalServerError, "Failed to create booking", nil))
		}

		return c.JSON(http.StatusOK, templates.Success(http.StatusOK, "Success create booking", res))
	}
}

func (cont *BookingController) GetById() echo.HandlerFunc {
	return func(c echo.Context) error {
		booking_uid := c.Param("booking_uid")

		res, err := cont.repo.GetById(booking_uid)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, templates.InternalServerError(http.StatusInternalServerError, "Your booking is not found", nil))
		}

		return c.JSON(http.StatusOK, templates.Success(http.StatusOK, "Success get booking", res))
	}
}

func (cont *BookingController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		v := validator.New()
		var book CreateBookingRequesFormat

		if err := c.Bind(&book); err != nil {
			return c.JSON(http.StatusBadRequest, templates.BadRequest(nil, "There is some problem from input", err))
		}

		if err := v.Struct(book); err != nil {
			return c.JSON(http.StatusBadRequest, templates.BadRequest(nil, "There is some problem from input", nil))
		}

		book.User_uid = middlewares.ExtractTokenId(c)

		//parse string tu date time.Time
		layoutFormat := "2006-01-02 15:04:05"
		start_date, _ := time.Parse(layoutFormat, book.Start_date)
		end_date, _ := time.Parse(layoutFormat, book.End_date)

		res, err := cont.repo.Update(book.User_uid, book.Room_uid, entities.Booking{Start_date: datatypes.Date(start_date), End_date: datatypes.Date(end_date)})

		if err != nil {
			return c.JSON(http.StatusInternalServerError, templates.InternalServerError(http.StatusInternalServerError, "Your booking is not found", nil))
		}

		return c.JSON(http.StatusOK, templates.Success(http.StatusOK, "Success update booking", res))
	}
}

func (cont *BookingController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		booking_uid := c.Param("booking_uid")
		user_uid := middlewares.ExtractTokenId(c)

		res, err := cont.repo.Delete(user_uid, booking_uid)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, templates.InternalServerError(http.StatusInternalServerError, "Your booking is not found", nil))
		}

		return c.JSON(http.StatusOK, templates.Success(http.StatusOK, "Success delete booking", res))
	}
}
