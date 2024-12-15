# Start the Organicfood network network

```
./startOrganicfoodNetwork.sh
```
### -------------Farmer Org-------

```
export FABRIC_CFG_PATH=./peercfg
export CHANNEL_NAME=organicchannel
export CORE_PEER_LOCALMSPID=FarmerMSP
export CORE_PEER_TLS_ENABLED=true
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/farmer.auto.com/peers/peer0.farmer.auto.com/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/farmer.auto.com/users/Admin@farmer.auto.com/msp
export CORE_PEER_ADDRESS=localhost:7051
export ORDERER_CA=${PWD}/organizations/ordererOrganizations/auto.com/orderers/orderer.auto.com/msp/tlscacerts/tlsca.auto.com-cert.pem
export FARMER_PEER_TLSROOTCERT=${PWD}/organizations/peerOrganizations/farmer.auto.com/peers/peer0.farmer.auto.com/tls/ca.crt
export DISTRIBUTOR_PEER_TLSROOTCERT=${PWD}/organizations/peerOrganizations/distributor.auto.com/peers/peer0.distributor.auto.com/tls/ca.crt
export RETAILER_PEER_TLSROOTCERT=${PWD}/organizations/peerOrganizations/retailer.auto.com/peers/peer0.retailer.auto.com/tls/ca.crt
```


### ------------Distributor Org------

```
export FABRIC_CFG_PATH=./peercfg
export CHANNEL_NAME=organicchannel 
export CORE_PEER_LOCALMSPID=DistributorMSP 
export CORE_PEER_TLS_ENABLED=true
export CORE_PEER_ADDRESS=localhost:9051 
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/distributor.auto.com/peers/peer0.distributor.auto.com/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/distributor.auto.com/users/Admin@distributor.auto.com/msp
export ORDERER_CA=${PWD}/organizations/ordererOrganizations/auto.com/orderers/orderer.auto.com/msp/tlscacerts/tlsca.auto.com-cert.pem
export FARMER_PEER_TLSROOTCERT=${PWD}/organizations/peerOrganizations/farmer.auto.com/peers/peer0.farmer.auto.com/tls/ca.crt
export DISTRIBUTOR_PEER_TLSROOTCERT=${PWD}/organizations/peerOrganizations/distributor.auto.com/peers/peer0.distributor.auto.com/tls/ca.crt
export RETAILER_PEER_TLSROOTCERT=${PWD}/organizations/peerOrganizations/retailer.auto.com/peers/peer0.retailer.auto.com/tls/ca.crt
```



### -----------Retailer Org----------

```
export FABRIC_CFG_PATH=./peercfg
export CHANNEL_NAME=organicchannel 
export CORE_PEER_LOCALMSPID=RetailerMSP 
export CORE_PEER_TLS_ENABLED=true
export CORE_PEER_ADDRESS=localhost:11051 
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/retailer.auto.com/peers/peer0.retailer.auto.com/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/retailer.auto.com/users/Admin@retailer.auto.com/msp
export ORDERER_CA=${PWD}/organizations/ordererOrganizations/auto.com/orderers/orderer.auto.com/msp/tlscacerts/tlsca.auto.com-cert.pem
export FARMER_PEER_TLSROOTCERT=${PWD}/organizations/peerOrganizations/farmer.auto.com/peers/peer0.farmer.auto.com/tls/ca.crt
export DISTRIBUTOR_PEER_TLSROOTCERT=${PWD}/organizations/peerOrganizations/distributor.auto.com/peers/peer0.distributor.auto.com/tls/ca.crt
export RETAILER_PEER_TLSROOTCERT=${PWD}/organizations/peerOrganizations/retailer.auto.com/peers/peer0.retailer.auto.com/tls/ca.crt
```

## ********************* Invovke CreateFood **********
```

peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.auto.com --tls --cafile $ORDERER_CA -C $CHANNEL_NAME -n organicfood --peerAddresses localhost:7051 --tlsRootCertFiles $FARMER_PEER_TLSROOTCERT --peerAddresses localhost:9051 --tlsRootCertFiles $DISTRIBUTOR_PEER_TLSROOTCERT --peerAddresses localhost:11051 --tlsRootCertFiles $RETAILER_PEER_TLSROOTCERT -c '{"function":"CreateFood","Args":["Food-1", "Vegetable", "India", "100"]}'

```

## ********************* ReadFood **********
```
peer chaincode query -C $CHANNEL_NAME -n organicfood --peerAddresses localhost:7051 --tlsRootCertFiles $FARMER_PEER_TLSROOTCERT -c '{"function":"ReadFood","Args":["Food-2"]}'
```

## *********************getallfoods **********
```
peer chaincode query -C $CHANNEL_NAME -n organicfood -c '{"Args":["GetAllFoods"]}'
```

## ********************** DeleteFood *****************************8888
```
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.auto.com --tls --cafile $ORDERER_CA -C $CHANNEL_NAME -n organicfood  -c '{"function":"DeleteFood","Args":["Food-1"]}'
```

## *********** GetFoodsByRange **********
```
peer chaincode query -C $CHANNEL_NAME -n organicfood -c '{"function":"GetFoodsByRange","Args":["Food-2", "Food-4"]}'
```
## ******************* To invoke the GetFoodHistory function ***************
```
peer chaincode query -C $CHANNEL_NAME -n organicfood -c '{"function":"GetFoodHistory","Args":["Food-1"]}'
```



## ****************************GetFoodsWithPagination***********************
```
peer chaincode query -C $CHANNEL_NAME -n organicfood -c '{"function":"GetFoodsWithPagination","Args":["5", ""]}'
```

	
## ***************** PDC Command Second Terminal ***********************
```
export QUANTITY=$(echo -n "100" | base64 | tr -d \\n)

export STATUS=$(echo -n "Disapaced" | base64 | tr -d \\n)
export DISTRIBUTOR_NAME=$(echo -n "XXX" | base64 | tr -d \\n)

export PRODUCTID=$(echo -n "Product-1" | base64 | tr -d \\n)
```
## ********************** Invoke CreateOrder PDC******************
```
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.auto.com --tls --cafile $ORDERER_CA -C $CHANNEL_NAME -n organicfood --peerAddresses localhost:7051 --tlsRootCertFiles $FARMER_PEER_TLSROOTCERT --peerAddresses localhost:9051 --tlsRootCertFiles $DISTRIBUTOR_PEER_TLSROOTCERT --peerAddresses localhost:11051 --tlsRootCertFiles $RETAILER_PEER_TLSROOTCERT -c '{"Args":["DistributorContract:CreateOrder","ORD101"]}' --transient "{\"quantity\":\"$QUANTITY\",\"status\":\"$STATUS\",\"distributorName\":\"$DISTRIBUTOR_NAME\",\"productId\":\"$PRODUCTID\"}"
```

## ************************ ReadOrder **************
```
peer chaincode query -C $CHANNEL_NAME -n organicfood -c '{"Args":["DistributorContract:ReadOrder","ORD201"]}'
```

## ********************* GetAllOrders **********
```
peer chaincode query -C $CHANNEL_NAME -n organicfood -c '{"Args":["DistributorContract:GetAllOrders"]}'
```
##  ********** GetOrdersByRange **********

```
peer chaincode query -C $CHANNEL_NAME -n organicfood -c '{"function":"GetOrdersByRange","Args":["ORD101", "ORD103"]}'
```

# **************** API ************************


### ************************ CreateFood *************************
POST  
```
http://localhost:3000/api/food
```
```
{
  "foodId": "food123",
  "type": "Organic Rice",
  "origin": "India",
  "quantity": "100kg",
  "status": "Harvested",
  "owner": "Farmer"
}
```

## ******************************** ReadFood ********************************
GET  
```
http://localhost:3000/api/food/food123
```

## ******************************** RegisterFood ********************************
POST 
```
http://localhost:3000/api/food/register
```
```
{
  "foodId": "food123",
  "ownerName": "Retailer Name",
  "registrationNumber": "KL-1234"
}
```

## ******************************** GetAllFoods ********************************

POST  
```
http://localhost:3000/api/food/all
```

## ******************************** Create Order ********************************

POST  
```
http://localhost:3000/api/order
```


```
{
  "OrderID": "12345",
  "DistributorName": "ABC Distributors",
  "ProductID": "P001",
  "Quantity": "100",
  "Status": "Pending"
}
```

## ******************************** Read Order ********************************

GET   
``` 
http://localhost:3000/api/order/2
```

## ******************************** GetAllOrders ********************************
GET     
``` 
http://localhost:3000/api/allorders

```




# ********** stop network *********
```
./stopOrganicfoodNetwork.sh
```














