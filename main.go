package main

import (
	"fmt"
	"github.com/spf13/viper"
	"nokowebapi/globals"
	"nokowebapi/nokocore"
	"nokowebapi/task"
	"pharma-cash-go/app"
)

func main() {
	var err error
	nokocore.KeepVoid(err)

	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	viper.SetConfigFile("nokowebapi.yaml")

	if err = viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("failed to read config, %w", err))
	}

	pTasksHandler := func(pTasks task.ProcessTasksImpl) {
		tasks := globals.GetTasks()
		nokocore.NoErr(pTasks.Execute(tasks))
		nokocore.NoErr(pTasks.Wait())
		fmt.Println("Done")

		//hwd.NewService()
		//reporting.NewService()
	}

	task.EntryPoint(app.Main, pTasksHandler)
}
