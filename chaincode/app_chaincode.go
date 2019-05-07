package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type DemandeChaincode struct {
}

type Organization struct {
	uuid               string
	name               string
	prettyname         string
	typeOfOrganization string
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
	if fc == "getOrganizationByID" {
		return t.getOrganizationByID(stub, args)
	}
	/* if fc == "ChangeFundAccessStatus" {
		return t.ChangeFundAccessStatus(stub, args)
	}
	if fc == "CreateInventory" {
		return t.CreateInventory(stub, args)
	} */

	return shim.Error("Called function is not defined in the chaincode ")
}

func (t *DemandeChaincode) InitiateDemand(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	/*
		uuid := "uuid:" + args[0]
		name := "name:" + args[1]
		prettyname := "pretty:" + args[2]
		typeOfOrganization := "type:" + args[3] */

	fmt.Println("Initiate Demand")

	organization := Organization{uuid: "org1", name: "", prettyname: "", typeOfOrganization: ""}
	organizationAsJSONBytes, _ := json.Marshal(organization)
	_ = t.PutOnLedger(stub, organization.uuid, organizationAsJSONBytes)

	jsonResp := "{\"OrganisationUUID\":\"" + organization.uuid + "\",\"Name\":\"" + organization.name + "\",\"PrettyName\":\"" + organization.prettyname + "\",\"Type\":\"" + organization.typeOfOrganization + "\"}"

	fmt.Printf("Demand:%s\n", jsonResp)

	return shim.Success([]byte("Assets created successfully."))

}

/* func (t *DemandeChaincode) ChangeFundAccessStatus(stub shim.ChaincodeStubInterface, args []string) pb.Response {

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
} */

func (t *DemandeChaincode) getOrganizationByID(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	uuid := "Org:" + args[0]

	fmt.Println("Query")
	organizationAsJSONBytes := t.GetFromLedger(stub, uuid)
	organization := Organization{}
	_ = json.Unmarshal(organizationAsJSONBytes.Payload, &organization)

	jsonResp := "{\"OrganisationUUID\":\"" + organization.uuid + "\",\"Name\":\"" + organization.name + "\",\"PrettyName\":\"" + organization.prettyname +
		"\",\"Type\":\"" + organization.typeOfOrganization + "\"}"

	fmt.Printf("Query Response:%:\n", jsonResp)
	resultAsBytes, _ := json.Marshal(organization)

	return shim.Success(resultAsBytes)
}

/* func (t *DemandeChaincode) getFundAccessByID(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	fundAccessID := "fa:" + args[0]

	fmt.Println("Query")

	fundAsJSONBytes := t.GetFromLedger(stub, fundAccessID)
	fundaccess := FundAccess{}
	_ = json.Unmarshal(fundAsJSONBytes.Payload, &fundaccess)

	jsonResp := "{\"FundID\":\"" + fundAccessID + "\",\"DemandStatus\":\"" + fundaccess.Status + "\",\"Investisseur\":\"" + fundaccess.Investisseur.UserID + "\"}"

	fmt.Printf("Query Response:%s\n", jsonResp)
	resultAsBytes, _ := json.Marshal(fundaccess)

	return shim.Success(resultAsBytes)
} */

/* func (t *DemandeChaincode) CreateInventory(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	inventoryFileID := "if:" + args[0]
	inventoryFile := InventoryFile{inventoryFileID: inventoryFileID}
	inventoryAsJSONBytes, _ := json.Marshal(inventoryFile)
	_ = t.PutOnLedger(stub, inventoryFile.inventoryFileID, inventoryAsJSONBytes)

	return shim.Success([]byte("Assets created successfully."))
} */

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
