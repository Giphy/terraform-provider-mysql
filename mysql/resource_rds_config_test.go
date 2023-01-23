package mysql

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccResourceRDS(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccRDSConfig_basic(10, 30),
				Check:  testAccRDSConfigExists("mysql_rds_config.doesntexist"),
			},
		},
	})
}

func testAccRDSConfig_basic(binlog int, replication int) string {
	return fmt.Sprintf(`
resource "mysql_rds_config" "test" {
                binlog_retention_period = %d
                replication_target_delay = %d
}`, binlog, replication)
}

func testAccRDSConfigExists(rn string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[rn]
		if !ok {
			return fmt.Errorf("resource not found: %s", rn)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("RDS config id not set")
		}

		return nil
	}
}
