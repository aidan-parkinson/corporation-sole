package main

import (
  "fmt"
  "encoding/json"
  "log"
  "github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for managing observations of service performance
type SmartContract struct {
  contractapi.Contract
}

// Service describes basic details of what makes up a simple service
// Insert struct field in alphabetic order => to achieve determinism across languages
// golang keeps the order when marshal to json but doesn't order automatically

type Service struct {       
        ActivityValue int `json:"activity.value"`
        ActivityUnit string `json:"activity.unit"`
        EnforcementValue int `json:"enforcement.value"`
        EnforcementUnit string `json:"enforcement.unit"`
        Guid string `json:"guid"`
        PerformanceValue int `json:"performance.value"`
        PerformanceUnit string `json:"performance.unit"`
        ProductionValue int `json:"production.value"`
        ProductionUnit string `json:"production.unit"`
        ExternalitiesValue int `json:"externalities.value"`
        ExternalitiesUnit string `json:"externalities.unit"`
        RatingValue int `json:"rating.value"`
        Principles array `json:"principles"`
        TimeTo int `json:"time.to"`,
        TimeFrom int `json:"time.from"`       
}

guid = crypto.randomUUID();

performance.value = enforcement.value / activity.value;

performance.unit = [enforcement.unit, activity.unit].join();

activity.unit = math.unit('kgCO2e');

externalities.unit = math.unit('kgCO2e');

rating.value = (externalities.value * performance.value) / production.value;

principles = [
    "Peoples are free and independent",
    "Peoples freedom and independence are to be respected by other Peoples",
    "Peoples are to observe treaties and undertakings",
    "Peoples are equal and are parties to the agreements that bind them",
    "Peoples are to observe a duty of non-intervention",
    "Peoples have the right of self-defense but no right to instigate war for reasons other than self-defense",
    "Peoples are to honour human rights",
    "Peoples are to observe certain specified restrictions in the conduct of war",
    "Peoples have a duty to assist other Peoples living under unfavourable conditions that prevent their having a just or decent political and social regime"
]

// InitLedger adds a base set of services to the ledger
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
  services := []Service{
    {
      ActivityValue 55266778604355.50,
      ActivityUnit math.unit('kgCO2e'),
      EnforcementValue 7713934834354.89,
      EnforcementUnit Intl.NumberFormat.currency(USD),
      Guid crypto.randomUUID(),
      PerformanceValue 0.13958,
      PerformanceUnit "$kgCO2e",
      ProductionValue 1367798.11
      ProductionUnit Intl.NumberFormat.currency(GBP)
      ExternalitiesValue 444474.72
      ExternalitiesUnit "kgCO2e"
      RatingValue 0.0365,
      Principles principles,
      TimeTo setUTCDate(January 01, 2020),
      TimeFrom setUTCDate(December 31, 2020)
    }
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
    activity.value int,
    enforcement.value int,
    enforcement.unit string,
    production.value int
    production.unit string
    externalities.value int
    time.to int,
    time.from int)
  
    service := Service{
        ActivityValue activity.value,
        ActivityUnit activity.unit,
        EnforcementValue enforcement.value,
        EnforcementUnit enforcement.unit,
        Guid guid,
        PerformanceValue performance.value,
        PerformanceUnit performance.unit,
        ProductionValue production.value,
        ProductionUnit production.unit,
        ExternalitiesValue externalities.value,
        ExternalitiesUnit externalities.unit,
        RatingValue rating.value,
        Principles principles,
        TimeTo time.to,
        TimeFrom time.from
    }
    ecosystemJSON, err := json.Marshal(service)
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