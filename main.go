package main

import (
	"fmt"
	"github.com/spf13/viper"
	"hwd/hwd"
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
		go hwd.NewWorker()

		tasks := globals.GetTasks()
		nokocore.NoErr(pTasks.Execute(tasks))
		nokocore.NoErr(pTasks.Wait())
		fmt.Println("Done")

		//reporting.NewService()
	}

	task.EntryPoint(app.Main, pTasksHandler)
}
