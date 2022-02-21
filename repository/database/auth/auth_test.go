package auth

// func TestLogin(t *testing.T) {
// 	config := configs.GetConfig()
// 	db := utils.InitDB(config)
// 	repo := New(db)
// 	db.Migrator().DropTable(&models.ProductType{})
// 	db.Migrator().DropTable(&models.PaymentMethod{})
// 	db.Migrator().DropTable(&models.User{})
// 	db.Migrator().DropTable(&models.Product{})
// 	db.Migrator().DropTable(&models.Cart{})
// 	db.Migrator().DropTable(&models.Order{})
// 	db.Migrator().DropTable(&models.OrderDetail{})
// 	db.AutoMigrate(&models.User{})

// 	t.Run("success run login", func(t *testing.T) {
// 		mockUser := models.User{Name: "anonim123", Email: "anonim@123", Password: "anonim123"}
// 		_, err := libUser.New(db).Create(mockUser)
// 		if err != nil {
// 			t.Fail()
// 		}
// 		mockLogin := templates.Userlogin{Email: "anonim@123", Password: "anonim123"}
// 		res, err := repo.Login(mockLogin)
// 		assert.Nil(t, err)
// 		assert.Equal(t, "anonim@123", res.Email)
// 		assert.Equal(t, "anonim123", res.Password)
// 	})

// 	t.Run("fail run login", func(t *testing.T) {
// 		mockLogin := templates.Userlogin{Email: "anonim@456", Password: "anonim456"}
// 		_, err := repo.Login(mockLogin)
// 		assert.NotNil(t, err)
// 	})

// }
