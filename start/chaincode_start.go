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

/*type Persona struct {
}*/

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
    if len(args) != 0 {
        return nil, errors.New("Incorrect number of arguments. Expecting 1")
    }

    err := stub.PutState("RamiroPombo", []byte("2000"))
    err = stub.PutState("GonzaloVarilla", []byte("2000"))
    if err != nil {
        return nil, err
    }

    return nil, nil
}

// Invoke is our entry point to invoke a chaincode function
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
    fmt.Println("invoke is running " + function)

    // Handle different functions
    if function == "init" {
        return t.Init(stub, "init", args)
    } else if function == "write" {
        return t.write(stub, args)
    }else if function == "restar" {
        return t.restar(stub, args)
    }else if function == "transferir" {
        return t.transferir(stub, args)
    }
    fmt.Println("invoke did not find func: " + function)

    return nil, errors.New("Received unknown function invocation")
}

//Agregado para el test
func (t *SimpleChaincode) transferir(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var a1, a2 []string
	var aux int
	var err error
	if len(args) != 3 {
        return nil, errors.New("Incorrect number of arguments. Expecting name of the var to query")
    }
    a1[0]=args[0]
    a1[1]=args[2]

    a2[0]=args[1]
    aux, err = strconv.Atoi(args[2])
    aux *=-1
    a2[1]=strconv.Itoa(aux)
    t.restar(stub,a1)
    t.restar(stub,a2)
    if err != nil {
        return nil, err
    }
    return nil, nil
}


func (t *SimpleChaincode) restar(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
    var name, jsonResp, valor string
    var Ivalor, sust int
    var err error

    if len(args) != 2 {
        return nil, errors.New("Incorrect number of arguments. Expecting name of the var to query")
    }

    name = args[0]
    valAsbytes, err := stub.GetState(name)
    valor = string(valAsbytes)
    Ivalor, err = strconv.Atoi(valor)
    sust, err=strconv.Atoi(args[1])
    Ivalor -=sust
    valor= strconv.Itoa(Ivalor)
    args[1]=valor
    //valAsbytes = []byte(strconv.Itoa(Ivalor))
    if err != nil {
        jsonResp = "{\"Error\":\"Failed to get state for " + name + "\"}"
        return nil, errors.New(jsonResp)
    }
    return t.write(stub, args)
}

// Agregado por mi
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
    //Ivalue, err = strconv.Atoi(value)
    //Ivalue -= 100
    //value= strconv.Itoa(Ivalue)
    /*num, err = string(stub.GetState(name))
    Ivalue, err = strconv.Atoi(value)
    Inum, err = strconv.Atoi(num)*/
    /*resta= Inum - Ivalue
    value= strconv.Itoa(resta)*/
    err = stub.PutState(name, []byte(value))  //write the variable into the chaincode state
    if err != nil {
        return nil, err
    }
    return nil, nil
}


// Query is our entry point for queries
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
    fmt.Println("query is running " + function)

    // Handle different functions
    if function == "read" {                            //read a variable
        return t.read(stub, args)
    }
    fmt.Println("query did not find func: " + function)

    return nil, errors.New("Received unknown function query")
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