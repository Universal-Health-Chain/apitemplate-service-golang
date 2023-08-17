package testData
const TestPayloadJWTId = "<not relevant JWT ID, just to identify the request>";
const TestBundleId = "<not relevant Bundle ID, just to identify the data exchange>";

const TestPatientIURI = "urn:es.unid:cds:v1:health:person:x-uhc-mid:uuid:<person-uuid-v4>";
const TestHashPatientIURI = "zQ<multiformat(Base58(SHA-256(TestPatientIURI)))>";

const TestCompositionIURI = TestPatientIURI + ":org.hl7.fhir.r4.Communication:uuid:<resource-uuid-v4>";
const TestHashCompositionIURI = "zQ<multiformat(Base58(SHA-256(TestCommunicationIURI)))>";

const TestImmunizationIURI = TestPatientIURI + ":org.hl7.fhir.r4.Immunization:uuid:<resource-uuid-v4>";
const TestHashImmunizationIURI = "zQ<multiformat(Base58(SHA-256(TestImmunizationIURI)))>";

//  1) FHIR data is created for the messaging service
//  1.a) Data for EHR is a FHIR Bundle transaction within the first resource object.
//  1.b) Data being sent to the end-user (patient, family member) is a FHIR Message.
//  As example, a URL with FHIR parameters are converted by the messaging service for the connections service
//  <tenantBaseUrl>/cds-<territory>/v1/edv/didcomm/<vaultID>/
//
//  Notes:
//  - The 1st resource object is the clinical Bundle "document",
//  and the fist entry in the embedded FHIR "resource" is the document "Composition".
//  - "fullUrl" is set by the messaging service but the base url can be replaced by the connections service.
//  - "id" is internal and it is the hash of the interoperable URI (for blockchain certification).
//  - "type" of the resource object can be  "org.openehr..." OR "org.hl7.fhir.<release>.bundle.document" OR "org.hl7.fhir.<release>.communication"
//  -  The "fullUrl" and "id" of Medication are not Patient-specific.
const TestMessagingServiceToEHR = `{
    "nbf": "",
    "exp": "",
    "jti": "TestPayloadJWTId",
    "body": {
        "data": [
        {
            "attributes": {
            },
            "id": "zQ<multiformat(Base58(SHA-256(TestCompositionIURI)))>", // for blockchain certification
            "resource": {
                "resourceType": "Bundle",
                "type": "document",
                "id": "TestBundleId",
                "entry": [
                    {
                        "fullUrl": "<baseUrl/tenant-1/cds-test/v1/health/fhir/r4/Composition/<composition-uuid-v4>",
                        "resource": {
                            "resourceType": "Composition",
                            "id": "urn:es.unid:cds:v1:health:person:x-uhc-mid:uuid:<person-uuid-v4>:org.hl7.fhir.r4.Composition:uuid:<composition-uuid-v4>"
                        }
                    },
                    {
                        "fullUrl": "<baseUrl/tenant-1/cds-test/v1/health/fhir/r4/directory/11369-6/Encounter/<encounter-uuid-v4>",
                        "resource": {
                            "resourceType": "Encounter",
                            "id": "urn:es.unid:cds:v1:health:person:x-uhc-mid:uuid:<person-uuid-v4>:org.hl7.fhir.r4.Encounter:uuid:<encounter-uuid-v4>"
                        }
                    },
                    {
                        "fullUrl": "<baseUrl/tenant-1/cds-test/v1/health/fhir/r4/directory/11369-6/Immunization/<immunization-uuid-v4>",
                        "resource": {
                            "resourceType": "Immunization",
                            "id": "urn:es.unid:cds:v1:health:person:x-uhc-mid:uuid:<person-uuid-v4>:org.hl7.fhir.r4.Immunization:uuid:<immunization-uuid-v4>"
                        }
                    },                   
                    {
                        "fullUrl": "<baseUrl/tenant-1/cds-test/v1/health/fhir/r4/Medication/<medication-uuid-v4>",
                        "resource": {
                            "resourceType": "Medication",
                            "id": "urn:es.unid:cds-<territory>:v1:health:org.hl7.fhir.r4.Medication:uuid:<medication-uuid-v4>"
                        }
                    },
                    {
                        "fullUrl": "<baseUrl/tenant-1/cds-test/v1/health/fhir/r4/directory/11369-6/ImmunizationRecommendation/<immunizationrecommendation-uuid-v4>",
                        "resource": {
                            "resourceType": "ImmunizationRecommendation",
                            "id": "urn:es.unid:cds:v1:health:person:x-uhc-mid:uuid:<person-uuid-v4>:org.hl7.fhir.r4.Immunization:uuid:<immunizationrecommendation-uuid-v4>"
                        }
                    },
                    {
                        "fullUrl": "<baseUrl/tenant-1/cds-test/v1/health/fhir/r4/directory/11369-6/AdverseEvent/<adverseevent-uuid-v4>",
                        "resource": {
                            "resourceType": "AdverseEvent",
                            "id": "urn:es.unid:cds-<territory>:v1:health:org.hl7.fhir.r4.AdverseEvent:uuid:<adverseevent-uuid-v4>"
                        }
                    }                                 
                ]
            },
            "type": "org.hl7.fhir.r4.bundle.document"
        },
        {
            "id": "zQ<multiformat(Base58(SHA-256(TestCommunicationIURI)))>",
            "type": "org.hl7.fhir.r4.communication",
            "fullUrl": "<baseUrl/tenant-1/cds-test/v1/health/fhir/r4/Composition/<composition-uuid-v4>",
            "resource": {
                "resourceType": "Communication",
                "id": "urn:es.unid:cds:v1:health:person:x-uhc-mid:uuid:<person-uuid-v4>:org.hl7.fhir.r4.Communication:uuid:<communication-uuid-v4>",
                "recipient": [{
                    "reference": "did:...:relatedperson:uuid:<mother>"
                },{
                    "reference": "did:...:relatedperson:uuid:<father>"
                }],
                "sender": {
                    "reference": "did:...:healthcareService"
                },
                "payload": [{
                    "note": [{
                        "author": {
                            "authorReference": {
                                "reference": "did:<employee>"
                            }
                        },
                        "text": "ha comido",
                        "time": "2017-01-01T00:00:00.000Z"
                    }, {
                        "author": {
                            "authorReference": {
                                "reference": "did:<employee>"
                            }
                        },
                        "text": "ha dormido",
                        "time": "2017-01-01T00:00:00.000Z"
                    }]
                }]
            }
        }]
    }
}`