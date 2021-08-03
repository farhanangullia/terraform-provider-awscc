// Code generated by generators/resource/main.go; DO NOT EDIT.

package appmesh_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-aws-cloudapi/internal/acctest"
)

type virtualServiceTest struct{}

func TestAccAWSAppMeshVirtualService_basic(t *testing.T) {
	data := acctest.NewTestData(t, "AWS::AppMesh::VirtualService", "aws_appmesh_virtual_service", "test")
	r := virtualServiceTest{}

	data.ResourceTest(t, []resource.TestStep{
		{
			Config: r.basic(data),

			ExpectError: regexp.MustCompile(`Missing required argument`),
		},
	})
}

func (r virtualServiceTest) basic(data acctest.TestData) string {
	return fmt.Sprintf(`
resource %[1]q %[2]q {
  provider = cloudapi
}
`, data.TerraformResourceType, data.ResourceLabel)
}
