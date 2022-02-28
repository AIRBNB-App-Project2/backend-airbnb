package booking

import (
	"be/configs"
	"be/delivery/controllers/auth"
	"be/entities"
	"be/repository/database/booking"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/midtrans/midtrans-go/coreapi"
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

func (m *MockBookingRepo) GetByIdMt(booking_uid string) (entities.Booking, error) {
	return entities.Booking{}, nil
}

type MockFailBookingRepo struct{}

func (m *MockFailBookingRepo) Create(user_uid string, room_uid string, newBooking booking.BookingReq) (booking.BookingCreateResp, error) {
	return booking.BookingCreateResp{}, errors.New("")
}

func (m *MockFailBookingRepo) Update(user_uid string, booking_uid string, upBooking booking.BookingReq) (booking.BookingCreateResp, error) {
	return booking.BookingCreateResp{}, errors.New("")
}

func (m *MockFailBookingRepo) GetById(booking_uid string) (booking.BookingGetByIdResp, error) {
	return booking.BookingGetByIdResp{}, errors.New("")
}

func (m *MockFailBookingRepo) Delete(booking_uid string) (entities.Booking, error) {
	return entities.Booking{}, errors.New("")
}

func (m *MockFailBookingRepo) GetByIdMt(booking_uid string) (entities.Booking, error) {
	return entities.Booking{}, errors.New("")
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
			"start_date": "2022-03-01",
			"end_date":   "2022-03-05",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/booking")

		taskController := New(&MockBookingRepo{}, coreapi.Client{})
		// taskController.GetById()(context)
		if err := middleware.JWT([]byte(configs.JWT_SECRET))(taskController.Create())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := GetBookingResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, 201, response.Code)

	})

	t.Run("bad request", func(t *testing.T) {
		e := echo.New()

		reqBody, _ := json.Marshal(map[string]interface{}{

			"user_uid": 1,
			"room_uid": 1,
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/booking")

		taskController := New(&MockBookingRepo{}, coreapi.Client{})
		// taskController.GetById()(context)
		if err := middleware.JWT([]byte(configs.JWT_SECRET))(taskController.Create())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := GetBookingResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, 400, response.Code)

	})

	t.Run("internal server error", func(t *testing.T) {
		e := echo.New()

		reqBody, _ := json.Marshal(map[string]interface{}{

			"user_uid":   "user_uid",
			"room_uid":   "room_uid",
			"start_date": "2022-03-01",
			"end_date":   "2022-03-05",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/booking")

		taskController := New(&MockFailBookingRepo{}, coreapi.Client{})
		// taskController.GetById()(context)
		if err := middleware.JWT([]byte(configs.JWT_SECRET))(taskController.Create())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := GetBookingResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, 500, response.Code)

	})

}

func TestGetByID(t *testing.T) {
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

	t.Run("success GetById", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(nil))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/booking/:booking_uid")

		taskController := New(&MockBookingRepo{}, coreapi.Client{})
		// taskController.GetById()(context)
		if err := middleware.JWT([]byte(configs.JWT_SECRET))(taskController.GetById())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := GetBookingResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, 200, response.Code)

	})

	t.Run("internal server error", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(nil))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/booking/:booking_uid")

		taskController := New(&MockFailBookingRepo{}, coreapi.Client{})
		// taskController.GetById()(context)
		if err := middleware.JWT([]byte(configs.JWT_SECRET))(taskController.GetById())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := GetBookingResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, 500, response.Code)

	})

}

func TestUpdate(t *testing.T) {

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

	t.Run("success Update", func(t *testing.T) {
		e := echo.New()

		reqBody, _ := json.Marshal(map[string]interface{}{

			"user_uid":   "user_uid",
			"room_uid":   "room_uid",
			"start_date": "2022-03-01",
			"end_date":   "2022-03-05",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/booking/:booking_uid")

		taskController := New(&MockBookingRepo{}, coreapi.Client{})
		// taskController.GetById()(context)
		if err := middleware.JWT([]byte(configs.JWT_SECRET))(taskController.Update())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := GetBookingResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, 202, response.Code)

	})

	t.Run("bad request", func(t *testing.T) {
		e := echo.New()

		reqBody, _ := json.Marshal(map[string]interface{}{

			"user_uid": 1,
			"room_uid": 1,
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/booking/:booking_uid")

		taskController := New(&MockBookingRepo{}, coreapi.Client{})
		// taskController.GetById()(context)
		if err := middleware.JWT([]byte(configs.JWT_SECRET))(taskController.Update())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := GetBookingResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, 400, response.Code)

	})

	t.Run("validator", func(t *testing.T) {
		e := echo.New()

		reqBody, _ := json.Marshal(map[string]interface{}{

			"user_uid":   "user_uid",
			"room_uid":   "room_uid",
			"start_date": "01 Mar 2022",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/booking/:booking_uid")

		taskController := New(&MockBookingRepo{}, coreapi.Client{})
		// taskController.GetById()(context)
		if err := middleware.JWT([]byte(configs.JWT_SECRET))(taskController.Update())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := GetBookingResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, 400, response.Code)

	})

	t.Run("internal server error", func(t *testing.T) {
		e := echo.New()

		reqBody, _ := json.Marshal(map[string]interface{}{

			"user_uid":   "user_uid",
			"room_uid":   "room_uid",
			"start_date": "2022-03-01",
			"end_date":   "2022-03-05",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/booking/:booking_uid")

		taskController := New(&MockFailBookingRepo{}, coreapi.Client{})
		// taskController.GetById()(context)
		if err := middleware.JWT([]byte(configs.JWT_SECRET))(taskController.Update())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := GetBookingResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, 500, response.Code)

	})
}

func TestDelete(t *testing.T) {
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

	t.Run("success Delete", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodDelete, "/", bytes.NewBuffer(nil))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/booking/:booking_uid")

		taskController := New(&MockBookingRepo{}, coreapi.Client{})
		// taskController.GetById()(context)
		if err := middleware.JWT([]byte(configs.JWT_SECRET))(taskController.Delete())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := GetBookingResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, 200, response.Code)

	})

	t.Run("internal server error", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(nil))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/booking/:booking_uid")

		taskController := New(&MockFailBookingRepo{}, coreapi.Client{})
		// taskController.GetById()(context)
		if err := middleware.JWT([]byte(configs.JWT_SECRET))(taskController.Delete())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := GetBookingResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, 500, response.Code)

	})

}
