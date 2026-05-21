package v1routes

import (
	v1handler "23.usersapi/internal/handler/v1"
	"23.usersapi/internal/routes"
	"github.com/gin-gonic/gin"
)

type RoutesUser struct {
	handlerUser *v1handler.HandlerUser
}

func NewRoutesUser(handlerUser *v1handler.HandlerUser) routes.Routers {
	return &RoutesUser{
		handlerUser: handlerUser,
	}
}

func (ur *RoutesUser) Register(r *gin.RouterGroup) {
	user := r.Group("/user")
	{
		user.GET("", ur.handlerUser.GetAllUsers)
		user.POST("", ur.handlerUser.PostUser)
		user.GET("/:uuid", ur.handlerUser.GetUserByUUID)
		user.PUT("/:uuid", ur.handlerUser.PutUser)
		user.DELETE("/:uuid", ur.handlerUser.DeleteUser)
		user.GET("/panic", ur.handlerUser.Panic)
	}
}
