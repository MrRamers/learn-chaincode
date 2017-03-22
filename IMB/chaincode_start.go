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
    P1.ID++

    return nil, nil
}

func (t *SimpleChaincode) acreditarCuentaCliente(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
 	var nameI, nameC string
 	var nameA, value int
    var IM IMB
    var P1 Per
    var err error

    fmt.Println("running acreditarCuentaCliente")

    if len(args) != 4 {
        return nil, errors.New("Incorrect number of arguments. Expecting 2. name of the variable and value to set")
    }

    nameI = args[0]
    nameC = args[1]
    nameA, err = strconv.Atoi(args[2])
    value, err = strconv.Atoi(args[3])
		
	if err != nil {
		return nil, errors.New("Failed to get int")
	}

    IM = IMBS[nameI]
    P1 = IM.Clientes[nameC]

    P1.Cuentas[nameA]+=value

    return nil, nil
}

func (t *SimpleChaincode) transaction(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Printf("Running Transaction")
	
 	var nameI1, nameI2, nameC1, nameC2 string
 	var nameA1,nameA2, value  int
    var IM IMB
    var P1, P2 Per
    var err error

	if len(args) != 7 {
		return nil, errors.New("Incorrect number of arguments. Expecting 3")
	}


    nameI1 = args[0]
    nameC1 = args[1]
    nameA1, err = strconv.Atoi(args[2])

    IM = IMBS[nameI1]
    P1 = IM.Clientes[nameC1]

	nameI2 = args[3]
    nameC2 = args[4]
    nameA2, err = strconv.Atoi(args[5])


    IM = IMBS[nameI2]
    P2 = IM.Clientes[nameC2]

    value, err = strconv.Atoi(args[6])

    if err != nil {
		return nil, errors.New("Failed to get int")
	}

	P1.Cuentas[nameA1]-=value
	P2.Cuentas[nameA2]+=value

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
	}else if function == "acreditarCuentaCliente" {
		fmt.Printf("Function is acreditarCuentaCliente")
		return t.acreditarCuentaCliente(stub,  args)
	}else if function == "transaction" {
		fmt.Printf("Function is transaction")
		return t.transaction(stub,  args)
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
