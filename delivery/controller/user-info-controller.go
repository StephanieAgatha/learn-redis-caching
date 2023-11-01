package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"redishop/delivery/middleware"
	"redishop/model"
	"redishop/usecase"
)

type UserController struct {
	UserUC usecase.UserUsecase
	gin    *gin.Engine
	redisC *redis.Client
}

func (u UserController) CreateUserHandler(c *gin.Context) {
	//bind json
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"Error": err.Error()})
		return
	}

	//usecase logic
	if err := u.UserUC.CreateUser(user); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"Message": "Successfully Register"})
}

// create route method
func (u UserController) Route() {
	u.gin.POST("/new", middleware.AuthMiddleware(u.redisC), u.CreateUserHandler)
}

func NewUserController(uc usecase.UserUsecase, g *gin.Engine, rediss *redis.Client) *UserController {
	return &UserController{
		UserUC: uc,
		gin:    g,
		redisC: rediss,
	}
}
