package main

import (
	"be/configs"
	"be/delivery/controllers/auth"
	"be/delivery/controllers/user"
	"be/delivery/routes"
	"fmt"

	authLib "be/repository/database/auth"

	userLib "be/repository/database/user"
	"be/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func main() {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	userRepo := userLib.New(db)
	userController := user.New(userRepo)

	authRepo := authLib.New(db)
	authController := auth.New(authRepo)

	e := echo.New()

	routes.RoutesPath(e, userController, authController)
	log.Fatal(e.Start(fmt.Sprintf(":%d", config.Port)))
}
