package delivery

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"redishop/config"
	"redishop/delivery/controller"
	"redishop/delivery/middleware"
	"redishop/manager"
)

type Server struct {
	um     manager.UsecaseManager
	gin    *gin.Engine
	host   string
	redisC *redis.Client
}

// middleware goes here
func (s *Server) InitMiddleware() {
	// Create a Zap logger
	logger, _ := zap.NewProduction()
	defer logger.Sync() // Ensure logs are flushed

	// Use the logger
	s.gin.Use(middleware.ZapLogger(logger))
}

// controller
func (s *Server) InitController() {
	controller.NewUserController(s.um.UserUsecase(), s.gin, s.redisC).Route()
	controller.NewUserCredentialController(s.um.UserCredUsecase(), s.gin, s.redisC).Route()
	controller.NewProductController(s.um.ProductUsecase(), s.gin, s.redisC).Route()
}

// run server
func (s *Server) Run() {
	s.InitMiddleware()
	s.InitController()
	err := s.gin.Run(s.host)
	if err != nil {
		fmt.Printf("Failed to run server %v ", err.Error())
	}
}

func NewServer() *Server {
	//define contrusctor from config
	cfg, err := config.NewConfig()
	if err != nil {
		fmt.Printf("Failed on config server %v", err.Error())
	}

	//constructor from infra
	im, err := manager.NewInfraManager(cfg)
	if err != nil {
		fmt.Printf("Failed on construct infra %v", err.Error())
	}

	//get the Redis client from the infraManager
	redisClient := im.RedisClient()

	//constructor from repomanager
	rm := manager.NewRepoManager(im)
	//contructor from usecase manager
	um := manager.NewUsecaseManager(rm)

	//set host for gin server
	host := fmt.Sprintf("%s:%s", cfg.ApiConfig.Host, cfg.ApiConfig.Port)
	//return gin instance
	g := gin.Default()

	return &Server{
		um:     um,
		gin:    g,
		host:   host,
		redisC: redisClient,
	}
}
