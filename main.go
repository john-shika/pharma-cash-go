package main

import (
	"fmt"
	"nokowebapi/globals"
	"github.com/spf13/viper"
	"nokowebapi/nokocore"
	"nokowebapi/task"
	"pharma-cash-go/app"
	// "reporting/reporting"
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
		tasks := globals.GetTasksConfig()
		nokocore.NoErr(pTasks.Execute(tasks))
		nokocore.NoErr(pTasks.Wait())
		fmt.Println("Done")

		//hwd.NewService()
		reporting.NewService()
	}

	task.EntryPoint(app.Main, pTasksHandler)
}
