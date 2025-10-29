package testimpl

import (
	"context"
	"encoding/json"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/gruntwork-io/terratest/modules/terraform"
	lcafTypes "github.com/launchbynttdata/lcaf-component-terratest/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
failedToGetKeyPolicyMsg   = "Failed to get KMS key policy"
failedToParseKeyPolicyMsg = "Failed to parse KMS key policy"
policyNotNilMsg           = "Key policy should not be nil"
statementArrayMsg         = "Policy should have at least one statement"
statementPrincipalMsg     = "Statement should have Principal"
expectedSid               = "EnableIAMUserPermissions"
expectedEffect            = "Allow"
defaultPolicyName         = "default"
)

func TestComposableComplete(t *testing.T, ctx lcafTypes.TestContext) {
	kmsClient := GetAWSKMSClient(t)

	// Get the policy ID output from the module
	policyId := terraform.Output(t, ctx.TerratestTerraformOptions(), "policy_id")

	t.Run("TestKMSKeyPolicyOutput", func(t *testing.T) {
testKMSKeyPolicyOutput(t, policyId)
})

	t.Run("TestKMSKeyPolicyExists", func(t *testing.T) {
testKMSKeyPolicyExists(t, kmsClient, policyId)
})

	t.Run("TestKMSKeyPolicyStructure", func(t *testing.T) {
testKMSKeyPolicyStructure(t, kmsClient, policyId)
})

	t.Run("TestKMSKeyPolicyPermissions", func(t *testing.T) {
testKMSKeyPolicyPermissions(t, kmsClient, policyId)
})
}

// testKMSKeyPolicyOutput verifies that the policy ID output is not empty and has the expected format
func testKMSKeyPolicyOutput(t *testing.T, policyId string) {
	assert.NotEmpty(t, policyId, "Policy ID should not be empty")
	// The policy ID should be the KMS key ID since aws_kms_key_policy resource ID is the key ID
	assert.NotEmpty(t, policyId, "Policy ID output should be returned from the module")
}

// testKMSKeyPolicyExists verifies that the KMS key policy exists and can be retrieved
func testKMSKeyPolicyExists(t *testing.T, kmsClient *kms.Client, keyId string) {
	policyOutput, err := kmsClient.GetKeyPolicy(context.TODO(), &kms.GetKeyPolicyInput{
		KeyId:      &keyId,
		PolicyName: aws.String(defaultPolicyName),
	})
	require.NoError(t, err, failedToGetKeyPolicyMsg)
	require.NotNil(t, policyOutput.Policy, policyNotNilMsg)
	assert.NotEmpty(t, *policyOutput.Policy, "Key policy should not be empty")
}

// testKMSKeyPolicyStructure verifies the basic structure of the KMS key policy
func testKMSKeyPolicyStructure(t *testing.T, kmsClient *kms.Client, keyId string) {
	policyOutput, err := kmsClient.GetKeyPolicy(context.TODO(), &kms.GetKeyPolicyInput{
		KeyId:      &keyId,
		PolicyName: aws.String(defaultPolicyName),
	})
	require.NoError(t, err, failedToGetKeyPolicyMsg)
	require.NotNil(t, policyOutput.Policy, policyNotNilMsg)

	// Parse the key policy JSON
	var keyPolicy map[string]interface{}
	err = json.Unmarshal([]byte(*policyOutput.Policy), &keyPolicy)
	require.NoError(t, err, failedToParseKeyPolicyMsg)

	// Verify basic IAM policy structure
	assert.Contains(t, keyPolicy, "Version", "Key policy should have Version field")
	assert.Contains(t, keyPolicy, "Statement", "Key policy should have Statement field")

	// Verify the Version is a valid IAM policy version
	version, ok := keyPolicy["Version"].(string)
	require.True(t, ok, "Version should be a string")
	assert.Contains(t, []string{"2012-10-17", "2008-10-17"}, version, "Version should be a valid IAM policy version")

	// Verify Statement is an array
	statements, ok := keyPolicy["Statement"].([]interface{})
	require.True(t, ok, "Statement should be an array")
	assert.Greater(t, len(statements), 0, statementArrayMsg)
}

// testKMSKeyPolicyPermissions verifies the specific permissions configured in the KMS key policy
func testKMSKeyPolicyPermissions(t *testing.T, kmsClient *kms.Client, keyId string) {
	policyOutput, err := kmsClient.GetKeyPolicy(context.TODO(), &kms.GetKeyPolicyInput{
		KeyId:      &keyId,
		PolicyName: aws.String(defaultPolicyName),
	})
	require.NoError(t, err, failedToGetKeyPolicyMsg)
	require.NotNil(t, policyOutput.Policy, policyNotNilMsg)

	// Parse the key policy JSON
	var keyPolicy map[string]interface{}
	err = json.Unmarshal([]byte(*policyOutput.Policy), &keyPolicy)
	require.NoError(t, err, failedToParseKeyPolicyMsg)

	statements, ok := keyPolicy["Statement"].([]interface{})
	require.True(t, ok, "Statement should be an array")
	require.Greater(t, len(statements), 0, statementArrayMsg)

	// Find the EnableIAMUserPermissions statement
	foundStatement := findStatementBySid(statements, expectedSid)
	require.NotNil(t, foundStatement, "Should find statement with Sid: %s", expectedSid)

	// Verify the statement structure
	verifyStatementStructure(t, foundStatement)

	// Verify Principal contains AWS account root
	verifyPrincipal(t, foundStatement)

	// Verify Action
	verifyAction(t, foundStatement)

	// Verify Resource
	verifyResource(t, foundStatement)
}

func findStatementBySid(statements []interface{}, sid string) map[string]interface{} {
	for _, stmt := range statements {
		stmtMap, ok := stmt.(map[string]interface{})
		if !ok {
			continue
		}
		stmtSid, exists := stmtMap["Sid"]
		if exists && stmtSid == sid {
			return stmtMap
		}
	}
	return nil
}

func verifyStatementStructure(t *testing.T, statement map[string]interface{}) {
	assert.Contains(t, statement, "Effect", "Statement should have Effect")
	assert.Contains(t, statement, "Principal", statementPrincipalMsg)
	assert.Contains(t, statement, "Action", "Statement should have Action")
	assert.Contains(t, statement, "Resource", "Statement should have Resource")

	// Verify Effect
	effect, ok := statement["Effect"].(string)
	require.True(t, ok, "Effect should be a string")
	assert.Equal(t, expectedEffect, effect, "Effect should be %s", expectedEffect)
}

func verifyPrincipal(t *testing.T, statement map[string]interface{}) {
	principal, exists := statement["Principal"]
	require.True(t, exists, statementPrincipalMsg)

	principalMap, ok := principal.(map[string]interface{})
	require.True(t, ok, "Principal should be a map")
	assert.Contains(t, principalMap, "AWS", "Principal should contain AWS field")

	awsPrincipal := principalMap["AWS"]
	require.NotNil(t, awsPrincipal, "AWS principal should not be nil")

	// AWS principal can be either a string or an array of strings
	foundRootPrincipal := checkForRootPrincipal(awsPrincipal)
	assert.True(t, foundRootPrincipal, "Expected to find account root principal in policy")
}

func checkForRootPrincipal(awsPrincipal interface{}) bool {
	switch v := awsPrincipal.(type) {
	case string:
		return strings.Contains(v, ":root")
	case []interface{}:
		for _, p := range v {
			if pStr, ok := p.(string); ok && strings.Contains(pStr, ":root") {
				return true
			}
		}
	}
	return false
}

func verifyAction(t *testing.T, statement map[string]interface{}) {
	action := statement["Action"]
	require.NotNil(t, action, "Action should not be nil")

	// Action can be a string or array of strings
	foundKMSAction := checkForKMSAction(action)
	assert.True(t, foundKMSAction, "Expected to find kms:* or kms: actions in policy")
}

func checkForKMSAction(action interface{}) bool {
	switch v := action.(type) {
	case string:
		return v == "kms:*" || v == "*"
	case []interface{}:
		for _, a := range v {
			if aStr, ok := a.(string); ok && (aStr == "kms:*" || strings.HasPrefix(aStr, "kms:")) {
				return true
			}
		}
	}
	return false
}

func verifyResource(t *testing.T, statement map[string]interface{}) {
	resource := statement["Resource"]
	require.NotNil(t, resource, "Resource should not be nil")

	// Resource can be a string or array of strings
	foundWildcardResource := checkForWildcardResource(resource)
	assert.True(t, foundWildcardResource, "Expected to find wildcard (*) resource in policy")
}

func checkForWildcardResource(resource interface{}) bool {
	switch v := resource.(type) {
	case string:
		return v == "*"
	case []interface{}:
		for _, r := range v {
			if rStr, ok := r.(string); ok && rStr == "*" {
				return true
			}
		}
	}
	return false
}

func GetAWSKMSClient(t *testing.T) *kms.Client {
	awsKMSClient := kms.NewFromConfig(GetAWSConfig(t))
	return awsKMSClient
}

func GetAWSConfig(t *testing.T) (cfg aws.Config) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	require.NoError(t, err, "Unable to load AWS SDK config")
	return cfg
}
