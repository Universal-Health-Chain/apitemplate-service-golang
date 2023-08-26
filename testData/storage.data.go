package testData

import "fmt"

const TestLocalHostPort = 8024

var TestLocalHostUrl = "http://localhost:" + fmt.Sprint(TestLocalHostPort)

var TestMongoConnStrWithoutPort = "mongodb://localhost" // no host port, no userpass
var TestConfidentialDocumentId1 = "confidential-doc-id-1"

// Note: lower case when creating the index
var TestIndexAttributeNamedType = "type"
var TestIndexAttributeNamedServiceDB = "ServiceDB"
var TestIndexAttributeNamedAlternateName = "alternateName"

var TestIndexAttributeValueOrganizationType = "Organization"
var TestIndexAttributeValueOccupationalCodeForDirectorICT = "1330"

var TestAlternateNameValue1 = "alternate-name-1"
var TestAlternateNameValue2 = "alternate-name-2"

// HMAC Keys
var TestHmacKeyForHost = "HmacKeyForHost"
var TestHmacKeyForStorage1 = "HmacKeyForStorage1"
var TestHmacKeyForStorage2 = "HmacKeyForStorage2"

// DEK Keys
var TestDekKeyForHost = "DekKeyForHost"
var TestDekKeyForStorage1 = "DekKeyForStorage1"
var TestDmacKeyForStorage2 = "DekKeyForStorage2"

// 0) Host's storage service
// 0.a) DB server for host and additional private storages
const TestPortServiceProviderMongo0 = 28017                                                               // default MongoDB port for testing
var TestConnStrServiceProviderMongo0 = "mongodb://localhost:" + fmt.Sprint(TestPortServiceProviderMongo0) // no userpass

// 1) Service Provider #1
const TestPortServiceProviderMongo1 = 3 + TestPortServiceProviderMongo0                                   // MongoDB server for vault1 at port 28020
var TestConnStrServiceProviderMongo1 = "mongodb://localhost:" + fmt.Sprint(TestPortServiceProviderMongo1) // 28020
