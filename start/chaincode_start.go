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
//var index int
type StudentInfo struct {
        StudentRollNo     string  `json:"Studentrollno"`
        StudentName        string   `json:"Studentname"`
        StudentBadge       []string  `json:"Studentbadge"`
        StudentMarks       []string   `json:"Studentmarks"`
        StudentSem         []string   `json:"Studentsem"`
        IssuedBy        []string   `json:"Issuedby"`
        
}
type BadgeInfo struct {

        BadgeName       []string   `json:"Badgeame"`
        BadgeUrl        []string `json:"Badgeurl"`
        BadgeIssuedBy   []string   `json:"Badgeissuedby"`
        BadgeIssuedTo   []string `json:"Badgeissuedto"`
        //time 
}

type Issuer struct {

        IssuerInfo      []string   `json:"Issuerinfo"`
        IssuerName        string `json:"Issuername"`
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

   
  //   if err!=nil {
    //                    return nil, err
      //          }
        // record := StudentInfo{}
       
       /*record.StudentRollNo ="MT2916"
       record.StudentName ="aarushi"
        
        record.StudentRollNo=append(record.StudentRollNo,"MT2016001");
        record.StudentName=append(record.StudentName,"Aarushi");
        record.StudentBadge=append(record.StudentBadge,"Mtech");
        record.StudentMarks=append(record.StudentMarks,"78");
        record.StudentSem=append(record.StudentMarks,"1st");
        record.IssuedBy=append(record.IssuedBy,"RC Sir");
        
        newrecordByte, err := json.Marshal(record);
        if err!=nil {

            return nil, err
        }
                err=stub.PutState("default",newrecordByte);
         if err!=nil {
                        return nil, err
            }
*/
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
record := StudentInfo{}

        record.StudentRollNo=args[0];
        record.StudentName=args[1];
        record.StudentBadge=append(record.StudentBadge,args[2]);
        record.StudentMarks=append(record.StudentMarks,args[3]);
        record.StudentSem=append(record.StudentMarks,args[4]);
        record.SssuedBy=append(record.IssuedBy,args[5]);
            
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
        record := StudentInfo{}
        if recordByte != nil {
        errrecordmarshal := json.Unmarshal(recordByte,&record);
        fmt.Printf(" the unmarshall function output is : %s" , errrecordmarshal)

        if errrecordmarshal != nil {
            return nil, errrecordmarshal
        }    
               
        }
       


        record.StudentRollNo=args[0];
        record.StudentName=args[1];
        record.StudentBadge=append(record.StudentBadge,args[2]);
        record.StudentMarks=append(record.StudentMarks,args[3]);
        record.StudentSem=append(record.StudentMarks,args[4]);
        record.IssuedBy=append(record.IssuedBy,args[5]);
            

            
        /*record.Rollno = append(record.Rollno,args[0]);
        record.Name = append(record.Name,args[1]);
        record.Sem=append(record.Sem,args[2]);
        record.Marks=append(record.Marks,args[3]);
*/
        fmt.Printf(" record structure rollno is : %s" ,  record.StudentRollNo)
        fmt.Printf(" record structure name is   : %s" ,  record.StudentName)
        fmt.Printf(" record structure badge is : %s" ,   record.StudentBadge)
        fmt.Printf(" record structure marks is : : %s" , record.StudentMarks)
        fmt.Printf(" record structure sem is : %s" ,     record.StudentSem)
        fmt.Printf(" record structure issuedby is : %s" ,record.IssuedBy)
        


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
else {
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
 record := StudentInfo{}
        if accountValueBytes != nil {
        errrecordmarshal := json.Unmarshal(accountValueBytes,&record);
        fmt.Printf(" the unmarshall function output is : %s" , errrecordmarshal)

        if errrecordmarshal != nil {
            return nil, errrecordmarshal
        }    
               
        }

        
    
        return accountValueBytes, nil
    }
}

func main() {
        err := shim.Start(new(CrowdFundChaincode))

        if err != nil {
                fmt.Printf("Error starting CrowdFundChaincode: %s", err)
        }
}
