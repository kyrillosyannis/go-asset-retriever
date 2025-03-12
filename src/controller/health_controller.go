package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type HealthController struct{
	db *gorm.DB
}

func NewHealthController(db *gorm.DB) *HealthController {
	return &HealthController{db: db}
}

func (controller *HealthController) Health(context *gin.Context) {
	sqlDB, _ := controller.db.DB()
	err := sqlDB.Ping()
	if err != nil {
		context.JSON(500, gin.H{"status": "database_down"})
		return
	}
	context.JSON(200, gin.H{"status": "up"})
}
