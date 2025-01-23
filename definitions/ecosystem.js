package main

import (
  "fmt"
  "encoding/json"
  "log"
  "github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for managing observations of ecosystem performance
type SmartContract struct {
  contractapi.Contract
}

// Ecosystem describes basic details of what makes up a simple ecosystem
// Insert struct field in alphabetic order => to achieve determinism across languages
// golang keeps the order when marshal to json but doesn't order automatically

type Ecosystem struct {       
        ActivityValue int `json:"activity.value"`
        ActivityUnit string `json:"activity.unit"`
        EnforcementValue int `json:"enforcement.value"`
        EnforcementUnit string `json:"enforcement.unit"`
        Guid string `json:"guid"`
        PerformanceValue int `json:"performance.value"`
        PerformanceUnit string `json:"performance.unit"`
        Principles array `json:"principles"`
        TimeTo int `json:"time.to"`,
        TimeFrom int `json:"time.from"`       
}

// InitLedger adds a base set of ecosystems to the ledger
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
  ecosystems := []Ecosystem{
    //Add any examples
  }

  for _, ecosystem := range ecosystems {
    ecosystemJSON, err := json.Marshal(ecosystem)
    if err != nil {
        return err
    }

    err = ctx.GetStub().PutState(ecosystem.Guid, ecosystemJSON)
    if err != nil {
        return fmt.Errorf("failed to put to world state. %v", err)
    }
  }

  return nil
}

// CreateEcosystem issues a new ecosystem to the world state with given details.
func (s *SmartContract) CreateEcosystem(ctx contractapi.TransactionContextInterface, 
    activity.value int,
    enforcement.value int,
    enforcement.unit string,
    guid string,
    performance.value int,
    performance.unit string,
    time.to int,
    time.from int)
    
    guid = crypto.randomUUID();

    performance.value = enforcement.value / activity.value;

    performance.unit = [enforcement.unit, activity.unit].join();

    activity.unit = math.unit('kgCO2e');

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
  
    ecosystem := Ecosystem{
        ActivityValue activity.value,
        ActivityUnit activity.unit,
        EnforcementValue enforcement.value,
        EnforcementUnit enforcement.unit,
        Guid guid,
        PerformanceValue performance.value,
        PerformanceUnit performance.unit,
        Principles principles,
        TimeTo time.to,
        TimeFrom time.from
    }
    ecosystemJSON, err := json.Marshal(ecosystem)
    if err != nil {
      return err
    }
  
    return ctx.GetStub().PutState(guid, ecosystemJSON)
  }

  // ReadEcosystem returns the ecosystem stored in the world state with given guid.
func (s *SmartContract) ReadEcosystem(ctx contractapi.TransactionContextInterface, guid string) (*Ecosystem, error) {
    ecosystemJSON, err := ctx.GetStub().GetState(guid)
    if err != nil {
      return nil, fmt.Errorf("failed to read from world state: %v", err)
    }
    if assetJSON == nil {
      return nil, fmt.Errorf("the ecosystem %s does not exist", guid)
    }
  
    var ecosystem Ecosystem
    err = json.Unmarshal(ecosystemJSON, &ecosystem)
    if err != nil {
      return nil, err
    }
  
    return &ecosystem, nil
  }

  // Eocsystem exists returns true when ecosystem with given GUID exists in world state
func (s *SmartContract) EcosystemExists(ctx contractapi.TransactionContextInterface, guid string) (bool, error) {
    ecosystemJSON, err := ctx.GetStub().GetState(guid)
    if err != nil {
      return false, fmt.Errorf("failed to read from world state: %v", err)
    }
  
    return ecosystemJSON != nil, nil
  }

  // GetAllEcosystems returns all ecosystems found in world state
func (s *SmartContract) GetAllEcosystems(ctx contractapi.TransactionContextInterface) ([]*Ecosystem, error) {
    // range query with empty string for startKey and endKey does an
    // open-ended query of all ecosystems in the chaincode namespace.
    resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
    if err != nil {
      return nil, err
    }
    defer resultsIterator.Close()
  
    var ecosystems []*Ecosystem
    for resultsIterator.HasNext() {
      queryResponse, err := resultsIterator.Next()
      if err != nil {
        return nil, err
      }
  
      var ecosystem Ecosystem
      err = json.Unmarshal(queryResponse.Value, &ecosystems)
      if err != nil {
        return nil, err
      }
      ecosystems = append(ecosystems, &ecosystem)
    }
  
    return ecosystems, nil
  }