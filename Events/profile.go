package main

// Config represents the configuration for a role.
type Config struct {
	CertPath     string `json:"certPath"`
	KeyDirectory string `json:"keyPath"`
	TLSCertPath  string `json:"tlsCertPath"`
	PeerEndpoint string `json:"peerEndpoint"`
	GatewayPeer  string `json:"gatewayPeer"`
	MSPID        string `json:"mspID"`
}

// Create a Profile map
var profile = map[string]Config{

	"farmer": {
		CertPath:     "../Organicfood-network/organizations/peerOrganizations/farmer.auto.com/users/User1@farmer.auto.com/msp/signcerts/cert.pem",
		KeyDirectory: "../Organicfood-network/organizations/peerOrganizations/farmer.auto.com/users/User1@farmer.auto.com/msp/keystore/",
		TLSCertPath:  "../Organicfood-network/organizations/peerOrganizations/farmer.auto.com/peers/peer0.farmer.auto.com/tls/ca.crt",
		PeerEndpoint: "localhost:7051",
		GatewayPeer:  "peer0.farmer.auto.com",
		MSPID:        "FarmerMSP",
	},

	"distributor": {
		CertPath:     "../Organicfood-network/organizations/peerOrganizations/distributor.auto.com/users/User1@distributor.auto.com/msp/signcerts/cert.pem",
		KeyDirectory: "../Organicfood-network/organizations/peerOrganizations/distributor.auto.com/users/User1@distributor.auto.com/msp/keystore/",
		TLSCertPath:  "../Organicfood-network/organizations/peerOrganizations/distributor.auto.com/peers/peer0.distributor.auto.com/tls/ca.crt",
		PeerEndpoint: "localhost:9051",
		GatewayPeer:  "peer0.distributor.auto.com",
		MSPID:        "DistributorMSP",
	},

	"retailer": {
		CertPath:     "../Organicfood-network/organizations/peerOrganizations/retailer.auto.com/users/User1@retailer.auto.com/msp/signcerts/cert.pem",
		KeyDirectory: "../Organicfood-network/organizations/peerOrganizations/retailer.auto.com/users/User1@retailer.auto.com/msp/keystore/",
		TLSCertPath:  "../Organicfood-network/organizations/peerOrganizations/retailer.auto.com/peers/peer0.retailer.auto.com/tls/ca.crt",
		PeerEndpoint: "localhost:11051",
		GatewayPeer:  "peer0.retailer.auto.com",
		MSPID:        "RetailerMSP",
	},

	"farmer2": {
		CertPath:     "../Organicfood-network/organizations/peerOrganizations/farmer.auto.com/users/User2@farmer.auto.com/msp/signcerts/cert.pem",
		KeyDirectory: "../Organicfood-network/organizations/peerOrganizations/farmer.auto.com/users/User2@farmer.auto.com/msp/keystore/",
		TLSCertPath:  "../Organicfood-network/organizations/peerOrganizations/farmer.auto.com/peers/peer0.farmer.auto.com/tls/ca.crt",
		PeerEndpoint: "localhost:7051",
		GatewayPeer:  "peer0.farmer.auto.com",
		MSPID:        "FarmerMSP",
	},

	"minifab-farmer": {
		CertPath:     "../Minifab_Network/vars/keyfiles/peerOrganizations/farmer.auto.com/users/Admin@farmer.auto.com/msp/signcerts/Admin@farmer.auto.com-cert.pem",
		KeyDirectory: "../Minifab_Network/vars/keyfiles/peerOrganizations/farmer.auto.com/users/Admin@farmer.auto.com/msp/keystore/",
		TLSCertPath:  "../Minifab_Network/vars/keyfiles/peerOrganizations/farmer.auto.com/peers/peer1.farmer.auto.com/tls/ca.crt",
		PeerEndpoint: "localhost:7003",
		GatewayPeer:  "peer1.farmer.auto.com",
		MSPID:        "farmer-auto-com",
	},

	"minifab-distributor": {
		CertPath:     "../Minifab_Network/vars/keyfiles/peerOrganizations/distributor.auto.com/users/Admin@distributor.auto.com/msp/signcerts/Admin@distributor.auto.com-cert.pem",
		KeyDirectory: "../Minifab_Network/vars/keyfiles/peerOrganizations/distributor.auto.com/users/Admin@distributor.auto.com/msp/keystore/",
		TLSCertPath:  "../Minifab_Network/vars/keyfiles/peerOrganizations/distributor.auto.com/peers/peer1.distributor.auto.com/tls/ca.crt",
		PeerEndpoint: "localhost:7004",
		GatewayPeer:  "peer1.distributor.auto.com",
		MSPID:        "distributor-auto-com",
	},

	"minifab-retailer": {
		CertPath:     "../Minifab_Network/vars/keyfiles/peerOrganizations/retailer.auto.com/users/Admin@retailer.auto.com/msp/signcerts/Admin@retailer.auto.com-cert.pem",
		KeyDirectory: "../Minifab_Network/vars/keyfiles/peerOrganizations/retailer.auto.com/users/Admin@retailer.auto.com/msp/keystore/",
		TLSCertPath:  "../Minifab_Network/vars/keyfiles/peerOrganizations/retailer.auto.com/peers/peer1.retailer.auto.com/tls/ca.crt",
		PeerEndpoint: "localhost:7005",
		GatewayPeer:  "peer1.retailer.auto.com",
		MSPID:        "retailer-auto-com",
	},
}
