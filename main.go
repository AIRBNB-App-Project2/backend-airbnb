package main

import (
	"be/configs"
	"be/delivery/controllers/auth"
	"be/delivery/controllers/image"
	"be/delivery/controllers/room"
	"be/delivery/controllers/user"
	"be/delivery/routes"
	"fmt"

	authLib "be/repository/database/auth"
	imageLib "be/repository/database/image"
	RoomRepo "be/repository/database/room"

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

	roomRepo := RoomRepo.New(db)
	roomController := room.New(roomRepo)

	e := echo.New()

	routes.RoutesPath(e, userController, authController, imageController, roomController)
	log.Fatal(e.Start(fmt.Sprintf(":%d", config.Port)))
}
