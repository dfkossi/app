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
	"testing"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func checkInit(t *testing.T, stub *shim.MockStub, args [][]byte) {
	res := stub.MockInit("1", args)

	if res.Status != shim.OK {
		fmt.Println("Init failed", string(res.Message))
		t.FailNow()
	}

}

func checkInvoke(t *testing.T, stub *shim.MockStub, args [][]byte) {
	res := stub.MockInvoke("1", args)
	if res.Status != shim.OK {
		fmt.Println("Invoke", args, "failed", string(res.Message))
		t.FailNow()
	}
}

func checkQuery(t *testing.T, stub *shim.MockStub, name string, value string) {
	res := stub.MockInvoke("1", [][]byte{[]byte("query"), []byte(name)})
	if res.Status != shim.OK {
		fmt.Println("Query", name, "failed", string(res.Message))
		t.FailNow()
	}
	if res.Payload == nil {
		fmt.Println("Query", name, "failed to get value")
		t.FailNow()
	}
	var Result string
	_ = json.Unmarshal(res.Payload, &Result)
	if string(Result) != value {

		fmt.Println("Query value", name, "was not", value, "as expected")
		t.FailNow()
	}
}

/* func checkChangeAccess(t *testing.T, stub *shim.MockStub, PortfolioID string, InvId string, Status string) {
	res := stub.MockInvoke("1", [][]byte{[]byte("ChangeStatus"), []byte(PortfolioID), []byte(InvId), []byte(Status)})
	if res.Status != shim.OK {
		fmt.Println("ChangeStatus", PortfolioID, "failed", string(res.Message))
		t.FailNow()
	}
	if res.Payload == nil {
		fmt.Println("ChangeStatus", PortfolioID, "failed to get value")
		t.FailNow()
	}

} */

func checkInitiateDemand(t *testing.T, stub *shim.MockStub, uuid string, name string, prettyName string) {

	res := stub.MockInvoke("1", [][]byte{[]byte("InitiateDemand"), []byte(uuid), []byte(name), []byte(prettyName)})
	if res.Status != shim.OK {
		fmt.Println("Demand not initiated from ", uuid, "for", name)
		t.FailNow()
	}
	if res.Payload == nil {
		fmt.Println("ChangeStatus", name, "failed to get value")
		t.FailNow()
	}
}

func TestExample(t *testing.T) {
	scc := new(DemandeChaincode)
	stub := shim.NewMockStub("ex02", scc)

	fmt.Println("Yes 1")
	checkInit(t, stub, [][]byte{[]byte("init")})
	fmt.Println("Yes 2")
	checkInitiateDemand(t, stub, "1", "DAD", "DADX")
	fmt.Println("Yes 3")
	checkQuery(t, stub, "1", "DAD")
	fmt.Println("Yes 4")
}

/*
func checkState(t *testing.T, stub *shim.MockStub, name string, value string) {
	bytes := stub.State[name]
	if bytes == nil {
		fmt.Println("State", name, "failed to get value")
		t.FailNow()
	}
	if string(bytes) != value {
		fmt.Println("State value", name, "was not", value, "as expected")
		t.FailNow()
	}
}


func TestExample02_Init(t *testing.T) {
	scc := new(SimpleChaincode)
	stub := shim.NewMockStub("ex02", scc)

	// Init A=123 B=234
	checkInit(t, stub, [][]byte{[]byte("init"), []byte("A"), []byte("123"), []byte("B"), []byte("234")})

	checkState(t, stub, "A", "123")
	checkState(t, stub, "B", "234")
}

func TestExample02_Query(t *testing.T) {
	scc := new(SimpleChaincode)
	stub := shim.NewMockStub("ex02", scc)

	// Init A=345 B=456
	checkInit(t, stub, [][]byte{[]byte("init"), []byte("A"), []byte("345"), []byte("B"), []byte("456")})

	// Query A
	checkQuery(t, stub, "A", "345")

	// Query B
	checkQuery(t, stub, "B", "456")
}


*/
