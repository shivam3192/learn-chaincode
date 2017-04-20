
package main

import (
        "errors"
        "fmt"
       
        "encoding/json"
        "github.com/hyperledger/fabric/core/chaincode/shim"
)

// CrowdFundChaincode implementation
type CrowdFundChaincode struct {
}
type Info struct {

        Rollno []string   `json:"rollno"`
        Name []string `json:"name"`
        Sem  []string   `json:"sem"`
        Marks []string `json:"marks"`

}
//
// Init creates the state variable with name "account" and stores the value
// from the incoming request into this variable. We now have a key/value pair
// for account --> accountValue.
//
func (t *CrowdFundChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
       
        var err error

        if len(args) != 2 {
                return nil, errors.New("Incorrect number of arguments. Expecting 2.")
        }

   
     if err!=nil {
                        return nil, err
                }
         record := Info{}
       
        record.Rollno=append(record.Rollno,"MT2016001");
        record.Name=append(record.Name,"Aarushi");
        record.Sem=append(record.Sem,"Ist");
        record.Marks=append(record.Marks,"78");
        newrecordByte, err := json.Marshal(record);
        if err!=nil {

            return nil, err
        }
                err=stub.PutState("default",newrecordByte);
         if err!=nil {
                        return nil, err
                }



        return nil, nil
}


func (t *CrowdFundChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
    
var account string

fmt.Printf(" the function which has been recieved as input is : %s" , function)
fmt.Printf(" the function which has been recieved as input is : %s" , args[0])
fmt.Printf(" the function which has been recieved as input is : %s" , args[1])
fmt.Printf(" the function which has been recieved as input is : %s" , args[2])
fmt.Printf(" the function which has been recieved as input is : %s" , args[3])

        var err error

        if len(args) != 4 {
                return nil, errors.New("Incorrect number of arguments. Expecting 2.")
        }
          account = args[0]
          fmt.Printf(" key is : %s" , account)

         recordByte, err := stub.GetState(account);
        fmt.Println(recordByte);
        if err != nil {

            return nil, err
        }
        record := Info{}
        if recordByte != nil {
        errrecordmarshal := json.Unmarshal(recordByte,&record);
        fmt.Printf(" the unmarshall function output is : %s" , errrecordmarshal)

        if errrecordmarshal != nil {
            return nil, errrecordmarshal
        }    
               
        }
       
            
        record.Rollno = append(record.Rollno,args[0]);
        record.Name = append(record.Name,args[1]);
        record.Sem=append(record.Sem,args[2]);
        record.Marks=append(record.Marks,args[3]);

        fmt.Printf(" record structure rollno is : %s" , record.Rollno)
        fmt.Printf(" record structure rollno is : %s" , record.Name)
        fmt.Printf(" record structure rollno is : %s" , record.Sem)
        fmt.Printf(" record structure rollno is : : %s" , record.Marks)

        newrecordByte, err := json.Marshal(record);

        stringNewRecordByte := string(newrecordByte)

        fmt.Printf(" the marshall function output is : %s" , stringNewRecordByte)

        if err!=nil {

            return nil, err
        }
        err =stub.PutState(account,newrecordByte);
        if err != nil {

            return nil, err;
        } 
        return nil, nil
}



func (t *CrowdFundChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
  if function != "query" {
                return nil, errors.New("Invalid query function name. Expecting \"query\".")
        }

       var err error

         if len(args) != 1 {
                return nil, errors.New("Incorrect number of arguments. Expecting name of the state variable to query.")
        }

     var   account = args[0]
   
        accountValueBytes ,err := stub.GetState(account)
        if err != nil {
              
                 return nil, err
        }
    
        return accountValueBytes, nil
}

func main() {
        err := shim.Start(new(CrowdFundChaincode))

        if err != nil {
                fmt.Printf("Error starting CrowdFundChaincode: %s", err)
        }
}

