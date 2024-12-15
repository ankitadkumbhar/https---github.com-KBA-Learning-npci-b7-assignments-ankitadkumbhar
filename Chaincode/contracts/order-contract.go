package contracts

// first write collection file

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// DistributorContract contract for managing distribution of organic food products
type DistributorContract struct {
	contractapi.Contract
}

type Order struct {
	AssetType       string `json:"assetType"`
	OrderID         string `json:"orderId"`
	DistributorName string `json:"distributorName"`
	ProductID       string `json:"productId"`
	Quantity        string `json:"quantity"`
	Status          string `json:"status"` // "Ordered", "Dispatched", "Delivered"
}

const collectionName string = "OrderCollection"

// OrderExists returns true when asset with given ID exists in private data collection
func (d *DistributorContract) OrderExists(ctx contractapi.TransactionContextInterface, orderId string) (bool, error) {
	data, err := ctx.GetStub().GetPrivateDataHash(collectionName, orderId)
	if err != nil {
		return false, fmt.Errorf("could not fetch private data hash. %s", err)
	}
	return data != nil, nil
}

func (d *DistributorContract) CreateOrder(ctx contractapi.TransactionContextInterface, orderId string) (string, error) {
	clientOrgID, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return "", fmt.Errorf("could not fetch client identity: %s", err)
	}

	// Check if the client is from the DistributorMSP
	if clientOrgID == "DistributorMSP" {
		exists, err := d.OrderExists(ctx, orderId)
		if err != nil {
			return "", fmt.Errorf("could not read from world state: %s", err)
		} else if exists {
			return "", fmt.Errorf("the order with ID %s already exists", orderId)
		}

		// Declare the order struct
		var order Order

		// Get transient data
		transientData, err := ctx.GetStub().GetTransient()
		if err != nil {
			return "", fmt.Errorf("could not fetch transient data: %s", err)
		}

		if len(transientData) == 0 {
			return "", fmt.Errorf("please provide the private data for the order")
		}

		// Parse transient data and set values
		quantity, exists := transientData["quantity"]
		if !exists {
			return "", fmt.Errorf("the quantity was not specified in transient data")
		}
		order.Quantity = string(quantity)

		status, exists := transientData["status"]
		if !exists {
			return "", fmt.Errorf("the status was not specified in transient data")
		}
		order.Status = string(status)

		distributorName, exists := transientData["distributorName"]
		if !exists {
			return "", fmt.Errorf("the distributorName was not specified in transient data")
		}
		order.DistributorName = string(distributorName)

		productID, exists := transientData["productId"]
		if !exists {
			return "", fmt.Errorf("the productId was not specified in transient data")
		}
		order.ProductID = string(productID)

		// Set asset type and order ID
		order.AssetType = "Order"
		order.OrderID = orderId

		// Marshal the order struct into bytes
		bytes, err := json.Marshal(order)
		if err != nil {
			return "", fmt.Errorf("could not marshal order data: %s", err)
		}

		// Store the order in the private data collection
		err = ctx.GetStub().PutPrivateData(collectionName, orderId, bytes)
		if err != nil {
			return "", fmt.Errorf("could not store the order: %s", err)
		}

		return fmt.Sprintf("Order with ID %v added successfully", orderId), nil
	} else {
		return "", fmt.Errorf("user from organisation with MSPID %v cannot create orders", clientOrgID)
	}
}

// ReadOrder retrieves an instance of Order from the private data collection
func (d *DistributorContract) ReadOrder(ctx contractapi.TransactionContextInterface, orderId string) (*Order, error) {
	exists, err := d.OrderExists(ctx, orderId)
	if err != nil {
		return nil, fmt.Errorf("could not read from world state. %s", err)
	} else if !exists {
		return nil, fmt.Errorf("the asset %s does not exist", orderId)
	}

	bytes, err := ctx.GetStub().GetPrivateData(collectionName, orderId)
	if err != nil {
		return nil, fmt.Errorf("could not get the private data. %s", err)
	}
	var order Order

	err = json.Unmarshal(bytes, &order)

	if err != nil {
		return nil, fmt.Errorf("could not unmarshal private data collection data to type Order")
	}

	return &order, nil

}

// DeleteOrder deletes an instance of Order from the private data collection
func (d *DistributorContract) DeleteOrder(ctx contractapi.TransactionContextInterface, orderId string) error {
	clientOrgID, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return fmt.Errorf("could not read the client identity. %s", err)
	}

	if clientOrgID == "DistributorMSP" {

		exists, err := d.OrderExists(ctx, orderId)

		if err != nil {
			return fmt.Errorf("could not read from world state. %s", err)
		} else if !exists {
			return fmt.Errorf("the asset %s does not exist", orderId)
		}

		return ctx.GetStub().DelPrivateData(collectionName, orderId)
	} else {
		return fmt.Errorf("organisation with %v cannot delete the order", clientOrgID)
	}
}

func (d *DistributorContract) GetAllOrders(ctx contractapi.TransactionContextInterface) ([]*Order, error) {
	queryString := `{"selector":{}}`
	resultsIterator, err := ctx.GetStub().GetPrivateDataQueryResult(collectionName, queryString)
	if err != nil {
		return nil, fmt.Errorf("could not fetch the query result. %s", err)
	}
	defer resultsIterator.Close()
	return OrderResultIteratorFunction(resultsIterator)
}

func (d *DistributorContract) GetOrdersByRange(ctx contractapi.TransactionContextInterface, startKey string, endKey string) ([]*Order, error) {
	resultsIterator, err := ctx.GetStub().GetPrivateDataByRange(collectionName, startKey, endKey)

	if err != nil {
		return nil, fmt.Errorf("could not fetch the private data by range. %s", err)
	}
	defer resultsIterator.Close()

	return OrderResultIteratorFunction(resultsIterator)

}

// iterator function

func OrderResultIteratorFunction(resultsIterator shim.StateQueryIteratorInterface) ([]*Order, error) {
	var orders []*Order
	for resultsIterator.HasNext() {
		queryResult, err := resultsIterator.Next()
		if err != nil {
			return nil, fmt.Errorf("could not fetch the details of result iterator. %s", err)
		}
		var order Order
		err = json.Unmarshal(queryResult.Value, &order)
		if err != nil {
			return nil, fmt.Errorf("could not unmarshal the data. %s", err)
		}
		orders = append(orders, &order)
	}

	return orders, nil
}
