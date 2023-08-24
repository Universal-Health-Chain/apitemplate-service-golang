package testData

import "strings"

// Organization 1
const TestOrg1LegalName = "Test Org 1 Legal Name"
const TestOrg1TaxID = "ESB0011223344"
const TestOrg1CountryCode2LetterISO = "ES"
const TestOrg1CountryCode3LetterISO = "ESP"

var TestOrg1UUID = "a4a09a0b-44a6-41f0-ab91-85bee9c71870"
var TestOrg1Territory = strings.ToLower(TestOrg1CountryCode3LetterISO)
var TestOrg1URNByUID = "urn:es.unid:cds-" + TestOrg1Territory + ":v1:identity:organization:" + "uuid" + ":" + TestOrg1UUID
var TestOrg1URNByTaxID = "urn:es.unid:cds-" + TestOrg1Territory + ":v1:identity:organization:" + "org.hl7.terminology.codesystem.v2-0203.tax" + ":" + TestOrg1TaxID

// Organization 2
const TestOrg2LegalName = "Test Org 2 Legal Name"
const TestOrg2TaxID = "001122334455"
const TestOrg2CountryCode2LetterISO = "CA"
const TestOrg2CountryCode3LetterISO = "CAN"

var TestOrg2UUID = "b1dec7bf-31b7-4d0c-9de0-63de1283fadc"
var TestOrg2Territory = strings.ToLower(TestOrg2CountryCode3LetterISO)
var TestOrg2URNByUID = "urn:es.unid:cds-" + TestOrg2Territory + ":v1:identity:organization:" + "uuid" + ":" + TestOrg2UUID
var TestOrg2URNByTaxID = "urn:es.unid:cds-" + TestOrg2Territory + ":v1:identity:organization:" + "org.hl7.terminology.codesystem.v2-0203.tax" + ":" + TestOrg2TaxID
