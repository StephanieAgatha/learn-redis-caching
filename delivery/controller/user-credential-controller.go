package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"redishop/model"
	"redishop/usecase"
	"time"
)

type UserCredentialController struct {
	userCredUc usecase.UserCredentialUsecase
	gin        *gin.Engine
	redisC     *redis.Client
}

func (u UserCredentialController) Register(c *gin.Context) {
	var userCred model.UserCredentials

	if err := c.ShouldBindJSON(&userCred); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"Error": "Bad JSON Format"})
		return
	}

	if err := u.userCredUc.Register(userCred); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"Message": "Successfully Register"})
}

func (u UserCredentialController) Login(c *gin.Context) {
	var userCred model.UserCredentials

	if err := c.ShouldBindJSON(&userCred); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"Error": "Bad JSON Format"})
		return
	}

	_, err := u.userCredUc.FindUserEMail(userCred.Email)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"Error": err.Error()})
		return
	}

	userToken, err := u.userCredUc.Login(userCred)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"Error": err.Error()})
		return
	}

	//set user info to redis (email+token) with 24 hours exp
	var ctx = context.Background()
	if err = u.redisC.Set(ctx, "userEmail:"+userCred.Email, userToken, 24*time.Hour).Err(); err != nil {
		c.AbortWithStatusJSON(500, gin.H{"Error": "Failed to save data to Redis"})
		return
	}

	c.JSON(200, gin.H{"Data": userToken})
}

func (u UserCredentialController) Logout(c *gin.Context) {
	var ctx = context.Background()

	var userLogout model.UserLogout

	if err := c.ShouldBindJSON(&userLogout); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"Error": "Wrong JSON Data"})
		return
	}

	//delete token
	err := u.redisC.Del(ctx, "userEmail:"+userLogout.Email).Err()
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"Error": "Failed to logout"})
		return
	}

	c.JSON(200, gin.H{"Message": "Logout successful"})
}

func (u UserCredentialController) Route() {
	authGroup := u.gin.Group("/auth")
	{
		authGroup.POST("/register", u.Register)
		authGroup.POST("/login", u.Login)
		authGroup.POST("/logout", u.Logout)
	}
}
func NewUserCredentialController(uc usecase.UserCredentialUsecase, g *gin.Engine, rediss *redis.Client) *UserCredentialController {
	return &UserCredentialController{
		userCredUc: uc,
		gin:        g,
		redisC:     rediss,
	}
}
