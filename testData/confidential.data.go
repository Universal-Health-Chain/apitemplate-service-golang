package testData

var TestConfidentialResourceObject1 = `{
	"id": "` + TestConfidentialDocumentId1 + `",
	"type": "confidential",
	"indexed": [
		{
			"attributes": [],
			"hmac": {
				"id": "did:method:<controllerId>#<keyId>",
				"type": "Sha256HmacKey2019"
			},
			"sequence": 0
		}
	],
	"jwe": {
		"protected": "",
		"recipients": [],
		"iv": "",
		"ciphertext": "",
		"tag": ""
	},
	"meta": {
		"created": "2019-06-18",
		"contentType": "video/mpeg",
		"contentLength": 56735817
	},
	"sequence": 0
}`

var TestPrimaryDocConfidential = `{
	"type": "batch",
	"data": [` + TestConfidentialResourceObject1 + `]
}`
