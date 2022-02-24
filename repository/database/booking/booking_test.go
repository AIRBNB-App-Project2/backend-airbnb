// func TestGetAll(t *testing.T) {

// config := configs.GetConfig()
// db := utils.InitDB(config)
// repo := New(db)
// db.Migrator().DropTable(&entities.User{})
// db.Migrator().DropTable(&entities.Room{})
// db.Migrator().DropTable(&entities.Image{})
// db.Migrator().DropTable(&entities.Booking{})
// db.AutoMigrate(&entities.Room{})

// t.Run("success run get all", func(t *testing.T) {

//mock User
// mockUser1 := entities.User{Name: "user1 name", Email: "user1 email", Password: "user1 password"}
// resu1, err1 := user.New(db).Create(mockUser1)
// if err1 != nil {
// 	t.Fatal()
// }
// mockUser2 := entities.User{Name: "user2 name", Email: "user2 email", Password: "user1 password"}
// resu2, err2 := user.New(db).Create(mockUser2)
// if err2 != nil {
// 	t.Fatal()
// }
// mockUser3 := entities.User{Name: "user3 name", Email: "user3 email", Password: "user1 password"}
// _, err3 := user.New(db).Create(mockUser3)
// if err3 != nil {
// 	t.Fatal()
// }
//==================

// city := "1"
// var category string = ""
// var name string = ""
// var length string = "10"
// var s string = "xo"

// var status string = ""
// mockroom1 := entities.Room{User_uid: resu1.User_uid, City_id: 1, Name: "room1 name", Price: 100, Detail: "room1 detail"}
// _, errroom1 := repo.Create(mockroom1)
// if errroom1 != nil {
// 	t.Fatal()
// }
// mockroom2 := entities.Room{User_uid: resu2.User_uid, City_id: 1, Name: "roxoom2 name", Price: 100, Detail: "room2 detail"}
// _, errroom2 := repo.Create(mockroom2)
// if errroom2 != nil {
// 	t.Fatal()
// }

// mockroom3 := entities.Room{User_uid: resu2.User_uid, City_id: 2, Name: "roxoxom3 name", Price: 100, Detail: "room1 detail"}
// _, errroom3 := repo.Create(mockroom3)
// if errroom3 != nil {
// 	t.Fatal()
// }
// mockroom4 := entities.Room{User_uid: resu2.User_uid, City_id: 2, Name: "room3 name", Price: 100, Detail: "room1 detail"}
// _, errroom4 := repo.Create(mockroom4)
// if errroom4 != nil {
// 	t.Fatal()
// }
// mockroom5 := entities.Room{User_uid: resu2.User_uid, City_id: 3, Name: "room3 name", Price: 100, Detail: "room1 detail"}
// _, errroom5 := repo.Create(mockroom5)
// if errroom5 != nil {
// 	t.Fatal()
// }

// res, _ := repo.GetAll(s, city, category, name, length, status)
// log.Info(res)

// // assert.Equal(t, "0", res[0].ID)
// assert.Equal(t, "roxoom2 name", res[0].Name)
// assert.Equal(t, res)
// log.Info(res)