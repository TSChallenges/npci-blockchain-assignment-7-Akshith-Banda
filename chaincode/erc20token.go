package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/hyperledger/fabric-chaincode-go/pkg/cid"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

const MaxTokensInCirculation int = 200000000
const tokenName = "BNB-Token"

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

	isadmin, clientId, err := isAdmin(ctx)
	if err != nil {
		return err
	}

	if !isadmin {
		return errors.New("client not admin. Only admins can init ledger")
	}

	token := &Token{
		Name:        "BNB-Token",
		Symbol:      "BNB",
		Decimals:    18,
		TotalSupply: 10000,
	}

	err = setToken(ctx, token)
	if err != nil {
		return err
	}

	balance := &Balance{
		Balance: token.TotalSupply,
	}

	err = setBalance(ctx, balance, clientId)
	if err != nil {
		return err
	}
	return nil
}

// TODO: MintTokens for minting tokens
func (tc *TokenContract) MintTokens(ctx contractapi.TransactionContextInterface, tokensToMint int) error {

	isadmin, clientId, err := isAdmin(ctx)
	if err != nil {
		return err
	}

	if !isadmin {
		return errors.New("client not admin. Only admin can mint tokens")
	}

	token, err := getToken(ctx)
	if err != nil {
		return err
	}

	if token.TotalSupply+tokensToMint > MaxTokensInCirculation {
		return errors.New("total supply of tokens exceeds the max tokens that can be in circulation")
	}

	token.TotalSupply += tokensToMint

	err = setToken(ctx, token)
	if err != nil {
		return err
	}

	balance, err := getBalance(ctx, clientId)
	if err != nil {
		return err
	}

	balance.Balance += tokensToMint

	err = setBalance(ctx, balance, clientId)
	if err != nil {
		return err
	}

	return nil
}

// TODO: TransferTokens for transferring tokens
func (tc *TokenContract) TransferTokens(ctx contractapi.TransactionContextInterface, transferFrom, transferTo string, amountToTransfer int) error {
	balance, err := getBalance(ctx, transferFrom)
	if err != nil {
		return err
	}

	if balance.Balance < amountToTransfer {
		return errors.New("insufficient balance")
	}

	balance.Balance -= amountToTransfer
	err = setBalance(ctx, balance, transferFrom)
	if err != nil {
		return err
	}

	recvBalance, err := getBalance(ctx, transferTo)
	if err != nil {
		return err
	}

	recvBalance.Balance += amountToTransfer

	err = setBalance(ctx, recvBalance, transferTo)
	if err != nil {
		return err
	}

	return nil
}

// TODO: GetBalance to check the balance
func (tc *TokenContract) GetBalance(ctx contractapi.TransactionContextInterface, clientId string) (*Balance, error) {
	return getBalance(ctx, clientId)
}

// TODO: ApproveSpender for approving spending
func (tc *TokenContract) ApproveSpender(ctx contractapi.TransactionContextInterface, owner string, tokens int) error {

	balance, err := getBalance(ctx, owner)
	if err != nil {
		return err
	}

	if balance.Balance < tokens {
		return errors.New("not enough tokens")
	}

	clientId, err := ctx.GetClientIdentity().GetID()
	if err != nil {
		return err
	}

	compositeKey, err := ctx.GetStub().CreateCompositeKey("ApproveSpender", []string{owner, clientId})
	if err != nil {
		return err
	}

	approvSpendBalance, err := getBalance(ctx, compositeKey)
	if err != nil {
		return err
	}

	if approvSpendBalance == nil {
		approvSpendBalance = &Balance{
			Balance: tokens,
		}
	} else {
		approvSpendBalance.Balance += tokens
	}

	approveSpendBalanceBytes, err := json.Marshal(approvSpendBalance)
	if err != nil {
		return err
	}

	err = ctx.GetStub().PutState(compositeKey, approveSpendBalanceBytes)
	if err != nil {
		return err
	}
	return nil
}

// TODO: TransferFrom for transferring from approved spenders
func (tc *TokenContract) TransferFrom(ctx contractapi.TransactionContextInterface, owner, sendTo string, tokens int) error {
	clientId, err := ctx.GetClientIdentity().GetID()
	if err != nil {
		return err
	}

	compositeKey, err := ctx.GetStub().CreateCompositeKey("ApproveSpender", []string{owner, clientId})
	if err != nil {
		return err
	}

	balance, err := getBalance(ctx, compositeKey)
	if err != nil {
		return err
	}

	if balance.Balance < tokens {
		return errors.New("not enough balance to transfer")
	}

	balance.Balance -= tokens

	err = setBalance(ctx, balance, compositeKey)
	if err != nil {
		return err
	}

	recvBalance, err := getBalance(ctx, sendTo)
	if err != nil {
		return err
	}

	recvBalance.Balance += tokens

	err = setBalance(ctx, recvBalance, sendTo)
	if err != nil {
		return err
	}
	return nil

}

// TODO: BurnTokens for burning tokens
func (tc *TokenContract) BurnTokens(ctx contractapi.TransactionContextInterface, burnTokensCount int) error {
	token, err := getToken(ctx)
	if err != nil {
		return err
	}

	if token.TotalSupply < burnTokensCount {
		return errors.New("not enough tokens to burn")
	}

	isadmin, clientId, err := isAdmin(ctx)
	if err != nil {
		return err
	}

	if !isadmin {
		return errors.New("client not admin. Only admin can burn tokens")
	}

	balance, err := getBalance(ctx, clientId)
	if err != nil {
		return err
	}

	if balance.Balance < burnTokensCount {
		return errors.New("not enough tokens at admin to burn")
	}

	token.TotalSupply -= burnTokensCount

	err = setToken(ctx, token)
	if err != nil {
		return err
	}

	balance.Balance -= burnTokensCount

	err = setBalance(ctx, balance, clientId)
	if err != nil {
		return err
	}
	return nil
}

func getToken(ctx contractapi.TransactionContextInterface) (*Token, error) {
	tokenBytes, err := ctx.GetStub().GetState(tokenName)
	if err != nil {
		return nil, err
	}

	token := &Token{}
	err = json.Unmarshal(tokenBytes, &token)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func setToken(ctx contractapi.TransactionContextInterface, token *Token) error {
	tokenBytes, err := json.Marshal(token)
	if err != nil {
		return err
	}

	err = ctx.GetStub().PutState(tokenName, tokenBytes)
	if err != nil {
		return err
	}

	return nil
}

func getBalance(ctx contractapi.TransactionContextInterface, clientId string) (*Balance, error) {

	balanceBytes, err := ctx.GetStub().GetState(clientId)
	if err != nil {
		return nil, err
	}

	if balanceBytes == nil {
		return nil, nil
	}

	balance := &Balance{}
	err = json.Unmarshal(balanceBytes, balance)
	if err != nil {
		return nil, err
	}
	return balance, nil
}

func setBalance(ctx contractapi.TransactionContextInterface, balance *Balance, clientId string) error {
	balanceBytes, err := json.Marshal(balance)
	if err != nil {
		return err
	}

	err = ctx.GetStub().PutState(clientId, balanceBytes)
	if err != nil {
		return err
	}
	return nil
}

func isAdmin(ctx contractapi.TransactionContextInterface) (bool, string, error) {
	isadmin, err := cid.HasOUValue(ctx.GetStub(), "admin")
	if err != nil {
		return false, "", err
	}
	if !isadmin {
		return false, "", errors.New("client is not an admin")
	}

	clientId, err := ctx.GetClientIdentity().GetID()
	if err != nil {
		return true, "", errors.New(fmt.Sprintf("error getting client id : %+v\n", err))
	}

	return true, clientId, nil
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
