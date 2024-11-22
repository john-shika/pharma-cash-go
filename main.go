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
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	pTasksHandler := func(pTasks task.ProcessTasksImpl) {
		tasks := globals.GetTasksConfig()
		nokocore.NoErr(pTasks.Execute(tasks))
		nokocore.NoErr(pTasks.Wait())
	}

	task.EntryPoint(app.Main, pTasksHandler)
}
