package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"main/db"
	"main/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// POST
func CreateOrder(c *gin.Context) {
	var orderRequest models.OrderRequest

	// Bind the request
	if err := c.ShouldBindJSON(&orderRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
    
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10 * time.Second)
	defer cancel()

	// Fetch the user 
	user , err := getUser(ctx, orderRequest.UserID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Failed to fetch user: %s", err.Error())})		
		return
	}

	// Create order
	order  := models.Order {
		UserID      : orderRequest.UserID,
        UserName    : user.Name,
        ProductName : orderRequest.ProductName,
        Quantity    : orderRequest.Quantity,
        Status      : models.Pending,
		CreatedAt   : time.Now(),
	}

    
	// Store in data base
	if result := db.DB.Create(&order); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})																																				
		return
	}

	// Send the response
	c.JSON(http.StatusCreated, order)								  
}																																				

// GET																																										
func GetOrder(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Order ID"})
		return
	}
	var order models.Order

	if result := db.DB.First(&order, id); result.Error != nil { 
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found!"})
		return
	}

	c.JSON(http.StatusOK, order)
}

// GET
func GetOrders(c *gin.Context) {
	var orders []models.Order

	db.DB.Find(&orders)

	if len(orders) == 0 {
		orders = []models.Order{}
	}
	
	c.JSON(http.StatusOK, orders)
}

// PUT
func UpdateOrder(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Order ID"})
		return
	}

	var order models.Order

	if result := db.DB.First(&order, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found!"})
		return
	}

	var updateOrder models.Order
	if err := c.ShouldBindJSON(&updateOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := db.DB.Model(&order).Updates(updateOrder)
	if result.Error != nil { 
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update Order"})
		return
	}

	c.JSON(http.StatusOK, order)
}

// DELETE
func DeleteOrder(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Order ID"})
		return
	}

	var order models.Order

	if result := db.DB.First(&order, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found!"})
		return
	}

	result := db.DB.Delete(&order)
	if result.Error != nil { 
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete Order"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Order deleted successfully"})
}

// DELETE
func DeleteOrders(c *gin.Context) {
	result := db.DB.Unscoped().Delete(&models.Order{})
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete Orders"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "All Orders deleted successfully"})
}

func getUser(ctx context.Context, userID uint) (*models.User,error) {
	bastUrl := "http://localhost:8080"
	url := fmt.Sprintf("%s/users/%d",bastUrl,userID)

	// Create request
	req , err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Custom client
	client := &http.Client {
		Timeout: 10 * time.Second,
	}

	// Send request
	response, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch user: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("user with ID %d not found", userID)
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code from user service: %d", response.StatusCode)
	}

	var user models.User
	if err := json.NewDecoder(response.Body).Decode(&user); err != nil {
		return nil, fmt.Errorf("failed to decode user response: %w", err)
	}

	return &user, nil
}