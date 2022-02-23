package user

import (
	"be/configs"
	"be/delivery/controllers/auth"
	"be/entities"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	m "github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

type MockAuthLib struct{}

func (m *MockAuthLib) Login(UserLogin entities.User) (entities.User, error) {
	return entities.User{Model: gorm.Model{ID: 1}, Email: UserLogin.Email, Password: UserLogin.Password}, nil
}

type MockUserLib struct{}

func (m *MockUserLib) Create(user entities.User) (entities.User, error) {
	return entities.User{}, nil
}

func (m *MockUserLib) GetById(userUid string) (entities.User, error) {
	return entities.User{}, nil
}
func (m *MockUserLib) Update(userUid string, user entities.User) (entities.User, error) {
	return entities.User{}, nil
}

func (m *MockUserLib) Delete(userUid string) error {
	return nil
}

///MockFailUserLib
type MockFailUserLib struct{}

func (m *MockFailUserLib) Create(user entities.User) (entities.User, error) {
	return entities.User{}, errors.New("")
}

func (m *MockFailUserLib) GetById(userUid string) (entities.User, error) {
	return entities.User{}, errors.New("")
}

func (m *MockFailUserLib) Update(userUid string, user entities.User) (entities.User, error) {
	return entities.User{}, errors.New("")
}
func (m *MockFailUserLib) Delete(userUid string) error {
	return errors.New("")
}

///END MOCK USER =======================================================

func TestCreate(t *testing.T) {
	t.Run("BadRequest", func(t *testing.T) {
		e := echo.New()

		reqBody, _ := json.Marshal(map[string]interface{}{
			"name":     1,
			"email":    "",
			"password": "anonim123",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/users")

		userController := New(&MockUserLib{})
		userController.Create()(context)

		response := GetUserResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, 400, response.Code)
		assert.Equal(t, "There is some problem from input", response.Message)
	})

	t.Run("InternalServerError", func(t *testing.T) {
		e := echo.New()

		reqBody, _ := json.Marshal(map[string]interface{}{
			"name":     "12345",
			"email":    "anonim@gmail.com",
			"password": "anonim123",
		})
		log.Info(reqBody)
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/users")
		log.Info(req)
		log.Info(context)
		userController := New(&MockFailUserLib{})
		userController.Create()(context)

		response := GetUserResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, 500, response.Code)
		assert.Equal(t, "error internal server error fo create new user", response.Message)
	})

	t.Run("success", func(t *testing.T) {
		e := echo.New()

		reqBody, _ := json.Marshal(map[string]string{
			"name":     "anonim123",
			"email":    "anonim@gmail.com",
			"password": "anonim123",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/users")

		userController := New(&MockUserLib{})
		userController.Create()(context)

		response := GetUserResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, 201, response.Code)
		assert.Equal(t, "Success create new user", response.Message)
	})
}

func TestGetById(t *testing.T) {

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

		jwtToken = responses.Data["Token"].(string)
		fmt.Println(jwtToken)
		assert.Equal(t, responses.Message, "success login")
	})

	t.Run("GetById", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/tasks/:uid")

		taskController := New(&MockUserLib{})
		// taskController.GetById()(context)
		if err := m.JWT([]byte(configs.JWT_SECRET))(taskController.GetById())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := GetUserResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, 200, response.Code)
		assert.Equal(t, "Success Get User", response.Message)

	})
	t.Run("ErorGetById", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		context := e.NewContext(req, res)
		context.SetPath("/users/:id")

		falseUserController := New(&MockFailUserLib{})
		falseUserController.GetById()(context)

		var response GetUserResponseFormat

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, response.Message, "not found")
	})

}

// func TestUserRegister(t *testing.T) {
// 	t.Run("UserRegister", func(t *testing.T) {
// 		e := echo.New()
// 		requestBody, _ := json.Marshal(map[string]interface{}{
// 			"nama":     "Adlan",
// 			"email":    "adlan@adlan.com",
// 			"password": "adlan123",
// 		})
// 		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
// 		res := httptest.NewRecorder()
// 		req.Header.Set("Content-Type", "application/json")
// 		context := e.NewContext(req, res)
// 		context.SetPath("/users")

// 		userController := New(MockUserRepository{})

// 		userController.UserRegister()(context)

// 		response := RegisterUserResponseFormat{}

// 		json.Unmarshal([]byte(res.Body.Bytes()), &response)
// 		// assert.Equal(t, 201, response.Code)
// 		assert.Equal(t, "Adlan", response.Data.Nama)
// 		assert.Equal(t, http.StatusCreated, response.Code)

// 	})
// 	t.Run("ErorUserRegister", func(t *testing.T) {
// 		e := echo.New()
// 		req := httptest.NewRequest(http.MethodPost, "/", nil)
// 		res := httptest.NewRecorder()
// 		context := e.NewContext(req, res)
// 		context.SetPath("/users")

// 		userController := New(&MockFalseUserRepository{})
// 		userController.UserRegister()(context)

// 		response := RegisterUserResponseFormat{}

// 		json.Unmarshal([]byte(res.Body.Bytes()), &response)

// 		assert.Equal(t, 500, response.Code)
// 		assert.Equal(t, "There is some error on server", response.Message)

// 	})

// 	t.Run("UserRegisterBind", func(t *testing.T) {
// 		e := echo.New()
// 		requestBody, _ := json.Marshal(map[string]interface{}{
// 			"nama":     "Adlan",
// 			"email":    "adlan@adlan.com",
// 			"password": 1,
// 		})
// 		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
// 		res := httptest.NewRecorder()
// 		req.Header.Set("Content-Type", "application/json")
// 		context := e.NewContext(req, res)
// 		context.SetPath("/users")

// 		userController := New(MockUserRepository{})
// 		userController.UserRegister()(context)

// 		response := RegisterUserResponseFormat{}

// 		json.Unmarshal([]byte(res.Body.Bytes()), &response)

// 		assert.Equal(t, 400, response.Code)

// 	})

// }

// func TestUpdate(t *testing.T) {
// 	t.Run("Update", func(t *testing.T) {
// 		e := echo.New()
// 		requestBody, _ := json.Marshal(map[string]interface{}{
// 			"nama":     "Adlan",
// 			"email":    "adlan@adlan.com",
// 			"password": "adlan123",
// 		})
// 		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))

// 		res := httptest.NewRecorder()
// 		req.Header.Set("Content-Type", "application/json")
// 		context := e.NewContext(req, res)
// 		context.SetPath("/users/:id")

// 		userController := New(&MockUserRepository{})
// 		userController.Update()(context)

// 		response := UpdateResponseFormat{}

// 		json.Unmarshal([]byte(res.Body.Bytes()), &response)

// 		assert.Equal(t, 200, response.Code)
// 		assert.Equal(t, "Adlan", response.Data.Nama)

// 	})

// 	t.Run("ErrorUpdate", func(t *testing.T) {
// 		e := echo.New()
// 		req := httptest.NewRequest(http.MethodPut, "/", nil)
// 		res := httptest.NewRecorder()
// 		context := e.NewContext(req, res)
// 		context.SetPath("/users/:id")

// 		userController := New(&MockFalseUserRepository{})
// 		userController.Update()(context)

// 		response := UpdateResponseFormat{}

// 		json.Unmarshal([]byte(res.Body.Bytes()), &response)

// 		assert.Equal(t, 500, response.Code)
// 		assert.Equal(t, "There is some error on server", response.Message)

// 	})
// 	t.Run("UpdateBind", func(t *testing.T) {
// 		e := echo.New()
// 		requestBody, _ := json.Marshal(map[string]interface{}{
// 			"nama":     "Adlan",
// 			"email":    "adlan@adlan.com",
// 			"password": 123,
// 		})
// 		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))

// 		res := httptest.NewRecorder()
// 		req.Header.Set("Content-Type", "application/json")
// 		context := e.NewContext(req, res)
// 		context.SetPath("/users/:id")

// 		userController := New(&MockUserRepository{})
// 		userController.Update()(context)

// 		response := UpdateResponseFormat{}

// 		json.Unmarshal([]byte(res.Body.Bytes()), &response)

// 		assert.Equal(t, 400, response.Code)

// 	})
// }

// func TestDelete(t *testing.T) {
// 	t.Run("DeleteUser", func(t *testing.T) {
// 		e := echo.New()
// 		req := httptest.NewRequest(http.MethodDelete, "/", nil)
// 		res := httptest.NewRecorder()
// 		context := e.NewContext(req, res)
// 		context.SetPath("/users/:id")

// 		userController := New(&MockUserRepository{})
// 		userController.Delete()(context)

// 		response := DeleteResponseFormat{}

// 		json.Unmarshal([]byte(res.Body.Bytes()), &response)

// 		assert.Equal(t, 200, response.Code)
// 		assert.Equal(t, nil, response.Data)

// 	})

// 	t.Run("ErrorDeleteUser", func(t *testing.T) {
// 		e := echo.New()
// 		req := httptest.NewRequest(http.MethodDelete, "/", nil)
// 		res := httptest.NewRecorder()
// 		context := e.NewContext(req, res)
// 		context.SetPath("/users/:id")

// 		userController := New(&MockFalseUserRepository{})
// 		userController.Delete()(context)

// 		response := DeleteResponseFormat{}

// 		json.Unmarshal([]byte(res.Body.Bytes()), &response)

// 		assert.Equal(t, 500, response.Code)
// 		assert.Equal(t, "There is some error on server", response.Message)

// 	})
// }

//====================

// func TestLogin(t *testing.T) {
// 	t.Run("UserLogin", func(t *testing.T) {
// 		e := echo.New()

// 		requestBody, _ := json.Marshal(map[string]interface{}{
// 			"email":    "adlan@adlan.com",
// 			"password": "adlan123",
// 		})
// 		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
// 		res := httptest.NewRecorder()
// 		req.Header.Set("Content-Type", "application/json")

// 		context := e.NewContext(req, res)
// 		context.SetPath("/users/login")

// 		userController := New(MockUserRepository{})
// 		userController.Login()(context)

// 		response := UserLoginResponseFormat{}
// 		json.Unmarshal([]byte(res.Body.Bytes()), &response)
// 		fmt.Println(response)

// 		assert.Equal(t, 200, response.Code)
// 		assert.Equal(t, "adlan@adlan.com", response.Data.Email)

// 	})

// 	t.Run("ErrorLogin", func(t *testing.T) {
// 		e := echo.New()
// 		req := httptest.NewRequest(http.MethodPost, "/", nil)
// 		res := httptest.NewRecorder()
// 		context := e.NewContext(req, res)
// 		context.SetPath("/users/login")

// 		userController := New(&MockFalseUserRepository{})
// 		userController.Login()(context)

// 		response := UserLoginResponseFormat{}

// 		json.Unmarshal([]byte(res.Body.Bytes()), &response)

// 		assert.Equal(t, 400, response.Code)
// 		assert.Equal(t, "There is some problem from input", response.Message)

// 	})

// 	t.Run("UserLoginBind", func(t *testing.T) {
// 		e := echo.New()

// 		requestBody, _ := json.Marshal(map[string]interface{}{
// 			"email":    "adlan@adlan.com",
// 			"password": 123,
// 		})
// 		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
// 		res := httptest.NewRecorder()
// 		req.Header.Set("Content-Type", "application/json")

// 		context := e.NewContext(req, res)
// 		context.SetPath("/users/login")

// 		userController := New(MockUserRepository{})
// 		userController.Login()(context)

// 		response := UserLoginResponseFormat{}

// 		json.Unmarshal([]byte(res.Body.Bytes()), &response)

// 		assert.Equal(t, 400, response.Code)

// 	})
// }
