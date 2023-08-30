package testData

const (
	TestBackup1UUID               = "c6978d7b-cc42-45f1-ac86-11a0b83b0bd4"
	TestBackup1AlternateNameValue = "backup-1"
	TestBackup2UUID               = "d49e0710-6201-4386-b754-1d941f225a8a"
	TestBackup2AlternateNameValue = "backup-2"
)

// Data for Backup #1 is stored using the external Mongo Service DB specified, distinct to the Host's Storage
var TestResourceObjectStringifiedToInitBackup1 = `{
	"type": "Backup",
    "attributes": {
        "@type": "SoftwareApplication",
        "identifier": "urn:uuid:` + TestBackup1UUID + `",
        "alternateName": "` + TestBackup1AlternateNameValue + `"
    }
    "included": [{
        "type": "ServiceDB",
        "attributes": {
            "@type": "Service",
            "additionalType": "org.mongodb",
            "alternateName": "` + TestOrgAlternateNameValue1 + `",
            "identifier": "urn:uuid:` + TestOrg1UUID + `",
            "url": "` + TestConnStrServiceProviderMongo1 + `"
        }
    }]
}`

// Data for Backup #2 is stored in a Private storage in the Host's Storage Provider, but without any specific HMAC or DEK
var TestResourceObjectStringifiedToInitBackup2 = `{
	"type": "Backup",
    "attributes": {
        "@type": "SoftwareApplication",
        "alternateName": "` + TestBackup2AlternateNameValue + `",
        "identifier": "urn:uuid:` + TestBackup2UUID + `"
    }
}`
