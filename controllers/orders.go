package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/Diva2504/Assignment-2-GLNG-KS03-003/models"
	"github.com/Diva2504/Assignment-2-GLNG-KS03-003/repository"
	"github.com/gin-gonic/gin"
)

type OrderResponse struct {
	OrderedAt    time.Time     `json:"ordered_at"`
	CustomerName string        `json:"customer_name"`
	Items        ItemsResponse `json:"items"`
}

type ItemsResponse struct {
	ItemCode    string `json:"item_code"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}

func (db Handler) GetAllOrder(c *gin.Context) {
	var (
		result   gin.H
		orderRes []OrderResponse
	)

	res, err := repository.GetAllOrders(db.Connect)
	for i := range res {
		orderRes[i].OrderedAt = res[i].CreatedAt
		orderRes[i].CustomerName = res[i].Customer_name
		orderRes[i].Items.ItemCode = res[i].Items.ItemCode
		orderRes[i].Items.Description = res[i].Items.Description
		orderRes[i].Items.Quantity = res[i].Items.Quantity
	}
	if err != nil {
		result = gin.H{
			"message": err,
		}
	}
	result = gin.H{
		"result": orderRes,
	}
	c.JSON(http.StatusOK, result)
}

func (db Handler) CreateOrder(c *gin.Context) {
	var (
		order  models.Orders
		result gin.H
	)
	if err := c.ShouldBindJSON(&order); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	err := repository.CreateOrder(&order, db.Connect)
	if err != nil {
		result = gin.H{
			"message": err,
		}
	}
	result = gin.H{
		"result": order,
	}
	c.JSON(http.StatusOK, result)
}

func (db Handler) UpdateOrder(c *gin.Context) {
	var (
		order  models.Orders
		result gin.H
	)
	if err := c.ShouldBindJSON(&order); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	orderId, _ := strconv.Atoi(c.Param("id"))
	_, err := repository.UpdateOrder(orderId, &order, db.Connect)
	if err != nil {
		result = gin.H{
			"message": err,
		}
	}
	result = gin.H{
		"result": order,
	}
	c.JSON(http.StatusOK, result)
}

func (db Handler) DeleteOrder(c *gin.Context) {
	var result gin.H
	requestId := c.Param("id")
	id, _ := strconv.Atoi(requestId)
	err := repository.DeleteOrder(id, db.Connect)
	if err != nil {
		result = gin.H{
			"message": err,
		}
	}
	result = gin.H{
		"message": "Order has been deleted",
	}
	c.JSON(http.StatusOK, result)
}
