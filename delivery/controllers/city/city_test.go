package city

import (
	"be/configs"
	"be/delivery/controllers/auth"
	"be/entities"
	"be/repository/database/city"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type MockAuthLib struct{}

func (m *MockAuthLib) Login(UserLogin entities.User) (entities.User, error) {
	return entities.User{Model: gorm.Model{ID: 1}, Email: UserLogin.Email, Password: UserLogin.Password}, nil
}

type MockCityRep struct{}

func (repo *MockCityRep) GetAll() ([]city.CityResp, error) {
	return []city.CityResp{}, nil
}

func TestGetAll(t *testing.T) {

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

	t.Run("success GetAll", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/city")

		controller := New(&MockCityRep{})

		if err := middleware.JWT([]byte(configs.JWT_SECRET))(controller.GetAll())(context); err != nil {
			log.Warn()
			return
		}

		resp := GetResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 200, resp.Code)
	})

}
