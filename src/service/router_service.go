package service

import "github.com/gin-gonic/gin"

type routerService struct {
	Create func() *gin.Engine
}

var Router = routerService{
	Create: func() *gin.Engine {
		router := gin.Default()
		router.GET("/water-bill", waterBill.Get)
		router.POST("/water-bill", waterBill.Post)
		return router
	},
}
