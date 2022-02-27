package booking

import (
	"be/delivery/controllers/templates"
	"be/delivery/middlewares"
	"be/entities"
	"be/repository/database/booking"
	"be/repository/database/user"
	"be/utils"
	"net/http"
	"time"

	"github.com/midtrans/midtrans-go"

	"github.com/midtrans/midtrans-go/coreapi"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"gorm.io/datatypes"
)

type BookingController struct {
	repo     booking.Booking
	mt       coreapi.Client
	userRepo user.User
}

func New(repo booking.Booking, mt coreapi.Client, users user.User) *BookingController {
	return &BookingController{
		repo:     repo,
		mt:       mt,
		userRepo: users,
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
		// user_uid := middlewares.ExtractTokenId(c)

		res, err := cont.repo.Delete(booking_uid)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, templates.InternalServerError(http.StatusInternalServerError, "Your booking is not found", nil))
		}

		return c.JSON(http.StatusOK, templates.Success(http.StatusOK, "Success delete booking", res))
	}
}
func (cont *BookingController) CreatePayment() echo.HandlerFunc {
	return func(c echo.Context) error {
		v := validator.New()
		booking_uid := c.Param("booking_uid")
		var method_payment_id PaymentTypeRequest
		// user := middlewares.ExtractTokenId(c)

		c.Bind(&method_payment_id)

		if err := v.Struct(method_payment_id); err != nil {
			return c.JSON(http.StatusBadRequest, templates.BadRequest(nil, "There is some problem from input", nil))
		}

		var result *coreapi.ChargeReq

		res_booking, err := cont.repo.GetById(booking_uid)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, templates.InternalServerError(http.StatusInternalServerError, "Your booking is not found", nil))
		}
		if err != nil {
			return c.JSON(http.StatusInternalServerError, templates.InternalServerError(http.StatusInternalServerError, "Your booking is not found", nil))
		}
		switch method_payment_id.Payment_id {
		case "1":
			result = &coreapi.ChargeReq{
				PaymentType: coreapi.PaymentTypeGopay,

				TransactionDetails: midtrans.TransactionDetails{
					OrderID:  booking_uid,
					GrossAmt: int64(res_booking.Price_total),
				},
				Items: &[]midtrans.ItemDetails{
					{
						ID:    booking_uid,
						Name:  res_booking.Name,
						Price: int64(res_booking.Price),
						Qty:   int32(res_booking.Days),
					},
				},
			}

		case "2":
			result = &coreapi.ChargeReq{
				PaymentType: coreapi.PaymentTypeShopeepay,

				TransactionDetails: midtrans.TransactionDetails{
					OrderID:  booking_uid,
					GrossAmt: int64(res_booking.Price_total),
				},
				Items: &[]midtrans.ItemDetails{
					{
						ID:    booking_uid,
						Name:  res_booking.Name,
						Price: int64(res_booking.Price),
						Qty:   int32(res_booking.Days),
					},
				},
				CustomerDetails: &midtrans.CustomerDetails{
					FName: "roger",
					LName: "san",
					Email: "dani@gmail.com",
					Phone: "089876543210",
				},
				ShopeePay: &coreapi.ShopeePayDetails{
					CallbackUrl: "https://plastic-cougar-32.loca.lt/booking/payment/callback",
				},
			}
		case "3":
			result = &coreapi.ChargeReq{
				PaymentType: coreapi.PaymentTypeQris,

				TransactionDetails: midtrans.TransactionDetails{
					OrderID:  booking_uid,
					GrossAmt: int64(res_booking.Price_total),
				},
				Items: &[]midtrans.ItemDetails{
					{
						ID:    booking_uid,
						Name:  res_booking.Name,
						Price: int64(res_booking.Price),
						Qty:   int32(res_booking.Days),
					},
				},
			}

		}

		apiRes, err := utils.CreateTransaction(cont.mt, result)
		// log.Info(apiRes)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, templates.InternalServerError(http.StatusInternalServerError, "Failed to create payment", nil))

		}

		var data PaymentResponse
		data.OrderID = apiRes.OrderID
		data.GrossAmount = apiRes.GrossAmount
		data.PaymentType = apiRes.PaymentType
		for i := range apiRes.Actions {
			data.Url = append(data.Url, apiRes.Actions[i].URL)
		}

		return c.JSON(http.StatusOK, templates.Success(http.StatusOK, "Success create payment booking", data))

	}
}

func (cont *BookingController) CallBack() echo.HandlerFunc {
	return func(c echo.Context) error {
		var request RequestCallBackMidtrans
		user_uid := middlewares.ExtractTokenId(c)

		if err := c.Bind(&request); err != nil {
			return c.JSON(http.StatusInternalServerError, templates.InternalServerError(http.StatusInternalServerError, "Failed to create payment", nil))
		}

		switch request.Transaction_status {
		case "settlement":
			cont.repo.Update(user_uid, request.Order_id, entities.Booking{Status: "paid"})
		case "failure":
			cont.repo.Update(user_uid, request.Order_id, entities.Booking{Status: "paid"})
		case "cancel":
			cont.repo.Update(user_uid, request.Order_id, entities.Booking{Status: "paid"})

		}

		// var strDebug string
		// strDebug = spew.Sdump(request)
		// ZapLogger.Info(`request: ` + strDebug)

		return c.JSON(http.StatusOK, templates.Success(http.StatusOK, "Success create payment booking", request))

	}
}
