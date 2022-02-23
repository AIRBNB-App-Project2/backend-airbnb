package main

import (
	"be/configs"
	"be/delivery/controllers/auth"
	"be/delivery/controllers/image"
	"be/delivery/controllers/user"
	"be/delivery/routes"
	"fmt"

	authLib "be/repository/database/auth"

	imageLib "be/repository/database/image"
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

	imageRepo := imageLib.New(db)
	imageController := image.New(imageRepo)

	e := echo.New()

	routes.RoutesPath(e, userController, authController)
	log.Fatal(e.Start(fmt.Sprintf(":%d", config.Port)))
}
