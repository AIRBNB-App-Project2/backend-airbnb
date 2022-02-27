package booking

import (
	"be/delivery/controllers/templates"
	"be/delivery/middlewares"
	"be/repository/database/booking"
	"net/http"
	"time"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
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
			return c.JSON(http.StatusBadRequest, templates.BadRequest(nil, "There is some problem from input", err))
		}

		book.User_uid = middlewares.ExtractTokenUserUid(c)
		// log.Info(book)
		//parse string tu date time.Time
		layoutFormat := "02 Jan 2006"
		start_date, _ := time.Parse(layoutFormat, book.Start_date)
		end_date, _ := time.Parse(layoutFormat, book.End_date)
		log.Info(start_date, end_date)

		res, err := cont.repo.Create(book.User_uid, book.Room_uid, booking.BookingReq{Start_date: start_date.String(), End_date: end_date.String()})

		if err != nil {
			return c.JSON(http.StatusInternalServerError, templates.InternalServerError(http.StatusInternalServerError, "Failed to add booking "+err.Error(), nil))
		}

		return c.JSON(http.StatusCreated, templates.Success(http.StatusCreated, "Success add booking", res))
	}
}

func (cont *BookingController) GetById() echo.HandlerFunc {
	return func(c echo.Context) error {
		booking_uid := c.Param("booking_uid")

		res, err := cont.repo.GetById(booking_uid)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, templates.InternalServerError(http.StatusInternalServerError, "internal server eror for get booking by id "+err.Error(), nil))
		}

		return c.JSON(http.StatusOK, templates.Success(http.StatusOK, "Success get booking", res))
	}
}

func (cont *BookingController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		booking_uid := c.Param("booking_uid")
		v := validator.New()
		var book CreateBookingRequesFormat

		if err := c.Bind(&book); err != nil {
			return c.JSON(http.StatusBadRequest, templates.BadRequest(nil, "There is some problem from input", err))
		}

		if err := v.Struct(book); err != nil {
			return c.JSON(http.StatusBadRequest, templates.BadRequest(nil, "There is some problem from input", nil))
		}

		book.User_uid = middlewares.ExtractTokenUserUid(c)

		//parse string tu date time.Time
		layoutFormat := "02 Jan 2006"
		start_date, _ := time.Parse(layoutFormat, book.Start_date)
		end_date, _ := time.Parse(layoutFormat, book.End_date)

		res, err := cont.repo.Update(book.User_uid, booking_uid, booking.BookingReq{Start_date: start_date.String(), End_date: end_date.String(), Status: book.Status, PaymentMethod: book.PaymentMethod})

		if err != nil {
			return c.JSON(http.StatusInternalServerError, templates.InternalServerError(http.StatusInternalServerError, "error internal server for update booking "+err.Error(), nil))
		}

		return c.JSON(http.StatusAccepted, templates.Success(http.StatusAccepted, "Success update booking", res))
	}
}

func (cont *BookingController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		booking_uid := c.Param("booking_uid")
		// user_uid := middlewares.ExtractTokenId(c)

		res, err := cont.repo.Delete(booking_uid)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, templates.InternalServerError(http.StatusInternalServerError, "error internal server for delete boooking "+err.Error(), nil))
		}

		return c.JSON(http.StatusOK, templates.Success(http.StatusOK, "Success delete booking", res))
	}
}
