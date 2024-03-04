package main

import (
	"github.com/Jonsslol/Go_prj1/config"
	"github.com/Jonsslol/Go_prj1/router"
)

var (
	logger *config.Logger
)

func main() {

	logger = config.Getlogger("main")
	//initialize COnfigs
	err := config.Inti()
	if err != nil {
		logger.Errorf("❌ config initialization error: %v ", err)
		return
	}

	//initialize Router
	router.Initialize()

}
