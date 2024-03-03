package router

import "github.com/gin-gonic/gin"

func Initialize() {

	//Initializes the Router using the Default functions of the gin
	router := gin.Default()

	//Initialize Routes
	initializeRoutes(router)

	//Run Server
	router.Run(":8080") // listen and serve on 0.0.0.0:8080

}
