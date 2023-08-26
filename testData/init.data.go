package testData

// Org1 data is stored using the external Mongo Service DB specified, distinct to the Host's Storage
var TestResourceObjectStringifiedToInitStorage1 = `{
	"type": "Organization",
    "attributes": {
        "@type": "Organization",
        "identifier": "urn:uuid:` + TestOrg1UUID + `",
        "alternateName": "` + TestAlternateNameValue1 + `"
    },
    "included": [{
        "type": "ServiceDB",
        "attributes": {
            "@type": "Service",
            "additionalType": "org.mongodb",
            "identifier": "urn:uuid:<organization-uuid-v4>",
            "alternateName": "` + TestAlternateNameValue1 + `",
            "url": "` + TestConnStrServiceProviderMongo1 + `"
        }
    }]
}`

// Org2 data is stored in a Host's private storage because no external Service DB is provided
var TestResourceObjectStringifiedToInitStorage2 = `{
	"type": "Organization",
    "attributes": {
        "@type": "Organization",
        "identifier": "urn:uuid:` + TestOrg2UUID + `",
        "alternateName": "` + TestAlternateNameValue2 + `"
    },
    "included": [{
        "type": "ServiceDB",
        "attributes": {
            "@type": "Service",
            "alternateName": "` + TestAlternateNameValue2 + `",
            "url": ""
        }
    }]
}`
