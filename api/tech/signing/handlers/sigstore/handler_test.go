package sigstore

import (
	"encoding/base64"
	"encoding/json"
	"testing"

	"github.com/sigstore/rekor/pkg/generated/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// testComponentV2 is a complete OCM component signed with Sigstore v2.6.x
// This serves as a reference for v3 compatibility testing
//
// Signing Details:
// - OCM CLI Version: 0.34.1
// - Sigstore Version: v2.6.1 (cosign/v2)
// - Component: ocm.software/podinfo-comp:1.0.0
// - Signature Name: my-ocm2
// - Hash Algorithm: SHA-256
// - Normalization: jsonNormalisation/v3
// - Signed: 2025-12-03T13:10:37Z
//
// The signature.value contains a base64-encoded JSON structure with the complete
// Rekor LogEntry, including:
// - body: base64-encoded hashedrekord
// - integratedTime: Unix timestamp
// - logID: Rekor log identifier
// - logIndex: Entry index in Rekor
// - verification: inclusionProof and signedEntryTimestamp
const testComponentV2 = `---
component:
  componentReferences: []
  creationTime: "2025-12-03T13:10:37Z"
  name: ocm.software/podinfo-comp
  provider: acme.org
  repositoryContexts: []
  resources:
  - access:
      imageReference: ghcr.io/stefanprodan/charts/podinfo:6.7.1
      type: ociArtifact
    digest:
      hashAlgorithm: SHA-256
      normalisationAlgorithm: ociArtifactDigest/v1
      value: 4d5bd3562f0b150bd6cdfdbfb149e7e4ac36e555e25baa1da9f511f4d9fb7391
    name: podinfo-image
    relation: external
    type: ociImage
    version: 1.0.0
  - access:
      imageReference: ghcr.io/stefanprodan/charts/podinfo:6.8.0
      type: ociArtifact
    digest:
      hashAlgorithm: SHA-256
      normalisationAlgorithm: ociArtifactDigest/v1
      value: 2360bdf32ddc50c05f8e128118173343b0a012a338daf145b16e0da9c80081a4
    name: podinfo-chart
    relation: external
    type: helmChart
    version: 6.8.0
  sources: []
  version: 1.0.0
meta:
  schemaVersion: v2
signatures:
- digest:
    hashAlgorithm: SHA-256
    normalisationAlgorithm: jsonNormalisation/v3
    value: fcdeec3e19adaaedb373b7a20731b784c42a8454c2cf5bdfe28b585f07784692
  name: my-ocm2
  signature:
    algorithm: sigstore
    mediaType: application/vnd.ocm.signature.sigstore
    value: eyIxMDhlOTE4NmU4YzU2NzdhNGM5OTUzODZmZTczOTg1ZDYxZjA2YjI0NGJlMzk5MDRhMWQyY2QzOGE0YjUzMzIxZWQwMWUzNGYyOTFmZjU4NiI6eyJib2R5IjoiZXlKaGNHbFdaWEp6YVc5dUlqb2lNQzR3TGpFaUxDSnJhVzVrSWpvaWFHRnphR1ZrY21WcmIzSmtJaXdpYzNCbFl5STZleUprWVhSaElqcDdJbWhoYzJnaU9uc2lZV3huYjNKcGRHaHRJam9pYzJoaE1qVTJJaXdpZG1Gc2RXVWlPaUptWTJSbFpXTXpaVEU1WVdSaFlXVmtZak0zTTJJM1lUSXdOek14WWpjNE5HTTBNbUU0TkRVMFl6SmpaalZpWkdabE1qaGlOVGcxWmpBM056ZzBOamt5SW4xOUxDSnphV2R1WVhSMWNtVWlPbnNpWTI5dWRHVnVkQ0k2SWsxRlZVTkpVVU5EY2podE4xVkJiSGRWWjBweE5XSkxTME5OYkhNcmJtUnRWRmhtUWxWeGIweElSVVV4TVc5RlZUaFJTV2RoUlRkUU4yTkhNelZCZVVwVlkyaEVXVVpvYlhGelNXMXFSbWt5YVVOaGRISkhZbXRFU0d3NFFWcFpQU0lzSW5CMVlteHBZMHRsZVNJNmV5SmpiMjUwWlc1MElqb2lURk13ZEV4VE1VTlNWV1JLVkdsQ1VWWlZTazFUVlUxblV6QldXa3hUTUhSTVV6QkxWRlZhY21Rd1ZqTlhWV2hNWWpGd1NtVnRiM2RSTUVaU1YxVnNUR0l4Y0VwbGJXOTNVa1ZHVWxrd1VsSmFNRVpHWW01a1VWTnRlRWhSV0VwRFVWaFdiMlZyWTNoVlNFWkZUVEl4U1dNd1pGWlZha3A0WlZGd1FtVnNhRVZNTUdReFpETmFjVlpIV25aYU1rNUlUMFpXY2xaRVVsQlVXRXA2VmtoYWQxWlVTbmxOZWtaSVlWaFdSMkV5YkV4bFJteFlWVVpvTlUxVVFuTldibG8xVmxod2FtTnJaRE5RVkRCTFRGTXdkRXhUTVVaVWExRm5WVVpXUTFSRmJFUkpSWFJHVjFNd2RFeFRNSFJEWnowOUluMTlmWDA9IiwiaW50ZWdyYXRlZFRpbWUiOjE3NjQ3Njc0NzMsImxvZ0lEIjoiYzBkMjNkNmFkNDA2OTczZjk1NTlmM2JhMmQxY2EwMWY4NDE0N2Q4ZmZjNWI4NDQ1YzIyNGY5OGI5NTkxODAxZCIsImxvZ0luZGV4Ijo3MzgwODQxODgsInZlcmlmaWNhdGlvbiI6eyJpbmNsdXNpb25Qcm9vZiI6eyJjaGVja3BvaW50IjoicmVrb3Iuc2lnc3RvcmUuZGV2IC0gMTE5MzA1MDk1OTkxNjY1NjUwNlxuNjE2MTc5OTI4XG5FNXNaR0ZvOXFWU2JqY2F1aEFlZkRpTy9HK3JUS1pxYzdVK1lGRCtMd3VZPVxuXG7igJQgcmVrb3Iuc2lnc3RvcmUuZGV2IHdOSTlhakJGQWlCRTdGd1A5QWlEM3RWOE56R0sweTVOQjBBL3QwTHdIdkhCL2pkSWRYemtoQUloQUtiSEkxWko1OUZDbnJLZmRFYUdpMjg1VVJoVjdNVVNHVEQ0MnRUTDlnRm1cbiIsImhhc2hlcyI6WyI1MDQ5MjNjMjgyYTFmMDJhYmMyNWJlNzY1ZmM4OGQ4OTA0YzE3ZmY3Njk0MzFkNjhiOGIzOTQyOTVjMzIwMzQyIiwiYzRlZDczMjIwOThiYzQxMjRhNjdhNDRjZDAxODA0NDIwZDVhZjI1MDQxMGQ0ZmI4ZjA0MTEyZGUwZDAyZDM4NiIsIjgxNmVjYTFjOTYzODFmNDgyNTRiMDU1MTI3MTY1OWFkMTIwYWZmN2M2NTA5M2M2OTA4Mjg0NmM3MTc2N2YyOTUiLCJiZWMwY2EwOGUxYzdlZGY2YTE5OTQyNTkzZTk4ZjIxZmUyNzJiMTkzNWFiYjczMzhkY2RmZWQ3ZWMwYTJkNWE0IiwiODM2YTcxZDYxNWYxNjA1OTM0OTkxM2Q1M2JjMGYxZmE1ZjRjNjlmM2M2MDAxZDdjY2Q5ZWJkYTRhN2FkNzE5OCIsIjU2ZmQzYmZiZWEwY2UyNGM5ZTQxMzNiZDE3ODA1ZWQ3YWMxYmJmMDdlMTNlNWRlYTIzNTc5MDBkNDk2ZDRhZDciLCJiOGJmMzBmMjc4ZGE3ZjEyYTg3YWM3N2MxNGMwODk2MWEyZjMxYzJmMTVjZGQ3MjkyNWMyNDZjMjZiMTdiZGYwIiwiMTQ2M2YxMjM2ZWFlNTE5OWJkMmYyYzI4MmZmZDkyOThkNzVmMGE1ODM1ZWIzOGRkYmZiYjE2YTM0ZGQ4ZmI5NiIsIjljZDg3OGNkNDY1YzRlN2UyNzkyM2M3ZjQ4NWJjY2QxNzRiYjlmMjg4YmJlOTUxNjA3NzM2YzIxNDczMDMxNWEiLCJlMmU2MDIyYjc5MjVkMzc3YTIwZjdlZjMwOWVmY2YxOGVkZTQwOGZmZjIyYzI0OTYxMWE0ZDg5MGIxYTI3ZTJhIiwiNmE4YjE2MzIzYjNhYWIwYTRiZjM0OGY4OTI4Y2I1NzcyM2JmYmNkNmI5NDNhMzY1YmYxMzEzZmI2NGY2MWIwMCIsIjU5ZmQxOWY5ODZjMTJmYzk4ZjlmZDAxZTMwMjc1Y2M3MWI5MmM1Y2VjM2JhMTUxNTNmNTg4NTdlNWM4MmZlMTciLCI4NmU2ODdkNTk0YTQ0NzIzZjhmYWZhNjYxYWRhNjY5MjQyZmNlMzk5NjE3MjYzN2Q4YThmMGMyMWVlOTJhMWJkIiwiNjY2NTI0NjI0MWMxY2I1MDdiZGI3MjZiMTIwODhhYmRlYTUzNzQ3NjJiM2ZhY2I2NmI4YTBlMGQ4YmUyZTU1NiIsIjRmODBlYTU4M2UzNjg0MGI0ZGZhZjVmYzhjYTA5NmFhODBiODk5ZTEzODI1ZTkwOGY0YmM1ODE4MjcwZmNiNTMiXSwibG9nSW5kZXgiOjYxNjE3OTkyNiwicm9vdEhhc2giOiIxMzliMTkxODVhM2RhOTU0OWI4ZGM2YWU4NDA3OWYwZTIzYmYxYmVhZDMyOTlhOWNlZDRmOTgxNDNmOGJjMmU2IiwidHJlZVNpemUiOjYxNjE3OTkyOH0sInNpZ25lZEVudHJ5VGltZXN0YW1wIjoiTUVRQ0lIZnhSOUdnOXBjRENRelNNVFNSUXRZVWFmY055bm1COEpqbzFxWk50TTgvQWlBUWdicjNwb3VWeWRzMUozNWRvaGNTUHhWdEpNWmx0ak1ic01mL2VzYndpdz09In19fQ==`

// Extract signature value and digest from the component YAML for easier access in tests
const testSignatureV2 = `eyIxMDhlOTE4NmU4YzU2NzdhNGM5OTUzODZmZTczOTg1ZDYxZjA2YjI0NGJlMzk5MDRhMWQyY2QzOGE0YjUzMzIxZWQwMWUzNGYyOTFmZjU4NiI6eyJib2R5IjoiZXlKaGNHbFdaWEp6YVc5dUlqb2lNQzR3TGpFaUxDSnJhVzVrSWpvaWFHRnphR1ZrY21WcmIzSmtJaXdpYzNCbFl5STZleUprWVhSaElqcDdJbWhoYzJnaU9uc2lZV3huYjNKcGRHaHRJam9pYzJoaE1qVTJJaXdpZG1Gc2RXVWlPaUptWTJSbFpXTXpaVEU1WVdSaFlXVmtZak0zTTJJM1lUSXdOek14WWpjNE5HTTBNbUU0TkRVMFl6SmpaalZpWkdabE1qaGlOVGcxWmpBM056ZzBOamt5SW4xOUxDSnphV2R1WVhSMWNtVWlPbnNpWTI5dWRHVnVkQ0k2SWsxRlZVTkpVVU5EY2podE4xVkJiSGRWWjBweE5XSkxTME5OYkhNcmJtUnRWRmhtUWxWeGIweElSVVV4TVc5RlZUaFJTV2RoUlRkUU4yTkhNelZCZVVwVlkyaEVXVVpvYlhGelNXMXFSbWt5YVVOaGRISkhZbXRFU0d3NFFWcFpQU0lzSW5CMVlteHBZMHRsZVNJNmV5SmpiMjUwWlc1MElqb2lURk13ZEV4VE1VTlNWV1JLVkdsQ1VWWlZTazFUVlUxblV6QldXa3hUTUhSTVV6QkxWRlZhY21Rd1ZqTlhWV2hNWWpGd1NtVnRiM2RSTUVaU1YxVnNUR0l4Y0VwbGJXOTNVa1ZHVWxrd1VsSmFNRVpHWW01a1VWTnRlRWhSV0VwRFVWaFdiMlZyWTNoVlNFWkZUVEl4U1dNd1pGWlZha3A0WlZGd1FtVnNhRVZNTUdReFpETmFjVlpIV25aYU1rNUlUMFpXY2xaRVVsQlVXRXA2VmtoYWQxWlVTbmxOZWtaSVlWaFdSMkV5YkV4bFJteFlWVVpvTlUxVVFuTldibG8xVmxod2FtTnJaRE5RVkRCTFRGTXdkRXhUTVVaVWExRm5WVVpXUTFSRmJFUkpSWFJHVjFNd2RFeFRNSFJEWnowOUluMTlmWDA9IiwiaW50ZWdyYXRlZFRpbWUiOjE3NjQ3Njc0NzMsImxvZ0lEIjoiYzBkMjNkNmFkNDA2OTczZjk1NTlmM2JhMmQxY2EwMWY4NDE0N2Q4ZmZjNWI4NDQ1YzIyNGY5OGI5NTkxODAxZCIsImxvZ0luZGV4Ijo3MzgwODQxODgsInZlcmlmaWNhdGlvbiI6eyJpbmNsdXNpb25Qcm9vZiI6eyJjaGVja3BvaW50IjoicmVrb3Iuc2lnc3RvcmUuZGV2IC0gMTE5MzA1MDk1OTkxNjY1NjUwNlxuNjE2MTc5OTI4XG5FNXNaR0ZvOXFWU2JqY2F1aEFlZkRpTy9HK3JUS1pxYzdVK1lGRCtMd3VZPVxuXG7igJQgcmVrb3Iuc2lnc3RvcmUuZGV2IHdOSTlhakJGQWlCRTdGd1A5QWlEM3RWOE56R0sweTVOQjBBL3QwTHdIdkhCL2pkSWRYemtoQUloQUtiSEkxWko1OUZDbnJLZmRFYUdpMjg1VVJoVjdNVVNHVEQ0MnRUTDlnRm1cbiIsImhhc2hlcyI6WyI1MDQ5MjNjMjgyYTFmMDJhYmMyNWJlNzY1ZmM4OGQ4OTA0YzE3ZmY3Njk0MzFkNjhiOGIzOTQyOTVjMzIwMzQyIiwiYzRlZDczMjIwOThiYzQxMjRhNjdhNDRjZDAxODA0NDIwZDVhZjI1MDQxMGQ0ZmI4ZjA0MTEyZGUwZDAyZDM4NiIsIjgxNmVjYTFjOTYzODFmNDgyNTRiMDU1MTI3MTY1OWFkMTIwYWZmN2M2NTA5M2M2OTA4Mjg0NmM3MTc2N2YyOTUiLCJiZWMwY2EwOGUxYzdlZGY2YTE5OTQyNTkzZTk4ZjIxZmUyNzJiMTkzNWFiYjczMzhkY2RmZWQ3ZWMwYTJkNWE0IiwiODM2YTcxZDYxNWYxNjA1OTM0OTkxM2Q1M2JjMGYxZmE1ZjRjNjlmM2M2MDAxZDdjY2Q5ZWJkYTRhN2FkNzE5OCIsIjU2ZmQzYmZiZWEwY2UyNGM5ZTQxMzNiZDE3ODA1ZWQ3YWMxYmJmMDdlMTNlNWRlYTIzNTc5MDBkNDk2ZDRhZDciLCJiOGJmMzBmMjc4ZGE3ZjEyYTg3YWM3N2MxNGMwODk2MWEyZjMxYzJmMTVjZGQ3MjkyNWMyNDZjMjZiMTdiZGYwIiwiMTQ2M2YxMjM2ZWFlNTE5OWJkMmYyYzI4MmZmZDkyOThkNzVmMGE1ODM1ZWIzOGRkYmZiYjE2YTM0ZGQ4ZmI5NiIsIjljZDg3OGNkNDY1YzRlN2UyNzkyM2M3ZjQ4NWJjY2QxNzRiYjlmMjg4YmJlOTUxNjA3NzM2YzIxNDczMDMxNWEiLCJlMmU2MDIyYjc5MjVkMzc3YTIwZjdlZjMwOWVmY2YxOGVkZTQwOGZmZjIyYzI0OTYxMWE0ZDg5MGIxYTI3ZTJhIiwiNmE4YjE2MzIzYjNhYWIwYTRiZjM0OGY4OTI4Y2I1NzcyM2JmYmNkNmI5NDNhMzY1YmYxMzEzZmI2NGY2MWIwMCIsIjU5ZmQxOWY5ODZjMTJmYzk4ZjlmZDAxZTMwMjc1Y2M3MWI5MmM1Y2VjM2JhMTUxNTNmNTg4NTdlNWM4MmZlMTciLCI4NmU2ODdkNTk0YTQ0NzIzZjhmYWZhNjYxYWRhNjY5MjQyZmNlMzk5NjE3MjYzN2Q4YThmMGMyMWVlOTJhMWJkIiwiNjY2NTI0NjI0MWMxY2I1MDdiZGI3MjZiMTIwODhhYmRlYTUzNzQ3NjJiM2ZhY2I2NmI4YTBlMGQ4YmUyZTU1NiIsIjRmODBlYTU4M2UzNjg0MGI0ZGZhZjVmYzhjYTA5NmFhODBiODk5ZTEzODI1ZTkwOGY0YmM1ODE4MjcwZmNiNTMiXSwibG9nSW5kZXgiOjYxNjE3OTkyNiwicm9vdEhhc2giOiIxMzliMTkxODVhM2RhOTU0OWI4ZGM2YWU4NDA3OWYwZTIzYmYxYmVhZDMyOTlhOWNlZDRmOTgxNDNmOGJjMmU2IiwidHJlZVNpemUiOjYxNjE3OTkyOH0sInNpZ25lZEVudHJ5VGltZXN0YW1wIjoiTUVRQ0lIZnhSOUdnOXBjRENRelNNVFNSUXRZVWFmY055bm1COEpqbzFxWk50TTgvQWlBUWdicjNwb3VWeWRzMUozNWRvaGNTUHhWdEpNWmx0ak1ic01mL2VzYndpdz09In19fQ==`

const testDigest = "fcdeec3e19adaaedb373b7a20731b784c42a8454c2cf5bdfe28b585f07784692"

// TestSignatureStructureV2 documents the structure of a v2.6.x signature
// This test serves as a baseline for v3 compatibility testing
func TestSignatureStructureV2(t *testing.T) {
	// Decode the base64 signature
	data, err := base64.StdEncoding.DecodeString(testSignatureV2)
	require.NoError(t, err, "Failed to decode base64 signature")

	// Parse as JSON
	var entries models.LogEntry
	err = json.Unmarshal(data, &entries)
	require.NoError(t, err, "Failed to unmarshal LogEntry")

	// Verify structure
	assert.NotEmpty(t, entries, "LogEntry should not be empty")

	// Document the structure for comparison with v3
	t.Logf("=== V2 Signature Structure ===")
	t.Logf("Number of entries: %d", len(entries))

	for logID, entry := range entries {
		t.Logf("\nLog ID: %s", logID)
		t.Logf("Log Index: %d", *entry.LogIndex)
		t.Logf("Integrated Time: %d", *entry.IntegratedTime)
		t.Logf("Log ID (from entry): %s", *entry.LogID)

		// Verify body exists
		assert.NotNil(t, entry.Body, "Entry body should not be nil")

		// Verify verification data exists
		assert.NotNil(t, entry.Verification, "Verification should not be nil")
		if entry.Verification != nil {
			assert.NotNil(t, entry.Verification.InclusionProof, "InclusionProof should not be nil")
			assert.NotNil(t, entry.Verification.SignedEntryTimestamp, "SignedEntryTimestamp should not be nil")
		}

		// Decode and verify the body contains the expected digest
		bodyData, err := base64.StdEncoding.DecodeString(entry.Body.(string))
		require.NoError(t, err, "Failed to decode entry body")

		var rekorEntry models.Hashedrekord
		err = json.Unmarshal(bodyData, &rekorEntry)
		require.NoError(t, err, "Failed to unmarshal Hashedrekord")

		// Verify the digest matches
		rekorEntrySpec := rekorEntry.Spec.(map[string]interface{})
		rekorHashValue := rekorEntrySpec["data"].(map[string]interface{})["hash"].(map[string]interface{})["value"]
		assert.Equal(t, testDigest, rekorHashValue, "Digest should match")

		t.Logf("Digest verified: %s", testDigest)
	}
}

// TestV2SignatureFormat documents the exact JSON structure for v3 comparison
func TestV2SignatureFormat(t *testing.T) {
	data, err := base64.StdEncoding.DecodeString(testSignatureV2)
	require.NoError(t, err)

	// Pretty print the structure
	var prettyJSON map[string]interface{}
	err = json.Unmarshal(data, &prettyJSON)
	require.NoError(t, err)

	formatted, err := json.MarshalIndent(prettyJSON, "", "  ")
	require.NoError(t, err)

	t.Logf("=== V2 Signature JSON Structure ===\n%s", string(formatted))

	// Save to file for reference
	t.Logf("\nThis structure should be compared with v3 output")
}

// TestRekorEntryFieldsV2 documents all fields present in v2 LogEntry
func TestRekorEntryFieldsV2(t *testing.T) {
	data, err := base64.StdEncoding.DecodeString(testSignatureV2)
	require.NoError(t, err)

	var entries models.LogEntry
	err = json.Unmarshal(data, &entries)
	require.NoError(t, err)

	for _, entry := range entries {
		t.Logf("=== V2 LogEntry Fields ===")
		t.Logf("Body: %v", entry.Body != nil)
		t.Logf("IntegratedTime: %v", entry.IntegratedTime != nil)
		t.Logf("LogID: %v", entry.LogID != nil)
		t.Logf("LogIndex: %v", entry.LogIndex != nil)
		t.Logf("Verification: %v", entry.Verification != nil)
		
		if entry.Verification != nil {
			t.Logf("  - InclusionProof: %v", entry.Verification.InclusionProof != nil)
			t.Logf("  - SignedEntryTimestamp: %v", entry.Verification.SignedEntryTimestamp != nil)
			
			if entry.Verification.InclusionProof != nil {
				proof := entry.Verification.InclusionProof
				t.Logf("    - Checkpoint: %v", proof.Checkpoint != nil)
				t.Logf("    - Hashes: %d items", len(proof.Hashes))
				t.Logf("    - LogIndex: %v", proof.LogIndex != nil)
				t.Logf("    - RootHash: %v", proof.RootHash != nil)
				t.Logf("    - TreeSize: %v", proof.TreeSize != nil)
			}
		}
	}
}
