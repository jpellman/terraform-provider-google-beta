// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
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

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccComputePacketMirroring_computePacketMirroringFullExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(10),
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckComputePacketMirroringDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccComputePacketMirroring_computePacketMirroringFullExample(context),
			},
		},
	})
}

func testAccComputePacketMirroring_computePacketMirroringFullExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_instance" "mirror" {
  name = "tf-test-my-instance%{random_suffix}"
  provider = google-beta
  machine_type = "n1-standard-1"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-9"
    }
  }

  network_interface {
    network = google_compute_network.default.self_link
    access_config {
    }
  }
}

resource "google_compute_packet_mirroring" "foobar" {
  name = "tf-test-my-mirroring%{random_suffix}"
  provider = google-beta
  description = "bar"
  network {
    url = google_compute_network.default.self_link
  }
  collector_ilb {
    url = google_compute_forwarding_rule.default.self_link
  }
  mirrored_resources {
    tags = ["foo"]
    instances {
      url = google_compute_instance.mirror.self_link
    }
  }
  filter {
    ip_protocols = ["tcp"]
    cidr_ranges = ["0.0.0.0/0"]
  }
}
resource "google_compute_network" "default" {
  name = "tf-test-my-network%{random_suffix}"
  provider = google-beta
}

resource "google_compute_subnetwork" "default" {
  name = "tf-test-my-subnetwork%{random_suffix}"
  provider = google-beta
  network       = google_compute_network.default.self_link
  ip_cidr_range = "10.2.0.0/16"

}

resource "google_compute_region_backend_service" "default" {
  name = "tf-test-my-service%{random_suffix}"
  provider = google-beta
  health_checks = ["${google_compute_health_check.default.self_link}"]
}

resource "google_compute_health_check" "default" {
  name = "tf-test-my-healthcheck%{random_suffix}"
  provider = google-beta
  check_interval_sec = 1
  timeout_sec        = 1
  tcp_health_check {
    port = "80"
  }
}

resource "google_compute_forwarding_rule" "default" {
  depends_on = [google_compute_subnetwork.default]
  provider = google-beta
  name       = "tf-test-my-ilb%{random_suffix}"

  is_mirroring_collector = true
  ip_protocol            = "TCP"
  load_balancing_scheme  = "INTERNAL"
  backend_service        = google_compute_region_backend_service.default.self_link
  all_ports              = true
  network                = google_compute_network.default.self_link
  subnetwork             = google_compute_subnetwork.default.self_link
  network_tier           = "PREMIUM"
}
`, context)
}

func testAccCheckComputePacketMirroringDestroy(s *terraform.State) error {
	for name, rs := range s.RootModule().Resources {
		if rs.Type != "google_compute_packet_mirroring" {
			continue
		}
		if strings.HasPrefix(name, "data.") {
			continue
		}

		config := testAccProvider.Meta().(*Config)

		url, err := replaceVarsForTest(config, rs, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/packetMirrorings/{{name}}")
		if err != nil {
			return err
		}

		_, err = sendRequest(config, "GET", "", url, nil)
		if err == nil {
			return fmt.Errorf("ComputePacketMirroring still exists at %s", url)
		}
	}

	return nil
}
