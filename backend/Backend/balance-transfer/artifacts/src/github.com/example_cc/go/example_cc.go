/*
Copyright IBM Corp. 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

var logger = shim.NewLogger("example_cc0")

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

type Parent struct {
	ObjectType  string `json:"doctype"`
	Id          string `json:"id"`
	Companyname string `json:"companyname"`
	TokenVal    string `json:"tokenval"`
}
type Child struct {
	ObjectType   string `json:"doctype"`
	Id           string `json:"id"`
	Companyname  string `json:"companyname"`
	Tokens       string `json:"tokens"`       //Tokens in hand
	PolicyTokens string `json:"policytokens"` // Tokens needed to buy policy
	RiskFactor   string `json:"riskFactor"`
	ClaimToken   string `json:"claimtoken"`
	PolicyPackage string `json:"policypackage"` // Tokens needed to buy policy
}
type Policy struct {
	ObjectType  string `json:"doctype"`
	Id          string `json:"id"`
	TotalTokens string `json:"totaltokens"`
	TotalRisk   string `json:"tokenrisk"`
	CreatedAt   string `json:"created_at"`
	Deadline    string `json:"deadline"`
	TokenPool   string `json:"tokenpool"`  
}

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	logger.Info("########### example_cc0 Init ###########")
	_, args := stub.GetFunctionAndParameters()
	Id := "Parent-1" // utils is a custom package. You can write your own too :P
	var parent = Parent{ObjectType: "Parent", Id: Id, Companyname: args[0], TokenVal: "10"}
	ParentBytes, _ := json.Marshal(parent)
	err := stub.PutState(Id, []byte(ParentBytes))
	if err != nil {
		return shim.Error("Failed during insert")
	}
	ChildId1 := "Child-1"
	var child1 = Child{ObjectType: "Child", Id: ChildId1, Companyname: args[1] + "_1", PolicyTokens: "0", Tokens: "0", RiskFactor: "6", ClaimToken: "0",PolicyPackage:"0"}
	ChildBytes1, _ := json.Marshal(child1)
	err = stub.PutState(ChildId1, []byte(ChildBytes1))
	if err != nil {
		return shim.Error("Failed during insert")
	}
	ChildId2 := "Child-2"
	var child2 = Child{ObjectType: "Child", Id: ChildId2, Companyname: args[2] + "_2", Tokens: "0", RiskFactor: "7", ClaimToken: "0",PolicyPackage: "0"}
	ChildBytes2, _ := json.Marshal(child2)
	err = stub.PutState(ChildId2, []byte(ChildBytes2))
	if err != nil {
		return shim.Error("Failed during insert")
	}
	return shim.Success(nil)
}

// Transaction makes payment of X units from A to B
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	logger.Info("########### example_cc0 Invoke ###########")

	function, args := stub.GetFunctionAndParameters()

	if function == "delete" {
		// Deletes an entity from its state
		// return t.delete(stub, args)
	}

	if function == "query" {
		// queries an entity state
		return t.query(stub, args)
	}
	if function == "move" {
		// Deletes an entity from its state
		// return t.move(stub, args)
	}
	if function == "newPolicy" {
		// Deletes an entity from its state
		return t.newPolicy(stub, args)
	}
	if function == "calculateRiscFactor" {
		// Deletes an entity from its state
		return t.calculateRiscFactor(stub, args)
	}
	if function == "buyPolicy" {
		// Deletes an entity from its state
		return t.buyPolicy(stub, args)
	}
	if function == "claim" {
		// Deletes an entity from its state
		return t.claim(stub, args)
	}
	if function == "buytokens" {
		// Deletes an entity from its state
		return t.buytokens(stub, args)
	}
	if function == "showPolicyTokens" {
		// Deletes an entity from its state
		return t.showPolicyTokens(stub, args)
	}
	logger.Errorf("Unknown action, check the first argument, must be one of 'delete', 'query', or 'move'. But got: %v", args[0])
	return shim.Error(fmt.Sprintf("Unknown action, check the first argument, must be one of 'delete', 'query', or 'move'. But got: %v", args[0]))
}

//Initialize new policy by the parent
func (t *SimpleChaincode) newPolicy(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	id := "policy1"

	var policy = Policy{ObjectType: "policy", Id: "policy1", TotalTokens: args[0], TotalRisk: "0", CreatedAt: args[1], Deadline: args[2],TokenPool: "0"}
	policyBytes, _ := json.Marshal(policy)
	err := stub.PutState(id, policyBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to update risk factor"))
	}
	return shim.Success(nil)
}

//Child buy tokens from parent 
func (t *SimpleChaincode) buytokens(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	id := args[0]
	tokenNo := args[1]
	pid := "policy1"
	parentCompanyBytes, _ := stub.GetState(pid)
	if parentCompanyBytes == nil {
		return shim.Error("Could not locate ")
	}

	parentCompany := &Policy{} //policy type struct
	json.Unmarshal(parentCompanyBytes, parentCompany)

	childCompanyBytes, _ := stub.GetState(id)
	if childCompanyBytes == nil {
		return shim.Error("Could not locate ")
	}

	childCompany := &Child{}

	json.Unmarshal(childCompanyBytes, childCompany)

	totTokens, _ := strconv.Atoi(childCompany.Tokens)    // Total tokens in hand of child
	n, _ := strconv.Atoi(parentCompany.TokenPool)     //Total tokens in hand of parent
		totTokens = totTokens + tokenNo 					//child compnay total tokend added by purchased tokens
		n = n - tokenNo               //parent total tokens reduced by x
		childCompany.TokenPool = strconv.Itoa(totTokens)
		parentCompany.TotalTokens = strconv.Itoa(n)
	childCompanyBytes, _ = json.Marshal(childCompany)
	err := stub.PutState(id, childCompanyBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to update risk factor"))
	}
	parentCompanyBytes, _ = json.Marshal(parentCompany)
	err = stub.PutState(pid, parentCompanyBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to update risk factor"))
	}
	return shim.Success(nil)

}

//Calculating Risk factor out of 10
func (t *SimpleChaincode) calculateRiscFactor(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 4, function followed by 2 names and 1 value")
	}
	pid := "policy1"
	parentCompanyBytes, _ := stub.GetState(pid)
	if parentCompanyBytes == nil {
		return shim.Error("Could not locate ")
	}

	parentCompany := &Policy{} //new Policy type
	json.Unmarshal(parentCompanyBytes, parentCompany)

	id := args[0]
	netGrowth, _ := strconv.Atoi(args[1])
	lastYearProfit, _ := strconv.Atoi(args[2])
	employeeExpen, _ := strconv.Atoi(args[3])
	stakeValue, _ := strconv.Atoi(args[4])

	childCompanyBytes, _ := stub.GetState(id)
	if childCompanyBytes == nil {
		return shim.Error("Could not locate ")
	}
	childCompany := &Child{}

	json.Unmarshal(childCompanyBytes, childCompany)

	currentRiskFactor := (netGrowth + lastYearProfit + stakeValue) / employeeExpen
	currentRiskFactor = currentRiskFactor * 10
	pTotRisk, _ := strconv.Atoi(parentCompany.TotalRisk)
	pTotRisk = pTotRisk + currentRiskFactor
	parentCompany.TotalRisk = strconv.Itoa(pTotRisk)
	childCompany.RiskFactor = strconv.Itoa(currentRiskFactor)

	childCompanyBytes, _ = json.Marshal(childCompany)
	err := stub.PutState(id, childCompanyBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to update risk factor"))
	}
	parentCompanyBytes, _ = json.Marshal(parentCompany)
	err = stub.PutState(pid, parentCompanyBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to update risk factor"))
	}
	return shim.Success(nil)
}

//Show the required policy tokens to buy
func (t *SimpleChaincode) showPolicyTokens(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	id := args[0]
	x:=0
	pid := "policy1"
	parentCompanyBytes, _ := stub.GetState(pid)
	if parentCompanyBytes == nil {
		return shim.Error("Could not locate ")
	}

	parentCompany := &Policy{} //policy type struct
	json.Unmarshal(parentCompanyBytes, parentCompany)

	childCompanyBytes, _ := stub.GetState(id)
	if childCompanyBytes == nil {
		return shim.Error("Could not locate ")
	}

	childCompany := &Child{}

	json.Unmarshal(childCompanyBytes, childCompany)

	tokens, _ := strconv.Atoi(childCompany.PolicyTokens) // Required policy tokens
	risk, _ := strconv.Atoi(childCompany.RiskFactor)
	totRisk, _ := strconv.Atoi(parentCompany.TotalRisk) // total parent company risk
	n, _ := strconv.Atoi(parentCompany.TotalTokens)     //Total tokens in hand of parent
	if tokens <= 0 {
		x = (risk * n) / totRisk
		totTokens = totTokens - x //child compnay total tokend reduced by x
		//n = n - x                 //parent total tokens reduced by x
		childCompany.PolicyPackage = strconv.Itoa(x)
		//parentCompany.TotalTokens = strconv.Itoa(n)
	}
	childCompanyBytes, _ = json.Marshal(childCompany)
	err := stub.PutState(id, childCompanyBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to update risk factor"))
	}
	return shim.Success(nil)			//return necessary tokens to purchase
}

//calculate Number of tokens for each child
func (t *SimpleChaincode) buyPolicy(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	id := args[0]
	pid := "policy1"
	parentCompanyBytes, _ := stub.GetState(pid)
	if parentCompanyBytes == nil {
		return shim.Error("Could not locate ")
	}

	parentCompany := &Policy{} //policy type struct
	json.Unmarshal(parentCompanyBytes, parentCompany)

	childCompanyBytes, _ := stub.GetState(id)
	if childCompanyBytes == nil {
		return shim.Error("Could not locate ")
	}

	childCompany := &Child{}

	json.Unmarshal(childCompanyBytes, childCompany)

	tokens, _ := strconv.Atoi(childCompany.PolicyTokens) // Required policy tokens
	totTokens, _ := strconv.Atoi(childCompany.Tokens)    // Total tokens in hand of child
	risk, _ := strconv.Atoi(childCompany.RiskFactor)
	n, _ := strconv.Atoi(parentCompany.TotalTokens)     //Total tokens in hand of parent
	totRisk, _ := strconv.Atoi(parentCompany.TotalRisk) // total parent company risk

	if tokens <= 0 {
		x := (risk * n) / totRisk
		totTokens = totTokens - x //child compnay total tokend reduced by x
		//n = n - x                 //parent total tokens reduced by x
		childCompany.PolicyTokens = strconv.Itoa(x)
		childCompany.Tokens = strconv.Itoa(totTokens)
		//parentCompany.TotalTokens = strconv.Itoa(n)
	}
	childCompanyBytes, _ = json.Marshal(childCompany)
	err := stub.PutState(id, childCompanyBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to update risk factor"))
	}
	parentCompanyBytes, _ = json.Marshal(parentCompany)
	err = stub.PutState(pid, parentCompanyBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to update risk factor"))
	}
	return shim.Success(nil)
}


//Policy claim function
func (t *SimpleChaincode) claim(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	id := args[0] //client id
	pid := "policy1"
	parentCompanyBytes, _ := stub.GetState(pid)
	if parentCompanyBytes == nil {
		return shim.Error("Could not locate ")
	}
	parentCompany := &Policy{}
	json.Unmarshal(parentCompanyBytes, parentCompany)

	childCompanyBytes, _ := stub.GetState(id)
	if childCompanyBytes == nil {
		return shim.Error("Could not locate ")
	}
	childCompany := &Child{}
	json.Unmarshal(childCompanyBytes, childCompany)

	claimtokens, _ := strconv.Atoi(args[1]) //claimed num of tokens
	newPolicyTokens, _ := strconv.Atoi(childCompany.PolicyTokens)
	newPolicyTokens = newPolicyTokens - claimtokens
	newTotTokens, _ := strconv.Atoi(parentCompany.TotalTokens)
	newTotTokens = newTotTokens + claimtokens
	childCompany.Tokens = strconv.Itoa(newPolicyTokens)
	parentCompany.TotalTokens = strconv.Itoa(newTotTokens)
	//Add to blockchain
	childCompanyBytes, _ = json.Marshal(childCompany)
	err := stub.PutState(id, childCompanyBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to update risk factor"))
	}
	parentCompanyBytes, _ = json.Marshal(parentCompany)
	err = stub.PutState(pid, parentCompanyBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to update risk factor"))
	}
	return shim.Success(nil)
}

// func (t *SimpleChaincode) move(stub shim.ChaincodeStubInterface, args []string) pb.Response {
// 	// must be an invoke
// 	var A, B string    // Entities
// 	var Aval, Bval int // Asset holdings
// 	var X int          // Transaction value
// 	var err error

// 	if len(args) != 3 {
// 		return shim.Error("Incorrect number of arguments. Expecting 4, function followed by 2 names and 1 value")
// 	}

// 	A = args[0]
// 	B = args[1]

// 	// Get the state from the ledger
// 	// TODO: will be nice to have a GetAllState call to ledger
// 	Avalbytes, err := stub.GetState(A)
// 	if err != nil {
// 		return shim.Error("Failed to get state")
// 	}
// 	if Avalbytes == nil {
// 		return shim.Error("Entity not found")
// 	}
// 	Aval, _ = strconv.Atoi(string(Avalbytes))

// 	Bvalbytes, err := stub.GetState(B)
// 	if err != nil {
// 		return shim.Error("Failed to get state")
// 	}
// 	if Bvalbytes == nil {
// 		return shim.Error("Entity not found")
// 	}
// 	Bval, _ = strconv.Atoi(string(Bvalbytes))

// 	// Perform the execution
// 	X, err = strconv.Atoi(args[2])
// 	if err != nil {
// 		return shim.Error("Invalid transaction amount, expecting a integer value")
// 	}
// 	Aval = Aval - X
// 	Bval = Bval + X
// 	logger.Infof("Aval = %d, Bval = %d\n", Aval, Bval)

// 	// Write the state back to the ledger
// 	err = stub.PutState(A, []byte(strconv.Itoa(Aval)))
// 	if err != nil {
// 		return shim.Error(err.Error())
// 	}

// 	err = stub.PutState(B, []byte(strconv.Itoa(Bval)))
// 	if err != nil {
// 		return shim.Error(err.Error())
// 	}

//         return shim.Success(nil);
// }

// // Deletes an entity from state
// func (t *SimpleChaincode) delete(stub shim.ChaincodeStubInterface, args []string) pb.Response {
// 	if len(args) != 1 {
// 		return shim.Error("Incorrect number of arguments. Expecting 1")
// 	}

// 	A := args[0]

// 	// Delete the key from the state in ledger
// 	err := stub.DelState(A)
// 	if err != nil {
// 		return shim.Error("Failed to delete state")
// 	}

// 	return shim.Success(nil)
// }

// // Query callback representing the query of a chaincode
func (t *SimpleChaincode) query(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting name of the person to query")
	}

	Id := args[0]

	// Get the state from the ledger
	databytes, err := stub.GetState(Id)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + Id + "\"}"
		return shim.Error(jsonResp)
	}
	// logger.Infof("Query Response:%s\n", jsonResp)
	return shim.Success(databytes)
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		logger.Errorf("Error starting Simple chaincode: %s", err)
	}
}
