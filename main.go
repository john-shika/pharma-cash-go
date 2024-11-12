package main

import (
	"fmt"
	"github.com/spf13/viper"
	"nokotan/app"
	"nokowebapi/console"
	"nokowebapi/globals"
	"nokowebapi/nokocore"
	"nokowebapi/task"
	"os"
)

func main() {
	var ok bool
	var err error
	var nokoWebApiSelfRunEnv string
	nokocore.KeepVoid(ok, err, nokoWebApiSelfRunEnv)

	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	viper.SetConfigFile("nokowebapi.yaml")

	if err = viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	
	console.Dir(globals.GetJwtConfig())

	return

	if nokoWebApiSelfRunEnv, ok = os.LookupEnv("NOKOWEBAPI_SELF_RUNNING"); ok {
		nokocore.NoErr(os.Setenv("NOKOWEBAPI_SELF_RUNNING", "1"))
		if nokocore.ParseEnvBool(nokoWebApiSelfRunEnv) {
			nokocore.ApplyMainFunc(app.Main)
			return
		}
	}

	processTasks := task.NewProcessTasks(nil)
	nokocore.NoErr(globals.GetTasks().ExecuteAsync(processTasks).Wait())
	//nokocore.NoErr(globals.GetTasks().Execute(processTasks))
	nokocore.NoErr(processTasks.Wait())

	return

	//var db *gorm.DB
	//cores.KeepVoid(db)
	//
	//e := echo.New()
	//config := &gorm.Config{}
	//
	//cores.NoErr(cores.EnsureDirAndFile("migrations/dev.db"))
	//
	//if db, err = gorm.Open(sqlite.Open("migrations/dev.db"), config); err != nil {
	//	panic("failed to connect database")
	//}
	//
	//if err = db.AutoMigrate(&models.User{}); err != nil {
	//	panic(fmt.Errorf("failed to migrate database: %w", err))
	//}
	//if err = db.AutoMigrate(&models.Session{}); err != nil {
	//	panic(fmt.Errorf("failed to migrate database: %w", err))
	//}
	//if err = db.AutoMigrate(&models.Todo{}); err != nil {
	//	panic(fmt.Errorf("failed to migrate database: %w", err))
	//}
	//
	//hash := sha256.Sum256([]byte("Admin@1234"))
	//password := cores.Base64Encode(hash[:])
	//db.Create(&models.User{
	//	UUID:     cores.NewUuid(),
	//	Email:    cores.NewNullString("admin@localhost"),
	//	Username: "admin",
	//	Password: password,
	//	Role:     "admin,user",
	//})
	//db.Create(&models.Todo{
	//	Title:       "Todo 1",
	//	Description: "Todo 1 description",
	//})
	//
	//router := e.Group("/api/v1")
	//cores.KeepVoid(router)
	//
	//e.HTTPErrorHandler = func(err error, c echo.Context) {
	//	// TODO
	//	fmt.Println(err)
	//	e.DefaultHTTPErrorHandler(err, c)
	//}
	//
	//if err = e.Start(":8080"); err != nil {
	//	e.Logger.Fatal(err)
	//}
}
