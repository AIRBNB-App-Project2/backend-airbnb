package booking

import (
	"be/delivery/controllers/templates"
	"be/delivery/middlewares"
	"be/entities"
	"be/repository/database/booking"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
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

		res, err := cont.repo.Create(book.User_uid, book.Room_uid, entities.Booking{Start_date: book.Start_date, End_date: book.End_date, PaymentMethod: book.PaymentMethod})

		if err != nil {
			return c.JSON(http.StatusInternalServerError, templates.InternalServerError(http.StatusInternalServerError, "Room not found", nil))
		}

		return c.JSON(http.StatusOK, templates.Success(http.StatusOK, "Success Get Room", res))
	}
}
func (cont *BookingController) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {

		userUid = middlewares.ExtractTokenId(c)

		res, err := cont.repo.GetAll(userUid)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, templates.InternalServerError(http.StatusInternalServerError, "Your booking is empty", nil))
		}

		return c.JSON(http.StatusOK, templates.Success(http.StatusOK, "Success get all booking", res))
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
