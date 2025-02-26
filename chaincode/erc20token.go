package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-chaincode-go/pkg/cid"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

const MaxTokensInCirculation int = 200000000

type Token struct {
	Name        string `json:"name"`
	Symbol      string `json:"symbol"`
	Decimals    int    `json:"decimals"`
	TotalSupply int    `json:"totalSupply"`
}

type TokenContract struct {
	contractapi.Contract
}

type Balance struct {
	Balance int `json:"balance"`
}

// TODO: InitLedger for token initialization
func (tc *TokenContract) InitLedger(ctx contractapi.TransactionContextInterface) error {

	clientID, err := ctx.GetClientIdentity().GetID()
	if err != nil {
		return err
	}

	found, err := cid.HasOUValue(ctx.GetStub(), "admin")
	if err != nil {
		return err
	}
	if !found {
		return fmt.Errorf("client is not an admin!")
	}

	token := Token{
		Name:        "BNB-Token",
		Symbol:      "BNB",
		Decimals:    18,
		TotalSupply: 10000,
	}

	tokenBytes, err := json.Marshal(token)
	if err != nil {
		return err
	}

	err = ctx.GetStub().PutState(clientID, tokenBytes)
	if err != nil {
		return err
	}
	return nil
}

// TODO: MintTokens for minting tokens
func (tc *TokenContract) MintTokens(ctx contractapi.TransactionContextInterface, tokensToMint int) error {
	clientID, err := ctx.GetClientIdentity().GetID()
	if err != nil {
		return err
	}

	found, err := cid.HasOUValue(ctx.GetStub(), "admin")
	if err != nil {
		return err
	}
	if !found {
		return fmt.Errorf("client is not an admin!")
	}

	tokenBytes, err := ctx.GetStub().GetState(clientID)
	if err != nil {
		return err
	}

	var token *Token
	err = json.Unmarshal(tokenBytes, token)
	if err != nil {
		return err
	}

	if token.TotalSupply+tokensToMint > MaxTokensInCirculation {
		return fmt.Errorf("total supply of tokens exceeds the max tokens that can be in circulation!")
	}

	token.TotalSupply += tokensToMint

	tokenBytes, err = json.Marshal(token)
	if err != nil {
		return err
	}

	err = ctx.GetStub().PutState(clientID, tokenBytes)
	if err != nil {
		return err
	}

	return nil
}

// TODO: TransferTokens for transferring tokens
func (tc *TokenContract) TransferTokens(ctx contractapi.TransactionContextInterface) {

}

// TODO: GetBalance to check the balance
func (tc *TokenContract) GetBalance(ctx contractapi.TransactionContextInterface) {

}

// TODO: ApproveSpender for approving spending
func (tc *TokenContract) ApproveSpender(ctx contractapi.TransactionContextInterface) {

}

// TODO: TransferFrom for transferring from approved spenders
func (tc *TokenContract) TransferFrom(ctx contractapi.TransactionContextInterface) {

}

// TODO: BurnTokens for burning tokens
func (tc *TokenContract) BurnTokens(ctx contractapi.TransactionContextInterface) {

}

func main() {
	chaincode, err := contractapi.NewChaincode(&TokenContract{})
	if err != nil {
		fmt.Printf("Error creating token chaincode: %v", err)
		return
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting token chaincode: %v", err)
	}
}
