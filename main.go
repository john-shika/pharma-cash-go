package main

import (
	"fmt"
<<<<<<< HEAD
=======
	"nokowebapi/globals"
>>>>>>> 1768c40f376afcea7e1c17a8c506cda94716f99c
	"github.com/spf13/viper"
	"nokowebapi/nokocore"
	"nokowebapi/task"
	"pharma-cash-go/app"
<<<<<<< HEAD
	"reporting/reporting"
=======
	// "reporting/reporting"
>>>>>>> 1768c40f376afcea7e1c17a8c506cda94716f99c
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
		//tasks := globals.GetTasks()
		//nokocore.NoErr(pTasks.Execute(tasks))
		//nokocore.NoErr(pTasks.Wait())
		fmt.Println("Done")

		//hwd.NewService()
		reporting.NewService()
	}

	task.EntryPoint(app.Main, pTasksHandler)
}
