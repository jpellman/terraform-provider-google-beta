// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccApiGatewayApi_apigatewayApiBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckApiGatewayApiDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccApiGatewayApi_apigatewayApiBasicExample(context),
			},
			{
				ResourceName:            "google_api_gateway_api.api",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"api_id"},
			},
		},
	})
}

func testAccApiGatewayApi_apigatewayApiBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_api_gateway_api" "api" {
  provider = google-beta
  api_id = "api%{random_suffix}"
}
`, context)
}

func TestAccApiGatewayApi_apigatewayApiFullExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckApiGatewayApiDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccApiGatewayApi_apigatewayApiFullExample(context),
			},
			{
				ResourceName:            "google_api_gateway_api.api",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"api_id"},
			},
		},
	})
}

func testAccApiGatewayApi_apigatewayApiFullExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_api_gateway_api" "api" {
  provider = google-beta
  api_id = "api%{random_suffix}"
  display_name = "MM Dev API"
  labels = {
    environment = "dev"
  }
}
`, context)
}

func testAccCheckApiGatewayApiDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_api_gateway_api" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := googleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{ApiGatewayBasePath}}projects/{{project}}/locations/global/apis/{{api_id}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = sendRequest(config, "GET", billingProject, url, config.userAgent, nil)
			if err == nil {
				return fmt.Errorf("ApiGatewayApi still exists at %s", url)
			}
		}

		return nil
	}
}
