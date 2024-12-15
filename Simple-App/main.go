package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Food struct {
	FoodID   string `json:"foodId"`
	Type     string `json:"type"`
	Origin   string `json:"origin"`
	Quantity string `json:"quantity"`
	Status   string `json:"status"`
	Owner    string `json:"owner"`
}

type FoodData struct {
	FoodID   string `json:"foodId"`
	Type     string `json:"type"`
	Origin   string `json:"origin"`
	Quantity string `json:"quantity"`
	Status   string `json:"status"`
	Owner    string `json:"owner"`
}

type Order struct {
	AssetType       string `json:"assetType"`
	OrderID         string `json:"orderId"`
	DistributorName string `json:"distributorName"`
	ProductID       string `json:"productId"`
	Quantity        string `json:"quantity"`
	Status          string `json:"status"`
}

type OrderData struct {
	AssetType       string `json:"assetType"`
	OrderID         string `json:"orderId"`
	DistributorName string `json:"distributorName"`
	ProductID       string `json:"productId"`
	Quantity        string `json:"quantity"`
	Status          string `json:"status"`
}

func main() {
	// Initialize Gin router
	router := gin.Default()

	// Serve static files (if needed)
	router.Static("/public", "./public")
	router.LoadHTMLGlob("templates/*")

	// Home route
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Manufacturer Dashboard",
		})
	})

	router.POST("/api/food", func(ctx *gin.Context) {
		var req Food
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
			return
		}

		result := submitTxnFn(
			"farmer",                // Client/Org (Farmer)
			"organicchannel",        // Channel name
			"organicfood",           // Chaincode name
			"FoodContract",          // Contract name
			"invoke",                // Action type
			make(map[string][]byte), // Transient data (none in this case)
			"CreateFood",            // Function to invoke
			req.FoodID,              // Food ID
			req.Type,                // Food type
			req.Origin,              // Origin of the food
			req.Quantity,            // Quantity of food
			req.Status,              // Status of food
		)

		// Send response
		ctx.JSON(http.StatusOK, gin.H{"message": "Created new food item", "result": result})
	})

	// Endpoint to read a food item by ID
	router.GET("/api/food/:id", func(ctx *gin.Context) {
		foodID := ctx.Param("id")

		// Call the "ReadFood" function in the contract (assuming submitTxnFn is correctly defined)
		result := submitTxnFn(
			"farmer",                // Client/Org (Farmer)
			"organicchannel",        // Channel name
			"organicfood",           // Chaincode name
			"FoodContract",          // Contract name
			"query",                 // Action type (query)
			make(map[string][]byte), // Transient data (none in this case)
			"ReadFood",              // Function to invoke
			foodID,                  // Food ID
		)

		// Send response
		ctx.JSON(http.StatusOK, gin.H{"data": result})
	})

	router.GET("/api/food/all", func(ctx *gin.Context) {
		result := submitTxnFn("farmer", "organicchannel", "organicfood", "FoodContract", "query", make(map[string][]byte), "GetAllFoods")

		var cars []FoodData

		if len(result) > 0 {
			// Unmarshal the JSON array string into the cars slice
			if err := json.Unmarshal([]byte(result), &cars); err != nil {
				fmt.Println("Error:", err)
				return
			}
		}
		ctx.JSON(http.StatusOK, gin.H{"data": result})

	})

	router.POST("/api/order", func(ctx *gin.Context) {
		var req Order
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
			return
		}

		// Check if all necessary fields are provided
		if req.OrderID == "" || req.DistributorName == "" || req.ProductID == "" || req.Quantity == "" || req.Status == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Missing required fields"})
			return
		}

		// Prepare private data (if needed)
		privateData := map[string][]byte{
			"orderId":         []byte(req.OrderID),
			"distributorName": []byte(req.DistributorName),
			"productId":       []byte(req.ProductID),
			"quantity":        []byte(req.Quantity),
			"status":          []byte(req.Status),
		}

		// Submit the private transaction to create the order using PDC
		result := submitTxnFn(
			"distributor",         // Client/Org (Distributor)
			"organicchannel",      // Channel name
			"organicfood",         // Chaincode name
			"DistributorContract", // Contract name
			"private",             // Action type (private)
			privateData,           // Private data for the transaction
			"CreateOrder",         // Function to invoke
			req.OrderID,           // Order ID
			req.DistributorName,   // Distributor name
			req.ProductID,         // Product ID
			req.Quantity,          // Quantity
			req.Status,            // Status
		)

		// Respond with success
		ctx.JSON(http.StatusOK, gin.H{"message": "Order created successfully", "order": req, "result": result})

	})

	router.GET("/api/order/:id", func(ctx *gin.Context) {
		// Retrieve the OrderID from the URL parameters
		orderId := ctx.Param("id")

		// Prepare the transaction to query the order using the submitTxnFn function
		result := submitTxnFn(
			"distributor",           // Client/Org (Distributor)
			"organicchannel",        // Channel name
			"organicfood",           // Chaincode name
			"DistributorContract",   // Contract name
			"query",                 // Action type (query)
			make(map[string][]byte), // Empty map for private data (if necessary)
			"ReadOrder",             // Function name to invoke
			orderId,                 // Order ID (parameter passed from the route)
		)

		// Return the result of the transaction (order data)
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Order retrieved successfully",
			"data":    result,
		})
	})

	router.GET("/api/allorders", func(ctx *gin.Context) {
		// Prepare the transaction to query all orders using the submitTxnFn function
		result := submitTxnFn(
			"distributor",           // Client/Org (Distributor)
			"organicchannel",        // Channel name
			"organicfood",           // Chaincode name
			"DistributorContract",   // Contract name
			"query",                 // Action type (query)
			make(map[string][]byte), // Empty map for private data (if necessary)
			"GetAllOrders",          // Function name to invoke (adjust according to your chaincode method)
			"",                      // No specific order ID, query all orders
		)

		// Return the result of the transaction (list of orders)
		ctx.JSON(http.StatusOK, gin.H{
			"message": "All Orders retrieved successfully",
			"data":    result,
		})
	})

	router.POST("/api/food/register", func(ctx *gin.Context) {
		var req struct {
			FoodID          string `json:"foodId"`
			OwnerName       string `json:"ownerName"`
			RegistrationNum string `json:"registrationNumber"`
		}

		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
			return
		}

		// Call the "RegisterFood" function in the contract (assuming submitTxnFn is correctly defined)
		result := submitTxnFn(
			"retailer",              // Client/Org (Retailer)
			"organicchannel",        // Channel name
			"organicfood",           // Chaincode name
			"FoodContract",          // Contract name
			"invoke",                // Action type
			make(map[string][]byte), // Transient data (none in this case)
			"RegisterFood",          // Function to invoke
			req.FoodID,              // Food ID
			req.OwnerName,           // Owner name
			req.RegistrationNum,     // Registration number
		)

		// Send response
		ctx.JSON(http.StatusOK, gin.H{"message": "Food shows successfully", "result": result})
	})

	// Start the Gin server
	router.Run(":3000")
}
