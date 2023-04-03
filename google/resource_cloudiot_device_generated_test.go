// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

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

func TestAccCloudIotDevice_cloudiotDeviceBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":       GetTestProjectFromEnv(),
		"region":        GetTestRegionFromEnv(),
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckCloudIotDeviceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudIotDevice_cloudiotDeviceBasicExample(context),
			},
			{
				ResourceName:            "google_cloudiot_device.test-device",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"registry"},
			},
		},
	})
}

func testAccCloudIotDevice_cloudiotDeviceBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_cloudiot_registry" "registry" {
  name     = "tf-test-cloudiot-device-registry%{random_suffix}"
}

resource "google_cloudiot_device" "test-device" {
  name     = "tf-test-cloudiot-device%{random_suffix}"
  registry = google_cloudiot_registry.registry.id
}
`, context)
}

func TestAccCloudIotDevice_cloudiotDeviceFullExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":       GetTestProjectFromEnv(),
		"region":        GetTestRegionFromEnv(),
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckCloudIotDeviceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudIotDevice_cloudiotDeviceFullExample(context),
			},
			{
				ResourceName:            "google_cloudiot_device.test-device",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"registry"},
			},
		},
	})
}

func testAccCloudIotDevice_cloudiotDeviceFullExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_cloudiot_registry" "registry" {
  name     = "tf-test-cloudiot-device-registry%{random_suffix}"
}

resource "google_cloudiot_device" "test-device" {
  name     = "tf-test-cloudiot-device%{random_suffix}"
  registry = google_cloudiot_registry.registry.id

  credentials {
    public_key {
        format = "RSA_PEM"
        key = file("test-fixtures/rsa_public.pem")
    }
  }

  blocked = false

  log_level = "INFO"

  metadata = {
    test_key_1 = "test_value_1"
  }

  gateway_config {
    gateway_type = "NON_GATEWAY"
  }
}
`, context)
}

func testAccCheckCloudIotDeviceDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_cloudiot_device" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{CloudIotBasePath}}{{registry}}/devices/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = SendRequest(config, "GET", billingProject, url, config.UserAgent, nil)
			if err == nil {
				return fmt.Errorf("CloudIotDevice still exists at %s", url)
			}
		}

		return nil
	}
}
