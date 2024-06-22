package router

import (
	"sosgo/config"
	"sosgo/handler"
	"sosgo/middleware"
	"sosgo/repository"
	"sosgo/service"

	"github.com/gin-gonic/gin"
)

func PostRouter(api *gin.RouterGroup) {
	postRepository := repository.NewPostRepository(config.DB)
	postService := service.NewPostService(postRepository)
	postHandler := handler.NewPostHandler(postService)

	r := api.Group("/tweets")

	r.GET("/", postHandler.Get)
	r.GET("/:id/detail", postHandler.Detail)

	r.Use(middleware.JWTMiddleware())

	r.POST("/", postHandler.Create)
	r.PUT("/:id/update", postHandler.Update)
	r.DELETE("/:id/delete", postHandler.Delete)
	r.GET("/me", postHandler.MyTweet)
}
