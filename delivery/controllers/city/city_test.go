package city

import (
	"be/entities"
	"be/repository/database/city"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/labstack/echo/v4"
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

type MockFailCityRep struct{}

func (repo *MockFailCityRep) GetAll() ([]city.CityResp, error) {
	return []city.CityResp{}, errors.New("")
}

func TestGetAll(t *testing.T) {

	t.Run("success GetAll", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/city")

		controller := New(&MockCityRep{})

		controller.GetAll()(context)

		resp := GetResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 200, resp.Code)
	})

	t.Run("fail GetAll", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/city")

		controller := New(&MockFailCityRep{})

		controller.GetAll()(context)

		resp := GetResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 500, resp.Code)
	})

}
