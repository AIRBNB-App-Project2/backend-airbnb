package booking

import (
	"be/delivery/controllers/templates"
	"be/delivery/middlewares"
	"be/repository/database/booking"
	"be/utils"
	"net/http"
	"time"

	"github.com/midtrans/midtrans-go"

	"github.com/midtrans/midtrans-go/coreapi"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type BookingController struct {
	repo booking.Booking
	mt   coreapi.Client
}

func New(repo booking.Booking, mt coreapi.Client) *BookingController {
	return &BookingController{
		repo: repo,
		mt:   mt,
	}
}

func (cont *BookingController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		var book CreateBookingRequesFormat

		if err := c.Bind(&book); err != nil {
			return c.JSON(http.StatusBadRequest, templates.BadRequest(nil, "There is some problem from input", err))
		}

		book.User_uid = middlewares.ExtractTokenUserUid(c)
		// log.Info(book)
		//parse string tu date time.Time
		layoutFormat := "2006-01-02"
		start_date, err1 := time.Parse(layoutFormat, book.Start_date)
		if err1 != nil {
			return c.JSON(http.StatusBadRequest, templates.BadRequest(nil, "There is some problem from parsing start date", err1))
		}
		end_date, err2 := time.Parse(layoutFormat, book.End_date)
		// log.Info(start_date, end_date)
		if err2 != nil {
			return c.JSON(http.StatusBadRequest, templates.BadRequest(nil, "There is some problem from parsing end date", err2))
		}
		// log.Info(start_date,end_date)
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
		// v := validator.New()
		var book CreateBookingRequesFormat

		if err := c.Bind(&book); err != nil {
			return c.JSON(http.StatusBadRequest, templates.BadRequest(nil, "There is some problem from input", err))
		}

		book.User_uid = middlewares.ExtractTokenUserUid(c)
		var start_date, end_date time.Time
		if book.Status == "" {
			//parse string tu date time.Time
			layoutFormat := "2006-01-02"
			start_date, err1 := time.Parse(layoutFormat, book.Start_date)
			if err1 != nil {
				return c.JSON(http.StatusBadRequest, templates.BadRequest(nil, "There is some problem from parsing start date", err1))
			}
			end_date, err2 := time.Parse(layoutFormat, book.End_date)
			log.Info(start_date, end_date)
			if err2 != nil {
				return c.JSON(http.StatusBadRequest, templates.BadRequest(nil, "There is some problem from parsing end date", err2))
			}
		}

		// log.Info(start_date,end_date)
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
func (cont *BookingController) CreatePayment() echo.HandlerFunc {
	return func(c echo.Context) error {
		booking_uid := c.Param("booking_uid")
		var payment_method PaymentTypeRequest
		// user := middlewares.ExtractTokenId(c)

		payment_method.Payment_method = "gopay"

		var result *coreapi.ChargeReq

		res_booking, err := cont.repo.GetById(booking_uid)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, templates.InternalServerError(http.StatusInternalServerError, "Your booking is not found", nil))
		}
		if err != nil {
			return c.JSON(http.StatusInternalServerError, templates.InternalServerError(http.StatusInternalServerError, "Your booking is not found", nil))
		}
		switch payment_method.Payment_method {
		case "gopay":
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

		case "shopeepay":
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
		case "qris":
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
		// for i := range apiRes.Actions {
		data.Url = /*  append(data.Url,  */ apiRes.Actions[1].URL /* ) */
		// }

		return c.JSON(http.StatusOK, templates.Success(http.StatusOK, "Success create payment booking", data))

	}
}

func (cont *BookingController) CallBack() echo.HandlerFunc {
	return func(c echo.Context) error {
		var request RequestCallBackMidtrans
		// user_uid := middlewares.ExtractTokenId(c)

		if err := c.Bind(&request); err != nil {
			return c.JSON(http.StatusInternalServerError, templates.InternalServerError(http.StatusInternalServerError, "Failed to create payment", nil))
		}

		res, err := cont.repo.GetByIdMt(request.Order_id)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, templates.InternalServerError(http.StatusInternalServerError, "internal server eror for get booking by id "+err.Error(), nil))
		}

		switch request.Transaction_status {
		case "settlement":
			cont.repo.Update(res.User_uid, request.Order_id, booking.BookingReq{Status: "paid"})
		case "failure":
			cont.repo.Update(res.User_uid, request.Order_id, booking.BookingReq{Status: "waiting"})
		case "cancel":
			cont.repo.Update(res.User_uid, request.Order_id, booking.BookingReq{Status: "waiting"})

		}

		// var strDebug string
		// strDebug = spew.Sdump(request)
		// ZapLogger.Info(`request: ` + strDebug)

		return c.JSON(http.StatusOK, templates.Success(http.StatusOK, "Success create payment booking", request))

	}
}
