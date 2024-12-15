package contracts

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type FoodContract struct {
	contractapi.Contract
}

type PaginatedQueryResult struct {
	Records             []*Food `json:"records"`
	FetchedRecordsCount int32   `json:"fetchedRecordsCount"`
	Bookmark            string  `json:"bookmark"`
}

type Food struct {
	AssetType string `json:"assetType"`
	FoodID    string `json:"foodId"`
	Type      string `json:"type"`
	Origin    string `json:"origin"`
	Quantity  string `json:"quantity"`
	Status    string `json:"status"`
	Owner     string `json:"owner"`
}
type HistoryQueryResult struct {
	Record    *Food  `json:"record"`
	TxId      string `json:"txId"`
	Timestamp string `json:"timestamp"`
	IsDelete  bool   `json:"isDelete"`
}

type EventData struct {
	Type string
}

// FoodExists returns true when asset with given ID exists in world state
func (c *FoodContract) FoodExists(ctx contractapi.TransactionContextInterface, foodID string) (bool, error) {
	data, err := ctx.GetStub().GetState(foodID)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}
	return data != nil, nil
}

// CreateFood creates a new instance of Food by Farmer
func (c *FoodContract) CreateFood(ctx contractapi.TransactionContextInterface, foodID string, foodType string, origin string, quantity string) (string, error) {
	clientOrgID, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return "", fmt.Errorf("could not fetch client identity. %s", err)
	}

	if clientOrgID == "FarmerMSP" {
		exists, err := c.FoodExists(ctx, foodID)
		if err != nil {
			return "", fmt.Errorf("could not fetch the details from world state.%s", err)
		} else if exists {
			return "", fmt.Errorf("the food item, %s already exists", foodID)
		}

		food := Food{
			AssetType: "food",
			FoodID:    foodID,
			Type:      foodType,
			Origin:    origin,
			Quantity:  quantity,
			Status:    "Harvested",
			Owner:     "Farmer",
		}

		bytes, _ := json.Marshal(food)
		err = ctx.GetStub().PutState(foodID, bytes)
		if err != nil {
			return "", fmt.Errorf("could not create food item. %s", err)
		} else {
			addFoodEventData := EventData{
				Type: "Food creation",
			}
			eventDataByte, _ := json.Marshal(addFoodEventData)
			ctx.GetStub().SetEvent("CreateFood", eventDataByte)

			return fmt.Sprintf("successfully added food %v", foodID), nil
		}
	} else {
		return "", fmt.Errorf("user under following MSPID: %v can't perform this action", clientOrgID)
	}
}

// ReadFood retrieves an instance of Food from the world state
func (c *FoodContract) ReadFood(ctx contractapi.TransactionContextInterface, foodID string) (*Food, error) {
	bytes, err := ctx.GetStub().GetState(foodID)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if bytes == nil {
		return nil, fmt.Errorf("the food item %s does not exist", foodID)
	}

	var food Food
	err = json.Unmarshal(bytes, &food)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal world state data to type Food")
	}

	return &food, nil
}

// DeleteFood removes the instance of Food from the world state
func (c *FoodContract) DeleteFood(ctx contractapi.TransactionContextInterface, FoodID string) (string, error) {

	clientOrgID, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return "", fmt.Errorf("could not fetch client identity. %s", err)
	}
	// if clientOrgID == "Org1MSP" {
	// if clientOrgID == "manufacturer-auto-com" {
	if clientOrgID == "FarmerMSP" {

		exists, err := c.FoodExists(ctx, FoodID)
		if err != nil {
			return "", fmt.Errorf("%s", err)
		} else if !exists {
			return "", fmt.Errorf("the food, %s does not exist", FoodID)
		}

		err = ctx.GetStub().DelState(FoodID)
		if err != nil {
			return "", err
		} else {
			return fmt.Sprintf("food with id %v is deleted from the world state.", FoodID), nil
		}

	} else {
		return "", fmt.Errorf("user under following MSPID: %v can't perform this action", clientOrgID)
	}
}

func (c *FoodContract) GetFoodsByRange(ctx contractapi.TransactionContextInterface, startKey, endKey string) ([]*Food, error) {
	resultsIterator, err := ctx.GetStub().GetStateByRange(startKey, endKey)
	if err != nil {
		return nil, fmt.Errorf("could not fetch the  data by range. %s", err)
	}
	defer resultsIterator.Close()

	return foodResultIteratorFunction(resultsIterator)
}

func (c *FoodContract) GetAllFoods(ctx contractapi.TransactionContextInterface) ([]*Food, error) {

	queryString := `{"selector":{}}`

	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return nil, fmt.Errorf("could not fetch the query result. %s", err)
	}
	defer resultsIterator.Close()
	return foodResultIteratorFunction(resultsIterator)
}

func foodResultIteratorFunction(resultsIterator shim.StateQueryIteratorInterface) ([]*Food, error) {
	var foods []*Food
	for resultsIterator.HasNext() {
		queryResult, err := resultsIterator.Next()
		if err != nil {
			return nil, fmt.Errorf("could not fetch the details of the result iterator. %s", err)
		}
		var food Food
		err = json.Unmarshal(queryResult.Value, &food)
		if err != nil {
			return nil, fmt.Errorf("could not unmarshal the data. %s", err)
		}
		foods = append(foods, &food)
	}

	return foods, nil
}

func (c *FoodContract) GetFoodHistory(ctx contractapi.TransactionContextInterface, FoodID string) ([]*HistoryQueryResult, error) {

	resultsIterator, err := ctx.GetStub().GetHistoryForKey(FoodID)
	if err != nil {
		return nil, fmt.Errorf("could not get the data. %s", err)
	}
	defer resultsIterator.Close()

	var records []*HistoryQueryResult
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return nil, fmt.Errorf("could not get the value of resultsIterator. %s", err)
		}

		var food Food
		if len(response.Value) > 0 {
			err = json.Unmarshal(response.Value, &food)
			if err != nil {
				return nil, err
			}
		} else {
			food = Food{
				FoodID: FoodID,
			}
		}

		timestamp := response.Timestamp.AsTime()

		formattedTime := timestamp.Format(time.RFC1123)

		record := HistoryQueryResult{
			TxId:      response.TxId,
			Timestamp: formattedTime,
			Record:    &food,
			IsDelete:  response.IsDelete,
		}
		records = append(records, &record)
	}

	return records, nil
}

func (c *FoodContract) GetFoodsWithPagination(ctx contractapi.TransactionContextInterface, pageSize int32, bookmark string) (*PaginatedQueryResult, error) {

	queryString := `{"selector":{"assetType":"food"}}`

	resultsIterator, responseMetadata, err := ctx.GetStub().GetQueryResultWithPagination(queryString, pageSize, bookmark)
	if err != nil {
		return nil, fmt.Errorf("could not get the food records. %s", err)
	}
	defer resultsIterator.Close()

	foods, err := foodResultIteratorFunction(resultsIterator)
	if err != nil {
		return nil, fmt.Errorf("could not return the food records %s", err)
	}

	return &PaginatedQueryResult{
		Records:             foods,
		FetchedRecordsCount: responseMetadata.FetchedRecordsCount,
		Bookmark:            responseMetadata.Bookmark,
	}, nil
}

// DONE
// // GetMatchingOrders get matching orders for food from the orders

// func (c *FoodContract) GetMatchingOrders(ctx contractapi.TransactionContextInterface, FoodID string) ([]*Order, error) {

// 	food, err := c.ReadFood(ctx, FoodID)
// 	if err != nil {
// 		return nil, fmt.Errorf("error reading food %v", err)
// 	}
// 	queryString := fmt.Sprintf(`{"selector":{"assetType":"Order","make":"%s", "model": "%s", "color":"%s"}}`, food.Make, food.Model, food.Color)
// 	resultsIterator, err := ctx.GetStub().GetPrivateDataQueryResult(collectionName, queryString)

// 	if err != nil {
// 		return nil, fmt.Errorf("could not get the data. %s", err)
// 	}
// 	defer resultsIterator.Close()

// 	return OrderResultIteratorFunction(resultsIterator)

// }

// // MatchOrder matches food with matching order
// func (c *FoodContract) MatchOrder(ctx contractapi.TransactionContextInterface, FoodID string, orderID string) (string, error) {
// 	clientOrgID, err := ctx.GetClientIdentity().GetMSPID()
// 	if err != nil {
// 		return "", fmt.Errorf("could not fetch client identity. %s", err)
// 	}
// 	// if clientOrgID == "Org1MSP" {
// 	// if clientOrgID == "manufacturer-auto-com" {
// 	if clientOrgID == "FarmerMSP" {
// 		bytes, err := ctx.GetStub().GetPrivateData(collectionName, orderID)
// 		if err != nil {
// 			return "", fmt.Errorf("could not get the private data: %s", err)
// 		}

// 		var order Order

// 		err = json.Unmarshal(bytes, &order)

// 		if err != nil {
// 			return "", fmt.Errorf("could not unmarshal the data. %s", err)
// 		}

// 		food, err := c.ReadFood(ctx, FoodID)
// 		if err != nil {
// 			return "", fmt.Errorf("could not read the data. %s", err)
// 		}

// 		if food.Make == order.Make && food.Color == order.Color && food.Model == order.Model {
// 			food.OwnedBy = order.DistributorName
// 			food.Status = "assigned to a distributor"

// 			bytes, _ := json.Marshal(food)

// 			ctx.GetStub().DelPrivateData(collectionName, orderID)

// 			err = ctx.GetStub().PutState(FoodID, bytes)

// 			if err != nil {
// 				return "", fmt.Errorf("could not add the data %s", err)
// 			} else {
// 				return fmt.Sprintf("Deleted order %v and Assigned %v to %v", orderID, food.FoodId, order.DistributorName), nil
// 			}
// 		} else {
// 			return "", fmt.Errorf("order is not matching")
// 		}
// 	} else {
// 		return "", fmt.Errorf("user under following MSPID: %v can't perform this action", clientOrgID)
// 	}
// }

// RegisterFood register food to the buyer
// func (c *FoodContract) RegisterFood(ctx contractapi.TransactionContextInterface, FoodID string, ownerName string, registrationNumber string) (string, error) {
// 	clientOrgID, err := ctx.GetClientIdentity().GetMSPID()
// 	if err != nil {
// 		return "", fmt.Errorf("could not get the MSPID %s", err)
// 	}

// 	// if clientOrgID == "Org3MSP" {
// 	// if clientOrgID == "retailer-auto-com" {
// 	if clientOrgID == "RetailerMSP" {
// 		food, _ := c.ReadFood(ctx, FoodID)
// 		food.Status = fmt.Sprintf("Registered to  %v with plate number %v", ownerName, registrationNumber)
// 		food.OwnedBy = ownerName

// 		bytes, _ := json.Marshal(food)
// 		err = ctx.GetStub().PutState(FoodID, bytes)
// 		if err != nil {
// 			return "", fmt.Errorf("could not add the updated food details %s", err)
// 		} else {
// 			return fmt.Sprintf("Food %v successfully registered to %v", FoodID, ownerName), nil
// 		}

// 	} else {
// 		return "", fmt.Errorf("user under following MSPID: %v cannot able to perform this action", clientOrgID)
// 	}

// }


// RegisterFood registers a food item to the buyer
func (c *FoodContract) RegisterFood(ctx contractapi.TransactionContextInterface, foodID string, ownerName string, registrationNumber string) (string, error) {
	clientOrgID, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return "", fmt.Errorf("could not get the MSPID: %s", err)
	}

	// Ensure that the client is from the RetailerMSP
	if clientOrgID == "RetailerMSP" {
		// Read the food item from the world state
		food, err := c.ReadFood(ctx, foodID)
		if err != nil {
			return "", fmt.Errorf("could not read food with ID %s: %v", foodID, err)
		}

		// Update the status and ownership of the food item
		food.Status = fmt.Sprintf("Registered to %v with plate number %v", ownerName, registrationNumber)
		food.Owner = ownerName // Update the owner

		// Marshal the updated food object to bytes
		updatedFoodBytes, err := json.Marshal(food)
		if err != nil {
			return "", fmt.Errorf("could not marshal updated food data: %s", err)
		}

		// Store the updated food object in the world state
		err = ctx.GetStub().PutState(foodID, updatedFoodBytes)
		if err != nil {
			return "", fmt.Errorf("could not update food item %s in world state: %s", foodID, err)
		}

		// Emit an event for the registration action
		eventData := EventData{
			Type: "Food registration",
		}
		eventDataByte, _ := json.Marshal(eventData)
		ctx.GetStub().SetEvent("RegisterFood", eventDataByte)

		// Return success message
		return fmt.Sprintf("Food %v successfully registered to %v", foodID, ownerName), nil
	} else {
		return "", fmt.Errorf("user with MSPID %v cannot perform the registration action", clientOrgID)
	}
}

