package auth

// import (
// 	"be/delivery/templates"
// 	"be/models"
// 	"bytes"
// 	"encoding/json"
// 	"errors"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/labstack/echo/v4"
// 	"github.com/stretchr/testify/assert"
// 	"gorm.io/gorm"
// )

// type MockAuthLib struct{}

// func (m *MockAuthLib) Login(UserLogin templates.Userlogin) (models.User, error) {
// 	return models.User{Model: gorm.Model{ID: 1}, Email: UserLogin.Email, Password: UserLogin.Password}, nil
// }

// type MockFailAuthLib struct{}

// func (m *MockFailAuthLib) Login(UserLogin templates.Userlogin) (models.User, error) {
// 	return models.User{}, errors.New("")
// }

// type MockAuthLibFailToken struct{}

// func (m *MockAuthLibFailToken) Login(UserLogin templates.Userlogin) (models.User, error) {
// 	return models.User{}, nil
// }

// func TestLogin(t *testing.T) {
// 	t.Run("error in request for login user", func(t *testing.T) {
// 		e := echo.New()

// 		reqBody, _ := json.Marshal(map[string]string{
// 			"email": "anonim@123",
// 		})

// 		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
// 		res := httptest.NewRecorder()
// 		req.Header.Set("Content-Type", "application/json")

// 		context := e.NewContext(req, res)
// 		context.SetPath("/login")

// 		authCont := New(&MockAuthLib{})
// 		authCont.Login()(context)

// 		resp := templates.LoginRespFormat{}

// 		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
// 		assert.Equal(t, 400, resp.Code)
// 		assert.Equal(t, "error in request for login user", resp.Message)
// 	})

// 	t.Run("error internal server error for login user", func(t *testing.T) {
// 		e := echo.New()

// 		reqBody, _ := json.Marshal(map[string]string{
// 			"email":    "anonim@123",
// 			"password": "anonim123",
// 		})

// 		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
// 		res := httptest.NewRecorder()
// 		req.Header.Set("Content-Type", "application/json")

// 		context := e.NewContext(req, res)
// 		context.SetPath("/login")

// 		authCont := New(&MockFailAuthLib{})
// 		authCont.Login()(context)

// 		resp := templates.LoginRespFormat{}

// 		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
// 		assert.Equal(t, 500, resp.Code)
// 		assert.Equal(t, "error internal server error for login user", resp.Message)
// 	})

// 	t.Run("error in process token", func(t *testing.T) {
// 		e := echo.New()

// 		reqBody, _ := json.Marshal(map[string]string{
// 			"email":    "anonim@123",
// 			"password": "anonim123",
// 		})

// 		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
// 		res := httptest.NewRecorder()
// 		req.Header.Set("Content-Type", "application/json")

// 		context := e.NewContext(req, res)
// 		context.SetPath("/login")

// 		authCont := New(&MockAuthLibFailToken{})
// 		authCont.Login()(context)

// 		resp := templates.LoginRespFormat{}

// 		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
// 		assert.Equal(t, 406, resp.Code)
// 		assert.Equal(t, "error in process token", resp.Message)
// 	})

// 	t.Run("success login", func(t *testing.T) {
// 		e := echo.New()

// 		reqBody, _ := json.Marshal(map[string]string{
// 			"email":    "anonim@123",
// 			"password": "anonim123",
// 		})

// 		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
// 		res := httptest.NewRecorder()
// 		req.Header.Set("Content-Type", "application/json")

// 		context := e.NewContext(req, res)
// 		context.SetPath("/login")

// 		authCont := New(&MockAuthLib{})
// 		authCont.Login()(context)

// 		resp := templates.LoginRespFormat{}

// 		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
// 		assert.Equal(t, 200, resp.Code)
// 		assert.Equal(t, "success login", resp.Message)
// 	})

// }
