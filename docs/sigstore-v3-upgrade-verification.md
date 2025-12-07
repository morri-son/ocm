# Sigstore Cosign v3 Upgrade Verification

**Upgrade PR**: [#1726 - Upgrade Sigstore Cosign from v2.6.1 to v3.0.2](https://github.com/open-component-model/ocm/pull/1726)

**Test Branch**: [feat/upgrade-sigstore-v3](https://github.com/morri-son/ocm/tree/feat/upgrade-sigstore-v3)

## Pre-signed Test Component

Since Cosign v3 API introduces changes that require additional attributes for OIDC during signing, the test flow cannot happen using Github Actions alone.
Therefore, a pre-signed test component with both v2 and v3 signatures was created manually.

### Component Details

**Registry Location**: `ghcr.io/morri-son/ocm-test//ocm.software/sigstore-test-comp:1.0.0`

**Component Definition**:
```yaml
components:
- name: ocm.software/sigstore-test-comp
  version: 1.0.0
  provider:
    name: ocm.software
  resources:
    - name: test-image
      version: 1.0.0
      type: ociImage
      access:
        type: ociArtifact
        imageReference: ghcr.io/stefanprodan/charts/podinfo:6.7.1
```

### Manual Creation Process

Create component with **two** keyless signatures, publicly accessible for verification testing.

#### Signature with Sigstore Cosign v2

```bash
# Using OCM CLI v0.34.1 (with Sigstore v2.6.1)
/tmp/ocm sign cv \
  --signature sigstore-v2 \
  --algorithm sigstore \
  --keyless \
  /tmp/ctf-sigstore//ocm.software/sigstore-test-comp:1.0.0
```

#### Signature with Sigstore Cosign v3
```bash
# Using OCM CLI from feat/upgrade-sigstore-v3 branch (with Sigstore v3.0.2)
~/go/bin/ocm sign cv \
  --signature sigstore-v3 \
  --algorithm sigstore \
  --keyless \
  /tmp/ctf-sigstore//ocm.software/sigstore-test-comp:1.0.0
```

#### Upload to Registry
```bash
# Transfer to public registry
ocm transfer cv \
  /tmp/ctf-sigstore//ocm.software/sigstore-test-comp:1.0.0 \
  ghcr.io/morri-son/ocm-test
```

## Automated Verification Tests

Verify that signature verification works across all combinations of CLI versions and signature types.

### Test Matrix

The workflow tests all 4 possible verification combinations:

| Test Case | CLI Version | Signature Type | Purpose |
|-----------|-------------|----------------|---------|
| 1 | v2 (v0.34.1) | sigstore-v2 | **Baseline** - Current production scenario |
| 2 | v2 (v0.34.1) | sigstore-v3 | **Forward Compatibility** - Old CLI verifies new signatures |
| 3 | v3 (feat/upgrade-sigstore-v3) | sigstore-v2 | **Backward Compatibility** - New CLI verifies old signatures |
| 4 | v3 (feat/upgrade-sigstore-v3) | sigstore-v3 | **Target State** - New CLI with new signatures |

## Test Results

Action run https://github.com/morri-son/ocm/actions/runs/20003483428 showed that all 4 test cases pass, proving:

✅ **Backward Compatibility**: v3 CLI can verify v2 signatures
- Existing signatures continue to work after upgrade
- No breaking changes for verification

✅ **Forward Compatibility**: v2 CLI can verify v3 signatures
- New signatures work with old CLI versions
- Smooth migration path

✅ **Signature Format Stability**: Signature format is compatible across versions
- No changes to signature structure
- Rekor entries remain valid

Based on the test results, the Sigstore v3 upgrade:
- ✅ Maintains full backward compatibility
- ✅ Maintains full forward compatibility
- ✅ Introduces no breaking changes
- ✅ Is safe to deploy

### Recommendation

**The upgrade from Sigstore v2.6.1 to v3.0.2 can be safely performed.**

All existing signatures will continue to work, and new signatures will be compatible with older OCM versions.
