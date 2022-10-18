package controllers

import (
	"assignment-02/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
Controllers for order
 1. Create Order -> method: POST;		path: "/orders"
 2. Get Orders   -> method: GET;		path: "/orders"
 3. Update Order -> method: PUT;		path: "/orders/:orderID"
 4. Delete Order -> method: DELETE; 	path: "/orders/:orderID"
*/

func (idb *InDB) CreateOrder(c *gin.Context) {
	var (
		order  structs.Order
		result gin.H
	)
	order.CustomerName = c.PostForm("customer_name")

	// create data in db
	err := idb.DB.Create(&order).Error
	if err != nil {
		result = gin.H{
			"result": "Error creating data",
		}
	} else {
		result = gin.H{
			"result": order,
		}
	}
	c.JSON(http.StatusOK, result)

}

func (idb *InDB) GetOrders(c *gin.Context) {
	var (
		orders []structs.Order
		result gin.H
	)

	idb.DB.Preload("Items").Find(&orders)
	if len(orders) <= 0 {
		result = gin.H{
			"result": nil,
			"count":  0}
	} else {
		result = gin.H{
			"result": orders,
			"count":  len(orders),
		}
	}
	c.JSON(http.StatusOK, result)
}

func (idb *InDB) UpdateOrder(c *gin.Context) {
	var (
		order    structs.Order
		newOrder structs.Order
		result   gin.H
	)

	order_id := c.Param("ID")

	// get the first data by its order_id
	err := idb.DB.First(&order, order_id).Error
	if err != nil {
		result = gin.H{
			"result": "Data not found",
		}
		c.JSON(http.StatusNotFound, result)

	}

	// set newOrder's CustomerName
	newOrder.CustomerName = c.PostForm("customer_name")

	err = idb.DB.Model(&order).Updates(newOrder).Error
	if err != nil {
		result = gin.H{
			"result": "Update failed",
		}
		c.JSON(http.StatusInternalServerError, result)

	} else {
		result = gin.H{
			"result": "Successfully updating data",
		}
	}
	c.JSON(http.StatusOK, result)

}

func (idb *InDB) DeleteOrder(c *gin.Context) {
	var (
		order  structs.Order
		result gin.H
	)
	order_id := c.Param("ID")
	err := idb.DB.First(&order, order_id).Error
	if err != nil {
		result = gin.H{
			"result": "Data not found",
		}
		c.JSON(http.StatusNotFound, result)

	}

	err = idb.DB.Delete(&order).Error
	if err != nil {
		result = gin.H{
			"result": "Failed deleting data",
		}
		c.JSON(http.StatusInternalServerError, result)

	} else {
		result = gin.H{
			"result": "Data has been successfully deleted",
		}
	}
	c.JSON(http.StatusOK, result)
}
