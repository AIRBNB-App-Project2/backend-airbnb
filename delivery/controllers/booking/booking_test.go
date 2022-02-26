package booking

import (
	"be/configs"
	"be/delivery/controllers/auth"
	"be/entities"
	"be/repository/database/booking"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

type MockAuthLib struct{}

func (m *MockAuthLib) Login(UserLogin entities.User) (entities.User, error) {
	return entities.User{Model: gorm.Model{ID: 1}, Email: UserLogin.Email, Password: UserLogin.Password}, nil
}

type MockBookingRepo struct{}

func (m *MockBookingRepo) Create(user_uid string, room_uid string, newBooking booking.BookingReq) (booking.BookingCreateResp, error) {
	return booking.BookingCreateResp{}, nil
}

func (m *MockBookingRepo) Update(user_uid string, booking_uid string, upBooking booking.BookingReq) (booking.BookingCreateResp, error) {
	return booking.BookingCreateResp{}, nil
}

func (m *MockBookingRepo) GetById(booking_uid string) (booking.BookingGetByIdResp, error) {
	return booking.BookingGetByIdResp{}, nil
}

func (m *MockBookingRepo) Delete(booking_uid string) (entities.Booking, error) {
	return entities.Booking{}, nil
}

func TestCreate(t *testing.T) {

	jwtToken := ""
	t.Run("Test Login", func(t *testing.T) {
		e := echo.New()

		requestBody, _ := json.Marshal(map[string]string{
			"email":    "test@gmail.com",
			"password": "xyz",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/login")

		authControl := auth.New(&MockAuthLib{})
		authControl.Login()(context)

		responses := auth.LoginRespFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		jwtToken = responses.Data["token"].(string)
		fmt.Println(jwtToken)
		assert.Equal(t, responses.Message, "success login")
	})

	t.Run("success Create", func(t *testing.T) {
		e := echo.New()

		reqBody, _ := json.Marshal(map[string]interface{}{

			"user_uid":   "user_uid",
			"room_uid":   "room_uid",
			"start_date": "01 Mar 2022",
			"end_date":   "03 Mar 2022",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/booking")

		taskController := New(&MockBookingRepo{})
		// taskController.GetById()(context)
		if err := middleware.JWT([]byte(configs.JWT_SECRET))(taskController.Create())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := GetBookingResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, 201, response.Code)

	})

}
