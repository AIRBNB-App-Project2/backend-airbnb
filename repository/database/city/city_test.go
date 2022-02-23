package city

import (
	"be/configs"
	"be/utils"
	"testing"

	"github.com/labstack/gommon/log"
	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)

	t.Run("success run GetAll", func(t *testing.T) {
		res, err := repo.GetAll()

		assert.Nil(t, err)
		assert.NotNil(t, res)
		log.Info(res)
	})
}
