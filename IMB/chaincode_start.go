/*
Copyright IBM Corp 2016 All Rights Reserved.

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
	"errors"
	"fmt"
	"strconv"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}
type Per struct{
	ID int
 	Cuentas map[int]int
}

type IMB struct{
	name string 
	Clientes map[string]Per
}

var IMBS map[string]IMB

var Member1, Member2, Bet1, Bet2, win string


// ============================================================================================================================
// Main
// ============================================================================================================================
func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

// Init resets all the things
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Printf("Init called, initializing chaincode")
	

	if len(args) != 0 {
		return nil, errors.New("Incorrect number of arguments. Expecting 4")
	}

	IMBS = make(map[string]IMB)

	return nil, nil
}


func (t *SimpleChaincode) createIMB(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
    var name string

    fmt.Println("running createIMB")

    if len(args) != 1 {
        return nil, errors.New("Incorrect number of arguments. Expecting 2. name of the variable and value to set")
    }

    name = args[0]                            
    IMBS[name]= IMB{name, make(map[string]Per)}

    return nil, nil
}

func (t *SimpleChaincode) createCli(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
    var nameI, nameC string
    var IM IMB

    fmt.Println("running createCli")

    if len(args) != 2 {
        return nil, errors.New("Incorrect number of arguments. Expecting 2. name of the variable and value to set")
    }

    nameI = args[0]
    nameC = args[1]                            
    IM =IMBS[nameI]
    IM.Clientes[nameC]=Per{0, make(map[int]int)}

    return nil, nil
}

func (t *SimpleChaincode) createCuenta(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
 	var nameI, nameC string
    var IM IMB
    var P1 Per

    fmt.Println("running createCuenta")

    if len(args) != 2 {
        return nil, errors.New("Incorrect number of arguments. Expecting 2. name of the variable and value to set")
    }

    nameI = args[0]
    nameC = args[1]

    IM = IMBS[nameI]
    P1 = IM.Clientes[nameC]
    P1.Cuentas[P1.ID]=0

    return nil, nil
}

func (t *SimpleChaincode) LeerPer(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
    var nameI, nameC, val string
    var IM IMB
    var P1 Per

    if len(args) != 2 {
        return nil, errors.New("Incorrect number of arguments. Expecting 2. name of the variable and value to set")
    }

    nameI = args[0]
    nameC = args[1]

    IM = IMBS[nameI]
    P1 = IM.Clientes[nameC]
    val= nameC
    for k := range P1.Cuentas {
    	val= val + " ID: " +  strconv.Itoa(k) + " Monto: " +   strconv.Itoa(P1.Cuentas[k])
	}
 
    valAsbytes := []byte(val)
   

    return valAsbytes, nil
}






// Deletes an entity from state
func (t *SimpleChaincode) delete(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Printf("Running delete")
	
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 3")
	}

	A := args[0]

	// Delete the key from the state in ledger
	err := stub.DelState(A)
	if err != nil {
		return nil, errors.New("Failed to delete state")
	}

	return nil, nil
}

// Invoke is our entry point to invoke a chaincode function
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Printf("Invoke called, determining function")
	
	// Handle different functions
	if function == "init" {
		fmt.Printf("Function is init")
		return t.Init(stub, function, args)
	} else if function == "createIMB" {
		fmt.Printf("Function is createIMB")
		return t.createIMB(stub,  args)
	}else if function == "createCli" {
		fmt.Printf("Function is createCli")
		return t.createCli(stub,  args)
	}else if function == "createCuenta" {
		fmt.Printf("Function is createCuenta")
		return t.createCuenta(stub,  args)
	}


	/*if function == "transaction" {
		// Transaction makes payment of X units from A to B
		fmt.Printf("Function is Transaction")
		return t.transaction(stub, args)
	} else if function == "init" {
		fmt.Printf("Function is init")
		return t.Init(stub, function, args)
	} else if function == "delete" {
		// Deletes an entity from its state
		fmt.Printf("Function is delete")
		return t.delete(stub, args)	
	}else if function == "create" {
		// Deletes an entity from its state
		fmt.Printf("Function is create")
		return t.create(stub, args)
	}else if function == "setBet" {
		// Deletes an entity from its state
		fmt.Printf("Function is setBet")
		return t.setBet(stub, args)
	}else if function == "Bet" {
		// Deletes an entity from its state
		fmt.Printf("Function is Bet")
		return t.Bet(stub, args)
	}*/

	return nil, errors.New("Received unknown function invocation")
}


func (t *SimpleChaincode) write(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
    var name, value string
    //var Ivalue int
    var err error
    fmt.Println("running write()")

    if len(args) != 2 {
        return nil, errors.New("Incorrect number of arguments. Expecting 2. name of the variable and value to set")
    }

    name = args[0]                            //rename for fun
    value = args[1]
    err = stub.PutState(name, []byte(value))  //write the variable into the chaincode state
    if err != nil {
        return nil, err
    }
    return nil, nil
}




// Query is our entry point for queries
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
    fmt.Println("query is running " + function)


	if function == "LeerPer" {                            //read a variable
        return t.LeerPer(stub, args)
    }
    // Handle different functions
    /*if function == "read" {                            //read a variable
        return t.read(stub, args)
    }else if function == "readVar" {                            //read
        return t.readVar(stub, args)
    }
    fmt.Println("query did not find func: " + function)
*/
    return nil, errors.New("Received unknown function query")
}

func (t *SimpleChaincode) readVar(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
    
    valAsbytes := []byte(Member1 +" "+ Member2 +" "+ Bet1 +" "+ Bet2 +" "+ win)
   

    return valAsbytes, nil
}


//Agregado por mi
func (t *SimpleChaincode) read(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
    var name, jsonResp string
    var err error

    if len(args) != 1 {
        return nil, errors.New("Incorrect number of arguments. Expecting name of the var to query")
    }

    name = args[0]
    valAsbytes, err := stub.GetState(name)
    if err != nil {
        jsonResp = "{\"Error\":\"Failed to get state for " + name + "\"}"
        return nil, errors.New(jsonResp)
    }

    return valAsbytes, nil
}
