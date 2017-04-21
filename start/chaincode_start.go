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
var index int
type studentInfo struct {
        studentRollNo     string  `json:"studentrollno"`
        studentName        string   `json:"studentname"`
        studentBadge       []string  `json:"studentbadge"`
        studentMarks       []string   `json:"studentmarks"`
        studentSem         []string   `json:"studentsem"`
        issuedBy        []string   `json:"issuedby"`
        
}
type BadgeInfo struct {

        badgeName       []string   `json:"rollno"`
        badgeUrl        []string `json:"name"`
        badgeIssuedBy   []string   `json:"sem"`
        badgeIssuedTo   []string `json:"marks"`
        //time 
}

type Issuer struct {

        issuerInfo      []string   `json:"rollno"`
        issuerName        string `json:"name"`
       // time            string   `json:"sem"`
        
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
         record := studentInfo{}
       
       record.studentRollNo ="MT2916"
       record.studentName ="aarushi"
        
        //record.studentRollNo=append(record.studentRollNo,"MT2016001");
        //record.studentName=append(record.studentName,"Aarushi");
        record.studentBadge=append(record.studentBadge,"Mtech");
        record.studentMarks=append(record.studentMarks,"78");
        record.studentSem=append(record.studentMarks,"1st");
        record.issuedBy=append(record.issuedBy,"RC Sir");
        
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
    if function == "write" {


fmt.Printf("-----------------------------------inside write function-------------------------------------------------------------");

var account string

fmt.Printf(" the function which has been recieved as input is : %s" , function)
fmt.Printf(" the function which has been recieved as input is : %s" , args[0])
fmt.Printf(" the function which has been recieved as input is : %s" , args[1])
fmt.Printf(" the function which has been recieved as input is : %s" , args[2])
fmt.Printf(" the function which has been recieved as input is : %s" , args[3])

        var err error

        if len(args) != 6 {
                return nil, errors.New("Incorrect number of arguments. Expecting 6.")
        }
          account = args[0]//got the roll no
          fmt.Printf(" key is : %s" , account)
record := studentInfo{}

        record.studentRollNo=args[0];
        record.studentName=args[1];
        record.studentBadge=append(record.studentBadge,args[2]);
        record.studentMarks=append(record.studentMarks,args[3]);
        record.studentSem=append(record.studentMarks,args[4]);
        record.issuedBy=append(record.issuedBy,args[5]);
            
            newrecordByte, err := json.Marshal(record);


 if err!=nil {

            return nil, err
        }
        err =stub.PutState(account,newrecordByte);
        if err != nil {

            return nil, err;
        } 
        return nil, nil



 } else {
if (function == "update") {
fmt.Printf("-----------------------------------inside update function-------------------------------------------------------------");

var account string

fmt.Printf(" the function which has been recieved as input is : %s" , function)
fmt.Printf(" the function which has been recieved as input is : %s" , args[0])
fmt.Printf(" the function which has been recieved as input is : %s" , args[1])
fmt.Printf(" the function which has been recieved as input is : %s" , args[2])
fmt.Printf(" the function which has been recieved as input is : %s" , args[3])

        var err error

        if len(args) != 6 {
                return nil, errors.New("Incorrect number of arguments. Expecting 6.")
        }
          account = args[0]//got the roll no
          fmt.Printf(" key is : %s" , account)

         recordByte, err := stub.GetState(account);
        fmt.Println(recordByte);
        if err != nil {

            return nil, err
        }
        record := studentInfo{}
        if recordByte != nil {
        errrecordmarshal := json.Unmarshal(recordByte,&record);
        fmt.Printf(" the unmarshall function output is : %s" , errrecordmarshal)

        if errrecordmarshal != nil {
            return nil, errrecordmarshal
        }    
               
        }
       


        record.studentRollNo=args[0];
        record.studentName=args[1];
        record.studentBadge=append(record.studentBadge,args[2]);
        record.studentMarks=append(record.studentMarks,args[3]);
        record.studentSem=append(record.studentMarks,args[4]);
        record.issuedBy=append(record.issuedBy,args[5]);
            

            
        /*record.Rollno = append(record.Rollno,args[0]);
        record.Name = append(record.Name,args[1]);
        record.Sem=append(record.Sem,args[2]);
        record.Marks=append(record.Marks,args[3]);
*/
        fmt.Printf(" record structure rollno is : %s" ,  record.studentRollNo)
        fmt.Printf(" record structure name is   : %s" ,  record.studentName)
        fmt.Printf(" record structure badge is : %s" ,   record.studentBadge)
        fmt.Printf(" record structure marks is : : %s" , record.studentMarks)
        fmt.Printf(" record structure sem is : %s" ,     record.studentSem)
        fmt.Printf(" record structure issuedby is : %s" ,record.issuedBy)
        


        newrecordByte, err := json.Marshal(record);//result comes in bytes

        stringNewRecordByte := string(newrecordByte)

        fmt.Printf(" the marshall function output is : %s" , stringNewRecordByte)

        if err!=nil {

            return nil, err
        }
        err =stub.PutState(account,newrecordByte);
        if err != nil {

            return nil, err;
        } 
}
}
        return nil, nil

}


func (t *CrowdFundChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
  if function != "read" {
                return nil, errors.New("Invalid query function name. Expecting \"query\".")
        }

        fmt.Printf("-----------------------------------inside read function-------------------------------------------------------------");


       var err error

         if len(args) != 1 {
                return nil, errors.New("Incorrect number of arguments. Expecting name of the state variable to query.")
        }

     var   account = args[0]


      fmt.Printf("----------------------------inside read function--------value of account is  %s--------------------------------, account");

   
        accountValueBytes ,err := stub.GetState(account)

if err != nil {
              
                 return nil, err
        }
 record := studentInfo{}
        if accountValueBytes != nil {
        errrecordmarshal := json.Unmarshal(accountValueBytes,&record);
        fmt.Printf(" the unmarshall function output is : %s" , errrecordmarshal)

        if errrecordmarshal != nil {
            return nil, errrecordmarshal
        }    
               
        }

        
    
        return accountValueBytes, nil
}

func main() {
        err := shim.Start(new(CrowdFundChaincode))

        if err != nil {
                fmt.Printf("Error starting CrowdFundChaincode: %s", err)
        }
}
