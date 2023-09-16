package testData

import "strings"

// Organization 1
const (
	TestOrg1LegalName             = "Test Org 1 Legal Name"
	TestOrg1TaxID                 = "ESB0011223344"
	TestOrg1CountryCode2LetterISO = "ES"
	TestOrg1CountryCode3LetterISO = "ESP"
)

var (
	TestOrg1UUID       = "a4a09a0b-44a6-41f0-ab91-85bee9c71870"
	TestOrg1Territory  = strings.ToLower(TestOrg1CountryCode3LetterISO)
	TestOrg1URNByUID   = "urn:es.unid:cds-" + TestOrg1Territory + ":v1:identity:organization:" + "uuid" + ":" + TestOrg1UUID
	TestOrg1URNByTaxID = "urn:es.unid:cds-" + TestOrg1Territory + ":v1:identity:organization:" + "org.hl7.terminology.codesystem.v2-0203.tax" + ":" + TestOrg1TaxID
)

// Organization 2
const (
	TestOrg2LegalName             = "Test Org 2 Legal Name"
	TestOrg2TaxID                 = "001122334455"
	TestOrg2CountryCode2LetterISO = "CA"
	TestOrg2CountryCode3LetterISO = "CAN"
)

var (
	TestOrg2UUID       = "b1dec7bf-31b7-4d0c-9de0-63de1283fadc"
	TestOrg2Territory  = strings.ToLower(TestOrg2CountryCode3LetterISO)
	TestOrg2URNByUID   = "urn:es.unid:cds-" + TestOrg2Territory + ":v1:identity:organization:" + "uuid" + ":" + TestOrg2UUID
	TestOrg2URNByTaxID = "urn:es.unid:cds-" + TestOrg2Territory + ":v1:identity:organization:" + "org.hl7.terminology.codesystem.v2-0203.tax" + ":" + TestOrg2TaxID
)
