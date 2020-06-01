// +build localstack

package test

import (
	"fmt"
	"path"
	// "regexp"
	"testing"

	"github.com/gruntwork-io/terratest/modules/files"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/require"
)

func TestSFN_basic(t *testing.T) {
	t.Parallel()

	sfnName := fmt.Sprintf("sfn-%s", random.UniqueId())
	exampleDir := "../examples/basic/"

	terraformOptions := SetupExample(t, sfnName, exampleDir)
	t.Logf("Terraform module inputs: %+v", *terraformOptions)
	// defer terraform.Destroy(t, terraformOptions)

	TerraformApplyAndValidateOutputs(t, terraformOptions)
	// require.Regexp(t, regexp.MustCompile("arn:aws:apigateway:us-east-1:lambda:path/*"), terraform.Output(t, terraformOptions, "invoke_arn"))
}

func TestSFN_lambda(t *testing.T) {
	t.Parallel()

	sfnName := fmt.Sprintf("sfn-%s", random.UniqueId())
	exampleDir := "../examples/lambda/"

	terraformOptions := SetupExample(t, sfnName, exampleDir)
	t.Logf("Terraform module inputs: %+v", *terraformOptions)
	// defer terraform.Destroy(t, terraformOptions)

	TerraformApplyAndValidateOutputs(t, terraformOptions)
	// require.Regexp(t, regexp.MustCompile("arn:aws:apigateway:us-east-1:lambda:path/*"), terraform.Output(t, terraformOptions, "invoke_arn"))
}

func SetupExample(t *testing.T, sfnName string, exampleDir string) *terraform.Options {

	localstackConfigDestination := path.Join(exampleDir, "localstack.tf")
	files.CopyFile("fixtures/localstack.tf", localstackConfigDestination)
	t.Logf("Copied localstack file to: %s", localstackConfigDestination)

	// lambdaFunctionDestination := path.Join(exampleDir, "foo.zip")
	// files.CopyFile("fixtures/foo.zip", lambdaFunctionDestination)
	// t.Logf("Copied lambda file to: %s", lambdaFunctionDestination)

	terraformOptions := &terraform.Options{
		TerraformDir: exampleDir,
		Vars: map[string]interface{}{
			"sfn_name": sfnName,
		},
	}
	return terraformOptions
}

func TerraformApplyAndValidateOutputs(t *testing.T, terraformOptions *terraform.Options) {
	terraformApplyOutput := terraform.InitAndApply(t, terraformOptions)
	resourceCount := terraform.GetResourceCount(t, terraformApplyOutput)

	require.Greater(t, resourceCount.Add, 0)
	require.Equal(t, resourceCount.Change, 0)
	require.Equal(t, resourceCount.Destroy, 0)

	// require.Regexp(t, terraform.Output(t, terraformOptions, "arn"), fmt.Sprintf("arn:aws:lambda:us-east-1:000000000000:function:%s", terraformOptions.Vars["function_name"]))
}

// func ValidateSQSTriggerOutputs(t *testing.T, terraformOptions *terraform.Options, isFifo bool) {
// 	dlq := terraform.OutputMap(t, terraformOptions, "dlq")
// 	queue := terraform.OutputMap(t, terraformOptions, "queue")

// 	expectedDlqName := fmt.Sprintf("%s-dlq", terraformOptions.Vars["function_name"])
// 	expectedQueueName := fmt.Sprintf("%s", terraformOptions.Vars["function_name"])

// 	if isFifo == true {
// 		expectedDlqName = fmt.Sprintf("%s.fifo", expectedDlqName)
// 		expectedQueueName = fmt.Sprintf("%s.fifo", expectedQueueName)
// 	}

// 	require.Equal(t, fmt.Sprintf("arn:aws:sqs:us-east-1:000000000000:%s", expectedDlqName), dlq["arn"])
// 	require.Regexp(t, regexp.MustCompile("http://*"), dlq["id"])
// 	require.Regexp(t, regexp.MustCompile("http://*"), dlq["url"])
// 	require.Equal(t, dlq["id"], dlq["url"])

// 	require.Equal(t, fmt.Sprintf("arn:aws:sqs:us-east-1:000000000000:%s", expectedQueueName), queue["arn"])
// 	require.Regexp(t, regexp.MustCompile("http://*"), queue["id"])
// 	require.Regexp(t, regexp.MustCompile("http://*"), queue["url"])
// 	require.Equal(t, queue["id"], queue["url"])

// 	require.NotEqual(t, queue["id"], dlq["id"])
// 	require.NotEqual(t, queue["arn"], dlq["arn"])

// }
