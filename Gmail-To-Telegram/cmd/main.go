package main

import (
	"fmt"
	Config "gomod/internal/config"
	Logger "gomod/internal/entities"
)

func main() {
	fmt.Println("Application running . . . ")
	Logger.LoggerIni()
	Logger.Log("application start")
	//We get the configuration structure from the file
	var config Config.Config
	err := Config.GetConfig(&config)
	if err != nil {
		panic(err)
	}

	fmt.Println(config)
}
