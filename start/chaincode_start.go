package main

import (
	"errors"
	"fmt"

	"encoding/json"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

var logger = shim.NewLogger("mylogger")

type StudentChaincode struct {
}

//custom data models


type StudentApplication struct {
	ID                     string        `json:"id"`
	Name 	               string        `json:"name"`
	Marks                  int           `json:"marks"`
	}

func GetStudentInfo(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	logger.Debug("Entering GetStudentInfo")

	if len(args) < 1 {
		logger.Error("Invalid number of arguments")
		return nil, errors.New("Missing student application ID")
	}

	var studentApplicationId = args[0]
	bytes, err := stub.GetState(studentApplicationId)
	if err != nil {
		logger.Error("Could not fetch student application with id "+studentApplicationId+" from ledger", err)
		return nil, err
	}
	return bytes, nil
}

func CreatestudentApplication(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	logger.Debug("Entering CreatestudentApplication")

	if len(args) < 2 {
		logger.Error("Invalid number of args")
		return nil, errors.New("Expected atleast two arguments for student application creation")
	}

	var studentApplicationId = args[0]
	var studentApplicationInput = args[1]

	err := stub.PutState(studentApplicationId, []byte(studentApplicationInput))
	if err != nil {
		logger.Error("Could not save student application to ledger", err)
		return nil, err
	}

	var customEvent = "{eventType: 'studentApplicationCreation', description:" + studentApplicationId + "' Successfully created'}"
	err = stub.SetEvent("evtSender", []byte(customEvent))
	if err != nil {
		return nil, err
	}
	logger.Info("Successfully saved student application")
	return nil, nil

}

/**
Updates the status of the student application
**/
func UpdatestudentApplication(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	logger.Debug("Entering UpdatestudentApplication")

	if len(args) < 2 {
		logger.Error("Invalid number of args")
		return nil, errors.New("Expected atleast two arguments for student application update")
	}

	var studentApplicationId = args[0]
	var status = args[1]

	laBytes, err := stub.GetState(studentApplicationId)
	if err != nil {
		logger.Error("Could not fetch student application from ledger", err)
		return nil, err
	}
	var studentApplication StudentApplication
	err = json.Unmarshal(laBytes, &studentApplication)
	studentApplication.Status = status

	laBytes, err = json.Marshal(&studentApplication)
	if err != nil {
		logger.Error("Could not marshal student application post update", err)
		return nil, err
	}

	err = stub.PutState(studentApplicationId, laBytes)
	if err != nil {
		logger.Error("Could not save student application post update", err)
		return nil, err
	}

	var customEvent = "{eventType: 'studentApplicationUpdate', description:" + studentApplicationId + "' Successfully updated status'}"
	err = stub.SetEvent("evtSender", []byte(customEvent))
	if err != nil {
		return nil, err
	}
	logger.Info("Successfully updated student application")
	return nil, nil

}

func (t *StudentChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	return nil, nil
}

func (t *StudentChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	if function == "GetStudentInfo" {
		return GetStudentInfo(stub, args)
	}
	return nil, nil
}

func GetCertAttribute(stub shim.ChaincodeStubInterface, attributeName string) (string, error) {
	logger.Debug("Entering GetCertAttribute")
	attr, err := stub.ReadCertAttribute(attributeName)
	if err != nil {
		return "", errors.New("Couldn't get attribute " + attributeName + ". Error: " + err.Error())
	}
	attrString := string(attr)
	return attrString, nil
}

func (t *StudentChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	if function == "CreatestudentApplication" {
		username, _ := GetCertAttribute(stub, "username")
		role, _ := GetCertAttribute(stub, "role")
		if role == "Bank_Home_student_Admin" {
			return CreatestudentApplication(stub, args)
		} else {
			return nil, errors.New(username + " with role " + role + " does not have access to create a student application")
		}

	}
	return nil, nil
}

type customEvent struct {
	Type       string `json:"type"`
	Decription string `json:"description"`
}

func main() {

	lld, _ := shim.LogLevel("DEBUG")
	fmt.Println(lld)

	logger.SetLevel(lld)
	fmt.Println(logger.IsEnabledFor(lld))

	err := shim.Start(new(StudentChaincode))
	if err != nil {
		logger.Error("Could not start StudentChaincode")
	} else {
		logger.Info("StudentChaincode successfully started")
	}

}
