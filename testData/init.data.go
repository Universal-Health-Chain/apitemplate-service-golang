package testData

// Org1 data is stored in another Storage Provider using the external Mongo Service DB specified
var TestResourceObjectStringifiedToInit = `{
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
            "name": "dbConnectionsVault1",
            "alternateName": "` + TestAlternateNameValue1 + `",
            "url": "` + TestConnStrServiceProviderMongo1 + `"
        }
    }]
}`

// Org2 data is stored in a Host's private storage because no external Service DB is provided
var TestResourceObjectStringifiedForInitializingVault2 = `{
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
