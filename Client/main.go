package main

import "fmt"

func main() {
	// result := submitTxnFn(
	// 	"farmer",
	// 	"organicchannel",
	// 	"organicfood",
	// 	"FoodContract",
	// 	"invoke",
	// 	make(map[string][]byte),
	// 	"CreateFood",
	// 	"Food-2",
	// 	"Greps",
	// 	"India",
	// 	"1000",
	// )

	// privateData := map[string][]byte{
	// 	"make":       []byte("Maruti"),
	// 	"model":      []byte("Alto"),
	// 	"color":      []byte("Red"),
	// 	"distributorName": []byte("Popular"),
	// }
	// 	privateData := map[string][]byte{
	// 	"quantity":       []byte("100"),
	// 	"status":      []byte("Delivered"),
	// 	"distributorName":      []byte("Maruti"),
	// 	"productId": []byte("Product-1"),
	// }

	// result := submitTxnFn("distributor", "organicchannel", "organicfood", "DistributorContract", "private", privateData, "CreateOrder", "ORD-1")

	// result := submitTxnFn("distributor", "organicchannel", "organicfood", "DistributorContract", "query", make(map[string][]byte), "ReadOrder", "ORD-1")

	// result := submitTxnFn("farmer", "organicchannel", "organicfood", "FoodContract", "query", make(map[string][]byte), "GetAllFoods")

	result := submitTxnFn("farmer", "organicchannel", "organicfood", "DistributorContract", "query", make(map[string][]byte), "GetAllOrders")

	// result := submitTxnFn("farmer", "organicchannel", "organicfood", "FoodContract", "query", make(map[string][]byte), "GetMatchingOrders", "Food-1")

	// result := submitTxnFn("farmer", "organicchannel", "organicfood", "FoodContract", "invoke", make(map[string][]byte), "MatchOrder", "Food-1", "ORD-03")

	// result := submitTxnFn("retailer", "organicchannel", "organicfood", "FoodContract", "invoke", make(map[string][]byte), "RegisterFood", "Food-1", "Dani", "KL-01-CD-01")

	// result := submitTxnFn("farmer", "organicchannel", "organicfood", "FoodContract", "query", make(map[string][]byte), "ReadFood", "Food-1")

	fmt.Println(result)

}
