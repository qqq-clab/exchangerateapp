package controllers

import (
	"exchangeapp/global"
	"exchangeapp/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateExchangeRate(ctx *gin.Context) {
	var exchangerate models.ExchangeRate

	if err := ctx.ShouldBindJSON(&exchangerate); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	exchangerate.Date = time.Now()
	if err := global.Db.AutoMigrate(&exchangerate); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := global.Db.Create(&exchangerate).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, exchangerate)
}

func GetExchangeRate(ctx *gin.Context) {
	var exchangerates []models.ExchangeRate = []models.ExchangeRate{}

	if err := global.Db.Find(&exchangerates).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, exchangerates)
}
