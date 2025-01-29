package main

import (
  "fmt"
  "encoding/json"
  "log"
  "github.com/hyperledger/fabric-contract-api-go/contractapi"
  "strings"
  "github.com/google/uuid"
  "time"
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
        TimeTo int `json:"time.to"`
        TimeFrom int `json:"time.from"`       
}

func GuId() {
  // Generate a new random UUID
  guid := uuid.New()
  fmt.Printf("Generated UUID: %s\n", u)
}

principles = []Principle{
  "Peoples are free and independent",
  "Peoples freedom and independence are to be respected by other Peoples",
  "Peoples are to observe treaties and undertakings",
  "Peoples are equal and are parties to the agreements that bind them",
  "Peoples are to observe a duty of non-intervention",
  "Peoples have the right of self-defense but no right to instigate war for reasons other than self-defense",
  "Peoples are to honour human rights",
  "Peoples are to observe certain specified restrictions in the conduct of war",
  "Peoples have a duty to assist other Peoples living under unfavourable conditions that prevent their having a just or decent political and social regime"
}

// InitLedger adds a base set of ecosystems to the ledger
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
  ecosystems := []Ecosystem{
    {
      ActivityValue: 42526464403980.70,
      ActivityUnit: "kgCO2e",
      EnforcementValue: 2796336506885.35,
      EnforcementUnit: "$",
      Guid: GuId(),
      PerformanceValue: 0.06576,
      PerformanceUnit: "$kgCO2e",
      Principles: principles,
      TimeTo: time.Date(2000, 1, 01, timeString),
      TimeFrom: time.Date(2000, 12, 31, timeString)
    },
    {
      ActivityValue: 41810195465417.20,
      ActivityUnit: "kgCO2e",
      EnforcementValue: 2955718062487.85,
      EnforcementUnit: "$",
      Guid: GuId(),
      PerformanceValue: 0.07069,
      PerformanceUnit: "$kgCO2e",
      Principles: principles,
      TimeTo: time.Date(2001, 1, 01, timeString),
      TimeFrom: time.Date(2001, 12, 31, timeString)
    },
    {
      ActivityValue: 43163264049581.50,
      ActivityUnit: "kgCO2e",
      EnforcementValue: 3251239119567.83,
      EnforcementUnit: "$",
      Guid: GuId(),
      PerformanceValue: 0.07532,
      PerformanceUnit: "$kgCO2e",
      Principles: principles,
      TimeTo: time.Date(2002, 1, 01, timeString),
      TimeFrom: time.Date(2002, 12, 31, timeString)
    },
    {
      ActivityValue: 46128239048293.40,
      ActivityUnit: "kgCO2e",
      EnforcementValue: 3705486991396.25,
      EnforcementUnit: "$",
      Guid: GuId(),
      PerformanceValue: 0.08033,
      PerformanceUnit: "$kgCO2e",
      Principles: principles,
      TimeTo: time.Date(2003, 1, 01, timeString),
      TimeFrom: time.Date(2003, 12, 31, timeString)
    },
    {
      ActivityValue: 48752271780026.20,
      ActivityUnit: "kgCO2e",
      EnforcementValue: 4184289916626.39,
      EnforcementUnit: "$",
      Guid: GuId(),
      PerformanceValue: 0.08583,
      PerformanceUnit: "$kgCO2e",
      Principles: principles,
      TimeTo: time.Date(2004, 1, 01, timeString),
      TimeFrom: time.Date(2004, 12, 31, timeString)
    },
    {
      ActivityValue: 48662428311438.10,
      ActivityUnit: "kgCO2e",
      EnforcementValue: 4392199284751.95,
      EnforcementUnit: "$",
      Guid: GuId(),
      PerformanceValue: 0.09026,
      PerformanceUnit: "$kgCO2e",
      Principles: principles,
      TimeTo: time.Date(2005, 1, 01, timeString),
      TimeFrom: time.Date(2005, 12, 31, timeString)
    },
    {
      ActivityValue: 50346662429817.70,
      ActivityUnit: "kgCO2e",
      EnforcementValue: 4713212363424.03,
      EnforcementUnit: "$",
      Guid: GuId(),
      PerformanceValue: 0.09362,
      PerformanceUnit: "$kgCO2e",
      Principles: principles,
      TimeTo: time.Date(2006, 1, 01, timeString),
      TimeFrom: time.Date(2006, 12, 31, timeString)
    },
    {
      ActivityValue: 49573652091628.70,
      ActivityUnit: "kgCO2e",
      EnforcementValue: 5073569564982.13,
      EnforcementUnit: "$",
      Guid: GuId(),
      PerformanceValue: 0.10234,
      PerformanceUnit: "$kgCO2e",
      Principles: principles,
      TimeTo: time.Date(2007, 1, 01, timeString),
      TimeFrom: time.Date(2007, 12, 31, timeString)
    },
    {
      ActivityValue: 49203903569354.90,
      ActivityUnit: "kgCO2e",
      EnforcementValue: 5861223736080.00,
      EnforcementUnit: "$",
      Guid: GuId(),
      PerformanceValue: 0.11912,
      PerformanceUnit: "$kgCO2e",
      Principles: principles,
      TimeTo: time.Date(2008, 1, 01, timeString),
      TimeFrom: time.Date(2008, 12, 31, timeString)
    },
    {
      ActivityValue: 48597710255510.80,
      ActivityUnit: "kgCO2e",
      EnforcementValue: 5929058323070.43,
      EnforcementUnit: "$",
      Guid: GuId(),
      PerformanceValue: 0.12200,
      PerformanceUnit: "$kgCO2e",
      Principles: principles,
      TimeTo: time.Date(2009, 1, 01, timeString),
      TimeFrom: time.Date(2009, 12, 31, timeString)
    },
    {
      ActivityValue: 49668982222572.40,
      ActivityUnit: "kgCO2e",
      EnforcementValue: 6228253670748.70,
      EnforcementUnit: "$",
      Guid: GuId(),
      PerformanceValue: 0.12540,
      PerformanceUnit: "$kgCO2e",
      Principles: principles,
      TimeTo: time.Date(2010, 1, 01, timeString),
      TimeFrom: time.Date(2010, 12, 31, timeString)
    },
    {
      ActivityValue: 52279637133072.60,
      ActivityUnit: "kgCO2e",
      EnforcementValue: 6844870022802.30,
      EnforcementUnit: "$",
      Guid: GuId(),
      PerformanceValue: 0.13093,
      PerformanceUnit: "$kgCO2e",
      Principles: principles,
      TimeTo: time.Date(2011, 1, 01, timeString),
      TimeFrom: time.Date(2011, 12, 31, timeString)
    },
    {
      ActivityValue: 53401099401947.90,
      ActivityUnit: "kgCO2e",
      EnforcementValue: 6814929419499.72,
      EnforcementUnit: "$",
      Guid: GuId(),
      PerformanceValue: 0.12762,
      PerformanceUnit: "$kgCO2e",
      Principles: principles,
      TimeTo: time.Date(2012, 1, 01, timeString),
      TimeFrom: time.Date(2012, 12, 31, timeString)
    },
    {
      ActivityValue: 53654461725180.10,
      ActivityUnit: "kgCO2e",
      EnforcementValue: 6834780717883.52,
      EnforcementUnit: "$",
      Guid: GuId(),
      PerformanceValue: 0.12739,
      PerformanceUnit: "$kgCO2e",
      Principles: principles,
      TimeTo: time.Date(2013, 1, 01, timeString),
      TimeFrom: time.Date(2013, 12, 31, timeString)
    },
    {
      ActivityValue: 54158329731188.20,
      ActivityUnit: "kgCO2e",
      EnforcementValue: 6875708851852.61,
      EnforcementUnit: "$",
      Guid: GuId(),
      PerformanceValue: 0.12696,
      PerformanceUnit: "$kgCO2e",
      Principles: principles,
      TimeTo: time.Date(2014, 1, 01, timeString),
      TimeFrom: time.Date(2014, 12, 31, timeString)
    },
    {
      ActivityValue: 54472369191803.50,
      ActivityUnit: "kgCO2e",
      EnforcementValue: 6413590627779.08,
      EnforcementUnit: "$",
      Guid: GuId(),
      PerformanceValue: 0.11774,
      PerformanceUnit: "$kgCO2e",
      Principles: principles,
      TimeTo: time.Date(2015, 1, 01, timeString),
      TimeFrom: time.Date(2015, 12, 31, timeString)
    },
    {
      ActivityValue: 53259932210750.80,
      ActivityUnit: "kgCO2e",
      EnforcementValue: 6452824525582.11,
      EnforcementUnit: "$",
      Guid: GuId(),
      PerformanceValue: 0.12116,
      PerformanceUnit: "$kgCO2e",
      Principles: principles,
      TimeTo: time.Date(2016, 1, 01, timeString),
      TimeFrom: time.Date(2016, 12, 31, timeString)
    },
    {
      ActivityValue: 54138630408721.60,
      ActivityUnit: "kgCO2e",
      EnforcementValue: 6780934554209.21,
      EnforcementUnit: "$",
      Guid: GuId(),
      PerformanceValue: 0.12525,
      PerformanceUnit: "$kgCO2e",
      Principles: principles,
      TimeTo: time.Date(2017, 1, 01, timeString),
      TimeFrom: time.Date(2017, 12, 31, timeString)
    },
    {
      ActivityValue: 55914879048865.20,
      ActivityUnit: "kgCO2e",
      EnforcementValue: 7338967125086.29,
      EnforcementUnit: "$",
      Guid: GuId(),
      PerformanceValue: 0.13125,
      PerformanceUnit: "$kgCO2e",
      Principles: principles,
      TimeTo: time.Date(2018, 1, 01, timeString),
      TimeFrom: time.Date(2018, 12, 31, timeString)
    },
    {
      ActivityValue: 56639566875488.90,
      ActivityUnit: "kgCO2e",
      EnforcementValue: 7528974444291.53,
      EnforcementUnit: "$",
      Guid: GuId(),
      PerformanceValue: 0.13293,
      PerformanceUnit: "$kgCO2e",
      Principles: principles,
      TimeTo: time.Date(2019, 1, 01, timeString),
      TimeFrom: time.Date(2019, 12, 31, timeString)
    },
    {
      ActivityValue: 55266778604355.50,
      ActivityUnit: "kgCO2e",
      EnforcementValue: 7713934834354.89,
      EnforcementUnit: "$",
      Guid: GuId(),
      PerformanceValue: 0.13958,
      PerformanceUnit: "$kgCO2e",
      Principles: principles,
      TimeTo: time.Date(2020, 1, 01, timeString),
      TimeFrom: time.Date(2020, 12, 31, timeString)
    } 
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
    time.to int,
    time.from int)
    
    var performance.value int

    var performance.unit string

    var rating.numerators int

    var rating.value int

    var activity.unit string
    
    // Function to compute performance valuation
    func PerformanceValue(
      enforcement.value int,
      activity.value int) {
      performance.value = enforcement.value / activity.value
    }
    
    // Function to compute performance unit
    func PerformanceUnit(
      enforcement.unit string, 
      activity.unit string) {
      str := []string{enforcement.unit, activity.unit}
      performance.unit = fmt.Println(strings.Join(str))
    }
    
    GuId()
    PerformanceValue()
    PerformanceUnit()
    
    activity.unit = "kgCO2e"
  
    ecosystem := Ecosystem{
        ActivityValue: activity.value,
        ActivityUnit: activity.unit,
        EnforcementValue: enforcement.value,
        EnforcementUnit: enforcement.unit,
        Guid: guid,
        PerformanceValue: performance.value,
        PerformanceUnit: performance.unit,
        Principles: principles,
        TimeTo: time.to,
        TimeFrom: time.from
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