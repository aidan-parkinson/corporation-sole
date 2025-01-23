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

// InitLedger adds a base set of ecosystems to the ledger
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
  ecosystems := []Ecosystem{
    {
      ActivityValue 42526464403980.70,
      ActivityUnit math.unit('kgCO2e'),
      EnforcementValue 2796336506885.35,
      EnforcementUnit Intl.NumberFormat.currency(USD),
      Guid crypto.randomUUID(),
      PerformanceValue 0.06576,
      PerformanceUnit "$kgCO2e",
      Principles principles,
      TimeTo setUTCDate(January 01, 2000),
      TimeFrom setUTCDate(December 31, 2000)
    },
    {
      ActivityValue 41810195465417.20,
      ActivityUnit math.unit('kgCO2e'),
      EnforcementValue 2955718062487.85,
      EnforcementUnit Intl.NumberFormat.currency(USD),
      Guid crypto.randomUUID(),
      PerformanceValue 0.07069,
      PerformanceUnit "$kgCO2e",
      Principles principles,
      TimeTo setUTCDate(January 01, 2001),
      TimeFrom setUTCDate(December 31, 2001)
    },
    {
      ActivityValue 43163264049581.50,
      ActivityUnit math.unit('kgCO2e'),
      EnforcementValue 3251239119567.83,
      EnforcementUnit Intl.NumberFormat.currency(USD),
      Guid crypto.randomUUID(),
      PerformanceValue 0.07532,
      PerformanceUnit "$kgCO2e",
      Principles principles,
      TimeTo setUTCDate(January 01, 2002),
      TimeFrom setUTCDate(December 31, 2002)
    },
    {
      ActivityValue 46128239048293.40,
      ActivityUnit math.unit('kgCO2e'),
      EnforcementValue 3705486991396.25,
      EnforcementUnit Intl.NumberFormat.currency(USD),
      Guid crypto.randomUUID(),
      PerformanceValue 0.08033,
      PerformanceUnit "$kgCO2e",
      Principles principles,
      TimeTo setUTCDate(January 01, 2003),
      TimeFrom setUTCDate(December 31, 2003)
    },
    {
      ActivityValue 48752271780026.20,
      ActivityUnit math.unit('kgCO2e'),
      EnforcementValue 4184289916626.39,
      EnforcementUnit Intl.NumberFormat.currency(USD),
      Guid crypto.randomUUID(),
      PerformanceValue 0.08583,
      PerformanceUnit "$kgCO2e",
      Principles principles,
      TimeTo setUTCDate(January 01, 2004),
      TimeFrom setUTCDate(December 31, 2004)
    },
    {
      ActivityValue 48662428311438.10,
      ActivityUnit math.unit('kgCO2e'),
      EnforcementValue 4392199284751.95,
      EnforcementUnit Intl.NumberFormat.currency(USD),
      Guid crypto.randomUUID(),
      PerformanceValue 0.09026,
      PerformanceUnit "$kgCO2e",
      Principles principles,
      TimeTo setUTCDate(January 01, 2005),
      TimeFrom setUTCDate(December 31, 2005)
    },
    {
      ActivityValue 50346662429817.70,
      ActivityUnit math.unit('kgCO2e'),
      EnforcementValue 4713212363424.03,
      EnforcementUnit Intl.NumberFormat.currency(USD),
      Guid crypto.randomUUID(),
      PerformanceValue 0.09362,
      PerformanceUnit "$kgCO2e",
      Principles principles,
      TimeTo setUTCDate(January 01, 2006),
      TimeFrom setUTCDate(December 31, 2006)
    },
    {
      ActivityValue 49573652091628.70,
      ActivityUnit math.unit('kgCO2e'),
      EnforcementValue 5073569564982.13,
      EnforcementUnit Intl.NumberFormat.currency(USD),
      Guid crypto.randomUUID(),
      PerformanceValue 0.10234,
      PerformanceUnit "$kgCO2e",
      Principles principles,
      TimeTo setUTCDate(January 01, 2007),
      TimeFrom setUTCDate(December 31, 2007)
    },
    {
      ActivityValue 49203903569354.90,
      ActivityUnit math.unit('kgCO2e'),
      EnforcementValue 5861223736080.00,
      EnforcementUnit Intl.NumberFormat.currency(USD),
      Guid crypto.randomUUID(),
      PerformanceValue 0.11912,
      PerformanceUnit "$kgCO2e",
      Principles principles,
      TimeTo setUTCDate(January 01, 2008),
      TimeFrom setUTCDate(December 31, 2008)
    },
    {
      ActivityValue 48597710255510.80,
      ActivityUnit math.unit('kgCO2e'),
      EnforcementValue 5929058323070.43,
      EnforcementUnit Intl.NumberFormat.currency(USD),
      Guid crypto.randomUUID(),
      PerformanceValue 0.12200,
      PerformanceUnit "$kgCO2e",
      Principles principles,
      TimeTo setUTCDate(January 01, 2009),
      TimeFrom setUTCDate(December 31, 2009)
    },
    {
      ActivityValue 49668982222572.40,
      ActivityUnit math.unit('kgCO2e'),
      EnforcementValue 6228253670748.70,
      EnforcementUnit Intl.NumberFormat.currency(USD),
      Guid crypto.randomUUID(),
      PerformanceValue 0.12540,
      PerformanceUnit "$kgCO2e",
      Principles principles,
      TimeTo setUTCDate(January 01, 2010),
      TimeFrom setUTCDate(December 31, 2010)
    },
    {
      ActivityValue 52279637133072.60,
      ActivityUnit math.unit('kgCO2e'),
      EnforcementValue 6844870022802.30,
      EnforcementUnit Intl.NumberFormat.currency(USD),
      Guid crypto.randomUUID(),
      PerformanceValue 0.13093,
      PerformanceUnit "$kgCO2e",
      Principles principles,
      TimeTo setUTCDate(January 01, 2011),
      TimeFrom setUTCDate(December 31, 2011)
    },
    {
      ActivityValue 53401099401947.90,
      ActivityUnit math.unit('kgCO2e'),
      EnforcementValue 6814929419499.72,
      EnforcementUnit Intl.NumberFormat.currency(USD),
      Guid crypto.randomUUID(),
      PerformanceValue 0.12762,
      PerformanceUnit "$kgCO2e",
      Principles principles,
      TimeTo setUTCDate(January 01, 2012),
      TimeFrom setUTCDate(December 31, 2012)
    },
    {
      ActivityValue 53654461725180.10,
      ActivityUnit math.unit('kgCO2e'),
      EnforcementValue 6834780717883.52,
      EnforcementUnit Intl.NumberFormat.currency(USD),
      Guid crypto.randomUUID(),
      PerformanceValue 0.12739,
      PerformanceUnit "$kgCO2e",
      Principles principles,
      TimeTo setUTCDate(January 01, 2013),
      TimeFrom setUTCDate(December 31, 2013)
    },
    {
      ActivityValue 54158329731188.20,
      ActivityUnit math.unit('kgCO2e'),
      EnforcementValue 6875708851852.61,
      EnforcementUnit Intl.NumberFormat.currency(USD),
      Guid crypto.randomUUID(),
      PerformanceValue 0.12696,
      PerformanceUnit "$kgCO2e",
      Principles principles,
      TimeTo setUTCDate(January 01, 2014),
      TimeFrom setUTCDate(December 31, 2014)
    },
    {
      ActivityValue 54472369191803.50,
      ActivityUnit math.unit('kgCO2e'),
      EnforcementValue 6413590627779.08,
      EnforcementUnit Intl.NumberFormat.currency(USD),
      Guid crypto.randomUUID(),
      PerformanceValue 0.11774,
      PerformanceUnit "$kgCO2e",
      Principles principles,
      TimeTo setUTCDate(January 01, 2015),
      TimeFrom setUTCDate(December 31, 2015)
    },
    {
      ActivityValue 53259932210750.80,
      ActivityUnit math.unit('kgCO2e'),
      EnforcementValue 6452824525582.11,
      EnforcementUnit Intl.NumberFormat.currency(USD),
      Guid crypto.randomUUID(),
      PerformanceValue 0.12116,
      PerformanceUnit "$kgCO2e",
      Principles principles,
      TimeTo setUTCDate(January 01, 2016),
      TimeFrom setUTCDate(December 31, 2016)
    },
    {
      ActivityValue 54138630408721.60,
      ActivityUnit math.unit('kgCO2e'),
      EnforcementValue 6780934554209.21,
      EnforcementUnit Intl.NumberFormat.currency(USD),
      Guid crypto.randomUUID(),
      PerformanceValue 0.12525,
      PerformanceUnit "$kgCO2e",
      Principles principles,
      TimeTo setUTCDate(January 01, 2017),
      TimeFrom setUTCDate(December 31, 2017)
    },
    {
      ActivityValue 55914879048865.20,
      ActivityUnit math.unit('kgCO2e'),
      EnforcementValue 7338967125086.29,
      EnforcementUnit Intl.NumberFormat.currency(USD),
      Guid crypto.randomUUID(),
      PerformanceValue 0.13125,
      PerformanceUnit "$kgCO2e",
      Principles principles,
      TimeTo setUTCDate(January 01, 2018),
      TimeFrom setUTCDate(December 31, 2018)
    },
    {
      ActivityValue 56639566875488.90,
      ActivityUnit math.unit('kgCO2e'),
      EnforcementValue 7528974444291.53,
      EnforcementUnit Intl.NumberFormat.currency(USD),
      Guid crypto.randomUUID(),
      PerformanceValue 0.13293,
      PerformanceUnit "$kgCO2e",
      Principles principles,
      TimeTo setUTCDate(January 01, 2019),
      TimeFrom setUTCDate(December 31, 2019)
    },
    {
      ActivityValue 55266778604355.50,
      ActivityUnit math.unit('kgCO2e'),
      EnforcementValue 7713934834354.89,
      EnforcementUnit Intl.NumberFormat.currency(USD),
      Guid crypto.randomUUID(),
      PerformanceValue 0.13958,
      PerformanceUnit "$kgCO2e",
      Principles principles,
      TimeTo setUTCDate(January 01, 2020),
      TimeFrom setUTCDate(December 31, 2020)
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
    performance.value int,
    performance.unit string,
    time.to int,
    time.from int)
  
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