package main

import (
	"log"

	"kbaauto/contracts"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func main() {
	foodContract := new(contracts.FoodContract)
	distributorContract := new(contracts.DistributorContract)

	chaincode, err := contractapi.NewChaincode(foodContract, distributorContract)

	if err != nil {
		log.Panicf("Could not create chaincode : %v", err)
	}

	err = chaincode.Start()

	if err != nil {
		log.Panicf("Failed to start chaincode : %v", err)
	}
}
