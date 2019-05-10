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
	Uuid               string
	Name               string
	PrettyName         string
	TypeOfOrganization string
	//Admin              Administrators
	AdminList  []Administrators
	DeviceList []Device
}

type Administrators struct {
	EmployeeID       string
	OrganizationName Organization
}

type PatchFile struct {
	Uuid       string
	Name       string
	PrettyName string
	fileURL    string
	fileHash   string
	version    string
}

type Patch struct {
	Uuid             string
	Name             string
	PrettyName       string
	ReleaseDate      string
	Status           string
	ReleaseURL       string
	Version          int
	Patches          []PatchFile
	OrganizationName Organization
}

type ConfigurationFile struct {
	ReleaseDate string
	Status      string
	FileURL     string
	FileHash    string
	Version     string
}

type Configuration struct {
	ReleaseDate        string
	Status             string
	ConfigurationFiles []ConfigurationFile
}

type DeviceClass struct {
	Uuid       string
	Name       string
	PrettyName string
	PatchList  []Patch
}

type Device struct {
	Uuid       string
	Name       string
	PrettyName string

	DeviceClassID        string
	Cpe                  string
	CurrentPatch         Patch
	EUCitizenID          string
	Provider             Organization
	CurrentConfiguration Configuration
	ConfigurationHistory []Configuration
}

type DeviceRights struct {
	DeviceID         Device
	OrganizationID   Organization
	Right            string
	DataSharingLevel string
}

type DeviceLog struct {
	Uuid       string
	Name       string
	PrettyName string

	DeviceID     string
	Format       string
	DateStart    string
	DateEnd      string
	DateOfficial string
	StorageURL   string
	Hash         string
}

type Evidence struct {
	Uuid       string
	Name       string
	PrettyName string

	EvidenceDate          string
	EvidenceURL           string
	TargetedVulnerability string
	//TypeOfAttack
}

type EvidenceFile struct {
	Uuid                     string
	Name                     string
	PrettyName               string
	AttackStatus             string
	EvidenceFileCreationDate string
	//EvidenceFileInformation
	EvidenceFileDataSourceTitle string
}

type OSEvidenceFile struct {
	EvidenceFile
	OSName    string
	OSVersion string
}

type SoftwareEvidenceFile struct {
	EvidenceFile
	SoftwareName    string
	SoftwareVersion string
	SoftwareType    string
}

type NetworkDiagnosticEvidenceFile struct {
	EvidenceFile
	DiagnosticApplicationName    string
	DiagnosticApplicationVersion string
	DiagnosticApplicationResult  string
}

type IncidentTrackingEvidenceFile struct {
	EvidenceFile
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
	if fc == "query" {
		return t.query(stub, args)
	}
	/* if fc == "getOrganizationByID" {
		return t.getOrganizationByID(stub, args)
	}
	if fc == "ChangeFundAccessStatus" {
		return t.ChangeFundAccessStatus(stub, args)
	}
	if fc == "CreateInventory" {
		return t.CreateInventory(stub, args)
	} */

	return shim.Error("Called function is not defined in the chaincode ")
}

func (t *DemandeChaincode) InitiateDemand(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	Uuid := args[0]
	Name := args[1]
	PrettyName := args[2]
	//TypeOfOrganization := args[3]

	//uuidAsJSONBytes, _ := json.Marshal
	fmt.Println("Initiate Demand 1")

	organization := Organization{Uuid: Uuid, Name: Name, PrettyName: PrettyName, TypeOfOrganization: "XXX"}
	organizationAsJSONBytes, _ := json.Marshal(organization)
	_ = t.PutOnLedger(stub, organization.Uuid, organizationAsJSONBytes)

	jsonResp := "{\"OrganisationUUID\":\"" + organization.Uuid + "\",\"Name\":\"" + organization.Name + "\",\"PrettyName\":\"" + organization.PrettyName + "\",\"Type\":\"" + organization.TypeOfOrganization + "\"}"

	fmt.Printf("Demand:%s\n", jsonResp)

	return shim.Success([]byte("Assets created successfully."))

}

func (t *DemandeChaincode) query(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	uuid := args[0]
	//InvestisseurID := args[1]
	fmt.Println("uuid:", uuid)
	fmt.Println("Query that")
	organizationAsJSONBytes := t.GetFromLedger(stub, uuid)
	fmt.Println(organizationAsJSONBytes)
	organization := Organization{}
	_ = json.Unmarshal(organizationAsJSONBytes.Payload, &organization)

	jsonResp := "{\"Name\":\"" + organization.Name + "\",\"OrganisationUUID\":\"" + organization.Uuid + "\",\"PrettyName\":\"" + organization.PrettyName + "\",\"Type\":\"" + organization.TypeOfOrganization + "\"}"

	fmt.Printf("Query Response:%s\n", jsonResp)
	resultAsBytes, _ := json.Marshal(organization.Name)

	return shim.Success(resultAsBytes)
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

/* func (t *DemandeChaincode) getOrganizationByID(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	uuid := args[0]

	fmt.Println("Query")
	organizationAsJSONBytes := t.GetFromLedger(stub, uuid)
	organization := Organization{}
	_ = json.Unmarshal(organizationAsJSONBytes.Payload, &organization)

	jsonResp := "{\"OrganisationUUID\":\"" + organization.uuid + "\",\"Name\":\"" + organization.name + "\",\"PrettyName\":\"" + organization.prettyName + "\",\"Type\":\"" + organization.typeOfOrganization + "\"}"

	fmt.Printf("Query Response:%s\n", jsonResp)
	resultAsBytes, _ := json.Marshal(organization)

	return shim.Success(resultAsBytes)
} */

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
