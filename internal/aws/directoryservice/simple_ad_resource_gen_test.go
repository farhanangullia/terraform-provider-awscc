// Code generated by generators/resource/main.go; DO NOT EDIT.

package directoryservice_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSDirectoryServiceSimpleAD_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::DirectoryService::SimpleAD", "awscc_directoryservice_simple_ad", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}