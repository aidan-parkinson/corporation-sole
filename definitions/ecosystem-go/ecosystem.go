package main

import (
  "fmt"
  "encoding/json"
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
        ActivityValue float64 `json:"activityValue"`
        ActivityUnit string `json:"activityUnit"`
        EnforcementValue float64 `json:"enforcementValue"`
        EnforcementUnit string `json:"enforcementUnit"`
        Guid string `json:"guid"`
        PerformanceValue float64 `json:"performanceValue"`
        PerformanceUnit string `json:"performanceUnit"`
        Principles []string `json:"principles"`
        TimeTo time.Time `json:"timeTo"`
        TimeFrom time.Time `json:"timeFrom"`       
}

// InitLedger adds a base set of ecosystems to the ledger
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
  

  ecosystems := []Ecosystem{
    {
      ActivityValue: 42526464403980.70,
      ActivityUnit: "kgCO2e",
      EnforcementValue: 2796336506885.35,
      EnforcementUnit: "$",
      Guid: "a7daf04daa5c4c508e8c7958f149b855",
      PerformanceValue: 0.06576,
      PerformanceUnit: "$kgCO2e",
      Principles: principles,
      TimeTo: time.Date(2000, time.January, 01, 0, 0, 0, 0, time.UTC),
      TimeFrom: time.Date(2000, time.December, 31, 24, 0, 0, 0, time.UTC),
    },
    {
      ActivityValue: 41810195465417.20,
      ActivityUnit: "kgCO2e",
      EnforcementValue: 2955718062487.85,
      EnforcementUnit: "$",
      Guid: "7b6f1640460c4290a2a18a7ddfbfb120",
      PerformanceValue: 0.07069,
      PerformanceUnit: "$kgCO2e",
      Principles: principles,
      TimeTo: time.Date(2001, time.January, 01, 0, 0, 0, 0, time.UTC),
      TimeFrom: time.Date(2001, time.December, 31, 24, 0, 0, 0, time.UTC),
    },
    {
      ActivityValue: 43163264049581.50,
      ActivityUnit: "kgCO2e",
      EnforcementValue: 3251239119567.83,
      EnforcementUnit: "$",
      Guid: "4bf6ab22ab574cdfb324edcba28ab8a9",
      PerformanceValue: 0.07532,
      PerformanceUnit: "$kgCO2e",
      Principles: principles,
      TimeTo: time.Date(2002, time.January, 01, 0, 0, 0, 0, time.UTC),
      TimeFrom: time.Date(2002, time.December, 31, 24, 0, 0, 0, time.UTC),
    },
    {
      ActivityValue: 46128239048293.40,
      ActivityUnit: "kgCO2e",
      EnforcementValue: 3705486991396.25,
      EnforcementUnit: "$",
      Guid: "94ae516c402f4857b12f416f1423248e",
      PerformanceValue: 0.08033,
      PerformanceUnit: "$kgCO2e",
      Principles: principles,
      TimeTo: time.Date(2003, time.January, 01, 0, 0, 0, 0, time.UTC),
      TimeFrom: time.Date(2003, time.December, 31, 24, 0, 0, 0, time.UTC),
    },
    {
      ActivityValue: 48752271780026.20,
      ActivityUnit: "kgCO2e",
      EnforcementValue: 4184289916626.39,
      EnforcementUnit: "$",
      Guid: "bcec5990fade4df9986986077142c217",
      PerformanceValue: 0.08583,
      PerformanceUnit: "$kgCO2e",
      Principles: principles,
      TimeTo: time.Date(2004, time.January, 01, 0, 0, 0, 0, time.UTC),
      TimeFrom: time.Date(2004, time.December, 31, 24, 0, 0, 0, time.UTC),
    },
    {
      ActivityValue: 48662428311438.10,
      ActivityUnit: "kgCO2e",
      EnforcementValue: 4392199284751.95,
      EnforcementUnit: "$",
      Guid: "991f0d5d9431424489469c45645ab98a",
      PerformanceValue: 0.09026,
      PerformanceUnit: "$kgCO2e",
      Principles: principles,
      TimeTo: time.Date(2005, time.January, 01, 0, 0, 0, 0, time.UTC),
      TimeFrom: time.Date(2005, time.December, 31, 24, 0, 0, 0, time.UTC),
    },
    {
      ActivityValue: 50346662429817.70,
      ActivityUnit: "kgCO2e",
      EnforcementValue: 4713212363424.03,
      EnforcementUnit: "$",
      Guid: "b48a8b6ce1a949a3bac0cb4d0764db2e",
      PerformanceValue: 0.09362,
      PerformanceUnit: "$kgCO2e",
      Principles: principles,
      TimeTo: time.Date(2006, time.January, 01, 0, 0, 0, 0, time.UTC),
      TimeFrom: time.Date(2006, time.December, 31, 24, 0, 0, 0, time.UTC),
    },
    {
      ActivityValue: 49573652091628.70,
      ActivityUnit: "kgCO2e",
      EnforcementValue: 5073569564982.13,
      EnforcementUnit: "$",
      Guid: "5acb02361ab742d1990919436c6a4fb6",
      PerformanceValue: 0.10234,
      PerformanceUnit: "$kgCO2e",
      Principles: principles,
      TimeTo: time.Date(2007, time.January, 01, 0, 0, 0, 0, time.UTC),
      TimeFrom: time.Date(2007, time.December, 31, 24, 0, 0, 0, time.UTC),
    },
    {
      ActivityValue: 49203903569354.90,
      ActivityUnit: "kgCO2e",
      EnforcementValue: 5861223736080.00,
      EnforcementUnit: "$",
      Guid: "9b02edf2c6414dc692aef39f2c6db79c",
      PerformanceValue: 0.11912,
      PerformanceUnit: "$kgCO2e",
      Principles: principles,
      TimeTo: time.Date(2008, time.January, 01, 0, 0, 0, 0, time.UTC),
      TimeFrom: time.Date(2008, time.December, 31, 24, 0, 0, 0, time.UTC),
    },
    {
      ActivityValue: 48597710255510.80,
      ActivityUnit: "kgCO2e",
      EnforcementValue: 5929058323070.43,
      EnforcementUnit: "$",
      Guid: "9855a3973fde4b88822bb43aeb3f2761",
      PerformanceValue: 0.12200,
      PerformanceUnit: "$kgCO2e",
      Principles: principles,
      TimeTo: time.Date(2009, time.January, 01, 0, 0, 0, 0, time.UTC),
      TimeFrom: time.Date(2009, time.December, 31, 24, 0, 0, 0, time.UTC),
    },
    {
      ActivityValue: 49668982222572.40,
      ActivityUnit: "kgCO2e",
      EnforcementValue: 6228253670748.70,
      EnforcementUnit: "$",
      Guid: "d926e9fa881e4bebb6aa6aa14f14d1ad",
      PerformanceValue: 0.12540,
      PerformanceUnit: "$kgCO2e",
      Principles: principles,
      TimeTo: time.Date(2010, time.January, 01, 0, 0, 0, 0, time.UTC),
      TimeFrom: time.Date(2010, time.December, 31, 24, 0, 0, 0, time.UTC),
    },
    {
      ActivityValue: 52279637133072.60,
      ActivityUnit: "kgCO2e",
      EnforcementValue: 6844870022802.30,
      EnforcementUnit: "$",
      Guid: "ad756ab40123467c8ddb55492f9c547a",
      PerformanceValue: 0.13093,
      PerformanceUnit: "$kgCO2e",
      Principles: principles,
      TimeTo: time.Date(2011, time.January, 01, 0, 0, 0, 0, time.UTC),
      TimeFrom: time.Date(2011, time.December, 31, 24, 0, 0, 0, time.UTC),
    },
    {
      ActivityValue: 53401099401947.90,
      ActivityUnit: "kgCO2e",
      EnforcementValue: 6814929419499.72,
      EnforcementUnit: "$",
      Guid: "61a76d2910fb4b4aae1c2fe300d71edb",
      PerformanceValue: 0.12762,
      PerformanceUnit: "$kgCO2e",
      Principles: principles,
      TimeTo: time.Date(2012, time.January, 01, 0, 0, 0, 0, time.UTC),
      TimeFrom: time.Date(2012, time.December, 31, 24, 0, 0, 0, time.UTC),
    },
    {
      ActivityValue: 53654461725180.10,
      ActivityUnit: "kgCO2e",
      EnforcementValue: 6834780717883.52,
      EnforcementUnit: "$",
      Guid: "f6df7b666c4c475d8df81a5670adfd9e",
      PerformanceValue: 0.12739,
      PerformanceUnit: "$kgCO2e",
      Principles: principles,
      TimeTo: time.Date(2013, time.January, 01, 0, 0, 0, 0, time.UTC),
      TimeFrom: time.Date(2013, time.December, 31, 24, 0, 0, 0, time.UTC),
    },
    {
      ActivityValue: 54158329731188.20,
      ActivityUnit: "kgCO2e",
      EnforcementValue: 6875708851852.61,
      EnforcementUnit: "$",
      Guid: "669a1dbf679e498a831585782dfbcc84",
      PerformanceValue: 0.12696,
      PerformanceUnit: "$kgCO2e",
      Principles: principles,
      TimeTo: time.Date(2014, time.January, 01, 0, 0, 0, 0, time.UTC),
      TimeFrom: time.Date(2014, time.December, 31, 24, 0, 0, 0, time.UTC),
    },
    {
      ActivityValue: 54472369191803.50,
      ActivityUnit: "kgCO2e",
      EnforcementValue: 6413590627779.08,
      EnforcementUnit: "$",
      Guid: "c00675f1e0ed4e5586d1901a8db6979f",
      PerformanceValue: 0.11774,
      PerformanceUnit: "$kgCO2e",
      Principles: principles,
      TimeTo: time.Date(2015, time.January, 01, 0, 0, 0, 0, time.UTC),
      TimeFrom: time.Date(2015, time.December, 31, 24, 0, 0, 0, time.UTC),
    },
    {
      ActivityValue: 53259932210750.80,
      ActivityUnit: "kgCO2e",
      EnforcementValue: 6452824525582.11,
      EnforcementUnit: "$",
      Guid: "d7ec76ce1d304c4da71caeac772fe24b",
      PerformanceValue: 0.12116,
      PerformanceUnit: "$kgCO2e",
      Principles: principles,
      TimeTo: time.Date(2016, time.January, 01, 0, 0, 0, 0, time.UTC),
      TimeFrom: time.Date(2016, time.December, 31, 24, 0, 0, 0, time.UTC),
    },
    {
      ActivityValue: 54138630408721.60,
      ActivityUnit: "kgCO2e",
      EnforcementValue: 6780934554209.21,
      EnforcementUnit: "$",
      Guid: "67968b46f41743a1ace1e1c90cdeb3c4",
      PerformanceValue: 0.12525,
      PerformanceUnit: "$kgCO2e",
      Principles: principles,
      TimeTo: time.Date(2017, time.January, 01, 0, 0, 0, 0, time.UTC),
      TimeFrom: time.Date(2017, time.December, 31, 24, 0, 0, 0, time.UTC),
    },
    {
      ActivityValue: 55914879048865.20,
      ActivityUnit: "kgCO2e",
      EnforcementValue: 7338967125086.29,
      EnforcementUnit: "$",
      Guid: "5c6294228ab743f2bc5c4a2d9e42e528",
      PerformanceValue: 0.13125,
      PerformanceUnit: "$kgCO2e",
      Principles: principles,
      TimeTo: time.Date(2018, time.January, 01, 0, 0, 0, 0, time.UTC),
      TimeFrom: time.Date(2018, time.December, 31, 24, 0, 0, 0, time.UTC),
    },
    {
      ActivityValue: 56639566875488.90,
      ActivityUnit: "kgCO2e",
      EnforcementValue: 7528974444291.53,
      EnforcementUnit: "$",
      Guid: "cc8893cc5ea84d44904de31e8c4a7bb7",
      PerformanceValue: 0.13293,
      PerformanceUnit: "$kgCO2e",
      Principles: principles,
      TimeTo: time.Date(2019, time.January, 01, 0, 0, 0, 0, time.UTC),
      TimeFrom: time.Date(2019, time.December, 31, 24, 0, 0, 0, time.UTC),
    },
    {
      ActivityValue: 55266778604355.50,
      ActivityUnit: "kgCO2e",
      EnforcementValue: 7713934834354.89,
      EnforcementUnit: "$",
      Guid: "5a6865982f1a4dbc9e452187cb2fc36c",
      PerformanceValue: 0.13958,
      PerformanceUnit: "$kgCO2e",
      Principles: principles,
      TimeTo: time.Date(2020, time.January, 01, 0, 0, 0, 0, time.UTC),
      TimeFrom: time.Date(2020, time.December, 31, 24, 0, 0, 0, time.UTC),
    }, 
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
    activityValue float64,
    enforcementValue float64,
    enforcementUnit string,
    guid string,
    activityUnit string,
    performanceValue float64,
    performanceUnit string,
    principles []string,
    timeTo time.Time,
    timeFrom time.Time) {

    id := uuid.New().String()
    
    guid = strings.Replace(id, "-", "", -1)
    
    activityUnit = "kgCO2e"

    performanceValue = enforcementValue / activityValue

    performanceUnit = enforcementUnit + activityUnit

    principles = []string {
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
  
    ecosystem := Ecosystem{
        ActivityValue: activityValue,
        ActivityUnit: activityUnit,
        EnforcementValue: enforcementValue,
        EnforcementUnit: enforcementUnit,
        Guid: guid,
        PerformanceValue: performanceValue,
        PerformanceUnit: performanceUnit,
        Principles: principles,
        TimeTo: timeTo,
        TimeFrom: timeFrom,
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
    if ecosystemJSON == nil {
      return nil, fmt.Errorf("the ecosystem %s does not exist", guid)
    }
  
    var ecosystem Ecosystem
    err = json.Unmarshal(ecosystemJSON, &ecosystem)
    if err != nil {
      return nil, err
    }
  
    return &ecosystem, nil
  }

  // Ecosystem exists returns true when ecosystem with given GUID exists in world state
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