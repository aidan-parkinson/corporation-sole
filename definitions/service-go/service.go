package main

import (
  "fmt"
  "encoding/json"
  "github.com/hyperledger/fabric-contract-api-go/contractapi"
  "strings"
  "github.com/google/uuid"
  "time"
  "log"
)

// SmartContract provides functions for managing observations of service performance
type SmartContract struct {
  contractapi.Contract
}

// Service describes basic details of what makes up a simple service
// Insert struct field in alphabetic order => to achieve determinism across languages
// golang keeps the order when marshal to json but doesn't order automatically

type Service struct {       
        ActivityValue float64 `json:"activityValue"`
        ActivityUnit string `json:"activityUnit"`
        EnforcementValue float64 `json:"enforcementValue"`
        EnforcementUnit string `json:"enforcementUnit"`
        Guid string `json:"guid"`
        PerformanceValue float64 `json:"performanceValue"`
        PerformanceUnit string `json:"performanceUnit"`
        ProductionValue float64 `json:"productionValue"`
        ProductionUnit string `json:"productionUnit"`
        ExternalitiesValue float64 `json:"externalitiesValue"`
        ExternalitiesUnit string `json:"externalitiesUnit"`
        RatingValue float64 `json:"ratingValue"`
        Principles []string `json:"principles"`
        TimeTo time.Time `json:"timeTo"`
        TimeFrom time.Time `json:"timeFrom"`       
}

// InitLedger adds a base set of services to the ledger
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {

  principles := []string {
    "Peoples are free and independent",
    "Peoples freedom and independence are to be respected by other Peoples",
    "Peoples are to observe treaties and undertakings",
    "Peoples are equal and are parties to the agreements that bind them",
    "Peoples are to observe a duty of non-intervention",
    "Peoples have the right of self-defense but no right to instigate war for reasons other than self-defense",
    "Peoples are to honour human rights",
    "Peoples are to observe certain specified restrictions in the conduct of war",
    "Peoples have a duty to assist other Peoples living under unfavourable conditions that prevent their having a just or decent political and social regime",
  }
  
  services := []Service{
    {
      ActivityValue: 55266778604355.50,
      ActivityUnit: "kgCO2e",
      EnforcementValue: 7713934834354.89,
      EnforcementUnit: "£",
      Guid: "e379bca5fe034e94bc843814d4e32280",
      PerformanceValue: 0.10873,
      PerformanceUnit: "£kgCO2e",
      ProductionValue: 1367798.11,
      ProductionUnit: "£",
      ExternalitiesValue: 48327.61,
      ExternalitiesUnit: "kgCO2e",
      RatingValue: 0.0353,
      Principles: principles,
      TimeTo: time.Date(2020, time.January, 01, 0, 0, 0, 0, time.UTC),
      TimeFrom: time.Date(2020, time.December, 31, 24, 0, 0, 0, time.UTC),
    },
  }

  for _, service := range services {
    serviceJSON, err := json.Marshal(service)
    if err != nil {
        return err
    }

    err = ctx.GetStub().PutState(service.Guid, serviceJSON)
    if err != nil {
        return fmt.Errorf("failed to put to world state. %v", err)
    }
  }

  return nil
}  

// CreateService issues a new service to the world state with given details.
func (s *SmartContract) CreateService(ctx contractapi.TransactionContextInterface, 
    activityValue float64,
    enforcementValue float64,
    enforcementUnit string,
    productionValue float64,
    productionUnit string,
    externalitiesValue float64,
    externalitiesUnit string,
    timeTo time.Time,
    timeFrom time.Time) error {

    id := uuid.New().String()
    
    guid := strings.Replace(id, "-", "", -1)
    
    activityUnit := "kgCO2e"

    performanceValue := enforcementValue / activityValue

    performanceUnit := enforcementUnit + activityUnit

    ratingNumerators := externalitiesValue * performanceValue

    ratingValue := ratingNumerators / productionValue

    principles := []string {
      "Peoples are free and independent",
      "Peoples freedom and independence are to be respected by other Peoples",
      "Peoples are to observe treaties and undertakings",
      "Peoples are equal and are parties to the agreements that bind them",
      "Peoples are to observe a duty of non-intervention",
      "Peoples have the right of self-defense but no right to instigate war for reasons other than self-defense",
      "Peoples are to honour human rights",
      "Peoples are to observe certain specified restrictions in the conduct of war",
      "Peoples have a duty to assist other Peoples living under unfavourable conditions that prevent their having a just or decent political and social regime",
    }

    service := Service{
        ActivityValue: activityValue,
        ActivityUnit: activityUnit,
        EnforcementValue: enforcementValue,
        EnforcementUnit: enforcementUnit,
        Guid: guid,
        PerformanceValue: performanceValue,
        PerformanceUnit: performanceUnit,
        ProductionValue: productionValue,
        ProductionUnit: productionUnit,
        ExternalitiesValue: externalitiesValue,
        ExternalitiesUnit: externalitiesUnit,
        RatingValue: ratingValue,
        Principles: principles,
        TimeTo: timeTo,
        TimeFrom: timeFrom,
    }
    serviceJSON, err := json.Marshal(service)
    if err != nil {
      return err
    }
  
    return ctx.GetStub().PutState(guid, serviceJSON)
  }

  // ReadService returns the service stored in the world state with given guid.
  func (s *SmartContract) ReadService(ctx contractapi.TransactionContextInterface, guid string) (*Service, error) {
    serviceJSON, err := ctx.GetStub().GetState(guid)
    if err != nil {
      return nil, fmt.Errorf("failed to read from world state: %v", err)
    }
    if serviceJSON == nil {
      return nil, fmt.Errorf("the service %s does not exist", guid)
    }
  
    var service Service
    err = json.Unmarshal(serviceJSON, &service)
    if err != nil {
      return nil, err
    }
  
    return &service, nil
  }

  // Service exists returns true when service with given GUID exists in world state
  func (s *SmartContract) ServiceExists(ctx contractapi.TransactionContextInterface, guid string) (bool, error) {
    serviceJSON, err := ctx.GetStub().GetState(guid)
    if err != nil {
      return false, fmt.Errorf("failed to read from world state: %v", err)
    }
  
    return serviceJSON != nil, nil
  }

  // GetAllServices returns all services found in world state
  func (s *SmartContract) GetAllServices(ctx contractapi.TransactionContextInterface) ([]*Service, error) {
    // range query with empty string for startKey and endKey does an
    // open-ended query of all services in the chaincode namespace.
    resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
    if err != nil {
      return nil, err
    }
    defer resultsIterator.Close()
  
    var services []*Service
    for resultsIterator.HasNext() {
      queryResponse, err := resultsIterator.Next()
      if err != nil {
        return nil, err
      }
  
      var service Service
      err = json.Unmarshal(queryResponse.Value, &services)
      if err != nil {
        return nil, err
      }
      services = append(services, &service)
    }
  
    return services, nil
  }


func main() {
  serviceChaincode, err := contractapi.NewChaincode(&SmartContract{})
  if err != nil {
    log.Panicf("Error creating service chaincode: %v", err)
  }

  if err := serviceChaincode.Start(); err != nil {
    log.Panicf("Error starting service chaincode: %v", err)
  }
}