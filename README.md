# NPCI-Blockchain-Assignment-7 : Create an ERC20-like Token on Hyperledger Fabric

## Assignment Objective
This assignment involves implementing an **ERC20-like token** on **Hyperledger Fabric**. The objective is to develop a chaincode that supports **minting**, **burning**, **transfers**, and **approvals**, ensuring that token transactions are securely recorded on a **Fabric blockchain network**.

**By completing this assignment, you will**:
- Set up a Hyperledger Fabric network.
- Implement token-like functionalities using chaincode.
- Deploy and interact with the token on the blockchain using CLI and Fabric SDK.

## Prerequisites
1. **GitHub Account** (for code submission).
2. **Docker and Docker** Compose installed.
3. **Go Programming Knowledge** (for writing chaincode).
4. **Basic Understanding of Hyperledger Fabric** (peer, orderer, chaincode concepts).
5. Access to `fabric-samples` repository: [Hyperledger Fabric Samples](https://github.com/hyperledger/fabric-samples).


## Assignment Tasks

### 1️⃣ Setup Hyperledger Fabric Network
Use Hyperledger Fabric's `test-network` to create a blockchain network.

The network should include:
- **Two peer organizations**: `Org1` and `Org2`.
- **One orderer organization** with a single orderer node.
- Each peer organization should have at least **one peer node**.

### 2️⃣ Implement Chaincode for ERC20-like Token
Create a chaincode in **Go** for token operations similar to an **ERC20 token**.

The chaincode should include:
- **Token Initialization** (`InitLedger`)
- **Minting tokens** (`MintTokens`) → Admin only
- **Transferring tokens** (`TransferTokens`) → Users send tokens
- **Checking balance** (`GetBalance`) → Retrieve holdings
- **Approving spending** (`ApproveSpender`) → Set spending limits
- **Transferring from approved spenders** (`TransferFrom`) → Allowed transfers
- **Burning tokens** (`BurnTokens`) → Admin only

### 3️⃣ Create and Deploy a Token
Deploy the chaincode to the Fabric network.

- Define one **admin address** (`User A`) and **three user addresses** (`User 2, User 3, User 4`).
- Define token properties similar to **BNB**:
  - **Token Name:** `BNB-Token`
  - **Symbol:** `BNB`
  - **Decimals:** `18`
  - **Total Supply:** `200,000,000 BNB`
  - Mint initial supply to an **admin wallet address**.

### 4️⃣ Execute Transactions
#### ✅ Mint Tokens
- Mint **10,000 BNB** for the admin account.

#### ✅ Transfer Tokens
- Transfer **100 BNB** from `User A` to `User B`.

#### ✅ Approve Token Spend
- Allow `User B` to spend **50 BNB** from `User A’s` balance.

#### ✅ Burn Tokens
- Burn **500 BNB** from the admin account.

### 5️⃣ Query and Validate Ledger Data
- Query `User A` and `User B` balances to verify transactions.
- Validate that the **total supply** reflects burned tokens.
- Ensure the Fabric ledger maintains **data integrity** across peers.

---

## Deliverables
### GitHub Repository Must Contain:
- **Fabric network setup scripts**
- **Chaincode implementation** (`erc20token.go`)
- **Transaction execution scripts** (`invoke.sh`, `query.sh`)
- **Screenshots or logs** demonstrating:
  - Successful chaincode deployment
  - Execution of **minting, transfers, approvals, and burning**
  - Ledger queries showing correct balances

---

##  Additional Notes
- Ensure all dependencies (**Docker, Fabric binaries, Go**) are installed before starting.
- Debug any issues using Fabric logs
  ```bash
  ( docker logs <peer/container_name> )
  ```
- Use Fabric event listeners to capture token transfers and approvals.
  
Happy coding! 🚀
