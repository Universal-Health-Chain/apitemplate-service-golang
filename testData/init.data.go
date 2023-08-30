package testData

// Org1 data is stored using the external Mongo Service DB specified, distinct to the Host's Storage
var TestResourceObjectStringifiedToInitStorage1 = `{
    "type": "Organization",
    "attributes": {
        "@type": "Organization",
        "alternateName": "` + TestOrgAlternateNameValue1 + `",
        "identifier": "urn:uuid:` + TestOrg1UUID + `"
    },
    "included": [{
        "type": "ServiceDB",
        "attributes": {
            "@type": "Service",
            "additionalType": "org.mongodb",
            "alternateName": "` + TestOrgAlternateNameValue1 + `",
            "identifier": "urn:uuid:` + TestOrg1UUID + `",
            "url": "` + TestConnStrServiceProviderMongo1 + `"
        }
    },
    {
        "type": "Key",
        "attributes": {
            "@type": "StructuredValue",
            "additionalType": "HS256",
            "value": "KeyHMAC"
        }
    },
    {
        "type": "Key",
        "attributes": {
            "@type": "StructuredValue",
            "additionalType": "AES-256-GCM",
            "value": "KeyAES"
        }
    }]
}`

// Org2 data is stored in a Host's private storage because no external Service DB is provided
var TestResourceObjectStringifiedToInitStorage2 = `{
	"type": "Organization",
    "attributes": {
        "@type": "Organization",
        "alternateName": "` + TestOrgAlternateNameValue2 + `",
        "identifier": "urn:uuid:` + TestOrg2UUID + `"
    },
    "included": [{
        "type": "ServiceDB",
        "attributes": {
            "@type": "Service",
            "alternateName": "` + TestOrgAlternateNameValue2 + `",
            "url": ""
        }
    },
    {
        "type": "Key",
        "attributes": {
            "@type": "StructuredValue",
            "additionalType": "HS256",
            "value": "KeyHMAC"
        }
    },
    {
        "type": "Key",
        "attributes": {
            "@type": "StructuredValue",
            "additionalType": "AES-256-GCM",
            "value": "KeyAES"
        }
    }]
}`
