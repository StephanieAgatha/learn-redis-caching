package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"redishop/model"
	"redishop/usecase"
	"redishop/util/helper"
	"strconv"
	"time"
)

type ProductController struct {
	prodUsecase usecase.ProductUsecase
	g           *gin.Engine
	redisC      *redis.Client
}

func (p ProductController) CreateNewProductHandler(c *gin.Context) {
	var product model.Products

	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"Error": "Wrong JSON Format"})
		return
	}

	if err := p.prodUsecase.CreateNewProduct(product); err != nil {
		c.AbortWithStatusJSON(500, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"Product": product})
}

func (p ProductController) GetAllProductHandler(c *gin.Context) {
	//get from redis
	var ctx = context.Background()

	val, err := p.redisC.Get(ctx, "alldata").Bytes()
	if err == nil {
		data := helper.JSONALLProductParse(val)
		c.JSON(200, gin.H{"Data (Cached)": data})
		return
	}

	prods, err := p.prodUsecase.GetAllProduct()
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"Error": err.Error()})
		return
	}

	//marshal dulu sebelum set ke redis karena kita perlu mengubahnya menjadi byte array menggunakan json.Marshal()
	val, err = json.Marshal(prods)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"Error": err.Error()})
		return
	}

	//set to redis
	if err = p.redisC.Set(ctx, "alldata", val, 10*time.Second).Err(); err != nil {
		fmt.Printf("Failed to set value %v", err.Error())
		return
	}

	c.JSON(200, gin.H{"Data": prods})
}

func (p ProductController) GetProductByIDHandler(c *gin.Context) {
	var prodid struct {
		ProductID int `json:"product_id"`
	}

	if err := c.ShouldBindJSON(&prodid); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"Error": err.Error()})
		return
	}

	prodidstr := strconv.Itoa(prodid.ProductID)

	//get from redis (if any)
	var ctx = context.Background()
	val, err := p.redisC.Get(ctx, "id:"+prodidstr).Bytes()
	if err != nil { //jika error pada get data from redis,dan jika data tidak ditemukan di redis,maka lempar request ke db
		if err != redis.Nil { //jika redis is nil
			c.AbortWithStatusJSON(500, gin.H{"Error": err.Error()})
			return
		}

		// Data not found in Redis, continue to fetch from other source
		prod, err := p.prodUsecase.GetProductByID(prodid.ProductID)
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{"Error": err.Error()})
			return
		}

		//marshal data prod
		val, err = json.Marshal(prod)
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{"Error": "Error to marshal data"})
			return
		}

		//set to redis
		if err = p.redisC.Set(ctx, "id:"+prodidstr, val, 10*time.Second).Err(); err != nil {
			fmt.Printf("Failed to set value %v", err.Error())
			return
		}

		c.JSON(200, gin.H{"Data": prod})
		return
	}

	//data cached
	var prod model.Products
	if err = json.Unmarshal(val, &prod); err != nil {
		c.AbortWithStatusJSON(500, gin.H{"Error": "Error to unmarshal data"})
		return
	}

	c.JSON(200, gin.H{"Data (Cached)": prod})
}

func (p ProductController) Route() {
	productRoute := p.g.Group("/products")
	{
		productRoute.GET("showall", p.GetAllProductHandler)
		productRoute.POST("create", p.CreateNewProductHandler)
		productRoute.POST("id", p.GetProductByIDHandler)
	}
}

func NewProductController(prodUsecase usecase.ProductUsecase, g *gin.Engine, redisC *redis.Client) *ProductController {
	return &ProductController{
		prodUsecase: prodUsecase,
		g:           g,
		redisC:      redisC,
	}
}
