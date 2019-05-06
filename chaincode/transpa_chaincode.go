package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type DemandeChaincode struct {
}

type Investisseur struct {
	UserID string
}

type AssetManager struct {
	UserID string
}

type FundAccess struct {
	fundAccessID string
	Investisseur Investisseur
	Assetmanager AssetManager
	PortfolioID  string
	Status       string
}

type InventoryFile struct {
	inventoryFileID string
}

func (t *DemandeChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {

	fmt.Println("Init")
	return shim.Success([]byte("Init success"))
}

func (t *DemandeChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {

	fc, args := stub.GetFunctionAndParameters()
	if fc == "InitiateDemand" {
		return t.InitiateDemand(stub, args)
	}
	if fc == "getFundAccessByID" {
		return t.getFundAccessByID(stub, args)
	}
	if fc == "ChangeFundAccessStatus" {
		return t.ChangeFundAccessStatus(stub, args)
	}
	if fc == "CreateInventory" {
		return t.CreateInventory(stub, args)
	}

	return shim.Error("Called function is not defined in the chaincode ")
}

func (t *DemandeChaincode) InitiateDemand(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	fundAccessID := "fa:" + args[0]
	InvestisseurID := "i:" + args[1]
	ManagerID := "am:" + args[2]
	PortfolioID := "ptfid:" + args[3]

	fmt.Println("Initiate Demand")

	investisseur := Investisseur{UserID: InvestisseurID}
	investisseurAsJSONBytes, _ := json.Marshal(investisseur)
	_ = t.PutOnLedger(stub, investisseur.UserID, investisseurAsJSONBytes)

	assetmanager := AssetManager{UserID: ManagerID}
	assetmanagerAsJSONBytes, _ := json.Marshal(assetmanager)
	_ = t.PutOnLedger(stub, assetmanager.UserID, assetmanagerAsJSONBytes)

	fundaccess := FundAccess{fundAccessID: fundAccessID, Investisseur: investisseur, Assetmanager: assetmanager, PortfolioID: PortfolioID, Status: "pending"}
	fundAsJSONBytes, _ := json.Marshal(fundaccess)
	_ = t.PutOnLedger(stub, fundaccess.fundAccessID, fundAsJSONBytes)

	jsonResp := "{\"FundID\":\"" + fundaccess.fundAccessID + "\",\"DemandStatus\":\"" + fundaccess.Status + "\",\"Investisseur\":\"" + fundaccess.Investisseur.UserID + "\"}"
	fmt.Printf("Demand:%s\n", jsonResp)

	return shim.Success([]byte("Assets created successfully."))

}

func (t *DemandeChaincode) ChangeFundAccessStatus(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	fundAccessID := "fa:" + args[0]
	Decision := args[1]

	fmt.Println("Change Demand Status")

	fundAsJSONBytes := t.GetFromLedger(stub, fundAccessID)
	fundaccess := FundAccess{}

	_ = json.Unmarshal(fundAsJSONBytes.Payload, &fundaccess)
	fundaccess.Status = Decision

	jsonResp := "{\"FundID\":\"" + fundAccessID + "\",\"DemandStatus\":\"" + fundaccess.Status + "\",\"Investisseur\":\"" + fundaccess.Investisseur.UserID + "\"}"
	fmt.Printf("Query Response:%s\n", jsonResp)

	NewfundAsJSONBytes, _ := json.Marshal(fundaccess)
	_ = t.PutOnLedger(stub, fundaccess.fundAccessID, NewfundAsJSONBytes)

	return shim.Success([]byte("Asset modified."))
}

func (t *DemandeChaincode) getFundAccessByID(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	fundAccessID := "fa:" + args[0]

	fmt.Println("Query")

	fundAsJSONBytes := t.GetFromLedger(stub, fundAccessID)
	fundaccess := FundAccess{}
	_ = json.Unmarshal(fundAsJSONBytes.Payload, &fundaccess)

	jsonResp := "{\"FundID\":\"" + fundAccessID + "\",\"DemandStatus\":\"" + fundaccess.Status + "\",\"Investisseur\":\"" + fundaccess.Investisseur.UserID + "\"}"

	fmt.Printf("Query Response:%s\n", jsonResp)
	resultAsBytes, _ := json.Marshal(fundaccess)

	return shim.Success(resultAsBytes)
}

func (t *DemandeChaincode) CreateInventory(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	inventoryFileID := "if:" + args[0]
	inventoryFile := InventoryFile{inventoryFileID: inventoryFileID}
	inventoryAsJSONBytes, _ := json.Marshal(inventoryFile)
	_ = t.PutOnLedger(stub, inventoryFile.inventoryFileID, inventoryAsJSONBytes)

	return shim.Success([]byte("Assets created successfully."))
}

func (t *DemandeChaincode) PutOnLedger(stub shim.ChaincodeStubInterface, key string, value []byte) pb.Response {

	err := stub.PutState(key, value)
	if err != nil {
		return shim.Error("Failed to save asset " + key)
	}
	return shim.Success(nil)
}

func (t *DemandeChaincode) GetFromLedger(stub shim.ChaincodeStubInterface, key string) pb.Response {

	resAsJSONBytes, err := stub.GetState(key)
	if err != nil {
		return shim.Error("Failed to create asset " + key)
	}
	if resAsJSONBytes == nil {
		return shim.Error("Asset related to key" + key + "not found")
	}
	return shim.Success(resAsJSONBytes)
}

func main() {
	err := shim.Start(new(DemandeChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
