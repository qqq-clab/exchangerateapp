package controllers

import (
	"exchangeapp/global"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

func LikeArticle(ctx *gin.Context) {
	articleID := ctx.Param("id")

	likeID := "article:" + articleID + ":likes"

	if err := global.RedisDB.Incr(likeID).Err(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Successfully liked the article",
	})
}

func GetArticleLikes(ctx *gin.Context) {
	articleID := ctx.Param("id")

	likeID := "article:" + articleID + ":likes"

	likes, err := global.RedisDB.Get(likeID).Result()

	if err == redis.Nil {
		likes = "0"
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"likes": likes,
	})
}
