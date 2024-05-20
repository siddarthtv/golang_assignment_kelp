package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/siddarthtv/golang_assignment_kelp/internal/services"
)

func InitServer() {
	service := services.InitService()
	router := gin.Default()
	router.GET("/api/financials", service.FinancialsHandler)
	router.GET("/api/sales", service.SalesHandler)
	router.GET("/api/employee", service.StatsHandler)
	router.Run()
}
