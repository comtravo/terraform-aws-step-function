// +build aws

package test

import (
	"fmt"
	"regexp"
	"testing"

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
	defer terraform.Destroy(t, terraformOptions)

	TerraformApplyAndValidateOutputs(t, terraformOptions)
}

func TestSFN_lambda(t *testing.T) {
	t.Parallel()

	sfnName := fmt.Sprintf("sfn-%s", random.UniqueId())
	exampleDir := "../examples/lambda/"

	terraformOptions := SetupExample(t, sfnName, exampleDir)
	t.Logf("Terraform module inputs: %+v", *terraformOptions)
	defer terraform.Destroy(t, terraformOptions)

	TerraformApplyAndValidateOutputs(t, terraformOptions)
}

func SetupExample(t *testing.T, sfnName string, exampleDir string) *terraform.Options {
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

	require.Regexp(t,
		regexp.MustCompile(fmt.Sprintf("arn:aws:states:us-east-1:\\d{12}:stateMachine:%s", terraformOptions.Vars["sfn_name"])),
		terraform.Output(t, terraformOptions, "arn"),
	)
}
