package user

// import (
// 	"be/delivery/controllers/auth"
// 	"be/entities"
// 	"bytes"
// 	"encoding/json"
// 	"errors"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/labstack/echo/v4"
// 	"github.com/labstack/gommon/log"
// 	"github.com/stretchr/testify/assert"
// 	"gorm.io/gorm"
// )

// type MockAuthLib struct{}

// func (m *MockAuthLib) Login(UserLogin auth.Userlogin) (entities.User, error) {
// 	return entities.User{Model: gorm.Model{ID: 1}, Email: UserLogin.Email, Password: UserLogin.Password}, nil
// }

// type MockUserLib struct{}

// func (m *MockUserLib) Create(user entities.User) (entities.User, error) {
// 	return entities.User{}, nil
// }

// type MockFailUserLib struct{}

// func (m *MockFailUserLib) Create(user entities.User) (entities.User, error) {
// 	return entities.User{}, errors.New("")
// }

// func TestCreate(t *testing.T) {
// 	t.Run("BadRequest", func(t *testing.T) {
// 		e := echo.New()

// 		reqBody, _ := json.Marshal(map[string]interface{}{
// 			"name":     1,
// 			"email":    "",
// 			"password": "anonim123",
// 		})

// 		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
// 		res := httptest.NewRecorder()
// 		req.Header.Set("Content-Type", "application/json")

// 		context := e.NewContext(req, res)
// 		context.SetPath("/users")

// 		userController := New(&MockUserLib{})
// 		userController.Create()(context)

// 		response := GetUserResponseFormat{}

// 		json.Unmarshal([]byte(res.Body.Bytes()), &response)

// 		assert.Equal(t, 400, response.Code)
// 		assert.Equal(t, "error in request for create new user", response.Message)
// 	})

// 	t.Run("InternalServerError", func(t *testing.T) {
// 		e := echo.New()

// 		reqBody, _ := json.Marshal(map[string]string{
// 			"name":     "",
// 			"email":    "anonim@123",
// 			"password": "anonim123",
// 		})
// 		log.Info(reqBody)
// 		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
// 		res := httptest.NewRecorder()
// 		req.Header.Set("Content-Type", "application/json")

// 		context := e.NewContext(req, res)
// 		context.SetPath("/users")
// 		log.Info(req)
// 		log.Info(context)
// 		userController := New(&MockFailUserLib{})
// 		userController.Create()(context)

// 		response := GetUserResponseFormat{}

// 		json.Unmarshal([]byte(res.Body.Bytes()), &response)

// 		assert.Equal(t, 500, response.Code)
// 		assert.Equal(t, "error internal server error fo create new user", response.Message)
// 	})

// 	t.Run("success", func(t *testing.T) {
// 		e := echo.New()

// 		reqBody, _ := json.Marshal(map[string]string{
// 			"name":     "anonim123",
// 			"email":    "anonim@123",
// 			"password": "anonim123",
// 		})

// 		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
// 		res := httptest.NewRecorder()
// 		req.Header.Set("Content-Type", "application/json")

// 		context := e.NewContext(req, res)
// 		context.SetPath("/users")

// 		userController := New(&MockUserLib{})
// 		userController.Create()(context)

// 		response := GetUserResponseFormat{}

// 		json.Unmarshal([]byte(res.Body.Bytes()), &response)

// 		assert.Equal(t, 201, response.Code)
// 		assert.Equal(t, "Success create new user", response.Message)
// 	})
// }
