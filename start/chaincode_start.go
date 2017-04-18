package main

import (
	"errors"
	"fmt"

	"encoding/json"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)


type StudentChaincode struct {
}

//custom data models


type StudentApplication struct {
	ID                     string        `json:"id"`
	Name 				   string        `json:"name"`
	Marks   		       int           `json:"marks"`
	}

func GetStudentInfo(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Printf("Entering GetStudentInfo")

	if len(args) < 1 {
		fmt.Printf("Invalid number of arguments")
		return nil, errors.New("Missing student application ID")
	}

	var studentApplicationId = args[0]
	bytes, err := stub.GetState(studentApplicationId)
	if err != nil {
		fmt.Printf("Could not fetch student application with id "+studentApplicationId+" from ledger", err)
		return nil, err
	}
	return bytes, nil
}

func CreateStudentApplication(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Printf("Entering CreatestudentApplication")

	if len(args) < 2 {
		fmt.Printf("Invalid number of args")
		return nil, errors.New("Expected atleast two arguments for student application creation")
	}

	var studentApplicationId = args[0]
	var studentApplicationInput = args[1]

	err := stub.PutState(studentApplicationId, []byte(studentApplicationInput))
	if err != nil {
		("Could not save student application to ledger", err)
		return nil, err
	}

	
	fmt.Printf("Successfully saved student application")
	return nil, nil

}

/**
Updates the status of the student application
**/
/*
func UpdatestudentApplication(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Printf("Entering UpdatestudentApplication")

	if len(args) < 2 {
		fmt.Printf("Invalid number of args")
		return nil, errors.New("Expected atleast two arguments for student application update")
	}

	var studentApplicationId = args[0]
	var status = args[1]

	laBytes, err := stub.GetState(studentApplicationId)
	if err != nil {
		fmt.Printf("Could not fetch student application from ledger", err)
		return nil, err
	}
	var studentApplication StudentApplication
	err = json.Unmarshal(laBytes, &studentApplication)
	studentApplication.Status = status

	laBytes, err = json.Marshal(&studentApplication)
	if err != nil {
		fmt.Printf("Could not marshal student application post update", err)
		return nil, err
	}

	err = stub.PutState(studentApplicationId, laBytes)
	if err != nil {
		fmt.Printf("Could not save student application post update", err)
		return nil, err
	}

	
	fmt.Printf("Successfully updated student application")
	return nil, nil

}
*/


func (t *StudentChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	
	return nil, nil
}

func (t *StudentChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	if function == "GetStudentInfo" {
		return GetStudentInfo(stub, args)
	}
	return nil, nil
}



func (t *StudentChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	if function == "CreateStudentApplication" {
		return CreateStudentApplication(stub, args)
		} /*else {
			if(function == "init")
			return nil, errors.New(username + " does not have access to create a student application")
		}
*/
	}
	return nil, nil
}


func main() {

	err := shim.Start(new(StudentChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	} else {
		fmt.Printf"StudentChaincode successfully started")
	}

}
