package datalineage_test

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/hashicorp/terraform-provider-google/google/acctest"
	"github.com/hashicorp/terraform-provider-google/google/envvar"
	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"

	"google.golang.org/api/googleapi"
)

var (
	_ = fmt.Sprintf
	_ = log.Print
	_ = strconv.Atoi
	_ = strings.Trim
	_ = time.Now
	_ = resource.TestMain
	_ = terraform.NewState
	_ = envvar.TestEnvVar
	_ = tpgresource.SetLabels
	_ = transport_tpg.Config{}
	_ = googleapi.Error{}
)

func TestAccDataLineageConfig_update(t *testing.T) {
	context := map[string]interface{}{
		"project": envvar.GetTestProjectFromEnv(),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataLineageConfigDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataLineageConfig_basic(context),
			},
			{
				ResourceName:            "google_data_lineage_config.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location", "parent"},
			},
			{
				Config: testAccDataLineageConfig_update(context),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction("google_data_lineage_config.default", plancheck.ResourceActionUpdate),
					},
				},
			},
			{
				ResourceName:            "google_data_lineage_config.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location", "parent"},
			},
		},
	})
}

func testAccDataLineageConfig_basic(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_data_lineage_config" "default" {
    parent = "projects/%{project}"
    location = "global"

    ingestion {
        rule {
            integration_selector {
                integration = "DATAPROC"
            }
            lineage_enablement {
                enabled = true
            }
        }
		rule {
            integration_selector {
                integration = "LOOKER_CORE"
            }
            lineage_enablement {
                enabled = true
            }
        }
    }
}
`, context)
}

func testAccDataLineageConfig_update(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_data_lineage_config" "default" {
    parent = "projects/%{project}"
    location = "global"

    ingestion {
        rule {
            integration_selector {
                integration = "DATAPROC"
            }
            lineage_enablement {
                enabled = true
            }
        }
    }
}
`, context)
}
