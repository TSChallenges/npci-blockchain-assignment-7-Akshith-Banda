#!/bin/bash

# Set the environment variables for Org1
export PATH=/workspaces/npci-blockchain-assignment-7-Akshith-Banda/fabric-samples/bin:$PATH
export FABRIC_CFG_PATH=/workspaces/npci-blockchain-assignment-7-Akshith-Banda/fabric-samples/config/
export CORE_PEER_TLS_ENABLED=true
export CORE_PEER_LOCALMSPID="Org1MSP"
export CORE_PEER_TLS_ROOTCERT_FILE=/workspaces/npci-blockchain-assignment-7-Akshith-Banda/fabric-samples/test-network/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt

export CORE_PEER_MSPCONFIGPATH=/workspaces/npci-blockchain-assignment-7-Akshith-Banda/fabric-samples/test-network/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
# export CORE_PEER_MSPCONFIGPATH=/workspaces/npci-blockchain-assignment-7-Akshith-Banda/fabric-samples/test-network/organizations/peerOrganizations/org1.example.com/users/User1@org1.example.com/msp
# export CORE_PEER_MSPCONFIGPATH=/workspaces/npci-blockchain-assignment-7-Akshith-Banda/fabric-samples/test-network/organizations/peerOrganizations/org1.example.com/users/User2@org1.example.com/msp
# export CORE_PEER_MSPCONFIGPATH=/workspaces/npci-blockchain-assignment-7-Akshith-Banda/fabric-samples/test-network/organizations/peerOrganizations/org1.example.com/users/User3@org1.example.com/msp

export CORE_PEER_ADDRESS=localhost:7051

# Query balances - admin
peer chaincode query -C erc20 -n erc20cc -c '{"Args":["GetBalance", "eDUwOTo6Q049QWRtaW5Ab3JnMS5leGFtcGxlLmNvbSxPVT1hZG1pbixMPVNhbiBGcmFuY2lzY28sU1Q9Q2FsaWZvcm5pYSxDPVVTOjpDTj1jYS5vcmcxLmV4YW1wbGUuY29tLE89b3JnMS5leGFtcGxlLmNvbSxMPVNhbiBGcmFuY2lzY28sU1Q9Q2FsaWZvcm5pYSxDPVVT"]}'
