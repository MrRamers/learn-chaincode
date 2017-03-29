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
	
 	Cuentas map[int]int
}
var IDPer int
type IMB struct{
	name string 
	Clientes map[string]Per
}

var IMBS map[string]IMB


type Adress struct{
	Banco string 
	Cliente string
	Cuenta string
}

type SmartContract struct{
	ID int
	Origen Adress
	Destino Adress
	Estado string
	Monto string
	Mensaje string
}
var SCs map[int]SmartContract
var SCID int

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
	SCs = make(map[int]SmartContract)
	SCID=0

	IDPer=0


	/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	IMBS["HSBC"]= IMB{"HSBC", make(map[string]Per)}

	IMBS["HSBC"].Clientes["Gonzalo"]=Per{make(map[int]int)}
	IM := IMBS["HSBC"]
    P1 := IM.Clientes["Gonzalo"]
    P1.Cuentas[IDPer]=2000
    IDPer++

	IMBS["HSBC"].Clientes["Ramiro"]=Per{make(map[int]int)}
    P1 = IM.Clientes["Ramiro"]
    P1.Cuentas[IDPer]=1000
    IDPer++
    /////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////


	      
	return nil, nil
}


func (t *SimpleChaincode) createSC(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
    var nameI1, nameI2, nameC1, nameC2,nameA1,nameA2, monto, mensaje string 	
    var A1, A2 Adress   

	fmt.Println("running createSC")

	if len(args) != 8 {
		return nil, errors.New("Incorrect number of arguments. Expecting 3")
	}


    nameI1 = args[0]
    nameC1 = args[1]
    nameA1 = args[2]
    A1=Adress{nameI1,nameC1,nameA1}

	nameI2 = args[3]
    nameC2 = args[4]
    nameA2 = args[5]
    A2=Adress{nameI2,nameC2,nameA2}

    monto = args[6]
    mensaje = args[7]
                            

    SCs[SCID]= SmartContract{SCID,A1,A2,"En Proceso",monto,mensaje}
	SCID+=1

    return nil, nil
}

func (t *SimpleChaincode) exjecutarSC(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {  
    //var A []string
    var nameI1, nameI2, nameC1, nameC2 string
 	var nameA1,nameA2, value  int
    var IM IMB
    var P1, P2 Per
    var err error


	fmt.Println("running createSC")

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 3")
	}

	ID, err:= strconv.Atoi(args[0])
	if err != nil {
		return nil, errors.New("Failed to get SmartContract")
	}
	SC := SCs[ID]
    
    nameI1 = SC.Origen.Banco
    nameC1 = SC.Origen.Cliente
    nameA1, err = strconv.Atoi(SC.Origen.Cuenta)

    IM = IMBS[nameI1]
    P1 = IM.Clientes[nameC1]

	nameI2 = SC.Destino.Banco
    nameC2 = SC.Destino.Cliente
    nameA2, err = strconv.Atoi(SC.Destino.Cuenta)


    IM = IMBS[nameI2]
    P2 = IM.Clientes[nameC2]

    value, err = strconv.Atoi(SC.Monto)

    if err != nil {
		return nil, errors.New("Failed to get int")
	}

	P1.Cuentas[nameA1]-=value
	P2.Cuentas[nameA2]+=value

    SC.Estado="Finalizado"                            
	

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
    IM.Clientes[nameC]=Per{make(map[int]int)}

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
    P1.Cuentas[IDPer]=0
    IDPer++

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
	}else if function == "createSC" {
		fmt.Printf("Function is createSC")
		return t.createSC(stub,  args)
	}else if function == "exjecutarSC" {
		fmt.Printf("Function is exjecutarSC")
		return t.exjecutarSC(stub,  args)
	}

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

	//val = strconv.Itoa(P1.ID) +"-----"+ val
 
    valAsbytes := []byte(val)
   

    return valAsbytes, nil
}

func (t *SimpleChaincode) LeerSC(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
    var  val string

    if len(args) != 1 {
        return nil, errors.New("Incorrect number of arguments. Expecting 2. name of the variable and value to set")
    }

	ID, err:= strconv.Atoi(args[0])
	if err != nil {
		return nil, errors.New("Failed to get SmartContract")
	}
    SC := SCs[ID]

    val = "ID: " + args[0] + "`"
    val = val + " Banco Origen: " + SC.Origen.Banco + "`"
    val = val + " Cliente Origen: " +SC.Origen.Cliente + "`"
    val = val + " Cuenta Origen: " +SC.Origen.Cuenta + "`"

	val = val + " Banco Destino: " +SC.Destino.Banco+ "`"
    val = val + " Cliente Destino: " +SC.Destino.Cliente+ "`"
    val = val + " Cuenta Destino: " +SC.Destino.Cuenta  + "`"

   	val = val + " Monto: " +SC.Monto+ "`"

   	val = val + " Estado: " +SC.Estado+ "`"

   	val = val + " Mensaje: " +SC.Mensaje+ "`"

    valAsbytes := []byte(val)
   

    return valAsbytes, nil
}


// Query is our entry point for queries
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
    fmt.Println("query is running " + function)


	if function == "LeerPer" {                            //read a variable
        return t.LeerPer(stub, args)
    }else if function == "LeerSC" {                            //read
        return t.LeerSC(stub, args)
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

